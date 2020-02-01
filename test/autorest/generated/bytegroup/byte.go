// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package bytegroup

import (
	"context"
	azinternal "generatortests/autorest/generated/bytegroup/internal/bytegroup"
)

// ByteOperations contains the methods for the Byte group.
type ByteOperations interface {
	// GetEmpty - Get empty byte value '' 
	GetEmpty(ctx context.Context) (*ByteGetEmptyResponse, error)
	// GetInvalid - Get invalid byte value ':::SWAGGER::::' 
	GetInvalid(ctx context.Context) (*ByteGetInvalidResponse, error)
	// GetNonASCII - Get non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6) 
	GetNonASCII(ctx context.Context) (*ByteGetNonASCIIResponse, error)
	// GetNull - Get null byte value 
	GetNull(ctx context.Context) (*ByteGetNullResponse, error)
	// PutNonASCII - Put non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6) 
	PutNonASCII(ctx context.Context, byteBody []byte) (*BytePutNonASCIIResponse, error)
}

type byteOperations struct {
	*Client
	azinternal.ByteOperations
}

	// GetEmpty - Get empty byte value '' 
func (client *byteOperations) GetEmpty(ctx context.Context) (*ByteGetEmptyResponse, error) {
	req, err := client.GetEmptyCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

	// GetInvalid - Get invalid byte value ':::SWAGGER::::' 
func (client *byteOperations) GetInvalid(ctx context.Context) (*ByteGetInvalidResponse, error) {
	req, err := client.GetInvalidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

	// GetNonASCII - Get non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6) 
func (client *byteOperations) GetNonASCII(ctx context.Context) (*ByteGetNonASCIIResponse, error) {
	req, err := client.GetNonASCIICreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetNonASCIIHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

	// GetNull - Get null byte value 
func (client *byteOperations) GetNull(ctx context.Context) (*ByteGetNullResponse, error) {
	req, err := client.GetNullCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

	// PutNonASCII - Put non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6) 
func (client *byteOperations) PutNonASCII(ctx context.Context, byteBody []byte) (*BytePutNonASCIIResponse, error) {
	req, err := client.PutNonASCIICreateRequest(*client.u, byteBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutNonASCIIHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ ByteOperations = (*byteOperations)(nil)
