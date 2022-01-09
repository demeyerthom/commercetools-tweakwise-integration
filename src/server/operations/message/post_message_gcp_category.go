// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostMessageGcpCategoryHandlerFunc turns a function with the right signature into a post message gcp category handler
type PostMessageGcpCategoryHandlerFunc func(PostMessageGcpCategoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMessageGcpCategoryHandlerFunc) Handle(params PostMessageGcpCategoryParams) middleware.Responder {
	return fn(params)
}

// PostMessageGcpCategoryHandler interface for that can handle valid post message gcp category params
type PostMessageGcpCategoryHandler interface {
	Handle(PostMessageGcpCategoryParams) middleware.Responder
}

// NewPostMessageGcpCategory creates a new http.Handler for the post message gcp category operation
func NewPostMessageGcpCategory(ctx *middleware.Context, handler PostMessageGcpCategoryHandler) *PostMessageGcpCategory {
	return &PostMessageGcpCategory{Context: ctx, Handler: handler}
}

/* PostMessageGcpCategory swagger:route POST /message/gcp/category message gcp postMessageGcpCategory

Receives a GCP message for a category update

*/
type PostMessageGcpCategory struct {
	Context *middleware.Context
	Handler PostMessageGcpCategoryHandler
}

func (o *PostMessageGcpCategory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostMessageGcpCategoryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
