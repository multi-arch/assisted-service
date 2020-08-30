// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// GetNextStepsOKCode is the HTTP code returned for type GetNextStepsOK
const GetNextStepsOKCode int = 200

/*GetNextStepsOK Success.

swagger:response getNextStepsOK
*/
type GetNextStepsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Steps `json:"body,omitempty"`
}

// NewGetNextStepsOK creates GetNextStepsOK with default headers values
func NewGetNextStepsOK() *GetNextStepsOK {

	return &GetNextStepsOK{}
}

// WithPayload adds the payload to the get next steps o k response
func (o *GetNextStepsOK) WithPayload(payload *models.Steps) *GetNextStepsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get next steps o k response
func (o *GetNextStepsOK) SetPayload(payload *models.Steps) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNextStepsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNextStepsUnauthorizedCode is the HTTP code returned for type GetNextStepsUnauthorized
const GetNextStepsUnauthorizedCode int = 401

/*GetNextStepsUnauthorized Unauthorized.

swagger:response getNextStepsUnauthorized
*/
type GetNextStepsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetNextStepsUnauthorized creates GetNextStepsUnauthorized with default headers values
func NewGetNextStepsUnauthorized() *GetNextStepsUnauthorized {

	return &GetNextStepsUnauthorized{}
}

// WithPayload adds the payload to the get next steps unauthorized response
func (o *GetNextStepsUnauthorized) WithPayload(payload *models.InfraError) *GetNextStepsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get next steps unauthorized response
func (o *GetNextStepsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNextStepsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNextStepsForbiddenCode is the HTTP code returned for type GetNextStepsForbidden
const GetNextStepsForbiddenCode int = 403

/*GetNextStepsForbidden Forbidden.

swagger:response getNextStepsForbidden
*/
type GetNextStepsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetNextStepsForbidden creates GetNextStepsForbidden with default headers values
func NewGetNextStepsForbidden() *GetNextStepsForbidden {

	return &GetNextStepsForbidden{}
}

// WithPayload adds the payload to the get next steps forbidden response
func (o *GetNextStepsForbidden) WithPayload(payload *models.InfraError) *GetNextStepsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get next steps forbidden response
func (o *GetNextStepsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNextStepsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNextStepsNotFoundCode is the HTTP code returned for type GetNextStepsNotFound
const GetNextStepsNotFoundCode int = 404

/*GetNextStepsNotFound Error.

swagger:response getNextStepsNotFound
*/
type GetNextStepsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNextStepsNotFound creates GetNextStepsNotFound with default headers values
func NewGetNextStepsNotFound() *GetNextStepsNotFound {

	return &GetNextStepsNotFound{}
}

// WithPayload adds the payload to the get next steps not found response
func (o *GetNextStepsNotFound) WithPayload(payload *models.Error) *GetNextStepsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get next steps not found response
func (o *GetNextStepsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNextStepsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNextStepsMethodNotAllowedCode is the HTTP code returned for type GetNextStepsMethodNotAllowed
const GetNextStepsMethodNotAllowedCode int = 405

/*GetNextStepsMethodNotAllowed Method Not Allowed.

swagger:response getNextStepsMethodNotAllowed
*/
type GetNextStepsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNextStepsMethodNotAllowed creates GetNextStepsMethodNotAllowed with default headers values
func NewGetNextStepsMethodNotAllowed() *GetNextStepsMethodNotAllowed {

	return &GetNextStepsMethodNotAllowed{}
}

// WithPayload adds the payload to the get next steps method not allowed response
func (o *GetNextStepsMethodNotAllowed) WithPayload(payload *models.Error) *GetNextStepsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get next steps method not allowed response
func (o *GetNextStepsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNextStepsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNextStepsInternalServerErrorCode is the HTTP code returned for type GetNextStepsInternalServerError
const GetNextStepsInternalServerErrorCode int = 500

/*GetNextStepsInternalServerError Error.

swagger:response getNextStepsInternalServerError
*/
type GetNextStepsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNextStepsInternalServerError creates GetNextStepsInternalServerError with default headers values
func NewGetNextStepsInternalServerError() *GetNextStepsInternalServerError {

	return &GetNextStepsInternalServerError{}
}

// WithPayload adds the payload to the get next steps internal server error response
func (o *GetNextStepsInternalServerError) WithPayload(payload *models.Error) *GetNextStepsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get next steps internal server error response
func (o *GetNextStepsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNextStepsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
