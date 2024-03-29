// Code generated by go-swagger; DO NOT EDIT.

package sql_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewSQLServiceGetUsersParams creates a new SQLServiceGetUsersParams object
//
// There are no default values defined in the spec.
func NewSQLServiceGetUsersParams() SQLServiceGetUsersParams {

	return SQLServiceGetUsersParams{}
}

// SQLServiceGetUsersParams contains all the bound params for the SQL service get users operation
// typically these are obtained from a http.Request
//
// swagger:parameters SQLService_getUsers
type SQLServiceGetUsersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	AuthKey *string
	/*
	  In: query
	*/
	MessageID *string
	/*
	  In: query
	*/
	UserID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSQLServiceGetUsersParams() beforehand.
func (o *SQLServiceGetUsersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAuthKey, qhkAuthKey, _ := qs.GetOK("authKey")
	if err := o.bindAuthKey(qAuthKey, qhkAuthKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qMessageID, qhkMessageID, _ := qs.GetOK("messageId")
	if err := o.bindMessageID(qMessageID, qhkMessageID, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserID, qhkUserID, _ := qs.GetOK("userId")
	if err := o.bindUserID(qUserID, qhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthKey binds and validates parameter AuthKey from query.
func (o *SQLServiceGetUsersParams) bindAuthKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.AuthKey = &raw

	return nil
}

// bindMessageID binds and validates parameter MessageID from query.
func (o *SQLServiceGetUsersParams) bindMessageID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.MessageID = &raw

	return nil
}

// bindUserID binds and validates parameter UserID from query.
func (o *SQLServiceGetUsersParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.UserID = &raw

	return nil
}
