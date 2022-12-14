// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListWebhooksParams creates a new ListWebhooksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListWebhooksParams() *ListWebhooksParams {
	return &ListWebhooksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListWebhooksParamsWithTimeout creates a new ListWebhooksParams object
// with the ability to set a timeout on a request.
func NewListWebhooksParamsWithTimeout(timeout time.Duration) *ListWebhooksParams {
	return &ListWebhooksParams{
		timeout: timeout,
	}
}

// NewListWebhooksParamsWithContext creates a new ListWebhooksParams object
// with the ability to set a context for a request.
func NewListWebhooksParamsWithContext(ctx context.Context) *ListWebhooksParams {
	return &ListWebhooksParams{
		Context: ctx,
	}
}

// NewListWebhooksParamsWithHTTPClient creates a new ListWebhooksParams object
// with the ability to set a custom HTTPClient for a request.
func NewListWebhooksParamsWithHTTPClient(client *http.Client) *ListWebhooksParams {
	return &ListWebhooksParams{
		HTTPClient: client,
	}
}

/*
ListWebhooksParams contains all the parameters to send to the API endpoint

	for the list webhooks operation.

	Typically these are written to a http.Request.
*/
type ListWebhooksParams struct {

	/* ScopeID.

	   ID of the scope being used (at the moment, only project ID is supported)

	   Format: uuid
	*/
	ScopeID strfmt.UUID

	/* ScopeType.

	   Type of the scope being used

	   Default: "project"
	*/
	ScopeType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list webhooks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListWebhooksParams) WithDefaults() *ListWebhooksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list webhooks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListWebhooksParams) SetDefaults() {
	var (
		scopeTypeDefault = string("project")
	)

	val := ListWebhooksParams{
		ScopeType: scopeTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list webhooks params
func (o *ListWebhooksParams) WithTimeout(timeout time.Duration) *ListWebhooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list webhooks params
func (o *ListWebhooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list webhooks params
func (o *ListWebhooksParams) WithContext(ctx context.Context) *ListWebhooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list webhooks params
func (o *ListWebhooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list webhooks params
func (o *ListWebhooksParams) WithHTTPClient(client *http.Client) *ListWebhooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list webhooks params
func (o *ListWebhooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScopeID adds the scopeID to the list webhooks params
func (o *ListWebhooksParams) WithScopeID(scopeID strfmt.UUID) *ListWebhooksParams {
	o.SetScopeID(scopeID)
	return o
}

// SetScopeID adds the scopeId to the list webhooks params
func (o *ListWebhooksParams) SetScopeID(scopeID strfmt.UUID) {
	o.ScopeID = scopeID
}

// WithScopeType adds the scopeType to the list webhooks params
func (o *ListWebhooksParams) WithScopeType(scopeType string) *ListWebhooksParams {
	o.SetScopeType(scopeType)
	return o
}

// SetScopeType adds the scopeType to the list webhooks params
func (o *ListWebhooksParams) SetScopeType(scopeType string) {
	o.ScopeType = scopeType
}

// WriteToRequest writes these params to a swagger request
func (o *ListWebhooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param scope-id
	qrScopeID := o.ScopeID
	qScopeID := qrScopeID.String()
	if qScopeID != "" {

		if err := r.SetQueryParam("scope-id", qScopeID); err != nil {
			return err
		}
	}

	// query param scope-type
	qrScopeType := o.ScopeType
	qScopeType := qrScopeType
	if qScopeType != "" {

		if err := r.SetQueryParam("scope-type", qScopeType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
