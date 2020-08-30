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

// DownloadClusterISOReader is a Reader for the DownloadClusterISO structure.
type DownloadClusterISOReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadClusterISOReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadClusterISOOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDownloadClusterISOBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDownloadClusterISOUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadClusterISOForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadClusterISONotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDownloadClusterISOMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDownloadClusterISOConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadClusterISOInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDownloadClusterISOOK creates a DownloadClusterISOOK with default headers values
func NewDownloadClusterISOOK(writer io.Writer) *DownloadClusterISOOK {
	return &DownloadClusterISOOK{
		Payload: writer,
	}
}

/*DownloadClusterISOOK handles this case with default header values.

Success.
*/
type DownloadClusterISOOK struct {
	Payload io.Writer
}

func (o *DownloadClusterISOOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image][%d] downloadClusterISOOK  %+v", 200, o.Payload)
}

func (o *DownloadClusterISOOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadClusterISOOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOBadRequest creates a DownloadClusterISOBadRequest with default headers values
func NewDownloadClusterISOBadRequest() *DownloadClusterISOBadRequest {
	return &DownloadClusterISOBadRequest{}
}

/*DownloadClusterISOBadRequest handles this case with default header values.

Error.
*/
type DownloadClusterISOBadRequest struct {
	Payload *models.Error
}

func (o *DownloadClusterISOBadRequest) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image][%d] downloadClusterISOBadRequest  %+v", 400, o.Payload)
}

func (o *DownloadClusterISOBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterISOBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOUnauthorized creates a DownloadClusterISOUnauthorized with default headers values
func NewDownloadClusterISOUnauthorized() *DownloadClusterISOUnauthorized {
	return &DownloadClusterISOUnauthorized{}
}

/*DownloadClusterISOUnauthorized handles this case with default header values.

Unauthorized.
*/
type DownloadClusterISOUnauthorized struct {
	Payload *models.InfraError
}

func (o *DownloadClusterISOUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image][%d] downloadClusterISOUnauthorized  %+v", 401, o.Payload)
}

func (o *DownloadClusterISOUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DownloadClusterISOUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOForbidden creates a DownloadClusterISOForbidden with default headers values
func NewDownloadClusterISOForbidden() *DownloadClusterISOForbidden {
	return &DownloadClusterISOForbidden{}
}

/*DownloadClusterISOForbidden handles this case with default header values.

Forbidden.
*/
type DownloadClusterISOForbidden struct {
	Payload *models.InfraError
}

func (o *DownloadClusterISOForbidden) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image][%d] downloadClusterISOForbidden  %+v", 403, o.Payload)
}

func (o *DownloadClusterISOForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DownloadClusterISOForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISONotFound creates a DownloadClusterISONotFound with default headers values
func NewDownloadClusterISONotFound() *DownloadClusterISONotFound {
	return &DownloadClusterISONotFound{}
}

/*DownloadClusterISONotFound handles this case with default header values.

Error.
*/
type DownloadClusterISONotFound struct {
	Payload *models.Error
}

func (o *DownloadClusterISONotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image][%d] downloadClusterISONotFound  %+v", 404, o.Payload)
}

func (o *DownloadClusterISONotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterISONotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOMethodNotAllowed creates a DownloadClusterISOMethodNotAllowed with default headers values
func NewDownloadClusterISOMethodNotAllowed() *DownloadClusterISOMethodNotAllowed {
	return &DownloadClusterISOMethodNotAllowed{}
}

/*DownloadClusterISOMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type DownloadClusterISOMethodNotAllowed struct {
}

func (o *DownloadClusterISOMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image][%d] downloadClusterISOMethodNotAllowed ", 405)
}

func (o *DownloadClusterISOMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadClusterISOConflict creates a DownloadClusterISOConflict with default headers values
func NewDownloadClusterISOConflict() *DownloadClusterISOConflict {
	return &DownloadClusterISOConflict{}
}

/*DownloadClusterISOConflict handles this case with default header values.

Error.
*/
type DownloadClusterISOConflict struct {
	Payload *models.Error
}

func (o *DownloadClusterISOConflict) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image][%d] downloadClusterISOConflict  %+v", 409, o.Payload)
}

func (o *DownloadClusterISOConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterISOConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOInternalServerError creates a DownloadClusterISOInternalServerError with default headers values
func NewDownloadClusterISOInternalServerError() *DownloadClusterISOInternalServerError {
	return &DownloadClusterISOInternalServerError{}
}

/*DownloadClusterISOInternalServerError handles this case with default header values.

Error.
*/
type DownloadClusterISOInternalServerError struct {
	Payload *models.Error
}

func (o *DownloadClusterISOInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image][%d] downloadClusterISOInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadClusterISOInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterISOInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
