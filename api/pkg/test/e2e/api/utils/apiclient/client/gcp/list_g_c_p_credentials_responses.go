// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListGCPCredentialsReader is a Reader for the ListGCPCredentials structure.
type ListGCPCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListGCPCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListGCPCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPCredentialsOK creates a ListGCPCredentialsOK with default headers values
func NewListGCPCredentialsOK() *ListGCPCredentialsOK {
	return &ListGCPCredentialsOK{}
}

/*ListGCPCredentialsOK handles this case with default header values.

CredentialList
*/
type ListGCPCredentialsOK struct {
	Payload *models.CredentialList
}

func (o *ListGCPCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/credentials][%d] listGCPCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListGCPCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPCredentialsDefault creates a ListGCPCredentialsDefault with default headers values
func NewListGCPCredentialsDefault(code int) *ListGCPCredentialsDefault {
	return &ListGCPCredentialsDefault{
		_statusCode: code,
	}
}

/*ListGCPCredentialsDefault handles this case with default header values.

ErrorResponse is the default representation of an error
*/
type ListGCPCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorDetails
}

// Code gets the status code for the list g c p credentials default response
func (o *ListGCPCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListGCPCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/credentials][%d] listGCPCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
