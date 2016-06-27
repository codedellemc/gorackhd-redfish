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

// GetSystemReader is a Reader for the GetSystem structure.
type GetSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSystemReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSystemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSystemOK creates a GetSystemOK with default headers values
func NewGetSystemOK() *GetSystemOK {
	return &GetSystemOK{}
}

/*GetSystemOK handles this case with default header values.

Success
*/
type GetSystemOK struct {
	Payload *models.ComputerSystem100ComputerSystem
}

func (o *GetSystemOK) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}][%d] getSystemOK  %+v", 200, o.Payload)
}

func (o *GetSystemOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputerSystem100ComputerSystem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemBadRequest creates a GetSystemBadRequest with default headers values
func NewGetSystemBadRequest() *GetSystemBadRequest {
	return &GetSystemBadRequest{}
}

/*GetSystemBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information  (such as validation error on an input field, a missing required value, and so on).  An extended error shall be returned in the response body, as defined in section Extended  Error Handling.

*/
type GetSystemBadRequest struct {
}

func (o *GetSystemBadRequest) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}][%d] getSystemBadRequest ", 400)
}

func (o *GetSystemBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemUnauthorized creates a GetSystemUnauthorized with default headers values
func NewGetSystemUnauthorized() *GetSystemUnauthorized {
	return &GetSystemUnauthorized{}
}

/*GetSystemUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetSystemUnauthorized struct {
}

func (o *GetSystemUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}][%d] getSystemUnauthorized ", 401)
}

func (o *GetSystemUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemForbidden creates a GetSystemForbidden with default headers values
func NewGetSystemForbidden() *GetSystemForbidden {
	return &GetSystemForbidden{}
}

/*GetSystemForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetSystemForbidden struct {
}

func (o *GetSystemForbidden) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}][%d] getSystemForbidden ", 403)
}

func (o *GetSystemForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemNotFound creates a GetSystemNotFound with default headers values
func NewGetSystemNotFound() *GetSystemNotFound {
	return &GetSystemNotFound{}
}

/*GetSystemNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetSystemNotFound struct {
}

func (o *GetSystemNotFound) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}][%d] getSystemNotFound ", 404)
}

func (o *GetSystemNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemInternalServerError creates a GetSystemInternalServerError with default headers values
func NewGetSystemInternalServerError() *GetSystemInternalServerError {
	return &GetSystemInternalServerError{}
}

/*GetSystemInternalServerError handles this case with default header values.

Error
*/
type GetSystemInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetSystemInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}][%d] getSystemInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystemInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}