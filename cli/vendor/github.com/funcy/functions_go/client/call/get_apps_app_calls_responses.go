package call

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/funcy/functions_go/models"
)

// GetAppsAppCallsReader is a Reader for the GetAppsAppCalls structure.
type GetAppsAppCallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppsAppCallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAppsAppCallsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetAppsAppCallsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppsAppCallsOK creates a GetAppsAppCallsOK with default headers values
func NewGetAppsAppCallsOK() *GetAppsAppCallsOK {
	return &GetAppsAppCallsOK{}
}

/*GetAppsAppCallsOK handles this case with default header values.

Calls found
*/
type GetAppsAppCallsOK struct {
	Payload *models.CallsWrapper
}

func (o *GetAppsAppCallsOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app}/calls/][%d] getAppsAppCallsOK  %+v", 200, o.Payload)
}

func (o *GetAppsAppCallsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallsWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppsAppCallsNotFound creates a GetAppsAppCallsNotFound with default headers values
func NewGetAppsAppCallsNotFound() *GetAppsAppCallsNotFound {
	return &GetAppsAppCallsNotFound{}
}

/*GetAppsAppCallsNotFound handles this case with default header values.

Calls not found.
*/
type GetAppsAppCallsNotFound struct {
	Payload *models.Error
}

func (o *GetAppsAppCallsNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app}/calls/][%d] getAppsAppCallsNotFound  %+v", 404, o.Payload)
}

func (o *GetAppsAppCallsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
