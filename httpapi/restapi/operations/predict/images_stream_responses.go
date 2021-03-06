// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rai-project/dlframework/httpapi/models"
)

// ImagesStreamOKCode is the HTTP code returned for type ImagesStreamOK
const ImagesStreamOKCode int = 200

/*ImagesStreamOK (streaming responses)

swagger:response imagesStreamOK
*/
type ImagesStreamOK struct {

	/*
	  In: Body
	*/
	Payload *models.DlframeworkFeatureResponse `json:"body,omitempty"`
}

// NewImagesStreamOK creates ImagesStreamOK with default headers values
func NewImagesStreamOK() *ImagesStreamOK {
	return &ImagesStreamOK{}
}

// WithPayload adds the payload to the images stream o k response
func (o *ImagesStreamOK) WithPayload(payload *models.DlframeworkFeatureResponse) *ImagesStreamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the images stream o k response
func (o *ImagesStreamOK) SetPayload(payload *models.DlframeworkFeatureResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImagesStreamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
