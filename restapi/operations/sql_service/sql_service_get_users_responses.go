// Code generated by go-swagger; DO NOT EDIT.

package sql_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/habib-web-go/biz-server/models"
)

// SQLServiceGetUsersOKCode is the HTTP code returned for type SQLServiceGetUsersOK
const SQLServiceGetUsersOKCode int = 200

/*
SQLServiceGetUsersOK A successful response.

swagger:response sqlServiceGetUsersOK
*/
type SQLServiceGetUsersOK struct {

	/*
	  In: Body
	*/
	Payload *models.SQLInjectionGetUsersResponse `json:"body,omitempty"`
}

// NewSQLServiceGetUsersOK creates SQLServiceGetUsersOK with default headers values
func NewSQLServiceGetUsersOK() *SQLServiceGetUsersOK {

	return &SQLServiceGetUsersOK{}
}

// WithPayload adds the payload to the sql service get users o k response
func (o *SQLServiceGetUsersOK) WithPayload(payload *models.SQLInjectionGetUsersResponse) *SQLServiceGetUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sql service get users o k response
func (o *SQLServiceGetUsersOK) SetPayload(payload *models.SQLInjectionGetUsersResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SQLServiceGetUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
SQLServiceGetUsersDefault An unexpected error response.

swagger:response sqlServiceGetUsersDefault
*/
type SQLServiceGetUsersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.RuntimeError `json:"body,omitempty"`
}

// NewSQLServiceGetUsersDefault creates SQLServiceGetUsersDefault with default headers values
func NewSQLServiceGetUsersDefault(code int) *SQLServiceGetUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &SQLServiceGetUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the SQL service get users default response
func (o *SQLServiceGetUsersDefault) WithStatusCode(code int) *SQLServiceGetUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the SQL service get users default response
func (o *SQLServiceGetUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the SQL service get users default response
func (o *SQLServiceGetUsersDefault) WithPayload(payload *models.RuntimeError) *SQLServiceGetUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the SQL service get users default response
func (o *SQLServiceGetUsersDefault) SetPayload(payload *models.RuntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SQLServiceGetUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
