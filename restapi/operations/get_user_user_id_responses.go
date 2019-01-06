// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sakata1222/go-rest-api-sample/models"
)

// GetUserUserIDOKCode is the HTTP code returned for type GetUserUserIDOK
const GetUserUserIDOKCode int = 200

/*GetUserUserIDOK Response

swagger:response getUserUserIdOK
*/
type GetUserUserIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserUserIDOK creates GetUserUserIDOK with default headers values
func NewGetUserUserIDOK() *GetUserUserIDOK {

	return &GetUserUserIDOK{}
}

// WithPayload adds the payload to the get user user Id o k response
func (o *GetUserUserIDOK) WithPayload(payload *models.User) *GetUserUserIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user user Id o k response
func (o *GetUserUserIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserUserIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
