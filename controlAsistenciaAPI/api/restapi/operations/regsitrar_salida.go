// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// RegsitrarSalidaHandlerFunc turns a function with the right signature into a regsitrar salida handler
type RegsitrarSalidaHandlerFunc func(RegsitrarSalidaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegsitrarSalidaHandlerFunc) Handle(params RegsitrarSalidaParams) middleware.Responder {
	return fn(params)
}

// RegsitrarSalidaHandler interface for that can handle valid regsitrar salida params
type RegsitrarSalidaHandler interface {
	Handle(RegsitrarSalidaParams) middleware.Responder
}

// NewRegsitrarSalida creates a new http.Handler for the regsitrar salida operation
func NewRegsitrarSalida(ctx *middleware.Context, handler RegsitrarSalidaHandler) *RegsitrarSalida {
	return &RegsitrarSalida{Context: ctx, Handler: handler}
}

/*RegsitrarSalida swagger:route POST /salidas regsitrarSalida

registra la salida del personal

registra la salida del personal por medio de su codigo (7 caracteres)

*/
type RegsitrarSalida struct {
	Context *middleware.Context
	Handler RegsitrarSalidaHandler
}

func (o *RegsitrarSalida) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRegsitrarSalidaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RegsitrarSalidaBody regsitrar salida body
// swagger:model RegsitrarSalidaBody
type RegsitrarSalidaBody struct {

	// codigo usuario
	// Required: true
	// Max Length: 7
	// Min Length: 7
	CodigoUsuario *string `json:"codigo_usuario"`
}

// Validate validates this regsitrar salida body
func (o *RegsitrarSalidaBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCodigoUsuario(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegsitrarSalidaBody) validateCodigoUsuario(formats strfmt.Registry) error {

	if err := validate.Required("usuario"+"."+"codigo_usuario", "body", o.CodigoUsuario); err != nil {
		return err
	}

	if err := validate.MinLength("usuario"+"."+"codigo_usuario", "body", string(*o.CodigoUsuario), 7); err != nil {
		return err
	}

	if err := validate.MaxLength("usuario"+"."+"codigo_usuario", "body", string(*o.CodigoUsuario), 7); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RegsitrarSalidaBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegsitrarSalidaBody) UnmarshalBinary(b []byte) error {
	var res RegsitrarSalidaBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
