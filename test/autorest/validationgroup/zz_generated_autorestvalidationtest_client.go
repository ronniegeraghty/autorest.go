//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package validationgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// AutoRestValidationTestClient contains the methods for the AutoRestValidationTest group.
// Don't use this type directly, use NewAutoRestValidationTestClient() instead.
type AutoRestValidationTestClient struct {
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAutoRestValidationTestClient creates a new instance of AutoRestValidationTestClient with the specified values.
// subscriptionID - Subscription ID.
// pl - the pipeline used for sending requests and handling responses.
func NewAutoRestValidationTestClient(subscriptionID string, pl runtime.Pipeline) *AutoRestValidationTestClient {
	client := &AutoRestValidationTestClient{
		subscriptionID: subscriptionID,
		pl:             pl,
	}
	return client
}

// GetWithConstantInPath -
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - AutoRestValidationTestClientGetWithConstantInPathOptions contains the optional parameters for the AutoRestValidationTestClient.GetWithConstantInPath
// method.
func (client *AutoRestValidationTestClient) GetWithConstantInPath(ctx context.Context, options *AutoRestValidationTestClientGetWithConstantInPathOptions) (AutoRestValidationTestClientGetWithConstantInPathResponse, error) {
	req, err := client.getWithConstantInPathCreateRequest(ctx, options)
	if err != nil {
		return AutoRestValidationTestClientGetWithConstantInPathResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutoRestValidationTestClientGetWithConstantInPathResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutoRestValidationTestClientGetWithConstantInPathResponse{}, runtime.NewResponseError(resp)
	}
	return AutoRestValidationTestClientGetWithConstantInPathResponse{}, nil
}

// getWithConstantInPathCreateRequest creates the GetWithConstantInPath request.
func (client *AutoRestValidationTestClient) getWithConstantInPathCreateRequest(ctx context.Context, options *AutoRestValidationTestClientGetWithConstantInPathOptions) (*policy.Request, error) {
	urlPath := "/validation/constantsInPath/{constantParam}/value"
	urlPath = strings.ReplaceAll(urlPath, "{constantParam}", url.PathEscape("constant"))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PostWithConstantInBody -
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - AutoRestValidationTestClientPostWithConstantInBodyOptions contains the optional parameters for the AutoRestValidationTestClient.PostWithConstantInBody
// method.
func (client *AutoRestValidationTestClient) PostWithConstantInBody(ctx context.Context, options *AutoRestValidationTestClientPostWithConstantInBodyOptions) (AutoRestValidationTestClientPostWithConstantInBodyResponse, error) {
	req, err := client.postWithConstantInBodyCreateRequest(ctx, options)
	if err != nil {
		return AutoRestValidationTestClientPostWithConstantInBodyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutoRestValidationTestClientPostWithConstantInBodyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutoRestValidationTestClientPostWithConstantInBodyResponse{}, runtime.NewResponseError(resp)
	}
	return client.postWithConstantInBodyHandleResponse(resp)
}

// postWithConstantInBodyCreateRequest creates the PostWithConstantInBody request.
func (client *AutoRestValidationTestClient) postWithConstantInBodyCreateRequest(ctx context.Context, options *AutoRestValidationTestClientPostWithConstantInBodyOptions) (*policy.Request, error) {
	urlPath := "/validation/constantsInPath/{constantParam}/value"
	urlPath = strings.ReplaceAll(urlPath, "{constantParam}", url.PathEscape("constant"))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// postWithConstantInBodyHandleResponse handles the PostWithConstantInBody response.
func (client *AutoRestValidationTestClient) postWithConstantInBodyHandleResponse(resp *http.Response) (AutoRestValidationTestClientPostWithConstantInBodyResponse, error) {
	result := AutoRestValidationTestClientPostWithConstantInBodyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Product); err != nil {
		return AutoRestValidationTestClientPostWithConstantInBodyResponse{}, err
	}
	return result, nil
}

// ValidationOfBody - Validates body parameters on the method. See swagger for details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// resourceGroupName - Required string between 3 and 10 chars with pattern [a-zA-Z0-9]+.
// id - Required int multiple of 10 from 100 to 1000.
// options - AutoRestValidationTestClientValidationOfBodyOptions contains the optional parameters for the AutoRestValidationTestClient.ValidationOfBody
// method.
func (client *AutoRestValidationTestClient) ValidationOfBody(ctx context.Context, resourceGroupName string, id int32, body Product, options *AutoRestValidationTestClientValidationOfBodyOptions) (AutoRestValidationTestClientValidationOfBodyResponse, error) {
	req, err := client.validationOfBodyCreateRequest(ctx, resourceGroupName, id, body, options)
	if err != nil {
		return AutoRestValidationTestClientValidationOfBodyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutoRestValidationTestClientValidationOfBodyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutoRestValidationTestClientValidationOfBodyResponse{}, runtime.NewResponseError(resp)
	}
	return client.validationOfBodyHandleResponse(resp)
}

// validationOfBodyCreateRequest creates the ValidationOfBody request.
func (client *AutoRestValidationTestClient) validationOfBodyCreateRequest(ctx context.Context, resourceGroupName string, id int32, body Product, options *AutoRestValidationTestClientValidationOfBodyOptions) (*policy.Request, error) {
	urlPath := "/fakepath/{subscriptionId}/{resourceGroupName}/{id}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("apiVersion", "1.0.0")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// validationOfBodyHandleResponse handles the ValidationOfBody response.
func (client *AutoRestValidationTestClient) validationOfBodyHandleResponse(resp *http.Response) (AutoRestValidationTestClientValidationOfBodyResponse, error) {
	result := AutoRestValidationTestClientValidationOfBodyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Product); err != nil {
		return AutoRestValidationTestClientValidationOfBodyResponse{}, err
	}
	return result, nil
}

// ValidationOfMethodParameters - Validates input parameters on the method. See swagger for details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// resourceGroupName - Required string between 3 and 10 chars with pattern [a-zA-Z0-9]+.
// id - Required int multiple of 10 from 100 to 1000.
// options - AutoRestValidationTestClientValidationOfMethodParametersOptions contains the optional parameters for the AutoRestValidationTestClient.ValidationOfMethodParameters
// method.
func (client *AutoRestValidationTestClient) ValidationOfMethodParameters(ctx context.Context, resourceGroupName string, id int32, options *AutoRestValidationTestClientValidationOfMethodParametersOptions) (AutoRestValidationTestClientValidationOfMethodParametersResponse, error) {
	req, err := client.validationOfMethodParametersCreateRequest(ctx, resourceGroupName, id, options)
	if err != nil {
		return AutoRestValidationTestClientValidationOfMethodParametersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutoRestValidationTestClientValidationOfMethodParametersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutoRestValidationTestClientValidationOfMethodParametersResponse{}, runtime.NewResponseError(resp)
	}
	return client.validationOfMethodParametersHandleResponse(resp)
}

// validationOfMethodParametersCreateRequest creates the ValidationOfMethodParameters request.
func (client *AutoRestValidationTestClient) validationOfMethodParametersCreateRequest(ctx context.Context, resourceGroupName string, id int32, options *AutoRestValidationTestClientValidationOfMethodParametersOptions) (*policy.Request, error) {
	urlPath := "/fakepath/{subscriptionId}/{resourceGroupName}/{id}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(strconv.FormatInt(int64(id), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("apiVersion", "1.0.0")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// validationOfMethodParametersHandleResponse handles the ValidationOfMethodParameters response.
func (client *AutoRestValidationTestClient) validationOfMethodParametersHandleResponse(resp *http.Response) (AutoRestValidationTestClientValidationOfMethodParametersResponse, error) {
	result := AutoRestValidationTestClientValidationOfMethodParametersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Product); err != nil {
		return AutoRestValidationTestClientValidationOfMethodParametersResponse{}, err
	}
	return result, nil
}
