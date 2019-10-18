// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/statusdev/status/api/v1/models"
)

// GetStatusOKCode is the HTTP code returned for type GetStatusOK
const GetStatusOKCode int = 200

/*GetStatusOK OK

swagger:response getStatusOK
*/
type GetStatusOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ProfileStatus `json:"body,omitempty"`
}

// NewGetStatusOK creates GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {

	return &GetStatusOK{}
}

// WithPayload adds the payload to the get status o k response
func (o *GetStatusOK) WithPayload(payload []*models.ProfileStatus) *GetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get status o k response
func (o *GetStatusOK) SetPayload(payload []*models.ProfileStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ProfileStatus, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetStatusUnauthorizedCode is the HTTP code returned for type GetStatusUnauthorized
const GetStatusUnauthorizedCode int = 401

/*GetStatusUnauthorized unauthorized

swagger:response getStatusUnauthorized
*/
type GetStatusUnauthorized struct {
}

// NewGetStatusUnauthorized creates GetStatusUnauthorized with default headers values
func NewGetStatusUnauthorized() *GetStatusUnauthorized {

	return &GetStatusUnauthorized{}
}

// WriteResponse to the client
func (o *GetStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetStatusInternalServerErrorCode is the HTTP code returned for type GetStatusInternalServerError
const GetStatusInternalServerErrorCode int = 500

/*GetStatusInternalServerError internal server error

swagger:response getStatusInternalServerError
*/
type GetStatusInternalServerError struct {
}

// NewGetStatusInternalServerError creates GetStatusInternalServerError with default headers values
func NewGetStatusInternalServerError() *GetStatusInternalServerError {

	return &GetStatusInternalServerError{}
}

// WriteResponse to the client
func (o *GetStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
