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

// GetEventServiceReader is a Reader for the GetEventService structure.
type GetEventServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetEventServiceReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetEventServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetEventServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewGetEventServiceNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventServiceOK creates a GetEventServiceOK with default headers values
func NewGetEventServiceOK() *GetEventServiceOK {
	return &GetEventServiceOK{}
}

/*GetEventServiceOK handles this case with default header values.

Success
*/
type GetEventServiceOK struct {
	Payload *models.AccountService100AccountService
}

func (o *GetEventServiceOK) Error() string {
	return fmt.Sprintf("[GET /EventService][%d] getEventServiceOK  %+v", 200, o.Payload)
}

func (o *GetEventServiceOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountService100AccountService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventServiceUnauthorized creates a GetEventServiceUnauthorized with default headers values
func NewGetEventServiceUnauthorized() *GetEventServiceUnauthorized {
	return &GetEventServiceUnauthorized{}
}

/*GetEventServiceUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetEventServiceUnauthorized struct {
}

func (o *GetEventServiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /EventService][%d] getEventServiceUnauthorized ", 401)
}

func (o *GetEventServiceUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEventServiceForbidden creates a GetEventServiceForbidden with default headers values
func NewGetEventServiceForbidden() *GetEventServiceForbidden {
	return &GetEventServiceForbidden{}
}

/*GetEventServiceForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetEventServiceForbidden struct {
}

func (o *GetEventServiceForbidden) Error() string {
	return fmt.Sprintf("[GET /EventService][%d] getEventServiceForbidden ", 403)
}

func (o *GetEventServiceForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEventServiceNotImplemented creates a GetEventServiceNotImplemented with default headers values
func NewGetEventServiceNotImplemented() *GetEventServiceNotImplemented {
	return &GetEventServiceNotImplemented{}
}

/*GetEventServiceNotImplemented handles this case with default header values.

The server does not (currently) support the functionality required to fulfill the request.  This is the appropriate response when the server does not recognize the request method  and is not capable of supporting the method for any resource.

*/
type GetEventServiceNotImplemented struct {
}

func (o *GetEventServiceNotImplemented) Error() string {
	return fmt.Sprintf("[GET /EventService][%d] getEventServiceNotImplemented ", 501)
}

func (o *GetEventServiceNotImplemented) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}