// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDeregisterHostParams creates a new DeregisterHostParams object
// no default values defined in spec.
func NewDeregisterHostParams() DeregisterHostParams {

	return DeregisterHostParams{}
}

// DeregisterHostParams contains all the bound params for the deregister host operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeregisterHost
type DeregisterHostParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ClusterID strfmt.UUID
	/*
	  Required: true
	  In: path
	*/
	HostID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeregisterHostParams() beforehand.
func (o *DeregisterHostParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rClusterID, rhkClusterID, _ := route.Params.GetOK("cluster_id")
	if err := o.bindClusterID(rClusterID, rhkClusterID, route.Formats); err != nil {
		res = append(res, err)
	}

	rHostID, rhkHostID, _ := route.Params.GetOK("host_id")
	if err := o.bindHostID(rHostID, rhkHostID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClusterID binds and validates parameter ClusterID from path.
func (o *DeregisterHostParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("cluster_id", "path", "strfmt.UUID", raw)
	}
	o.ClusterID = *(value.(*strfmt.UUID))

	if err := o.validateClusterID(formats); err != nil {
		return err
	}

	return nil
}

// validateClusterID carries on validations for parameter ClusterID
func (o *DeregisterHostParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.FormatOf("cluster_id", "path", "uuid", o.ClusterID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindHostID binds and validates parameter HostID from path.
func (o *DeregisterHostParams) bindHostID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("host_id", "path", "strfmt.UUID", raw)
	}
	o.HostID = *(value.(*strfmt.UUID))

	if err := o.validateHostID(formats); err != nil {
		return err
	}

	return nil
}

// validateHostID carries on validations for parameter HostID
func (o *DeregisterHostParams) validateHostID(formats strfmt.Registry) error {

	if err := validate.FormatOf("host_id", "path", "uuid", o.HostID.String(), formats); err != nil {
		return err
	}
	return nil
}
