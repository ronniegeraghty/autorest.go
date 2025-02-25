//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package objectgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ObjectTypeClient contains the methods for the ObjectTypeClient group.
// Don't use this type directly, use NewObjectTypeClient() instead.
type ObjectTypeClient struct {
	pl runtime.Pipeline
}

// NewObjectTypeClient creates a new instance of ObjectTypeClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewObjectTypeClient(pl runtime.Pipeline) *ObjectTypeClient {
	client := &ObjectTypeClient{
		pl: pl,
	}
	return client
}

// Get - Basic get that returns an object. Returns object { 'message': 'An object was successfully returned' }
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - ObjectTypeClientGetOptions contains the optional parameters for the ObjectTypeClient.Get method.
func (client *ObjectTypeClient) Get(ctx context.Context, options *ObjectTypeClientGetOptions) (ObjectTypeClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return ObjectTypeClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectTypeClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ObjectTypeClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ObjectTypeClient) getCreateRequest(ctx context.Context, options *ObjectTypeClientGetOptions) (*policy.Request, error) {
	urlPath := "/objectType/get"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ObjectTypeClient) getHandleResponse(resp *http.Response) (ObjectTypeClientGetResponse, error) {
	result := ObjectTypeClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Interface); err != nil {
		return ObjectTypeClientGetResponse{}, err
	}
	return result, nil
}

// Put - Basic put that puts an object. Pass in {'foo': 'bar'} to get a 200 and anything else to get an object error.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// putObject - Pass in {'foo': 'bar'} for a 200, anything else for an object error
// options - ObjectTypeClientPutOptions contains the optional parameters for the ObjectTypeClient.Put method.
func (client *ObjectTypeClient) Put(ctx context.Context, putObject interface{}, options *ObjectTypeClientPutOptions) (ObjectTypeClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, putObject, options)
	if err != nil {
		return ObjectTypeClientPutResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectTypeClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ObjectTypeClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return ObjectTypeClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *ObjectTypeClient) putCreateRequest(ctx context.Context, putObject interface{}, options *ObjectTypeClientPutOptions) (*policy.Request, error) {
	urlPath := "/objectType/put"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, putObject)
}
