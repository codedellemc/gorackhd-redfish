package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd-redfish/models"
)

// DoBootImageReader is a Reader for the DoBootImage structure.
type DoBootImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DoBootImageReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDoBootImageAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDoBootImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDoBootImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDoBootImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDoBootImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDoBootImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDoBootImageAccepted creates a DoBootImageAccepted with default headers values
func NewDoBootImageAccepted() *DoBootImageAccepted {
	return &DoBootImageAccepted{}
}

/*DoBootImageAccepted handles this case with default header values.

The request has been accepted for processing, but the processing has not been completed.  The Location header shall be set to the URI of a Task resource that can later be queried  to determine the status of the operation. A representation of the Task resource may be  included in the response body.

*/
type DoBootImageAccepted struct {
}

func (o *DoBootImageAccepted) Error() string {
	return fmt.Sprintf("[POST /Systems/{identifier}/Actions/RackHD.BootImage][%d] doBootImageAccepted ", 202)
}

func (o *DoBootImageAccepted) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDoBootImageBadRequest creates a DoBootImageBadRequest with default headers values
func NewDoBootImageBadRequest() *DoBootImageBadRequest {
	return &DoBootImageBadRequest{}
}

/*DoBootImageBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information  (such as validation error on an input field, a missing required value, and so on).  An extended error shall be returned in the response body, as defined in section Extended  Error Handling.

*/
type DoBootImageBadRequest struct {
}

func (o *DoBootImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /Systems/{identifier}/Actions/RackHD.BootImage][%d] doBootImageBadRequest ", 400)
}

func (o *DoBootImageBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDoBootImageUnauthorized creates a DoBootImageUnauthorized with default headers values
func NewDoBootImageUnauthorized() *DoBootImageUnauthorized {
	return &DoBootImageUnauthorized{}
}

/*DoBootImageUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type DoBootImageUnauthorized struct {
}

func (o *DoBootImageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /Systems/{identifier}/Actions/RackHD.BootImage][%d] doBootImageUnauthorized ", 401)
}

func (o *DoBootImageUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDoBootImageForbidden creates a DoBootImageForbidden with default headers values
func NewDoBootImageForbidden() *DoBootImageForbidden {
	return &DoBootImageForbidden{}
}

/*DoBootImageForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type DoBootImageForbidden struct {
}

func (o *DoBootImageForbidden) Error() string {
	return fmt.Sprintf("[POST /Systems/{identifier}/Actions/RackHD.BootImage][%d] doBootImageForbidden ", 403)
}

func (o *DoBootImageForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDoBootImageNotFound creates a DoBootImageNotFound with default headers values
func NewDoBootImageNotFound() *DoBootImageNotFound {
	return &DoBootImageNotFound{}
}

/*DoBootImageNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type DoBootImageNotFound struct {
}

func (o *DoBootImageNotFound) Error() string {
	return fmt.Sprintf("[POST /Systems/{identifier}/Actions/RackHD.BootImage][%d] doBootImageNotFound ", 404)
}

func (o *DoBootImageNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDoBootImageInternalServerError creates a DoBootImageInternalServerError with default headers values
func NewDoBootImageInternalServerError() *DoBootImageInternalServerError {
	return &DoBootImageInternalServerError{}
}

/*DoBootImageInternalServerError handles this case with default header values.

Error
*/
type DoBootImageInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DoBootImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /Systems/{identifier}/Actions/RackHD.BootImage][%d] doBootImageInternalServerError  %+v", 500, o.Payload)
}

func (o *DoBootImageInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}