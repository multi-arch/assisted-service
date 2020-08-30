// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// SetDebugStepReader is a Reader for the SetDebugStep structure.
type SetDebugStepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDebugStepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSetDebugStepNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetDebugStepUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetDebugStepForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetDebugStepNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewSetDebugStepMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetDebugStepInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetDebugStepNoContent creates a SetDebugStepNoContent with default headers values
func NewSetDebugStepNoContent() *SetDebugStepNoContent {
	return &SetDebugStepNoContent{}
}

/*SetDebugStepNoContent handles this case with default header values.

Success.
*/
type SetDebugStepNoContent struct {
}

func (o *SetDebugStepNoContent) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts/{host_id}/actions/debug][%d] setDebugStepNoContent ", 204)
}

func (o *SetDebugStepNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetDebugStepUnauthorized creates a SetDebugStepUnauthorized with default headers values
func NewSetDebugStepUnauthorized() *SetDebugStepUnauthorized {
	return &SetDebugStepUnauthorized{}
}

/*SetDebugStepUnauthorized handles this case with default header values.

Unauthorized.
*/
type SetDebugStepUnauthorized struct {
	Payload *models.InfraError
}

func (o *SetDebugStepUnauthorized) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts/{host_id}/actions/debug][%d] setDebugStepUnauthorized  %+v", 401, o.Payload)
}

func (o *SetDebugStepUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *SetDebugStepUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDebugStepForbidden creates a SetDebugStepForbidden with default headers values
func NewSetDebugStepForbidden() *SetDebugStepForbidden {
	return &SetDebugStepForbidden{}
}

/*SetDebugStepForbidden handles this case with default header values.

Forbidden.
*/
type SetDebugStepForbidden struct {
	Payload *models.InfraError
}

func (o *SetDebugStepForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts/{host_id}/actions/debug][%d] setDebugStepForbidden  %+v", 403, o.Payload)
}

func (o *SetDebugStepForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *SetDebugStepForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDebugStepNotFound creates a SetDebugStepNotFound with default headers values
func NewSetDebugStepNotFound() *SetDebugStepNotFound {
	return &SetDebugStepNotFound{}
}

/*SetDebugStepNotFound handles this case with default header values.

Error.
*/
type SetDebugStepNotFound struct {
	Payload *models.Error
}

func (o *SetDebugStepNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts/{host_id}/actions/debug][%d] setDebugStepNotFound  %+v", 404, o.Payload)
}

func (o *SetDebugStepNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetDebugStepNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDebugStepMethodNotAllowed creates a SetDebugStepMethodNotAllowed with default headers values
func NewSetDebugStepMethodNotAllowed() *SetDebugStepMethodNotAllowed {
	return &SetDebugStepMethodNotAllowed{}
}

/*SetDebugStepMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type SetDebugStepMethodNotAllowed struct {
	Payload *models.Error
}

func (o *SetDebugStepMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts/{host_id}/actions/debug][%d] setDebugStepMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SetDebugStepMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetDebugStepMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDebugStepInternalServerError creates a SetDebugStepInternalServerError with default headers values
func NewSetDebugStepInternalServerError() *SetDebugStepInternalServerError {
	return &SetDebugStepInternalServerError{}
}

/*SetDebugStepInternalServerError handles this case with default header values.

Error.
*/
type SetDebugStepInternalServerError struct {
	Payload *models.Error
}

func (o *SetDebugStepInternalServerError) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts/{host_id}/actions/debug][%d] setDebugStepInternalServerError  %+v", 500, o.Payload)
}

func (o *SetDebugStepInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetDebugStepInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
