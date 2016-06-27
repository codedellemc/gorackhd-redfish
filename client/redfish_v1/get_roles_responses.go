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

// GetRolesReader is a Reader for the GetRoles structure.
type GetRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetRolesReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRolesOK creates a GetRolesOK with default headers values
func NewGetRolesOK() *GetRolesOK {
	return &GetRolesOK{}
}

/*GetRolesOK handles this case with default header values.

Success
*/
type GetRolesOK struct {
	Payload *models.RoleCollectionRoleCollection
}

func (o *GetRolesOK) Error() string {
	return fmt.Sprintf("[GET /Roles][%d] getRolesOK  %+v", 200, o.Payload)
}

func (o *GetRolesOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleCollectionRoleCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesBadRequest creates a GetRolesBadRequest with default headers values
func NewGetRolesBadRequest() *GetRolesBadRequest {
	return &GetRolesBadRequest{}
}

/*GetRolesBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information  (such as validation error on an input field, a missing required value, and so on).  An extended error shall be returned in the response body, as defined in section Extended  Error Handling.

*/
type GetRolesBadRequest struct {
}

func (o *GetRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /Roles][%d] getRolesBadRequest ", 400)
}

func (o *GetRolesBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRolesUnauthorized creates a GetRolesUnauthorized with default headers values
func NewGetRolesUnauthorized() *GetRolesUnauthorized {
	return &GetRolesUnauthorized{}
}

/*GetRolesUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetRolesUnauthorized struct {
}

func (o *GetRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Roles][%d] getRolesUnauthorized ", 401)
}

func (o *GetRolesUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRolesForbidden creates a GetRolesForbidden with default headers values
func NewGetRolesForbidden() *GetRolesForbidden {
	return &GetRolesForbidden{}
}

/*GetRolesForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetRolesForbidden struct {
}

func (o *GetRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /Roles][%d] getRolesForbidden ", 403)
}

func (o *GetRolesForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRolesNotFound creates a GetRolesNotFound with default headers values
func NewGetRolesNotFound() *GetRolesNotFound {
	return &GetRolesNotFound{}
}

/*GetRolesNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetRolesNotFound struct {
}

func (o *GetRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /Roles][%d] getRolesNotFound ", 404)
}

func (o *GetRolesNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRolesInternalServerError creates a GetRolesInternalServerError with default headers values
func NewGetRolesInternalServerError() *GetRolesInternalServerError {
	return &GetRolesInternalServerError{}
}

/*GetRolesInternalServerError handles this case with default header values.

Error
*/
type GetRolesInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Roles][%d] getRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRolesInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}