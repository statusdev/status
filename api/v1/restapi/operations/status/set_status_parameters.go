// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewSetStatusParams creates a new SetStatusParams object
// no default values defined in spec.
func NewSetStatusParams() SetStatusParams {

	return SetStatusParams{}
}

// SetStatusParams contains all the bound params for the set status operation
// typically these are obtained from a http.Request
//
// swagger:parameters setStatus
type SetStatusParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The media to upload
	  In: formData
	*/
	Media io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSetStatusParams() beforehand.
func (o *SetStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	media, mediaHeader, err := r.FormFile("media")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "media", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindMedia(media, mediaHeader); err != nil {
		res = append(res, err)
	} else {
		o.Media = &runtime.File{Data: media, Header: mediaHeader}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMedia binds file parameter Media.
//
// The only supported validations on files are MinLength and MaxLength
func (o *SetStatusParams) bindMedia(file multipart.File, header *multipart.FileHeader) error {
	return nil
}
