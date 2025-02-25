//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// SkipURLEncodingClient contains the methods for the SkipURLEncoding group.
// Don't use this type directly, use NewSkipURLEncodingClient() instead.
type SkipURLEncodingClient struct {
	pl runtime.Pipeline
}

// NewSkipURLEncodingClient creates a new instance of SkipURLEncodingClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewSkipURLEncodingClient(pl runtime.Pipeline) *SkipURLEncodingClient {
	client := &SkipURLEncodingClient{
		pl: pl,
	}
	return client
}

// GetMethodPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// unencodedPathParam - Unencoded path parameter with value 'path1/path2/path3'
// options - SkipURLEncodingClientGetMethodPathValidOptions contains the optional parameters for the SkipURLEncodingClient.GetMethodPathValid
// method.
func (client *SkipURLEncodingClient) GetMethodPathValid(ctx context.Context, unencodedPathParam string, options *SkipURLEncodingClientGetMethodPathValidOptions) (SkipURLEncodingClientGetMethodPathValidResponse, error) {
	req, err := client.getMethodPathValidCreateRequest(ctx, unencodedPathParam, options)
	if err != nil {
		return SkipURLEncodingClientGetMethodPathValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SkipURLEncodingClientGetMethodPathValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SkipURLEncodingClientGetMethodPathValidResponse{}, runtime.NewResponseError(resp)
	}
	return SkipURLEncodingClientGetMethodPathValidResponse{}, nil
}

