// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SetStatusHandlerFunc turns a function with the right signature into a set status handler
type SetStatusHandlerFunc func(SetStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SetStatusHandlerFunc) Handle(params SetStatusParams) middleware.Responder {
	return fn(params)
}

// SetStatusHandler interface for that can handle valid set status params
type SetStatusHandler interface {
	Handle(SetStatusParams) middleware.Responder
}

// NewSetStatus creates a new http.Handler for the set status operation
func NewSetStatus(ctx *middleware.Context, handler SetStatusHandler) *SetStatus {
	return &SetStatus{Context: ctx, Handler: handler}
}

/*SetStatus swagger:route POST /status status setStatus

add a new status to the instance

*/
type SetStatus struct {
	Context *middleware.Context
	Handler SetStatusHandler
}

func (o *SetStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSetStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
