package booleangroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// BoolClient is the test Infrastructure for AutoRest
type BoolClient struct {
	BaseClient
}

// NewBoolClient creates an instance of the BoolClient client.
func NewBoolClient() BoolClient {
	return NewBoolClientWithBaseURI(DefaultBaseURI)
}

// NewBoolClientWithBaseURI creates an instance of the BoolClient client.
func NewBoolClientWithBaseURI(baseURI string) BoolClient {
	return BoolClient{NewWithBaseURI(baseURI)}
}

// GetFalse get false Boolean value
func (client BoolClient) GetFalse(ctx context.Context) (result BoolModel, err error) {
	req, err := client.GetFalsePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetFalse", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFalseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetFalse", resp, "Failure sending request")
		return
	}

	result, err = client.GetFalseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetFalse", resp, "Failure responding to request")
	}

	return
}

// GetFalsePreparer prepares the GetFalse request.
func (client BoolClient) GetFalsePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/bool/false"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFalseSender sends the GetFalse request. The method will close the
// http.Response Body if it receives an error.
func (client BoolClient) GetFalseSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetFalseResponder handles the response to the GetFalse request. The method always
// closes the http.Response Body.
func (client BoolClient) GetFalseResponder(resp *http.Response) (result BoolModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetInvalid get invalid Boolean value
func (client BoolClient) GetInvalid(ctx context.Context) (result BoolModel, err error) {
	req, err := client.GetInvalidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetInvalid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetInvalid", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetInvalid", resp, "Failure responding to request")
	}

	return
}

// GetInvalidPreparer prepares the GetInvalid request.
func (client BoolClient) GetInvalidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/bool/invalid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInvalidSender sends the GetInvalid request. The method will close the
// http.Response Body if it receives an error.
func (client BoolClient) GetInvalidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetInvalidResponder handles the response to the GetInvalid request. The method always
// closes the http.Response Body.
func (client BoolClient) GetInvalidResponder(resp *http.Response) (result BoolModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get null Boolean value
func (client BoolClient) GetNull(ctx context.Context) (result BoolModel, err error) {
	req, err := client.GetNullPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client BoolClient) GetNullPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/bool/null"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client BoolClient) GetNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client BoolClient) GetNullResponder(resp *http.Response) (result BoolModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTrue get true Boolean value
func (client BoolClient) GetTrue(ctx context.Context) (result BoolModel, err error) {
	req, err := client.GetTruePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetTrue", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTrueSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetTrue", resp, "Failure sending request")
		return
	}

	result, err = client.GetTrueResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "GetTrue", resp, "Failure responding to request")
	}

	return
}

// GetTruePreparer prepares the GetTrue request.
func (client BoolClient) GetTruePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/bool/true"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTrueSender sends the GetTrue request. The method will close the
// http.Response Body if it receives an error.
func (client BoolClient) GetTrueSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetTrueResponder handles the response to the GetTrue request. The method always
// closes the http.Response Body.
func (client BoolClient) GetTrueResponder(resp *http.Response) (result BoolModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutFalse set Boolean value false
func (client BoolClient) PutFalse(ctx context.Context) (result autorest.Response, err error) {
	req, err := client.PutFalsePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "PutFalse", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutFalseSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "PutFalse", resp, "Failure sending request")
		return
	}

	result, err = client.PutFalseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "PutFalse", resp, "Failure responding to request")
	}

	return
}

// PutFalsePreparer prepares the PutFalse request.
func (client BoolClient) PutFalsePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/bool/false"),
		autorest.WithJSON(false))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutFalseSender sends the PutFalse request. The method will close the
// http.Response Body if it receives an error.
func (client BoolClient) PutFalseSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutFalseResponder handles the response to the PutFalse request. The method always
// closes the http.Response Body.
func (client BoolClient) PutFalseResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutTrue set Boolean value true
func (client BoolClient) PutTrue(ctx context.Context) (result autorest.Response, err error) {
	req, err := client.PutTruePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "PutTrue", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutTrueSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "PutTrue", resp, "Failure sending request")
		return
	}

	result, err = client.PutTrueResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "booleangroup.BoolClient", "PutTrue", resp, "Failure responding to request")
	}

	return
}

// PutTruePreparer prepares the PutTrue request.
func (client BoolClient) PutTruePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/bool/true"),
		autorest.WithJSON(true))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutTrueSender sends the PutTrue request. The method will close the
// http.Response Body if it receives an error.
func (client BoolClient) PutTrueSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutTrueResponder handles the response to the PutTrue request. The method always
// closes the http.Response Body.
func (client BoolClient) PutTrueResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
