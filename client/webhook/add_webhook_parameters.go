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

	"github.com/kelvintaywl/circleci-webhook-go-sdk/models"
)

// NewAddWebhookParams creates a new AddWebhookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddWebhookParams() *AddWebhookParams {
	return &AddWebhookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddWebhookParamsWithTimeout creates a new AddWebhookParams object
// with the ability to set a timeout on a request.
func NewAddWebhookParamsWithTimeout(timeout time.Duration) *AddWebhookParams {
	return &AddWebhookParams{
		timeout: timeout,
	}
}

// NewAddWebhookParamsWithContext creates a new AddWebhookParams object
// with the ability to set a context for a request.
func NewAddWebhookParamsWithContext(ctx context.Context) *AddWebhookParams {
	return &AddWebhookParams{
		Context: ctx,
	}
}

// NewAddWebhookParamsWithHTTPClient creates a new AddWebhookParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddWebhookParamsWithHTTPClient(client *http.Client) *AddWebhookParams {
	return &AddWebhookParams{
		HTTPClient: client,
	}
}

/*
AddWebhookParams contains all the parameters to send to the API endpoint

	for the add webhook operation.

	Typically these are written to a http.Request.
*/
type AddWebhookParams struct {

	/* Body.

	   Webhook information (payload)
	*/
	Body *models.WebhookPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add webhook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddWebhookParams) WithDefaults() *AddWebhookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add webhook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddWebhookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add webhook params
func (o *AddWebhookParams) WithTimeout(timeout time.Duration) *AddWebhookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add webhook params
func (o *AddWebhookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add webhook params
func (o *AddWebhookParams) WithContext(ctx context.Context) *AddWebhookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add webhook params
func (o *AddWebhookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add webhook params
func (o *AddWebhookParams) WithHTTPClient(client *http.Client) *AddWebhookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add webhook params
func (o *AddWebhookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add webhook params
func (o *AddWebhookParams) WithBody(body *models.WebhookPayload) *AddWebhookParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add webhook params
func (o *AddWebhookParams) SetBody(body *models.WebhookPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddWebhookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
