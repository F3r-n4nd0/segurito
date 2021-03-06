// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// RegsitrarSalidaCreatedCode is the HTTP code returned for type RegsitrarSalidaCreated
const RegsitrarSalidaCreatedCode int = 201

/*RegsitrarSalidaCreated confirmacion creacion de registro

swagger:response regsitrarSalidaCreated
*/
type RegsitrarSalidaCreated struct {
}

// NewRegsitrarSalidaCreated creates RegsitrarSalidaCreated with default headers values
func NewRegsitrarSalidaCreated() *RegsitrarSalidaCreated {

	return &RegsitrarSalidaCreated{}
}

// WriteResponse to the client
func (o *RegsitrarSalidaCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// RegsitrarSalidaBadRequestCode is the HTTP code returned for type RegsitrarSalidaBadRequest
const RegsitrarSalidaBadRequestCode int = 400

/*RegsitrarSalidaBadRequest parametro invalido

swagger:response regsitrarSalidaBadRequest
*/
type RegsitrarSalidaBadRequest struct {
}

// NewRegsitrarSalidaBadRequest creates RegsitrarSalidaBadRequest with default headers values
func NewRegsitrarSalidaBadRequest() *RegsitrarSalidaBadRequest {

	return &RegsitrarSalidaBadRequest{}
}

// WriteResponse to the client
func (o *RegsitrarSalidaBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// RegsitrarSalidaUnauthorizedCode is the HTTP code returned for type RegsitrarSalidaUnauthorized
const RegsitrarSalidaUnauthorizedCode int = 401

/*RegsitrarSalidaUnauthorized codigo de usuario incorrecto

swagger:response regsitrarSalidaUnauthorized
*/
type RegsitrarSalidaUnauthorized struct {
}

// NewRegsitrarSalidaUnauthorized creates RegsitrarSalidaUnauthorized with default headers values
func NewRegsitrarSalidaUnauthorized() *RegsitrarSalidaUnauthorized {

	return &RegsitrarSalidaUnauthorized{}
}

// WriteResponse to the client
func (o *RegsitrarSalidaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// RegsitrarSalidaConflictCode is the HTTP code returned for type RegsitrarSalidaConflict
const RegsitrarSalidaConflictCode int = 409

/*RegsitrarSalidaConflict usuario en estado incorrecto para registro de salida

swagger:response regsitrarSalidaConflict
*/
type RegsitrarSalidaConflict struct {
}

// NewRegsitrarSalidaConflict creates RegsitrarSalidaConflict with default headers values
func NewRegsitrarSalidaConflict() *RegsitrarSalidaConflict {

	return &RegsitrarSalidaConflict{}
}

// WriteResponse to the client
func (o *RegsitrarSalidaConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// RegsitrarSalidaInternalServerErrorCode is the HTTP code returned for type RegsitrarSalidaInternalServerError
const RegsitrarSalidaInternalServerErrorCode int = 500

/*RegsitrarSalidaInternalServerError error interno

swagger:response regsitrarSalidaInternalServerError
*/
type RegsitrarSalidaInternalServerError struct {
}

// NewRegsitrarSalidaInternalServerError creates RegsitrarSalidaInternalServerError with default headers values
func NewRegsitrarSalidaInternalServerError() *RegsitrarSalidaInternalServerError {

	return &RegsitrarSalidaInternalServerError{}
}

// WriteResponse to the client
func (o *RegsitrarSalidaInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
