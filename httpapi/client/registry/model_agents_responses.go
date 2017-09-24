// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rai-project/dlframework/httpapi/models"
)

// ModelAgentsReader is a Reader for the ModelAgents structure.
type ModelAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModelAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModelAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModelAgentsOK creates a ModelAgentsOK with default headers values
func NewModelAgentsOK() *ModelAgentsOK {
	return &ModelAgentsOK{}
}

/*ModelAgentsOK handles this case with default header values.

ModelAgentsOK model agents o k
*/
type ModelAgentsOK struct {
	Payload *models.DlframeworkAgents
}

func (o *ModelAgentsOK) Error() string {
	return fmt.Sprintf("[GET /registry/models/agent][%d] modelAgentsOK  %+v", 200, o.Payload)
}

func (o *ModelAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkAgents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
