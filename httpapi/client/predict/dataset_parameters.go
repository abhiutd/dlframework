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

// NewDatasetParams creates a new DatasetParams object
// with the default values initialized.
func NewDatasetParams() *DatasetParams {
	var ()
	return &DatasetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDatasetParamsWithTimeout creates a new DatasetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDatasetParamsWithTimeout(timeout time.Duration) *DatasetParams {
	var ()
	return &DatasetParams{

		timeout: timeout,
	}
}

// NewDatasetParamsWithContext creates a new DatasetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDatasetParamsWithContext(ctx context.Context) *DatasetParams {
	var ()
	return &DatasetParams{

		Context: ctx,
	}
}

// NewDatasetParamsWithHTTPClient creates a new DatasetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDatasetParamsWithHTTPClient(client *http.Client) *DatasetParams {
	var ()
	return &DatasetParams{
		HTTPClient: client,
	}
}

/*DatasetParams contains all the parameters to send to the API endpoint
for the dataset operation typically these are written to a http.Request
*/
type DatasetParams struct {

	/*Body*/
	Body *models.DlframeworkDatasetRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dataset params
func (o *DatasetParams) WithTimeout(timeout time.Duration) *DatasetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dataset params
func (o *DatasetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dataset params
func (o *DatasetParams) WithContext(ctx context.Context) *DatasetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dataset params
func (o *DatasetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dataset params
func (o *DatasetParams) WithHTTPClient(client *http.Client) *DatasetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dataset params
func (o *DatasetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the dataset params
func (o *DatasetParams) WithBody(body *models.DlframeworkDatasetRequest) *DatasetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the dataset params
func (o *DatasetParams) SetBody(body *models.DlframeworkDatasetRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DatasetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.DlframeworkDatasetRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
