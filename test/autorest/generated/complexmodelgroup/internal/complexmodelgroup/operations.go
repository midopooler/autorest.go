// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexmodelgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Operations contains the methods for the Operations group.
type Operations struct{}

// CreateCreateRequest creates the Create request.
func (Operations) CreateCreateRequest(u url.URL, subscriptionId string, resourceGroupName string, bodyParameter CatalogDictionaryOfArray) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/Microsoft.Cache/Redis"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	u.Path = path.Join(u.Path, urlPath)
	query := u.Query()
	query.Set("apiVersion", "2014-04-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, u)
	err := req.MarshalAsJSON(bodyParameter)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CreateHandleResponse handles the Create response.
func (Operations) CreateHandleResponse(resp *azcore.Response) (*CreateResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := CreateResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CatalogDictionary)
}

// ListCreateRequest creates the List request.
func (Operations) ListCreateRequest(u url.URL, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/Microsoft.Cache/Redis"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape("123456"))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	u.Path = path.Join(u.Path, urlPath)
	query := u.Query()
	query.Set("apiVersion", "2014-04-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// ListHandleResponse handles the List response.
func (Operations) ListHandleResponse(resp *azcore.Response) (*ListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := ListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CatalogArray)
}

// UpdateCreateRequest creates the Update request.
func (Operations) UpdateCreateRequest(u url.URL, subscriptionId string, resourceGroupName string, bodyParameter CatalogArrayOfDictionary) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/Microsoft.Cache/Redis"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	u.Path = path.Join(u.Path, urlPath)
	query := u.Query()
	query.Set("apiVersion", "2014-04-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(bodyParameter)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateHandleResponse handles the Update response.
func (Operations) UpdateHandleResponse(resp *azcore.Response) (*UpdateResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := UpdateResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CatalogArray)
}