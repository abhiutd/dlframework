// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DatasetStreamHandlerFunc turns a function with the right signature into a dataset stream handler
type DatasetStreamHandlerFunc func(DatasetStreamParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DatasetStreamHandlerFunc) Handle(params DatasetStreamParams) middleware.Responder {
	return fn(params)
}

// DatasetStreamHandler interface for that can handle valid dataset stream params
type DatasetStreamHandler interface {
	Handle(DatasetStreamParams) middleware.Responder
}

// NewDatasetStream creates a new http.Handler for the dataset stream operation
func NewDatasetStream(ctx *middleware.Context, handler DatasetStreamHandler) *DatasetStream {
	return &DatasetStream{Context: ctx, Handler: handler}
}

/*DatasetStream swagger:route POST /predict/stream/dataset Predict datasetStream

Dataset method receives a single dataset and runs
the predictor on all elements of the dataset.

The result is a prediction feature stream.

*/
type DatasetStream struct {
	Context *middleware.Context
	Handler DatasetStreamHandler
}

func (o *DatasetStream) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDatasetStreamParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
