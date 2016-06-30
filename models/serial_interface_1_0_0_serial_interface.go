package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*SerialInterface100SerialInterface This schema defines an asynchronous serial interface resource.

swagger:model SerialInterface.1.0.0_SerialInterface
*/
type SerialInterface100SerialInterface struct {

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

	/* The receive and transmit rate of data flow, typically in bits-per-second (bps), over the serial connection.
	 */
	BitRate string `json:"BitRate,omitempty"`

	/* The type of connector used for this interface.

	Read Only: true
	*/
	ConnectorType string `json:"ConnectorType,omitempty"`

	/* The number of data bits that will follow the start bit over the serial connection.
	 */
	DataBits string `json:"DataBits,omitempty"`

	/* Provides a description of this resource and is used for commonality  in the schema definitions.

	Read Only: true
	*/
	Description string `json:"Description,omitempty"`

	/* The type of flow control, if any, that will be imposed on the serial connection.
	 */
	FlowControl string `json:"FlowControl,omitempty"`

	/* Uniquely identifies the resource within the collection of like resources.

	Read Only: true
	*/
	ID string `json:"Id,omitempty"`

	/* This indicates whether this interface is enabled.
	 */
	InterfaceEnabled bool `json:"InterfaceEnabled,omitempty"`

	/* The name of the resource or array element.

	Read Only: true
	*/
	Name string `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* The type of parity used by the sender and receiver in order to detect errors over the serial connection.
	 */
	Parity string `json:"Parity,omitempty"`

	/* The physical pin configuration needed for a serial connector.

	Read Only: true
	*/
	PinOut string `json:"PinOut,omitempty"`

	/* The type of signal used for the communication connection - RS232 or RS485.

	Read Only: true
	*/
	SignalType string `json:"SignalType,omitempty"`

	/* The period of time before the next start bit is transmitted.
	 */
	StopBits string `json:"StopBits,omitempty"`
}

// Validate validates this serial interface 1 0 0 serial interface
func (m *SerialInterface100SerialInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBitRate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectorType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDataBits(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFlowControl(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePinOut(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSignalType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStopBits(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serialInterface100SerialInterfaceTypeBitRatePropEnum []interface{}

// prop value enum
func (m *SerialInterface100SerialInterface) validateBitRateEnum(path, location string, value string) error {
	if serialInterface100SerialInterfaceTypeBitRatePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["1200","2400","4800","9600","19200","38400","57600","115200","230400"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serialInterface100SerialInterfaceTypeBitRatePropEnum = append(serialInterface100SerialInterfaceTypeBitRatePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serialInterface100SerialInterfaceTypeBitRatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SerialInterface100SerialInterface) validateBitRate(formats strfmt.Registry) error {

	if swag.IsZero(m.BitRate) { // not required
		return nil
	}

	// value enum
	if err := m.validateBitRateEnum("BitRate", "body", m.BitRate); err != nil {
		return err
	}

	return nil
}

var serialInterface100SerialInterfaceTypeConnectorTypePropEnum []interface{}

// prop value enum
func (m *SerialInterface100SerialInterface) validateConnectorTypeEnum(path, location string, value string) error {
	if serialInterface100SerialInterfaceTypeConnectorTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["RJ45","RJ11","DB9 Female","DB9 Male","DB25 Female","DB25 Male","USB","mUSB","uUSB"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serialInterface100SerialInterfaceTypeConnectorTypePropEnum = append(serialInterface100SerialInterfaceTypeConnectorTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serialInterface100SerialInterfaceTypeConnectorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SerialInterface100SerialInterface) validateConnectorType(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectorType) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectorTypeEnum("ConnectorType", "body", m.ConnectorType); err != nil {
		return err
	}

	return nil
}

var serialInterface100SerialInterfaceTypeDataBitsPropEnum []interface{}

// prop value enum
func (m *SerialInterface100SerialInterface) validateDataBitsEnum(path, location string, value string) error {
	if serialInterface100SerialInterfaceTypeDataBitsPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["5","6","7","8"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serialInterface100SerialInterfaceTypeDataBitsPropEnum = append(serialInterface100SerialInterfaceTypeDataBitsPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serialInterface100SerialInterfaceTypeDataBitsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SerialInterface100SerialInterface) validateDataBits(formats strfmt.Registry) error {

	if swag.IsZero(m.DataBits) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataBitsEnum("DataBits", "body", m.DataBits); err != nil {
		return err
	}

	return nil
}

var serialInterface100SerialInterfaceTypeFlowControlPropEnum []interface{}

// prop value enum
func (m *SerialInterface100SerialInterface) validateFlowControlEnum(path, location string, value string) error {
	if serialInterface100SerialInterfaceTypeFlowControlPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["None","Software","Hardware"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serialInterface100SerialInterfaceTypeFlowControlPropEnum = append(serialInterface100SerialInterfaceTypeFlowControlPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serialInterface100SerialInterfaceTypeFlowControlPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SerialInterface100SerialInterface) validateFlowControl(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowControl) { // not required
		return nil
	}

	// value enum
	if err := m.validateFlowControlEnum("FlowControl", "body", m.FlowControl); err != nil {
		return err
	}

	return nil
}

var serialInterface100SerialInterfaceTypeParityPropEnum []interface{}

// prop value enum
func (m *SerialInterface100SerialInterface) validateParityEnum(path, location string, value string) error {
	if serialInterface100SerialInterfaceTypeParityPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["None","Even","Odd","Mark","Space"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serialInterface100SerialInterfaceTypeParityPropEnum = append(serialInterface100SerialInterfaceTypeParityPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serialInterface100SerialInterfaceTypeParityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SerialInterface100SerialInterface) validateParity(formats strfmt.Registry) error {

	if swag.IsZero(m.Parity) { // not required
		return nil
	}

	// value enum
	if err := m.validateParityEnum("Parity", "body", m.Parity); err != nil {
		return err
	}

	return nil
}

var serialInterface100SerialInterfaceTypePinOutPropEnum []interface{}

// prop value enum
func (m *SerialInterface100SerialInterface) validatePinOutEnum(path, location string, value string) error {
	if serialInterface100SerialInterfaceTypePinOutPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Cisco","Cyclades","Digi"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serialInterface100SerialInterfaceTypePinOutPropEnum = append(serialInterface100SerialInterfaceTypePinOutPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serialInterface100SerialInterfaceTypePinOutPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SerialInterface100SerialInterface) validatePinOut(formats strfmt.Registry) error {

	if swag.IsZero(m.PinOut) { // not required
		return nil
	}

	// value enum
	if err := m.validatePinOutEnum("PinOut", "body", m.PinOut); err != nil {
		return err
	}

	return nil
}

var serialInterface100SerialInterfaceTypeSignalTypePropEnum []interface{}

// prop value enum
func (m *SerialInterface100SerialInterface) validateSignalTypeEnum(path, location string, value string) error {
	if serialInterface100SerialInterfaceTypeSignalTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Rs232","Rs485"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serialInterface100SerialInterfaceTypeSignalTypePropEnum = append(serialInterface100SerialInterfaceTypeSignalTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serialInterface100SerialInterfaceTypeSignalTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SerialInterface100SerialInterface) validateSignalType(formats strfmt.Registry) error {

	if swag.IsZero(m.SignalType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSignalTypeEnum("SignalType", "body", m.SignalType); err != nil {
		return err
	}

	return nil
}

var serialInterface100SerialInterfaceTypeStopBitsPropEnum []interface{}

// prop value enum
func (m *SerialInterface100SerialInterface) validateStopBitsEnum(path, location string, value string) error {
	if serialInterface100SerialInterfaceTypeStopBitsPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["1","2"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serialInterface100SerialInterfaceTypeStopBitsPropEnum = append(serialInterface100SerialInterfaceTypeStopBitsPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serialInterface100SerialInterfaceTypeStopBitsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SerialInterface100SerialInterface) validateStopBits(formats strfmt.Registry) error {

	if swag.IsZero(m.StopBits) { // not required
		return nil
	}

	// value enum
	if err := m.validateStopBitsEnum("StopBits", "body", m.StopBits); err != nil {
		return err
	}

	return nil
}