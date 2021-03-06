// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkPredictionOptions dlframework prediction options
// swagger:model dlframeworkPredictionOptions

type DlframeworkPredictionOptions struct {

	// agent
	Agent string `json:"agent,omitempty"`

	// batch size
	BatchSize int64 `json:"batch_size,omitempty"`

	// execution options
	ExecutionOptions *DlframeworkExecutionOptions `json:"execution_options,omitempty"`

	// feature limit
	FeatureLimit int32 `json:"feature_limit,omitempty"`

	// request id
	RequestID string `json:"request_id,omitempty"`
}

/* polymorph dlframeworkPredictionOptions agent false */

/* polymorph dlframeworkPredictionOptions batch_size false */

/* polymorph dlframeworkPredictionOptions execution_options false */

/* polymorph dlframeworkPredictionOptions feature_limit false */

/* polymorph dlframeworkPredictionOptions request_id false */

// Validate validates this dlframework prediction options
func (m *DlframeworkPredictionOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutionOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DlframeworkPredictionOptions) validateExecutionOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutionOptions) { // not required
		return nil
	}

	if m.ExecutionOptions != nil {

		if err := m.ExecutionOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execution_options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkPredictionOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkPredictionOptions) UnmarshalBinary(b []byte) error {
	var res DlframeworkPredictionOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
