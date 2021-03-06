// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/statusdev/status/api/v1/models"
)

// SetStatusOKCode is the HTTP code returned for type SetStatusOK
const SetStatusOKCode int = 200

/*SetStatusOK OK

swagger:response setStatusOK
*/
type SetStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.Status `json:"body,omitempty"`
}

// NewSetStatusOK creates SetStatusOK with default headers values
func NewSetStatusOK() *SetStatusOK {

	return &SetStatusOK{}
}

// WithPayload adds the payload to the set status o k response
func (o *SetStatusOK) WithPayload(payload *models.Status) *SetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set status o k response
func (o *SetStatusOK) SetPayload(payload *models.Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetStatusUnauthorizedCode is the HTTP code returned for type SetStatusUnauthorized
const SetStatusUnauthorizedCode int = 401

/*SetStatusUnauthorized unauthorized

swagger:response setStatusUnauthorized
*/
type SetStatusUnauthorized struct {
}

// NewSetStatusUnauthorized creates SetStatusUnauthorized with default headers values
func NewSetStatusUnauthorized() *SetStatusUnauthorized {

	return &SetStatusUnauthorized{}
}

// WriteResponse to the client
func (o *SetStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SetStatusInternalServerErrorCode is the HTTP code returned for type SetStatusInternalServerError
const SetStatusInternalServerErrorCode int = 500

/*SetStatusInternalServerError internal server error

swagger:response setStatusInternalServerError
*/
type SetStatusInternalServerError struct {
}

// NewSetStatusInternalServerError creates SetStatusInternalServerError with default headers values
func NewSetStatusInternalServerError() *SetStatusInternalServerError {

	return &SetStatusInternalServerError{}
}

// WriteResponse to the client
func (o *SetStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
