// Code generated by go-swagger; DO NOT EDIT.

package subscribers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddSubscriberHandlerFunc turns a function with the right signature into a add subscriber handler
type AddSubscriberHandlerFunc func(AddSubscriberParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddSubscriberHandlerFunc) Handle(params AddSubscriberParams) middleware.Responder {
	return fn(params)
}

// AddSubscriberHandler interface for that can handle valid add subscriber params
type AddSubscriberHandler interface {
	Handle(AddSubscriberParams) middleware.Responder
}

// NewAddSubscriber creates a new http.Handler for the add subscriber operation
func NewAddSubscriber(ctx *middleware.Context, handler AddSubscriberHandler) *AddSubscriber {
	return &AddSubscriber{Context: ctx, Handler: handler}
}

/*AddSubscriber swagger:route POST /subscribers subscribers addSubscriber

subscribes to the instance

*/
type AddSubscriber struct {
	Context *middleware.Context
	Handler AddSubscriberHandler
}

func (o *AddSubscriber) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddSubscriberParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
