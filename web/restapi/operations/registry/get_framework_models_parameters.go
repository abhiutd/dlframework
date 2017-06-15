package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFrameworkModelsParams creates a new GetFrameworkModelsParams object
// with the default values initialized.
func NewGetFrameworkModelsParams() GetFrameworkModelsParams {
	var ()
	return GetFrameworkModelsParams{}
}

// GetFrameworkModelsParams contains all the bound params for the get framework models operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetFrameworkModels
type GetFrameworkModelsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	FrameworkName string
	/*
	  Required: true
	  In: path
	*/
	FrameworkVersion string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetFrameworkModelsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rFrameworkName, rhkFrameworkName, _ := route.Params.GetOK("framework_name")
	if err := o.bindFrameworkName(rFrameworkName, rhkFrameworkName, route.Formats); err != nil {
		res = append(res, err)
	}

	rFrameworkVersion, rhkFrameworkVersion, _ := route.Params.GetOK("framework_version")
	if err := o.bindFrameworkVersion(rFrameworkVersion, rhkFrameworkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFrameworkModelsParams) bindFrameworkName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.FrameworkName = raw

	return nil
}

func (o *GetFrameworkModelsParams) bindFrameworkVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.FrameworkVersion = raw

	return nil
}
