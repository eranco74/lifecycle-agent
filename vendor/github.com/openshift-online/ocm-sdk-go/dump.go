/*
Copyright (c) 2018 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file contains the implementations of the methods of the connection that are used to dump to
// the log the details of HTTP requests and responses.

package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"gitlab.com/c0b/go-ordered-json"
)

// dumpRoundTripper is a round tripper that dumps the details of the requests and the responses to
// the log.
type dumpRoundTripper struct {
	logger Logger
	next   http.RoundTripper
}

// Make sure that we implement the http.RoundTripper interface:
var _ http.RoundTripper = &dumpRoundTripper{}

// RoundTrip is he implementation of the http.RoundTripper interface.
func (d *dumpRoundTripper) RoundTrip(request *http.Request) (response *http.Response, err error) {
	// Get the context:
	ctx := request.Context()

	// Read the complete body in memory, in order to send it to the log, and replace it with a
	// reader that reads it from memory:
	if request.Body != nil {
		var body []byte
		body, err = ioutil.ReadAll(request.Body)
		if err != nil {
			return
		}
		err = request.Body.Close()
		if err != nil {
			return
		}
		d.dumpRequest(ctx, request, body)
		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	} else {
		d.dumpRequest(ctx, request, nil)
	}

	// Call the next round tripper:
	response, err = d.next.RoundTrip(request)
	if err != nil {
		return
	}

	// Read the complete response body in memory, in order to send it the log, and replace it
	// with a reader that reads it from memory:
	if response.Body != nil {
		var body []byte
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
		err = response.Body.Close()
		if err != nil {
			return
		}
		d.dumpResponse(ctx, response, body)
		response.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	} else {
		d.dumpResponse(ctx, response, nil)
	}

	return
}

const (
	// redactionStr replaces sensitive values in output.
	redactionStr = "***"
)

// redactFields are removed from log output when dumped.
var redactFields = map[string]bool{
	"access_token":  true,
	"admin":         true,
	"id_token":      true,
	"refresh_token": true,
	"password":      true,
	"client_secret": true,
	"kubeconfig":    true,
	"ssh":           true,
}

// dumpRequest dumps to the log, in debug level, the details of the given HTTP request.
func (d *dumpRoundTripper) dumpRequest(ctx context.Context, request *http.Request, body []byte) {
	d.logger.Debug(ctx, "Request method is %s", request.Method)
	d.logger.Debug(ctx, "Request URL is '%s'", request.URL)
	header := request.Header
	names := make([]string, len(header))
	i := 0
	for name := range header {
		names[i] = name
		i++
	}
	sort.Strings(names)
	for _, name := range names {
		values := header[name]
		for _, value := range values {
			if strings.ToLower(name) == "authorization" {
				d.logger.Debug(ctx, "Request header '%s' is omitted", name)
			} else {
				d.logger.Debug(ctx, "Request header '%s' is '%s'", name, value)
			}
		}
	}
	if body != nil {
		d.logger.Debug(ctx, "Request body follows")
		d.dumpBody(ctx, header, body)
	}
}

// dumpResponse dumps to the log, in debug level, the details of the given HTTP response.
func (d *dumpRoundTripper) dumpResponse(ctx context.Context, response *http.Response, body []byte) {
	d.logger.Debug(ctx, "Response status is '%s'", response.Status)
	header := response.Header
	names := make([]string, len(header))
	i := 0
	for name := range header {
		names[i] = name
		i++
	}
	sort.Strings(names)
	for _, name := range names {
		values := header[name]
		for _, value := range values {
			d.logger.Debug(ctx, "Response header '%s' is '%s'", name, value)
		}
	}
	if body != nil {
		d.logger.Debug(ctx, "Response body follows")
		d.dumpBody(ctx, header, body)
	}
}

// dumpBody checks the content type used in the given header and then it dumps the given body in a
// format suitable for that content type.
func (d *dumpRoundTripper) dumpBody(ctx context.Context, header http.Header, body []byte) {
	// Try to parse the content type:
	var mediaType string
	contentType := header.Get("Content-Type")
	if contentType != "" {
		var err error
		mediaType, _, err = mime.ParseMediaType(contentType)
		if err != nil {
			d.logger.Error(ctx, "Can't parse content type '%s': %v", contentType, err)
		}
	} else {
		mediaType = contentType
	}

	// Dump the body according to the content type:
	switch mediaType {
	case "application/x-www-form-urlencoded":
		d.dumpForm(ctx, body)
	case "application/json", "":
		d.dumpJSON(ctx, body)
	default:
		d.dumpBytes(ctx, body)
	}
}

// dumpForm sends to the log the contents of the given form data, excluding security sensitive
// fields.
func (d *dumpRoundTripper) dumpForm(ctx context.Context, data []byte) {
	// Parse the form:
	form, err := url.ParseQuery(string(data))
	if err != nil {
		d.dumpBytes(ctx, data)
		return
	}

	// Redact values corresponding to security sensitive fields:
	for name, values := range form {
		if redactFields[name] {
			for i := range values {
				values[i] = redactionStr
			}
		}
	}

	// Get and sort the names of the fields of the form, so that the generated output will be
	// predictable:
	names := make([]string, 0, len(form))
	for name := range form {
		names = append(names, name)
	}
	sort.Strings(names)

	// Write the names and values to the buffer while redacting the sensitive fields:
	buffer := &bytes.Buffer{}
	for _, name := range names {
		key := url.QueryEscape(name)
		values := form[name]
		for _, value := range values {
			var redacted string
			if redactFields[name] {
				redacted = "***"
			} else {
				redacted = url.QueryEscape(value)
			}
			if buffer.Len() > 0 {
				buffer.WriteByte('&') // #nosec G104
			}
			buffer.WriteString(key)      // #nosec G104
			buffer.WriteByte('=')        // #nosec G104
			buffer.WriteString(redacted) // #nosec G104
		}
	}

	// Send the redacted data to the log:
	d.dumpBytes(ctx, buffer.Bytes())
}

// dumpJSON tries to parse the given data as a JSON document. If that works, then it dumps it
// indented, otherwise dumps it as is.
func (d *dumpRoundTripper) dumpJSON(ctx context.Context, data []byte) {
	parsed := ordered.NewOrderedMap()
	err := json.Unmarshal(data, parsed)
	if err != nil {
		d.logger.Debug(ctx, "%s", data)
	} else {
		// remove sensitive information
		d.redactSensitive(parsed)

		indented, err := json.MarshalIndent(parsed, "", "  ")
		if err != nil {
			d.logger.Debug(ctx, "%s", data)
		} else {
			d.logger.Debug(ctx, "%s", indented)
		}
	}
}

// dumpBytes dump the given data as an array of bytes.
func (d *dumpRoundTripper) dumpBytes(ctx context.Context, data []byte) {
	d.logger.Debug(ctx, "%s", data)
}

// redactSensitive replaces sensitive fields within a response with redactionStr.
func (d *dumpRoundTripper) redactSensitive(body *ordered.OrderedMap) {
	iterator := body.EntriesIter()
	for {
		pair, ok := iterator()
		if !ok {
			break
		}
		if redactFields[pair.Key] {
			body.Set(pair.Key, redactionStr)
		}
	}
}
