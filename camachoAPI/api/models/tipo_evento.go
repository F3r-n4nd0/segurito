// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TipoEvento tipo evento
// swagger:model TipoEvento
type TipoEvento string

const (

	// TipoEventoEntrada captures enum value "entrada"
	TipoEventoEntrada TipoEvento = "entrada"

	// TipoEventoSalida captures enum value "salida"
	TipoEventoSalida TipoEvento = "salida"
)

// for schema
var tipoEventoEnum []interface{}

func init() {
	var res []TipoEvento
	if err := json.Unmarshal([]byte(`["entrada","salida"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tipoEventoEnum = append(tipoEventoEnum, v)
	}
}

func (m TipoEvento) validateTipoEventoEnum(path, location string, value TipoEvento) error {
	if err := validate.Enum(path, location, value, tipoEventoEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tipo evento
func (m TipoEvento) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTipoEventoEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}