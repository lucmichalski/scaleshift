// Code generated by go-swagger; DO NOT EDIT.

package rescale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/rescale-labs/scaleshift/api/src/generated/models"
)

// GetRescaleApplicationVersionOKCode is the HTTP code returned for type GetRescaleApplicationVersionOK
const GetRescaleApplicationVersionOKCode int = 200

/*GetRescaleApplicationVersionOK OK

swagger:response getRescaleApplicationVersionOK
*/
type GetRescaleApplicationVersionOK struct {

	/*
	  In: Body
	*/
	Payload *models.RescaleApplicationVersion `json:"body,omitempty"`
}

// NewGetRescaleApplicationVersionOK creates GetRescaleApplicationVersionOK with default headers values
func NewGetRescaleApplicationVersionOK() *GetRescaleApplicationVersionOK {

	return &GetRescaleApplicationVersionOK{}
}

// WithPayload adds the payload to the get rescale application version o k response
func (o *GetRescaleApplicationVersionOK) WithPayload(payload *models.RescaleApplicationVersion) *GetRescaleApplicationVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rescale application version o k response
func (o *GetRescaleApplicationVersionOK) SetPayload(payload *models.RescaleApplicationVersion) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRescaleApplicationVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRescaleApplicationVersionDefault Unexpected error

swagger:response getRescaleApplicationVersionDefault
*/
type GetRescaleApplicationVersionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRescaleApplicationVersionDefault creates GetRescaleApplicationVersionDefault with default headers values
func NewGetRescaleApplicationVersionDefault(code int) *GetRescaleApplicationVersionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRescaleApplicationVersionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get rescale application version default response
func (o *GetRescaleApplicationVersionDefault) WithStatusCode(code int) *GetRescaleApplicationVersionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get rescale application version default response
func (o *GetRescaleApplicationVersionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get rescale application version default response
func (o *GetRescaleApplicationVersionDefault) WithPayload(payload *models.Error) *GetRescaleApplicationVersionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rescale application version default response
func (o *GetRescaleApplicationVersionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRescaleApplicationVersionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}