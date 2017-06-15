package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rai-project/dlframework/web/models"
)

// GetFrameworkModelManifestReader is a Reader for the GetFrameworkModelManifest structure.
type GetFrameworkModelManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFrameworkModelManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFrameworkModelManifestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFrameworkModelManifestOK creates a GetFrameworkModelManifestOK with default headers values
func NewGetFrameworkModelManifestOK() *GetFrameworkModelManifestOK {
	return &GetFrameworkModelManifestOK{}
}

/*GetFrameworkModelManifestOK handles this case with default header values.

GetFrameworkModelManifestOK get framework model manifest o k
*/
type GetFrameworkModelManifestOK struct {
	Payload *models.DlframeworkModelManifest
}

func (o *GetFrameworkModelManifestOK) Error() string {
	return fmt.Sprintf("[POST /v1/framework/{framework_name}/{framework_version}/model/{model_name}/{model_version}/info][%d] getFrameworkModelManifestOK  %+v", 200, o.Payload)
}

func (o *GetFrameworkModelManifestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkModelManifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
