// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NotifyHandlerFunc turns a function with the right signature into a notify handler
type NotifyHandlerFunc func(NotifyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NotifyHandlerFunc) Handle(params NotifyParams) middleware.Responder {
	return fn(params)
}

// NotifyHandler interface for that can handle valid notify params
type NotifyHandler interface {
	Handle(NotifyParams) middleware.Responder
}

// NewNotify creates a new http.Handler for the notify operation
func NewNotify(ctx *middleware.Context, handler NotifyHandler) *Notify {
	return &Notify{Context: ctx, Handler: handler}
}

/*Notify swagger:route POST /notifications status notify

get the current status to the instance

*/
type Notify struct {
	Context *middleware.Context
	Handler NotifyHandler
}

func (o *Notify) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNotifyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
