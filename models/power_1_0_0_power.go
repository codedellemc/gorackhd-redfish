package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
)

/*Power100Power This is the schema definition for the Power Metrics.  It represents the properties for Power Consumption and Power Limiting.

swagger:model Power.1.0.0_Power
*/
type Power100Power struct {

	/* at odata context

	Read Only: true
	*/
	AtOdataContext strfmt.URI `json:"@odata.context,omitempty"`

	/* at odata id

	Read Only: true
	*/
	AtOdataID strfmt.URI `json:"@odata.id,omitempty"`

	/* at odata type

	Read Only: true
	*/
	AtOdataType string `json:"@odata.type,omitempty"`

	/* Provides a description of this resource and is used for commonality  in the schema definitions.

	Read Only: true
	*/
	Description string `json:"Description,omitempty"`

	/* Uniquely identifies the resource within the collection of like resources.

	Read Only: true
	*/
	ID string `json:"Id,omitempty"`

	/* The name of the resource or array element.

	Read Only: true
	*/
	Name string `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* This is the definition for power control function (power reading/limiting).

	Read Only: true
	*/
	PowerControl []*Power100PowerControl `json:"PowerControl,omitempty"`

	/* power control at odata count

	Read Only: true
	*/
	PowerControlAtOdataCount float64 `json:"PowerControl@odata.count,omitempty"`

	/* power control at odata navigation link
	 */
	PowerControlAtOdataNavigationLink *Odata400IDRef `json:"PowerControl@odata.navigationLink,omitempty"`

	/* Details of the power supplies associated with this system or device

	Read Only: true
	*/
	PowerSupplies []*Power100PowerSupply `json:"PowerSupplies,omitempty"`

	/* power supplies at odata count

	Read Only: true
	*/
	PowerSuppliesAtOdataCount float64 `json:"PowerSupplies@odata.count,omitempty"`

	/* power supplies at odata navigation link
	 */
	PowerSuppliesAtOdataNavigationLink *Odata400IDRef `json:"PowerSupplies@odata.navigationLink,omitempty"`

	/* Redundancy information for the power subsystem of this system or device

	Read Only: true
	*/
	Redundancy []*Odata400IDRef `json:"Redundancy,omitempty"`

	/* redundancy at odata count

	Read Only: true
	*/
	RedundancyAtOdataCount float64 `json:"Redundancy@odata.count,omitempty"`

	/* redundancy at odata navigation link
	 */
	RedundancyAtOdataNavigationLink *Odata400IDRef `json:"Redundancy@odata.navigationLink,omitempty"`

	/* This is the definition for voltage sensors.

	Read Only: true
	*/
	Voltages []*Power100Voltage `json:"Voltages,omitempty"`

	/* voltages at odata count

	Read Only: true
	*/
	VoltagesAtOdataCount float64 `json:"Voltages@odata.count,omitempty"`

	/* voltages at odata navigation link
	 */
	VoltagesAtOdataNavigationLink *Odata400IDRef `json:"Voltages@odata.navigationLink,omitempty"`
}

// Validate validates this power 1 0 0 power
func (m *Power100Power) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePowerControl(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePowerSupplies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRedundancy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVoltages(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Power100Power) validatePowerControl(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerControl) { // not required
		return nil
	}

	return nil
}

func (m *Power100Power) validatePowerSupplies(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerSupplies) { // not required
		return nil
	}

	return nil
}

func (m *Power100Power) validateRedundancy(formats strfmt.Registry) error {

	if swag.IsZero(m.Redundancy) { // not required
		return nil
	}

	return nil
}

func (m *Power100Power) validateVoltages(formats strfmt.Registry) error {

	if swag.IsZero(m.Voltages) { // not required
		return nil
	}

	return nil
}
