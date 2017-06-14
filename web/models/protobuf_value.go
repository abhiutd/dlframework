package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProtobufValue `Value` represents a dynamically typed value which can be either
// null, a number, a string, a boolean, a recursive struct value, or a
// list of values. A producer of value is expected to set one of that
// variants, absence of any variant indicates an error.
//
// The JSON representation for `Value` is JSON value.
// swagger:model protobufValue
type ProtobufValue struct {

	// Represents a boolean value.
	BoolValue bool `json:"bool_value,omitempty"`

	// Represents a repeated `Value`.
	ListValue *ProtobufListValue `json:"list_value,omitempty"`

	// Represents a null value.
	NullValue ProtobufNullValue `json:"null_value,omitempty"`

	// Represents a double value.
	NumberValue float64 `json:"number_value,omitempty"`

	// Represents a string value.
	StringValue string `json:"string_value,omitempty"`

	// Represents a structured value.
	StructValue *ProtobufStruct `json:"struct_value,omitempty"`
}

// Validate validates this protobuf value
func (m *ProtobufValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateListValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNullValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStructValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtobufValue) validateListValue(formats strfmt.Registry) error {

	if swag.IsZero(m.ListValue) { // not required
		return nil
	}

	if m.ListValue != nil {

		if err := m.ListValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("list_value")
			}
			return err
		}
	}

	return nil
}

func (m *ProtobufValue) validateNullValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NullValue) { // not required
		return nil
	}

	if err := m.NullValue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("null_value")
		}
		return err
	}

	return nil
}

func (m *ProtobufValue) validateStructValue(formats strfmt.Registry) error {

	if swag.IsZero(m.StructValue) { // not required
		return nil
	}

	if m.StructValue != nil {

		if err := m.StructValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("struct_value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtobufValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtobufValue) UnmarshalBinary(b []byte) error {
	var res ProtobufValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
