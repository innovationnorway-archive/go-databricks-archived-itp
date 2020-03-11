// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/innovationnorway/go-databricks/models"
)

// ListNodeTypesReader is a Reader for the ListNodeTypes structure.
type ListNodeTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNodeTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNodeTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListNodeTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListNodeTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNodeTypesOK creates a ListNodeTypesOK with default headers values
func NewListNodeTypesOK() *ListNodeTypesOK {
	return &ListNodeTypesOK{}
}

/*ListNodeTypesOK handles this case with default header values.

OK
*/
type ListNodeTypesOK struct {
	Payload []*models.NodeType
}

func (o *ListNodeTypesOK) Error() string {
	return fmt.Sprintf("[GET /clusters/list-node-types][%d] listNodeTypesOK  %+v", 200, o.Payload)
}

func (o *ListNodeTypesOK) GetPayload() []*models.NodeType {
	return o.Payload
}

func (o *ListNodeTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNodeTypesForbidden creates a ListNodeTypesForbidden with default headers values
func NewListNodeTypesForbidden() *ListNodeTypesForbidden {
	return &ListNodeTypesForbidden{}
}

/*ListNodeTypesForbidden handles this case with default header values.

invalid access token
*/
type ListNodeTypesForbidden struct {
	Payload string
}

func (o *ListNodeTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /clusters/list-node-types][%d] listNodeTypesForbidden  %+v", 403, o.Payload)
}

func (o *ListNodeTypesForbidden) GetPayload() string {
	return o.Payload
}

func (o *ListNodeTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNodeTypesDefault creates a ListNodeTypesDefault with default headers values
func NewListNodeTypesDefault(code int) *ListNodeTypesDefault {
	return &ListNodeTypesDefault{
		_statusCode: code,
	}
}

/*ListNodeTypesDefault handles this case with default header values.

error
*/
type ListNodeTypesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list node types default response
func (o *ListNodeTypesDefault) Code() int {
	return o._statusCode
}

func (o *ListNodeTypesDefault) Error() string {
	return fmt.Sprintf("[GET /clusters/list-node-types][%d] listNodeTypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListNodeTypesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListNodeTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
