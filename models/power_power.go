package models

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*PowerPower power power

swagger:model Power_Power
*/
type PowerPower struct {
	Power100Power
}

// Validate validates this power power
func (m *PowerPower) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Power100Power.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}