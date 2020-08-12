package stringgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "tests/generated/stringgroup"

// Base64URL ...
type Base64URL struct {
	autorest.Response `json:"-"`
	// Value - a URL-encoded base64 string
	Value *string `json:"value,omitempty"`
}

// Error ...
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// RefColorConstant ...
type RefColorConstant struct {
	autorest.Response `json:"-"`
	// ColorConstant - Referenced Color Constant Description.
	ColorConstant *string `json:"ColorConstant,omitempty"`
	// Field1 - Sample string.
	Field1 *string `json:"field1,omitempty"`
}

// StringModel ...
type StringModel struct {
	autorest.Response `json:"-"`
	// Value - Possible values include: ''
	Value *string `json:"value,omitempty"`
}
