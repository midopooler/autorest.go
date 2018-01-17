package urlgroup

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

// PathItemsClient is the test Infrastructure for AutoRest
type PathItemsClient struct {
	BaseClient
}

// NewPathItemsClient creates an instance of the PathItemsClient client.
func NewPathItemsClient(globalStringPath string, globalStringQuery string) PathItemsClient {
	return NewPathItemsClientWithBaseURI(DefaultBaseURI, globalStringPath, globalStringQuery)
}

// NewPathItemsClientWithBaseURI creates an instance of the PathItemsClient client.
func NewPathItemsClientWithBaseURI(baseURI string, globalStringPath string, globalStringQuery string) PathItemsClient {
	return PathItemsClient{NewWithBaseURI(baseURI, globalStringPath, globalStringQuery)}
}

// GetAllWithValues send globalStringPath='globalStringPath', pathItemStringPath='pathItemStringPath',
// localStringPath='localStringPath', globalStringQuery='globalStringQuery', pathItemStringQuery='pathItemStringQuery',
// localStringQuery='localStringQuery'
//
// localStringPath is should contain value 'localStringPath' pathItemStringPath is a string value
// 'pathItemStringPath' that appears in the path localStringQuery is should contain value 'localStringQuery'
// pathItemStringQuery is a string value 'pathItemStringQuery' that appears as a query parameter
func (client PathItemsClient) GetAllWithValues(ctx context.Context, localStringPath string, pathItemStringPath string, localStringQuery string, pathItemStringQuery string) (result autorest.Response, err error) {
	req, err := client.GetAllWithValuesPreparer(ctx, localStringPath, pathItemStringPath, localStringQuery, pathItemStringQuery)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetAllWithValues", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllWithValuesSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetAllWithValues", resp, "Failure sending request")
		return
	}

	result, err = client.GetAllWithValuesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetAllWithValues", resp, "Failure responding to request")
	}

	return
}

// GetAllWithValuesPreparer prepares the GetAllWithValues request.
func (client PathItemsClient) GetAllWithValuesPreparer(ctx context.Context, localStringPath string, pathItemStringPath string, localStringQuery string, pathItemStringQuery string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"globalStringPath":   autorest.Encode("path", client.GlobalStringPath),
		"localStringPath":    autorest.Encode("path", localStringPath),
		"pathItemStringPath": autorest.Encode("path", pathItemStringPath),
	}

	queryParameters := map[string]interface{}{}
	if len(localStringQuery) > 0 {
		queryParameters["localStringQuery"] = autorest.Encode("query", localStringQuery)
	}
	if len(pathItemStringQuery) > 0 {
		queryParameters["pathItemStringQuery"] = autorest.Encode("query", pathItemStringQuery)
	}
	if len(client.GlobalStringQuery) > 0 {
		queryParameters["globalStringQuery"] = autorest.Encode("query", client.GlobalStringQuery)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/globalStringQuery/pathItemStringQuery/localStringQuery", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllWithValuesSender sends the GetAllWithValues request. The method will close the
// http.Response Body if it receives an error.
func (client PathItemsClient) GetAllWithValuesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAllWithValuesResponder handles the response to the GetAllWithValues request. The method always
// closes the http.Response Body.
func (client PathItemsClient) GetAllWithValuesResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetGlobalAndLocalQueryNull send globalStringPath=globalStringPath, pathItemStringPath='pathItemStringPath',
// localStringPath='localStringPath', globalStringQuery=null, pathItemStringQuery='pathItemStringQuery',
// localStringQuery=null
//
// localStringPath is should contain value 'localStringPath' pathItemStringPath is a string value
// 'pathItemStringPath' that appears in the path localStringQuery is should contain null value pathItemStringQuery
// is a string value 'pathItemStringQuery' that appears as a query parameter
func (client PathItemsClient) GetGlobalAndLocalQueryNull(ctx context.Context, localStringPath string, pathItemStringPath string, localStringQuery string, pathItemStringQuery string) (result autorest.Response, err error) {
	req, err := client.GetGlobalAndLocalQueryNullPreparer(ctx, localStringPath, pathItemStringPath, localStringQuery, pathItemStringQuery)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetGlobalAndLocalQueryNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetGlobalAndLocalQueryNullSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetGlobalAndLocalQueryNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetGlobalAndLocalQueryNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetGlobalAndLocalQueryNull", resp, "Failure responding to request")
	}

	return
}

