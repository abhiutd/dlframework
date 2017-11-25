// Code generated by go-swagger; DO NOT EDIT.

package predict

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

	"github.com/rai-project/dlframework/httpapi/models"
)

// NewResetParams creates a new ResetParams object
// with the default values initialized.
func NewResetParams() *ResetParams {
	var ()
	return &ResetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResetParamsWithTimeout creates a new ResetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResetParamsWithTimeout(timeout time.Duration) *ResetParams {
	var ()
	return &ResetParams{

		timeout: timeout,
	}
}

// NewResetParamsWithContext creates a new ResetParams object
// with the default values initialized, and the ability to set a context for a request
func NewResetParamsWithContext(ctx context.Context) *ResetParams {
	var ()
	return &ResetParams{

		Context: ctx,
	}
}

// NewResetParamsWithHTTPClient creates a new ResetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResetParamsWithHTTPClient(client *http.Client) *ResetParams {
	var ()
	return &ResetParams{
		HTTPClient: client,
	}
}

/*ResetParams contains all the parameters to send to the API endpoint
for the reset operation typically these are written to a http.Request
*/
type ResetParams struct {

	/*Body*/
	Body *models.DlframeworkResetRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reset params
func (o *ResetParams) WithTimeout(timeout time.Duration) *ResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset params
func (o *ResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset params
func (o *ResetParams) WithContext(ctx context.Context) *ResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset params
func (o *ResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset params
func (o *ResetParams) WithHTTPClient(client *http.Client) *ResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset params
func (o *ResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reset params
func (o *ResetParams) WithBody(body *models.DlframeworkResetRequest) *ResetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reset params
func (o *ResetParams) SetBody(body *models.DlframeworkResetRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
