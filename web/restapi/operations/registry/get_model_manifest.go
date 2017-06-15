package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetModelManifestHandlerFunc turns a function with the right signature into a get model manifest handler
type GetModelManifestHandlerFunc func(GetModelManifestParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetModelManifestHandlerFunc) Handle(params GetModelManifestParams) middleware.Responder {
	return fn(params)
}

// GetModelManifestHandler interface for that can handle valid get model manifest params
type GetModelManifestHandler interface {
	Handle(GetModelManifestParams) middleware.Responder
}

// NewGetModelManifest creates a new http.Handler for the get model manifest operation
func NewGetModelManifest(ctx *middleware.Context, handler GetModelManifestHandler) *GetModelManifest {
	return &GetModelManifest{Context: ctx, Handler: handler}
}

/*GetModelManifest swagger:route POST /v1/model/{model_name}/{model_version}/info registry getModelManifest

GetModelManifest get model manifest API

*/
type GetModelManifest struct {
	Context *middleware.Context
	Handler GetModelManifestHandler
}

func (o *GetModelManifest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetModelManifestParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
