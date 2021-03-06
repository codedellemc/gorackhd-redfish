package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListLocalEthernetInterfacesParams creates a new ListLocalEthernetInterfacesParams object
// with the default values initialized.
func NewListLocalEthernetInterfacesParams() *ListLocalEthernetInterfacesParams {

	return &ListLocalEthernetInterfacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListLocalEthernetInterfacesParamsWithTimeout creates a new ListLocalEthernetInterfacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListLocalEthernetInterfacesParamsWithTimeout(timeout time.Duration) *ListLocalEthernetInterfacesParams {

	return &ListLocalEthernetInterfacesParams{

		timeout: timeout,
	}
}

/*ListLocalEthernetInterfacesParams contains all the parameters to send to the API endpoint
for the list local ethernet interfaces operation typically these are written to a http.Request
*/
type ListLocalEthernetInterfacesParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *ListLocalEthernetInterfacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
