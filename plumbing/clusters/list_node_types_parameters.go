// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListNodeTypesParams creates a new ListNodeTypesParams object
// with the default values initialized.
func NewListNodeTypesParams() *ListNodeTypesParams {

	return &ListNodeTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListNodeTypesParamsWithTimeout creates a new ListNodeTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListNodeTypesParamsWithTimeout(timeout time.Duration) *ListNodeTypesParams {

	return &ListNodeTypesParams{

		timeout: timeout,
	}
}

// NewListNodeTypesParamsWithContext creates a new ListNodeTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListNodeTypesParamsWithContext(ctx context.Context) *ListNodeTypesParams {

	return &ListNodeTypesParams{

		Context: ctx,
	}
}

// NewListNodeTypesParamsWithHTTPClient creates a new ListNodeTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListNodeTypesParamsWithHTTPClient(client *http.Client) *ListNodeTypesParams {

	return &ListNodeTypesParams{
		HTTPClient: client,
	}
}

/*ListNodeTypesParams contains all the parameters to send to the API endpoint
for the list node types operation typically these are written to a http.Request
*/
type ListNodeTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list node types params
func (o *ListNodeTypesParams) WithTimeout(timeout time.Duration) *ListNodeTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list node types params
func (o *ListNodeTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list node types params
func (o *ListNodeTypesParams) WithContext(ctx context.Context) *ListNodeTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list node types params
func (o *ListNodeTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list node types params
func (o *ListNodeTypesParams) WithHTTPClient(client *http.Client) *ListNodeTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list node types params
func (o *ListNodeTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListNodeTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
