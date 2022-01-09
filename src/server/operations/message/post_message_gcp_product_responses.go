// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostMessageGcpProductCreatedCode is the HTTP code returned for type PostMessageGcpProductCreated
const PostMessageGcpProductCreatedCode int = 201

/*PostMessageGcpProductCreated OK

swagger:response postMessageGcpProductCreated
*/
type PostMessageGcpProductCreated struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostMessageGcpProductCreated creates PostMessageGcpProductCreated with default headers values
func NewPostMessageGcpProductCreated() *PostMessageGcpProductCreated {

	return &PostMessageGcpProductCreated{}
}

// WithPayload adds the payload to the post message gcp product created response
func (o *PostMessageGcpProductCreated) WithPayload(payload interface{}) *PostMessageGcpProductCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post message gcp product created response
func (o *PostMessageGcpProductCreated) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMessageGcpProductCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}