// getMethodPathValidCreateRequest creates the GetMethodPathValid request.
func (client *SkipURLEncodingClient) getMethodPathValidCreateRequest(ctx context.Context, unencodedPathParam string, options *SkipURLEncodingClientGetMethodPathValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/method/path/valid/{unencodedPathParam}"
	urlPath = strings.ReplaceAll(urlPath, "{unencodedPathParam}", unencodedPathParam)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetMethodQueryNull - Get method with unencoded query parameter with value null
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// options - SkipURLEncodingClientGetMethodQueryNullOptions contains the optional parameters for the SkipURLEncodingClient.GetMethodQueryNull
// method.
func (client *SkipURLEncodingClient) GetMethodQueryNull(ctx context.Context, options *SkipURLEncodingClientGetMethodQueryNullOptions) (SkipURLEncodingClientGetMethodQueryNullResponse, error) {
	req, err := client.getMethodQueryNullCreateRequest(ctx, options)
	if err != nil {
		return SkipURLEncodingClientGetMethodQueryNullResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SkipURLEncodingClientGetMethodQueryNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SkipURLEncodingClientGetMethodQueryNullResponse{}, runtime.NewResponseError(resp)
	}
	return SkipURLEncodingClientGetMethodQueryNullResponse{}, nil
}

// getMethodQueryNullCreateRequest creates the GetMethodQueryNull request.
func (client *SkipURLEncodingClient) getMethodQueryNullCreateRequest(ctx context.Context, options *SkipURLEncodingClientGetMethodQueryNullOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/method/query/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	unencodedParams := []string{}
	if options != nil && options.Q1 != nil {
		unencodedParams = append(unencodedParams, "q1="+*options.Q1)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetMethodQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// q1 - Unencoded query parameter with value 'value1&q2=value2&q3=value3'
// options - SkipURLEncodingClientGetMethodQueryValidOptions contains the optional parameters for the SkipURLEncodingClient.GetMethodQueryValid
// method.
func (client *SkipURLEncodingClient) GetMethodQueryValid(ctx context.Context, q1 string, options *SkipURLEncodingClientGetMethodQueryValidOptions) (SkipURLEncodingClientGetMethodQueryValidResponse, error) {
	req, err := client.getMethodQueryValidCreateRequest(ctx, q1, options)
	if err != nil {
		return SkipURLEncodingClientGetMethodQueryValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SkipURLEncodingClientGetMethodQueryValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SkipURLEncodingClientGetMethodQueryValidResponse{}, runtime.NewResponseError(resp)
	}
	return SkipURLEncodingClientGetMethodQueryValidResponse{}, nil
}

// getMethodQueryValidCreateRequest creates the GetMethodQueryValid request.
func (client *SkipURLEncodingClient) getMethodQueryValidCreateRequest(ctx context.Context, q1 string, options *SkipURLEncodingClientGetMethodQueryValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/method/query/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	unencodedParams := []string{}
	unencodedParams = append(unencodedParams, "q1="+q1)
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetPathQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// q1 - Unencoded query parameter with value 'value1&q2=value2&q3=value3'
// options - SkipURLEncodingClientGetPathQueryValidOptions contains the optional parameters for the SkipURLEncodingClient.GetPathQueryValid
// method.
func (client *SkipURLEncodingClient) GetPathQueryValid(ctx context.Context, q1 string, options *SkipURLEncodingClientGetPathQueryValidOptions) (SkipURLEncodingClientGetPathQueryValidResponse, error) {
	req, err := client.getPathQueryValidCreateRequest(ctx, q1, options)
	if err != nil {
		return SkipURLEncodingClientGetPathQueryValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SkipURLEncodingClientGetPathQueryValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SkipURLEncodingClientGetPathQueryValidResponse{}, runtime.NewResponseError(resp)
	}
	return SkipURLEncodingClientGetPathQueryValidResponse{}, nil
}

// getPathQueryValidCreateRequest creates the GetPathQueryValid request.
func (client *SkipURLEncodingClient) getPathQueryValidCreateRequest(ctx context.Context, q1 string, options *SkipURLEncodingClientGetPathQueryValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/path/query/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	unencodedParams := []string{}
	unencodedParams = append(unencodedParams, "q1="+q1)
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// unencodedPathParam - Unencoded path parameter with value 'path1/path2/path3'
// options - SkipURLEncodingClientGetPathValidOptions contains the optional parameters for the SkipURLEncodingClient.GetPathValid
// method.
func (client *SkipURLEncodingClient) GetPathValid(ctx context.Context, unencodedPathParam string, options *SkipURLEncodingClientGetPathValidOptions) (SkipURLEncodingClientGetPathValidResponse, error) {
	req, err := client.getPathValidCreateRequest(ctx, unencodedPathParam, options)
	if err != nil {
		return SkipURLEncodingClientGetPathValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SkipURLEncodingClientGetPathValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SkipURLEncodingClientGetPathValidResponse{}, runtime.NewResponseError(resp)
	}
	return SkipURLEncodingClientGetPathValidResponse{}, nil
}

// getPathValidCreateRequest creates the GetPathValid request.
func (client *SkipURLEncodingClient) getPathValidCreateRequest(ctx context.Context, unencodedPathParam string, options *SkipURLEncodingClientGetPathValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/path/path/valid/{unencodedPathParam}"
	urlPath = strings.ReplaceAll(urlPath, "{unencodedPathParam}", unencodedPathParam)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetSwaggerPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// options - SkipURLEncodingClientGetSwaggerPathValidOptions contains the optional parameters for the SkipURLEncodingClient.GetSwaggerPathValid
// method.
func (client *SkipURLEncodingClient) GetSwaggerPathValid(ctx context.Context, options *SkipURLEncodingClientGetSwaggerPathValidOptions) (SkipURLEncodingClientGetSwaggerPathValidResponse, error) {
	req, err := client.getSwaggerPathValidCreateRequest(ctx, options)
	if err != nil {
		return SkipURLEncodingClientGetSwaggerPathValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SkipURLEncodingClientGetSwaggerPathValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SkipURLEncodingClientGetSwaggerPathValidResponse{}, runtime.NewResponseError(resp)
	}
	return SkipURLEncodingClientGetSwaggerPathValidResponse{}, nil
}

// getSwaggerPathValidCreateRequest creates the GetSwaggerPathValid request.
func (client *SkipURLEncodingClient) getSwaggerPathValidCreateRequest(ctx context.Context, options *SkipURLEncodingClientGetSwaggerPathValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/swagger/path/valid/{unencodedPathParam}"
	urlPath = strings.ReplaceAll(urlPath, "{unencodedPathParam}", "path1/path2/path3")
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetSwaggerQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01-preview
// options - SkipURLEncodingClientGetSwaggerQueryValidOptions contains the optional parameters for the SkipURLEncodingClient.GetSwaggerQueryValid
// method.
func (client *SkipURLEncodingClient) GetSwaggerQueryValid(ctx context.Context, options *SkipURLEncodingClientGetSwaggerQueryValidOptions) (SkipURLEncodingClientGetSwaggerQueryValidResponse, error) {
	req, err := client.getSwaggerQueryValidCreateRequest(ctx, options)
	if err != nil {
		return SkipURLEncodingClientGetSwaggerQueryValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SkipURLEncodingClientGetSwaggerQueryValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SkipURLEncodingClientGetSwaggerQueryValidResponse{}, runtime.NewResponseError(resp)
	}
	return SkipURLEncodingClientGetSwaggerQueryValidResponse{}, nil
}

// getSwaggerQueryValidCreateRequest creates the GetSwaggerQueryValid request.
func (client *SkipURLEncodingClient) getSwaggerQueryValidCreateRequest(ctx context.Context, options *SkipURLEncodingClientGetSwaggerQueryValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/swagger/query/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	unencodedParams := []string{}
	unencodedParams = append(unencodedParams, "q1="+"value1&q2=value2&q3=value3")
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
