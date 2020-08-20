// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func TestReadonlypropertyGetValid(t *testing.T) {
	client := complexgroup.NewDefaultClient(nil).ReadonlypropertyOperations()
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.ReadonlyObj, &complexgroup.ReadonlyObj{ID: to.StringPtr("1234"), Size: to.Int32Ptr(2)})
}

func TestReadonlypropertyPutValid(t *testing.T) {
	client := complexgroup.NewDefaultClient(nil).ReadonlypropertyOperations()
	id, size := "1234", int32(2)
	result, err := client.PutValid(context.Background(), complexgroup.ReadonlyObj{ID: &id, Size: &size})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
