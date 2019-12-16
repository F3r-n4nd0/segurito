// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// RegistrarEntradaCreatedCode is the HTTP code returned for type RegistrarEntradaCreated
const RegistrarEntradaCreatedCode int = 201

/*RegistrarEntradaCreated confirmacion creacion de registro

swagger:response registrarEntradaCreated
*/
type RegistrarEntradaCreated struct {
}

// NewRegistrarEntradaCreated creates RegistrarEntradaCreated with default headers values
func NewRegistrarEntradaCreated() *RegistrarEntradaCreated {

	return &RegistrarEntradaCreated{}
}

// WriteResponse to the client
func (o *RegistrarEntradaCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// RegistrarEntradaBadRequestCode is the HTTP code returned for type RegistrarEntradaBadRequest
const RegistrarEntradaBadRequestCode int = 400

/*RegistrarEntradaBadRequest parametro invalido

swagger:response registrarEntradaBadRequest
*/
type RegistrarEntradaBadRequest struct {
}

// NewRegistrarEntradaBadRequest creates RegistrarEntradaBadRequest with default headers values
func NewRegistrarEntradaBadRequest() *RegistrarEntradaBadRequest {

	return &RegistrarEntradaBadRequest{}
}

// WriteResponse to the client
func (o *RegistrarEntradaBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// RegistrarEntradaUnauthorizedCode is the HTTP code returned for type RegistrarEntradaUnauthorized
const RegistrarEntradaUnauthorizedCode int = 401

/*RegistrarEntradaUnauthorized codigo de usuario incorrecto

swagger:response registrarEntradaUnauthorized
*/
type RegistrarEntradaUnauthorized struct {
}

// NewRegistrarEntradaUnauthorized creates RegistrarEntradaUnauthorized with default headers values
func NewRegistrarEntradaUnauthorized() *RegistrarEntradaUnauthorized {

	return &RegistrarEntradaUnauthorized{}
}

// WriteResponse to the client
func (o *RegistrarEntradaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// RegistrarEntradaConflictCode is the HTTP code returned for type RegistrarEntradaConflict
const RegistrarEntradaConflictCode int = 409

/*RegistrarEntradaConflict usuario en estado incorrecto para registro de controlasistencia

swagger:response registrarEntradaConflict
*/
type RegistrarEntradaConflict struct {
}

// NewRegistrarEntradaConflict creates RegistrarEntradaConflict with default headers values
func NewRegistrarEntradaConflict() *RegistrarEntradaConflict {

	return &RegistrarEntradaConflict{}
}

// WriteResponse to the client
func (o *RegistrarEntradaConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// RegistrarEntradaInternalServerErrorCode is the HTTP code returned for type RegistrarEntradaInternalServerError
const RegistrarEntradaInternalServerErrorCode int = 500

/*RegistrarEntradaInternalServerError error interno

swagger:response registrarEntradaInternalServerError
*/
type RegistrarEntradaInternalServerError struct {
}

// NewRegistrarEntradaInternalServerError creates RegistrarEntradaInternalServerError with default headers values
func NewRegistrarEntradaInternalServerError() *RegistrarEntradaInternalServerError {

	return &RegistrarEntradaInternalServerError{}
}

// WriteResponse to the client
func (o *RegistrarEntradaInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}