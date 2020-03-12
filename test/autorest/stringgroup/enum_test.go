// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgrouptest

import (
	"context"
	"generatortests/autorest/generated/stringgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getEnumClient(t *testing.T) stringgroup.EnumOperations {
	client, err := stringgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create enum client: %v", err)
	}
	return client.EnumOperations()
}

func TestEnumGetNotExpandable(t *testing.T) {
	client := getEnumClient(t)
	result, err := client.GetNotExpandable(context.Background())
	if err != nil {
		t.Fatalf("GetNotExpandable: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, stringgroup.ColorsRedColor.ToPtr())
}

func TestEnumGetReferenced(t *testing.T) {
	client := getEnumClient(t)
	result, err := client.GetReferenced(context.Background())
	if err != nil {
		t.Fatalf("GetReferenced: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, stringgroup.ColorsRedColor.ToPtr())
}

func TestEnumGetReferencedConstant(t *testing.T) {
	client := getEnumClient(t)
	result, err := client.GetReferencedConstant(context.Background())
	if err != nil {
		t.Fatalf("GetReferencedConstant: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val := "Sample String"
	helpers.DeepEqualOrFatal(t, result.RefColorConstant, &stringgroup.RefColorConstant{Field1: &val})
}

func TestEnumPutNotExpandable(t *testing.T) {
	client := getEnumClient(t)
	result, err := client.PutNotExpandable(context.Background(), stringgroup.ColorsRedColor)
	if err != nil {
		t.Fatalf("PutNotExpandable: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestEnumPutReferenced(t *testing.T) {
	client := getEnumClient(t)
	result, err := client.PutReferenced(context.Background(), stringgroup.ColorsRedColor)
	if err != nil {
		t.Fatalf("PutReferenced: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestEnumPutReferencedConstant(t *testing.T) {
	client := getEnumClient(t)
	val := string(stringgroup.ColorsGreenColor)
	result, err := client.PutReferencedConstant(context.Background(), stringgroup.RefColorConstant{ColorConstant: &val})
	if err != nil {
		t.Fatalf("PutReferencedConstant: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)

}