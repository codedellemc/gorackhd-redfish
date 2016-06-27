package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Task100TaskState task 1 0 0 task state

swagger:model Task.1.0.0_TaskState
*/
type Task100TaskState string

// for schema
var task100TaskStateEnum []interface{}

func (m Task100TaskState) validateTask100TaskStateEnum(path, location string, value Task100TaskState) error {
	if task100TaskStateEnum == nil {
		var res []Task100TaskState
		if err := json.Unmarshal([]byte(`["New","Starting","Running","Suspended","Interrupted","Pending","Stopping","Completed","Killed","Exception","Service"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			task100TaskStateEnum = append(task100TaskStateEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, task100TaskStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this task 1 0 0 task state
func (m Task100TaskState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTask100TaskStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}