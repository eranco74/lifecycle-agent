// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// UpdateHostInstallProgressReader is a Reader for the UpdateHostInstallProgress structure.
type UpdateHostInstallProgressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHostInstallProgressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateHostInstallProgressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateHostInstallProgressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateHostInstallProgressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateHostInstallProgressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdateHostInstallProgressMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateHostInstallProgressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateHostInstallProgressOK creates a UpdateHostInstallProgressOK with default headers values
func NewUpdateHostInstallProgressOK() *UpdateHostInstallProgressOK {
	return &UpdateHostInstallProgressOK{}
}

/*UpdateHostInstallProgressOK handles this case with default header values.

Update install progress
*/
type UpdateHostInstallProgressOK struct {
}

func (o *UpdateHostInstallProgressOK) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/hosts/{host_id}/progress][%d] updateHostInstallProgressOK ", 200)
}

func (o *UpdateHostInstallProgressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateHostInstallProgressUnauthorized creates a UpdateHostInstallProgressUnauthorized with default headers values
func NewUpdateHostInstallProgressUnauthorized() *UpdateHostInstallProgressUnauthorized {
	return &UpdateHostInstallProgressUnauthorized{}
}

/*UpdateHostInstallProgressUnauthorized handles this case with default header values.

Unauthorized.
*/
type UpdateHostInstallProgressUnauthorized struct {
	Payload *models.InfraError
}

func (o *UpdateHostInstallProgressUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/hosts/{host_id}/progress][%d] updateHostInstallProgressUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateHostInstallProgressUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateHostInstallProgressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHostInstallProgressForbidden creates a UpdateHostInstallProgressForbidden with default headers values
func NewUpdateHostInstallProgressForbidden() *UpdateHostInstallProgressForbidden {
	return &UpdateHostInstallProgressForbidden{}
}

/*UpdateHostInstallProgressForbidden handles this case with default header values.

Forbidden.
*/
type UpdateHostInstallProgressForbidden struct {
	Payload *models.InfraError
}

func (o *UpdateHostInstallProgressForbidden) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/hosts/{host_id}/progress][%d] updateHostInstallProgressForbidden  %+v", 403, o.Payload)
}

func (o *UpdateHostInstallProgressForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateHostInstallProgressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHostInstallProgressNotFound creates a UpdateHostInstallProgressNotFound with default headers values
func NewUpdateHostInstallProgressNotFound() *UpdateHostInstallProgressNotFound {
	return &UpdateHostInstallProgressNotFound{}
}

/*UpdateHostInstallProgressNotFound handles this case with default header values.

Error.
*/
type UpdateHostInstallProgressNotFound struct {
	Payload *models.Error
}

func (o *UpdateHostInstallProgressNotFound) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/hosts/{host_id}/progress][%d] updateHostInstallProgressNotFound  %+v", 404, o.Payload)
}

func (o *UpdateHostInstallProgressNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHostInstallProgressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHostInstallProgressMethodNotAllowed creates a UpdateHostInstallProgressMethodNotAllowed with default headers values
func NewUpdateHostInstallProgressMethodNotAllowed() *UpdateHostInstallProgressMethodNotAllowed {
	return &UpdateHostInstallProgressMethodNotAllowed{}
}

/*UpdateHostInstallProgressMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type UpdateHostInstallProgressMethodNotAllowed struct {
	Payload *models.Error
}

func (o *UpdateHostInstallProgressMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/hosts/{host_id}/progress][%d] updateHostInstallProgressMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UpdateHostInstallProgressMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHostInstallProgressMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHostInstallProgressInternalServerError creates a UpdateHostInstallProgressInternalServerError with default headers values
func NewUpdateHostInstallProgressInternalServerError() *UpdateHostInstallProgressInternalServerError {
	return &UpdateHostInstallProgressInternalServerError{}
}

/*UpdateHostInstallProgressInternalServerError handles this case with default header values.

Error.
*/
type UpdateHostInstallProgressInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateHostInstallProgressInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/hosts/{host_id}/progress][%d] updateHostInstallProgressInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateHostInstallProgressInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHostInstallProgressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
