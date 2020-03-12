// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexmodelgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/complexmodelgroup/internal/complexmodelgroup"
)

// Operations contains the methods for the Operations group.
type Operations interface {
	// Create - Resets products.
	Create(ctx context.Context, subscriptionId string, resourceGroupName string, bodyParameter CatalogDictionaryOfArray) (*CreateResponse, error)
	// List - The Products endpoint returns information about the Uber products offered at a given location. The response includes the display name and other details about each product, and lists the products in the proper display order.
	List(ctx context.Context, resourceGroupName string) (*ListResponse, error)
	// Update - Resets products.
	Update(ctx context.Context, subscriptionId string, resourceGroupName string, bodyParameter CatalogArrayOfDictionary) (*UpdateResponse, error)
}

type operations struct {
	*Client
	azinternal.Operations
}

// Create - Resets products.
func (client *operations) Create(ctx context.Context, subscriptionId string, resourceGroupName string, bodyParameter CatalogDictionaryOfArray) (*CreateResponse, error) {
	req, err := client.CreateCreateRequest(*client.u, subscriptionId, resourceGroupName, bodyParameter)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// List - The Products endpoint returns information about the Uber products offered at a given location. The response includes the display name and other details about each product, and lists the products in the proper display order.
func (client *operations) List(ctx context.Context, resourceGroupName string) (*ListResponse, error) {
	req, err := client.ListCreateRequest(*client.u, resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update - Resets products.
func (client *operations) Update(ctx context.Context, subscriptionId string, resourceGroupName string, bodyParameter CatalogArrayOfDictionary) (*UpdateResponse, error) {
	req, err := client.UpdateCreateRequest(*client.u, subscriptionId, resourceGroupName, bodyParameter)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.UpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ Operations = (*operations)(nil)