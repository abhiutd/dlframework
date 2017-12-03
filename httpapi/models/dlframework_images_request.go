// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkImagesRequest dlframework images request
// swagger:model dlframeworkImagesRequest

type DlframeworkImagesRequest struct {

	// A list of Base64 encoded images
	Images []*ImagesRequestImage `json:"images"`

	// options
	Options *DlframeworkPredictionOptions `json:"options,omitempty"`

	// predictor
	Predictor *DlframeworkPredictor `json:"predictor,omitempty"`
}

/* polymorph dlframeworkImagesRequest images false */

/* polymorph dlframeworkImagesRequest options false */

/* polymorph dlframeworkImagesRequest predictor false */

// Validate validates this dlframework images request
func (m *DlframeworkImagesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePredictor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DlframeworkImagesRequest) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {

		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {

			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DlframeworkImagesRequest) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {

		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *DlframeworkImagesRequest) validatePredictor(formats strfmt.Registry) error {

	if swag.IsZero(m.Predictor) { // not required
		return nil
	}

	if m.Predictor != nil {

		if err := m.Predictor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("predictor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkImagesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkImagesRequest) UnmarshalBinary(b []byte) error {
	var res DlframeworkImagesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
