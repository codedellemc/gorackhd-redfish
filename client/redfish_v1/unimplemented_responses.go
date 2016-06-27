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

// UnimplementedReader is a Reader for the Unimplemented structure.
type UnimplementedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UnimplementedReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUnimplementedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewUnimplementedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUnimplementedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewUnimplementedNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUnimplementedOK creates a UnimplementedOK with default headers values
func NewUnimplementedOK() *UnimplementedOK {
	return &UnimplementedOK{}
}

/*UnimplementedOK handles this case with default header values.

Success
*/
type UnimplementedOK struct {
	Payload *models.AccountService100AccountService
}

func (o *UnimplementedOK) Error() string {
	return fmt.Sprintf("[GET /AccountService][%d] unimplementedOK  %+v", 200, o.Payload)
}

func (o *UnimplementedOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountService100AccountService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnimplementedUnauthorized creates a UnimplementedUnauthorized with default headers values
func NewUnimplementedUnauthorized() *UnimplementedUnauthorized {
	return &UnimplementedUnauthorized{}
}

/*UnimplementedUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type UnimplementedUnauthorized struct {
}

func (o *UnimplementedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /AccountService][%d] unimplementedUnauthorized ", 401)
}

func (o *UnimplementedUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnimplementedForbidden creates a UnimplementedForbidden with default headers values
func NewUnimplementedForbidden() *UnimplementedForbidden {
	return &UnimplementedForbidden{}
}

/*UnimplementedForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type UnimplementedForbidden struct {
}

func (o *UnimplementedForbidden) Error() string {
	return fmt.Sprintf("[GET /AccountService][%d] unimplementedForbidden ", 403)
}

func (o *UnimplementedForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnimplementedNotImplemented creates a UnimplementedNotImplemented with default headers values
func NewUnimplementedNotImplemented() *UnimplementedNotImplemented {
	return &UnimplementedNotImplemented{}
}

/*UnimplementedNotImplemented handles this case with default header values.

The server does not (currently) support the functionality required to fulfill the request.  This is the appropriate response when the server does not recognize the request method  and is not capable of supporting the method for any resource.

*/
type UnimplementedNotImplemented struct {
}

func (o *UnimplementedNotImplemented) Error() string {
	return fmt.Sprintf("[GET /AccountService][%d] unimplementedNotImplemented ", 501)
}

func (o *UnimplementedNotImplemented) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}