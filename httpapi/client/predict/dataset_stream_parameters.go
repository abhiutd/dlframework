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

// NewDatasetStreamParams creates a new DatasetStreamParams object
// with the default values initialized.
func NewDatasetStreamParams() *DatasetStreamParams {
	var ()
	return &DatasetStreamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDatasetStreamParamsWithTimeout creates a new DatasetStreamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDatasetStreamParamsWithTimeout(timeout time.Duration) *DatasetStreamParams {
	var ()
	return &DatasetStreamParams{

		timeout: timeout,
	}
}

// NewDatasetStreamParamsWithContext creates a new DatasetStreamParams object
// with the default values initialized, and the ability to set a context for a request
func NewDatasetStreamParamsWithContext(ctx context.Context) *DatasetStreamParams {
	var ()
	return &DatasetStreamParams{

		Context: ctx,
	}
}

// NewDatasetStreamParamsWithHTTPClient creates a new DatasetStreamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDatasetStreamParamsWithHTTPClient(client *http.Client) *DatasetStreamParams {
	var ()
	return &DatasetStreamParams{
		HTTPClient: client,
	}
}

/*DatasetStreamParams contains all the parameters to send to the API endpoint
for the dataset stream operation typically these are written to a http.Request
*/
type DatasetStreamParams struct {

	/*Body*/
	Body *models.DlframeworkDatasetRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dataset stream params
func (o *DatasetStreamParams) WithTimeout(timeout time.Duration) *DatasetStreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dataset stream params
func (o *DatasetStreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dataset stream params
func (o *DatasetStreamParams) WithContext(ctx context.Context) *DatasetStreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dataset stream params
func (o *DatasetStreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dataset stream params
func (o *DatasetStreamParams) WithHTTPClient(client *http.Client) *DatasetStreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dataset stream params
func (o *DatasetStreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the dataset stream params
func (o *DatasetStreamParams) WithBody(body *models.DlframeworkDatasetRequest) *DatasetStreamParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the dataset stream params
func (o *DatasetStreamParams) SetBody(body *models.DlframeworkDatasetRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DatasetStreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