// GetGlobalAndLocalQueryNullPreparer prepares the GetGlobalAndLocalQueryNull request.
func (client PathItemsClient) GetGlobalAndLocalQueryNullPreparer(ctx context.Context, localStringPath string, pathItemStringPath string, localStringQuery string, pathItemStringQuery string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"globalStringPath":   autorest.Encode("path", client.GlobalStringPath),
		"localStringPath":    autorest.Encode("path", localStringPath),
		"pathItemStringPath": autorest.Encode("path", pathItemStringPath),
	}

	queryParameters := map[string]interface{}{}
	if len(localStringQuery) > 0 {
		queryParameters["localStringQuery"] = autorest.Encode("query", localStringQuery)
	}
	if len(pathItemStringQuery) > 0 {
		queryParameters["pathItemStringQuery"] = autorest.Encode("query", pathItemStringQuery)
	}
	if len(client.GlobalStringQuery) > 0 {
		queryParameters["globalStringQuery"] = autorest.Encode("query", client.GlobalStringQuery)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/null/pathItemStringQuery/null", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetGlobalAndLocalQueryNullSender sends the GetGlobalAndLocalQueryNull request. The method will close the
// http.Response Body if it receives an error.
func (client PathItemsClient) GetGlobalAndLocalQueryNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetGlobalAndLocalQueryNullResponder handles the response to the GetGlobalAndLocalQueryNull request. The method always
// closes the http.Response Body.
func (client PathItemsClient) GetGlobalAndLocalQueryNullResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetGlobalQueryNull send globalStringPath='globalStringPath', pathItemStringPath='pathItemStringPath',
// localStringPath='localStringPath', globalStringQuery=null, pathItemStringQuery='pathItemStringQuery',
// localStringQuery='localStringQuery'
//
// localStringPath is should contain value 'localStringPath' pathItemStringPath is a string value
// 'pathItemStringPath' that appears in the path localStringQuery is should contain value 'localStringQuery'
// pathItemStringQuery is a string value 'pathItemStringQuery' that appears as a query parameter
func (client PathItemsClient) GetGlobalQueryNull(ctx context.Context, localStringPath string, pathItemStringPath string, localStringQuery string, pathItemStringQuery string) (result autorest.Response, err error) {
	req, err := client.GetGlobalQueryNullPreparer(ctx, localStringPath, pathItemStringPath, localStringQuery, pathItemStringQuery)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetGlobalQueryNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetGlobalQueryNullSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetGlobalQueryNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetGlobalQueryNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetGlobalQueryNull", resp, "Failure responding to request")
	}

	return
}

// GetGlobalQueryNullPreparer prepares the GetGlobalQueryNull request.
func (client PathItemsClient) GetGlobalQueryNullPreparer(ctx context.Context, localStringPath string, pathItemStringPath string, localStringQuery string, pathItemStringQuery string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"globalStringPath":   autorest.Encode("path", client.GlobalStringPath),
		"localStringPath":    autorest.Encode("path", localStringPath),
		"pathItemStringPath": autorest.Encode("path", pathItemStringPath),
	}

	queryParameters := map[string]interface{}{}
	if len(localStringQuery) > 0 {
		queryParameters["localStringQuery"] = autorest.Encode("query", localStringQuery)
	}
	if len(pathItemStringQuery) > 0 {
		queryParameters["pathItemStringQuery"] = autorest.Encode("query", pathItemStringQuery)
	}
	if len(client.GlobalStringQuery) > 0 {
		queryParameters["globalStringQuery"] = autorest.Encode("query", client.GlobalStringQuery)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/null/pathItemStringQuery/localStringQuery", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetGlobalQueryNullSender sends the GetGlobalQueryNull request. The method will close the
// http.Response Body if it receives an error.
func (client PathItemsClient) GetGlobalQueryNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetGlobalQueryNullResponder handles the response to the GetGlobalQueryNull request. The method always
// closes the http.Response Body.
func (client PathItemsClient) GetGlobalQueryNullResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetLocalPathItemQueryNull send globalStringPath='globalStringPath', pathItemStringPath='pathItemStringPath',
// localStringPath='localStringPath', globalStringQuery='globalStringQuery', pathItemStringQuery=null,
// localStringQuery=null
//
// localStringPath is should contain value 'localStringPath' pathItemStringPath is a string value
// 'pathItemStringPath' that appears in the path localStringQuery is should contain value null pathItemStringQuery
// is should contain value null
func (client PathItemsClient) GetLocalPathItemQueryNull(ctx context.Context, localStringPath string, pathItemStringPath string, localStringQuery string, pathItemStringQuery string) (result autorest.Response, err error) {
	req, err := client.GetLocalPathItemQueryNullPreparer(ctx, localStringPath, pathItemStringPath, localStringQuery, pathItemStringQuery)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetLocalPathItemQueryNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLocalPathItemQueryNullSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetLocalPathItemQueryNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetLocalPathItemQueryNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlgroup.PathItemsClient", "GetLocalPathItemQueryNull", resp, "Failure responding to request")
	}

	return
}

// GetLocalPathItemQueryNullPreparer prepares the GetLocalPathItemQueryNull request.
func (client PathItemsClient) GetLocalPathItemQueryNullPreparer(ctx context.Context, localStringPath string, pathItemStringPath string, localStringQuery string, pathItemStringQuery string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"globalStringPath":   autorest.Encode("path", client.GlobalStringPath),
		"localStringPath":    autorest.Encode("path", localStringPath),
		"pathItemStringPath": autorest.Encode("path", pathItemStringPath),
	}

	queryParameters := map[string]interface{}{}
	if len(localStringQuery) > 0 {
		queryParameters["localStringQuery"] = autorest.Encode("query", localStringQuery)
	}
	if len(pathItemStringQuery) > 0 {
		queryParameters["pathItemStringQuery"] = autorest.Encode("query", pathItemStringQuery)
	}
	if len(client.GlobalStringQuery) > 0 {
		queryParameters["globalStringQuery"] = autorest.Encode("query", client.GlobalStringQuery)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/globalStringQuery/null/null", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLocalPathItemQueryNullSender sends the GetLocalPathItemQueryNull request. The method will close the
// http.Response Body if it receives an error.
func (client PathItemsClient) GetLocalPathItemQueryNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetLocalPathItemQueryNullResponder handles the response to the GetLocalPathItemQueryNull request. The method always
// closes the http.Response Body.
func (client PathItemsClient) GetLocalPathItemQueryNullResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
