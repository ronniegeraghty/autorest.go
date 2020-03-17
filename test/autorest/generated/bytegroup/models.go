// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package bytegroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ByteGetEmptyResponse contains the response from method Byte.GetEmpty.
type ByteGetEmptyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The null byte value
	Value *[]byte
}

// ByteGetInvalidResponse contains the response from method Byte.GetInvalid.
type ByteGetInvalidResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The null byte value
	Value *[]byte
}

// ByteGetNonASCIIResponse contains the response from method Byte.GetNonASCII.
type ByteGetNonASCIIResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The null byte value
	Value *[]byte
}

// ByteGetNullResponse contains the response from method Byte.GetNull.
type ByteGetNullResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The null byte value
	Value *[]byte
}

// BytePutNonASCIIResponse contains the response from method Byte.PutNonASCII.
type BytePutNonASCIIResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

func newError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (e Error) Error() string {
	msg := ""
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if e.Status != nil {
		msg += fmt.Sprintf("Status: %v\n", *e.Status)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}
