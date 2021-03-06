package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteEventReader is a Reader for the DeleteEvent structure.
type DeleteEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteEventForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewDeleteEventNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEventOK creates a DeleteEventOK with default headers values
func NewDeleteEventOK() *DeleteEventOK {
	return &DeleteEventOK{}
}

/*DeleteEventOK handles this case with default header values.

Success
*/
type DeleteEventOK struct {
}

func (o *DeleteEventOK) Error() string {
	return fmt.Sprintf("[DELETE /EventService/Subscriptions/{index}][%d] deleteEventOK ", 200)
}

func (o *DeleteEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEventUnauthorized creates a DeleteEventUnauthorized with default headers values
func NewDeleteEventUnauthorized() *DeleteEventUnauthorized {
	return &DeleteEventUnauthorized{}
}

/*DeleteEventUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type DeleteEventUnauthorized struct {
}

func (o *DeleteEventUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /EventService/Subscriptions/{index}][%d] deleteEventUnauthorized ", 401)
}

func (o *DeleteEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEventForbidden creates a DeleteEventForbidden with default headers values
func NewDeleteEventForbidden() *DeleteEventForbidden {
	return &DeleteEventForbidden{}
}

/*DeleteEventForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type DeleteEventForbidden struct {
}

func (o *DeleteEventForbidden) Error() string {
	return fmt.Sprintf("[DELETE /EventService/Subscriptions/{index}][%d] deleteEventForbidden ", 403)
}

func (o *DeleteEventForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEventNotImplemented creates a DeleteEventNotImplemented with default headers values
func NewDeleteEventNotImplemented() *DeleteEventNotImplemented {
	return &DeleteEventNotImplemented{}
}

/*DeleteEventNotImplemented handles this case with default header values.

The server does not (currently) support the functionality required to fulfill the request.  This is the appropriate response when the server does not recognize the request method  and is not capable of supporting the method for any resource.

*/
type DeleteEventNotImplemented struct {
}

func (o *DeleteEventNotImplemented) Error() string {
	return fmt.Sprintf("[DELETE /EventService/Subscriptions/{index}][%d] deleteEventNotImplemented ", 501)
}

func (o *DeleteEventNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
