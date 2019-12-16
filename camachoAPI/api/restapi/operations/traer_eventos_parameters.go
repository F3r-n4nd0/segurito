// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTraerEventosParams creates a new TraerEventosParams object
// no default values defined in spec.
func NewTraerEventosParams() TraerEventosParams {

	return TraerEventosParams{}
}

// TraerEventosParams contains all the bound params for the traer eventos operation
// typically these are obtained from a http.Request
//
// swagger:parameters traerEventos
type TraerEventosParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Id del usuario
	  Required: true
	  In: path
	*/
	IDUsuario string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewTraerEventosParams() beforehand.
func (o *TraerEventosParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rIDUsuario, rhkIDUsuario, _ := route.Params.GetOK("id_usuario")
	if err := o.bindIDUsuario(rIDUsuario, rhkIDUsuario, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIDUsuario binds and validates parameter IDUsuario from path.
func (o *TraerEventosParams) bindIDUsuario(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.IDUsuario = raw

	return nil
}
