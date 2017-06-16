package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFrameworkManifestParams creates a new GetFrameworkManifestParams object
// with the default values initialized.
func NewGetFrameworkManifestParams() *GetFrameworkManifestParams {
	var ()
	return &GetFrameworkManifestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFrameworkManifestParamsWithTimeout creates a new GetFrameworkManifestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFrameworkManifestParamsWithTimeout(timeout time.Duration) *GetFrameworkManifestParams {
	var ()
	return &GetFrameworkManifestParams{

		timeout: timeout,
	}
}

// NewGetFrameworkManifestParamsWithContext creates a new GetFrameworkManifestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFrameworkManifestParamsWithContext(ctx context.Context) *GetFrameworkManifestParams {
	var ()
	return &GetFrameworkManifestParams{

		Context: ctx,
	}
}

// NewGetFrameworkManifestParamsWithHTTPClient creates a new GetFrameworkManifestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFrameworkManifestParamsWithHTTPClient(client *http.Client) *GetFrameworkManifestParams {
	var ()
	return &GetFrameworkManifestParams{
		HTTPClient: client,
	}
}

/*GetFrameworkManifestParams contains all the parameters to send to the API endpoint
for the get framework manifest operation typically these are written to a http.Request
*/
type GetFrameworkManifestParams struct {

	/*FrameworkName*/
	FrameworkName string
	/*FrameworkVersion*/
	FrameworkVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get framework manifest params
func (o *GetFrameworkManifestParams) WithTimeout(timeout time.Duration) *GetFrameworkManifestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get framework manifest params
func (o *GetFrameworkManifestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get framework manifest params
func (o *GetFrameworkManifestParams) WithContext(ctx context.Context) *GetFrameworkManifestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get framework manifest params
func (o *GetFrameworkManifestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get framework manifest params
func (o *GetFrameworkManifestParams) WithHTTPClient(client *http.Client) *GetFrameworkManifestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get framework manifest params
func (o *GetFrameworkManifestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrameworkName adds the frameworkName to the get framework manifest params
func (o *GetFrameworkManifestParams) WithFrameworkName(frameworkName string) *GetFrameworkManifestParams {
	o.SetFrameworkName(frameworkName)
	return o
}

// SetFrameworkName adds the frameworkName to the get framework manifest params
func (o *GetFrameworkManifestParams) SetFrameworkName(frameworkName string) {
	o.FrameworkName = frameworkName
}

// WithFrameworkVersion adds the frameworkVersion to the get framework manifest params
func (o *GetFrameworkManifestParams) WithFrameworkVersion(frameworkVersion *string) *GetFrameworkManifestParams {
	o.SetFrameworkVersion(frameworkVersion)
	return o
}

// SetFrameworkVersion adds the frameworkVersion to the get framework manifest params
func (o *GetFrameworkManifestParams) SetFrameworkVersion(frameworkVersion *string) {
	o.FrameworkVersion = frameworkVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetFrameworkManifestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param framework_name
	if err := r.SetPathParam("framework_name", o.FrameworkName); err != nil {
		return err
	}

	if o.FrameworkVersion != nil {

		// query param framework_version
		var qrFrameworkVersion string
		if o.FrameworkVersion != nil {
			qrFrameworkVersion = *o.FrameworkVersion
		}
		qFrameworkVersion := qrFrameworkVersion
		if qFrameworkVersion != "" {
			if err := r.SetQueryParam("framework_version", qFrameworkVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}