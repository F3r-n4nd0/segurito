// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// RegsitrarEntradaCreatedCode is the HTTP code returned for type RegsitrarEntradaCreated
const RegsitrarEntradaCreatedCode int = 201

/*RegsitrarEntradaCreated confirmacion creacion de registro

swagger:response regsitrarEntradaCreated
*/
type RegsitrarEntradaCreated struct {
}

// NewRegsitrarEntradaCreated creates RegsitrarEntradaCreated with default headers values
func NewRegsitrarEntradaCreated() *RegsitrarEntradaCreated {

	return &RegsitrarEntradaCreated{}
}

// WriteResponse to the client
func (o *RegsitrarEntradaCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// RegsitrarEntradaBadRequestCode is the HTTP code returned for type RegsitrarEntradaBadRequest
const RegsitrarEntradaBadRequestCode int = 400

/*RegsitrarEntradaBadRequest parametro invalido

swagger:response regsitrarEntradaBadRequest
*/
type RegsitrarEntradaBadRequest struct {
}

// NewRegsitrarEntradaBadRequest creates RegsitrarEntradaBadRequest with default headers values
func NewRegsitrarEntradaBadRequest() *RegsitrarEntradaBadRequest {

	return &RegsitrarEntradaBadRequest{}
}

// WriteResponse to the client
func (o *RegsitrarEntradaBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// RegsitrarEntradaUnauthorizedCode is the HTTP code returned for type RegsitrarEntradaUnauthorized
const RegsitrarEntradaUnauthorizedCode int = 401

/*RegsitrarEntradaUnauthorized codigo de usuario incorrecto

swagger:response regsitrarEntradaUnauthorized
*/
type RegsitrarEntradaUnauthorized struct {
}

// NewRegsitrarEntradaUnauthorized creates RegsitrarEntradaUnauthorized with default headers values
func NewRegsitrarEntradaUnauthorized() *RegsitrarEntradaUnauthorized {

	return &RegsitrarEntradaUnauthorized{}
}

// WriteResponse to the client
func (o *RegsitrarEntradaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// RegsitrarEntradaConflictCode is the HTTP code returned for type RegsitrarEntradaConflict
const RegsitrarEntradaConflictCode int = 409

/*RegsitrarEntradaConflict usuario en estado incorrecto para registro de controlasistencia

swagger:response regsitrarEntradaConflict
*/
type RegsitrarEntradaConflict struct {
}

// NewRegsitrarEntradaConflict creates RegsitrarEntradaConflict with default headers values
func NewRegsitrarEntradaConflict() *RegsitrarEntradaConflict {

	return &RegsitrarEntradaConflict{}
}

// WriteResponse to the client
func (o *RegsitrarEntradaConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// RegsitrarEntradaInternalServerErrorCode is the HTTP code returned for type RegsitrarEntradaInternalServerError
const RegsitrarEntradaInternalServerErrorCode int = 500

/*RegsitrarEntradaInternalServerError error interno

swagger:response regsitrarEntradaInternalServerError
*/
type RegsitrarEntradaInternalServerError struct {
}

// NewRegsitrarEntradaInternalServerError creates RegsitrarEntradaInternalServerError with default headers values
func NewRegsitrarEntradaInternalServerError() *RegsitrarEntradaInternalServerError {

	return &RegsitrarEntradaInternalServerError{}
}

// WriteResponse to the client
func (o *RegsitrarEntradaInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}