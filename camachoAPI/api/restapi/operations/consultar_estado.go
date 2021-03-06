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

	models "camachoAPI/api/models"
)

// ConsultarEstadoHandlerFunc turns a function with the right signature into a consultar estado handler
type ConsultarEstadoHandlerFunc func(ConsultarEstadoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConsultarEstadoHandlerFunc) Handle(params ConsultarEstadoParams) middleware.Responder {
	return fn(params)
}

// ConsultarEstadoHandler interface for that can handle valid consultar estado params
type ConsultarEstadoHandler interface {
	Handle(ConsultarEstadoParams) middleware.Responder
}

// NewConsultarEstado creates a new http.Handler for the consultar estado operation
func NewConsultarEstado(ctx *middleware.Context, handler ConsultarEstadoHandler) *ConsultarEstado {
	return &ConsultarEstado{Context: ctx, Handler: handler}
}

/*ConsultarEstado swagger:route GET /eventos/{id_usuario}/estado consultarEstado

devuelve el estado actual del usuario

*/
type ConsultarEstado struct {
	Context *middleware.Context
	Handler ConsultarEstadoHandler
}

func (o *ConsultarEstado) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConsultarEstadoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ConsultarEstadoOKBody consultar estado o k body
// swagger:model ConsultarEstadoOKBody
type ConsultarEstadoOKBody struct {

	// estado
	Estado models.EstadoUsuario `json:"estado,omitempty"`
}

// Validate validates this consultar estado o k body
func (o *ConsultarEstadoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEstado(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConsultarEstadoOKBody) validateEstado(formats strfmt.Registry) error {

	if swag.IsZero(o.Estado) { // not required
		return nil
	}

	if err := o.Estado.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("consultarEstadoOK" + "." + "estado")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConsultarEstadoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConsultarEstadoOKBody) UnmarshalBinary(b []byte) error {
	var res ConsultarEstadoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
