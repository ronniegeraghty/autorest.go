//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// InheritanceClient contains the methods for the Inheritance group.
// Don't use this type directly, use NewInheritanceClient() instead.
type InheritanceClient struct {
	pl runtime.Pipeline
}

// NewInheritanceClient creates a new instance of InheritanceClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewInheritanceClient(pl runtime.Pipeline) *InheritanceClient {
	client := &InheritanceClient{
		pl: pl,
	}
	return client
}

// GetValid - Get complex types that extend others
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// options - InheritanceClientGetValidOptions contains the optional parameters for the InheritanceClient.GetValid method.
func (client *InheritanceClient) GetValid(ctx context.Context, options *InheritanceClientGetValidOptions) (InheritanceClientGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return InheritanceClientGetValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InheritanceClientGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InheritanceClientGetValidResponse{}, runtime.NewResponseError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *InheritanceClient) getValidCreateRequest(ctx context.Context, options *InheritanceClientGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/inheritance/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *InheritanceClient) getValidHandleResponse(resp *http.Response) (InheritanceClientGetValidResponse, error) {
	result := InheritanceClientGetValidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Siamese); err != nil {
		return InheritanceClientGetValidResponse{}, err
	}
	return result, nil
}

// PutValid - Put complex types that extend others
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-02-29
// complexBody - Please put a siamese with id=2, name="Siameee", color=green, breed=persion, which hates 2 dogs, the 1st one
// named "Potato" with id=1 and food="tomato", and the 2nd one named "Tomato" with id=-1 and
// food="french fries".
// options - InheritanceClientPutValidOptions contains the optional parameters for the InheritanceClient.PutValid method.
func (client *InheritanceClient) PutValid(ctx context.Context, complexBody Siamese, options *InheritanceClientPutValidOptions) (InheritanceClientPutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return InheritanceClientPutValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InheritanceClientPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InheritanceClientPutValidResponse{}, runtime.NewResponseError(resp)
	}
	return InheritanceClientPutValidResponse{}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *InheritanceClient) putValidCreateRequest(ctx context.Context, complexBody Siamese, options *InheritanceClientPutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/inheritance/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}
