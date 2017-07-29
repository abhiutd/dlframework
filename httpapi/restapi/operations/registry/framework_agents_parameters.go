// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFrameworkAgentsParams creates a new FrameworkAgentsParams object
// with the default values initialized.
func NewFrameworkAgentsParams() FrameworkAgentsParams {
	var ()
	return FrameworkAgentsParams{}
}

// FrameworkAgentsParams contains all the bound params for the framework agents operation
// typically these are obtained from a http.Request
//
// swagger:parameters FrameworkAgents
type FrameworkAgentsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  In: query
	*/
	FrameworkName *string
	/*
	  In: query
	*/
	FrameworkVersion *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FrameworkAgentsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFrameworkName, qhkFrameworkName, _ := qs.GetOK("framework_name")
	if err := o.bindFrameworkName(qFrameworkName, qhkFrameworkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qFrameworkVersion, qhkFrameworkVersion, _ := qs.GetOK("framework_version")
	if err := o.bindFrameworkVersion(qFrameworkVersion, qhkFrameworkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FrameworkAgentsParams) bindFrameworkName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FrameworkName = &raw

	return nil
}

func (o *FrameworkAgentsParams) bindFrameworkVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FrameworkVersion = &raw

	return nil
}