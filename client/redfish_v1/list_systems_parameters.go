package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewListSystemsParams creates a new ListSystemsParams object
// with the default values initialized.
func NewListSystemsParams() *ListSystemsParams {

	return &ListSystemsParams{}
}

/*ListSystemsParams contains all the parameters to send to the API endpoint
for the list systems operation typically these are written to a http.Request
*/
type ListSystemsParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *ListSystemsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}