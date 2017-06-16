package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFrameworkManifestHandlerFunc turns a function with the right signature into a get framework manifest handler
type GetFrameworkManifestHandlerFunc func(GetFrameworkManifestParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFrameworkManifestHandlerFunc) Handle(params GetFrameworkManifestParams) middleware.Responder {
	return fn(params)
}

// GetFrameworkManifestHandler interface for that can handle valid get framework manifest params
type GetFrameworkManifestHandler interface {
	Handle(GetFrameworkManifestParams) middleware.Responder
}

// NewGetFrameworkManifest creates a new http.Handler for the get framework manifest operation
func NewGetFrameworkManifest(ctx *middleware.Context, handler GetFrameworkManifestHandler) *GetFrameworkManifest {
	return &GetFrameworkManifest{Context: ctx, Handler: handler}
}

/*GetFrameworkManifest swagger:route GET /v1/framework/{framework_name}/info registry getFrameworkManifest

GetFrameworkManifest get framework manifest API

*/
type GetFrameworkManifest struct {
	Context *middleware.Context
	Handler GetFrameworkManifestHandler
}

func (o *GetFrameworkManifest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetFrameworkManifestParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}