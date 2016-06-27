package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
)

/*TaskService100TaskService This is the schema definition for the Task Service.  It represents the properties for the service itself and has links to the actual list of tasks.

swagger:model TaskService.1.0.0_TaskService
*/
type TaskService100TaskService struct {

	/* at odata context
	 */
	AtOdataContext Odata400Context `json:"@odata.context,omitempty"`

	/* at odata id
	 */
	AtOdataID Odata400ID `json:"@odata.id,omitempty"`

	/* at odata type
	 */
	AtOdataType Odata400Type `json:"@odata.type,omitempty"`

	/* Overwrite policy of completed tasks

	Read Only: true
	*/
	CompletedTaskOverWritePolicy TaskService100OverWritePolicy `json:"CompletedTaskOverWritePolicy,omitempty"`

	/* The current DateTime (with offset) setting that the task service is using.

	Read Only: true
	*/
	DateTime strfmt.DateTime `json:"DateTime,omitempty"`

	/* description
	 */
	Description ResourceDescription `json:"Description,omitempty"`

	/* Id
	 */
	ID ResourceID `json:"Id,omitempty"`

	/* Send an Event upon Task State Change.

	Read Only: true
	*/
	LifeCycleEventOnTaskStateChange *bool `json:"LifeCycleEventOnTaskStateChange,omitempty"`

	/* name
	 */
	Name ResourceName `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* This indicates whether this service is enabled.
	 */
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	/* status
	 */
	Status *ResourceStatus `json:"Status,omitempty"`

	/* References to the Tasks collection.

	Read Only: true
	*/
	Tasks *TaskCollectionTaskCollection `json:"Tasks,omitempty"`
}

// Validate validates this task service 1 0 0 task service
func (m *TaskService100TaskService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedTaskOverWritePolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskService100TaskService) validateCompletedTaskOverWritePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedTaskOverWritePolicy) { // not required
		return nil
	}

	if err := m.CompletedTaskOverWritePolicy.Validate(formats); err != nil {
		return err
	}

	return nil
}