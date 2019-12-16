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

// RegistrarEntradaHandlerFunc turns a function with the right signature into a registrar entrada handler
type RegistrarEntradaHandlerFunc func(RegistrarEntradaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegistrarEntradaHandlerFunc) Handle(params RegistrarEntradaParams) middleware.Responder {
	return fn(params)
}

// RegistrarEntradaHandler interface for that can handle valid registrar entrada params
type RegistrarEntradaHandler interface {
	Handle(RegistrarEntradaParams) middleware.Responder
}

// NewRegistrarEntrada creates a new http.Handler for the registrar entrada operation
func NewRegistrarEntrada(ctx *middleware.Context, handler RegistrarEntradaHandler) *RegistrarEntrada {
	return &RegistrarEntrada{Context: ctx, Handler: handler}
}

/*RegistrarEntrada swagger:route POST /entradas registrarEntrada

registra la controlasistencia del personal

registra la controlasistencia del personal por medio de su codigo (7 caracteres)

*/
type RegistrarEntrada struct {
	Context *middleware.Context
	Handler RegistrarEntradaHandler
}

func (o *RegistrarEntrada) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRegistrarEntradaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RegistrarEntradaBody registrar entrada body
// swagger:model RegistrarEntradaBody
type RegistrarEntradaBody struct {

	// codigo usuario
	// Required: true
	// Max Length: 7
	// Min Length: 7
	CodigoUsuario *string `json:"codigo_usuario"`
}

// Validate validates this registrar entrada body
func (o *RegistrarEntradaBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCodigoUsuario(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegistrarEntradaBody) validateCodigoUsuario(formats strfmt.Registry) error {

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
func (o *RegistrarEntradaBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegistrarEntradaBody) UnmarshalBinary(b []byte) error {
	var res RegistrarEntradaBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}