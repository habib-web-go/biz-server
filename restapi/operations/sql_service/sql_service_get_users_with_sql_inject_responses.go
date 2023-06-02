// Code generated by go-swagger; DO NOT EDIT.

package sql_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/habib/biz/models"
)

// SQLServiceGetUsersWithSQLInjectOKCode is the HTTP code returned for type SQLServiceGetUsersWithSQLInjectOK
const SQLServiceGetUsersWithSQLInjectOKCode int = 200

/*
SQLServiceGetUsersWithSQLInjectOK A successful response.

swagger:response sqlServiceGetUsersWithSqlInjectOK
*/
type SQLServiceGetUsersWithSQLInjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.SQLInjectionGetUsersResponse `json:"body,omitempty"`
}

// NewSQLServiceGetUsersWithSQLInjectOK creates SQLServiceGetUsersWithSQLInjectOK with default headers values
func NewSQLServiceGetUsersWithSQLInjectOK() *SQLServiceGetUsersWithSQLInjectOK {

	return &SQLServiceGetUsersWithSQLInjectOK{}
}

// WithPayload adds the payload to the sql service get users with Sql inject o k response
func (o *SQLServiceGetUsersWithSQLInjectOK) WithPayload(payload *models.SQLInjectionGetUsersResponse) *SQLServiceGetUsersWithSQLInjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sql service get users with Sql inject o k response
func (o *SQLServiceGetUsersWithSQLInjectOK) SetPayload(payload *models.SQLInjectionGetUsersResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SQLServiceGetUsersWithSQLInjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
SQLServiceGetUsersWithSQLInjectDefault An unexpected error response.

swagger:response sqlServiceGetUsersWithSqlInjectDefault
*/
type SQLServiceGetUsersWithSQLInjectDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.RuntimeError `json:"body,omitempty"`
}

// NewSQLServiceGetUsersWithSQLInjectDefault creates SQLServiceGetUsersWithSQLInjectDefault with default headers values
func NewSQLServiceGetUsersWithSQLInjectDefault(code int) *SQLServiceGetUsersWithSQLInjectDefault {
	if code <= 0 {
		code = 500
	}

	return &SQLServiceGetUsersWithSQLInjectDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the SQL service get users with Sql inject default response
func (o *SQLServiceGetUsersWithSQLInjectDefault) WithStatusCode(code int) *SQLServiceGetUsersWithSQLInjectDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the SQL service get users with Sql inject default response
func (o *SQLServiceGetUsersWithSQLInjectDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the SQL service get users with Sql inject default response
func (o *SQLServiceGetUsersWithSQLInjectDefault) WithPayload(payload *models.RuntimeError) *SQLServiceGetUsersWithSQLInjectDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the SQL service get users with Sql inject default response
func (o *SQLServiceGetUsersWithSQLInjectDefault) SetPayload(payload *models.RuntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SQLServiceGetUsersWithSQLInjectDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}