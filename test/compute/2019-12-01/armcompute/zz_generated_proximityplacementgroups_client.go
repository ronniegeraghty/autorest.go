// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ProximityPlacementGroupsClient contains the methods for the ProximityPlacementGroups group.
// Don't use this type directly, use NewProximityPlacementGroupsClient() instead.
type ProximityPlacementGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewProximityPlacementGroupsClient creates a new instance of ProximityPlacementGroupsClient with the specified values.
func NewProximityPlacementGroupsClient(con *armcore.Connection, subscriptionID string) *ProximityPlacementGroupsClient {
	return &ProximityPlacementGroupsClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create or update a proximity placement group.
func (client *ProximityPlacementGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup, options *ProximityPlacementGroupsCreateOrUpdateOptions) (ProximityPlacementGroupResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, parameters, options)
	if err != nil {
		return ProximityPlacementGroupResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ProximityPlacementGroupResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return ProximityPlacementGroupResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProximityPlacementGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup, options *ProximityPlacementGroupsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProximityPlacementGroupsClient) createOrUpdateHandleResponse(resp *azcore.Response) (ProximityPlacementGroupResponse, error) {
	var val *ProximityPlacementGroup
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ProximityPlacementGroupResponse{}, err
	}
	return ProximityPlacementGroupResponse{RawResponse: resp.Response, ProximityPlacementGroup: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ProximityPlacementGroupsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Delete - Delete a proximity placement group.
func (client *ProximityPlacementGroupsClient) Delete(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProximityPlacementGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ProximityPlacementGroupsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Retrieves information about a proximity placement group .
func (client *ProximityPlacementGroupsClient) Get(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsGetOptions) (ProximityPlacementGroupResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, options)
	if err != nil {
		return ProximityPlacementGroupResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ProximityPlacementGroupResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ProximityPlacementGroupResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProximityPlacementGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.IncludeColocationStatus != nil {
		reqQP.Set("includeColocationStatus", *options.IncludeColocationStatus)
	}
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProximityPlacementGroupsClient) getHandleResponse(resp *azcore.Response) (ProximityPlacementGroupResponse, error) {
	var val *ProximityPlacementGroup
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ProximityPlacementGroupResponse{}, err
	}
	return ProximityPlacementGroupResponse{RawResponse: resp.Response, ProximityPlacementGroup: val}, nil
}

// getHandleError handles the Get error response.
func (client *ProximityPlacementGroupsClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByResourceGroup - Lists all proximity placement groups in a resource group.
func (client *ProximityPlacementGroupsClient) ListByResourceGroup(resourceGroupName string, options *ProximityPlacementGroupsListByResourceGroupOptions) ProximityPlacementGroupListResultPager {
	return &proximityPlacementGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ProximityPlacementGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ProximityPlacementGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ProximityPlacementGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ProximityPlacementGroupsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ProximityPlacementGroupsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ProximityPlacementGroupListResultResponse, error) {
	var val *ProximityPlacementGroupListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ProximityPlacementGroupListResultResponse{}, err
	}
	return ProximityPlacementGroupListResultResponse{RawResponse: resp.Response, ProximityPlacementGroupListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ProximityPlacementGroupsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListBySubscription - Lists all proximity placement groups in a subscription.
func (client *ProximityPlacementGroupsClient) ListBySubscription(options *ProximityPlacementGroupsListBySubscriptionOptions) ProximityPlacementGroupListResultPager {
	return &proximityPlacementGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		responder: client.listBySubscriptionHandleResponse,
		errorer:   client.listBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp ProximityPlacementGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ProximityPlacementGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ProximityPlacementGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ProximityPlacementGroupsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/proximityPlacementGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ProximityPlacementGroupsClient) listBySubscriptionHandleResponse(resp *azcore.Response) (ProximityPlacementGroupListResultResponse, error) {
	var val *ProximityPlacementGroupListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ProximityPlacementGroupListResultResponse{}, err
	}
	return ProximityPlacementGroupListResultResponse{RawResponse: resp.Response, ProximityPlacementGroupListResult: val}, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ProximityPlacementGroupsClient) listBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Update - Update a proximity placement group.
func (client *ProximityPlacementGroupsClient) Update(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroupUpdate, options *ProximityPlacementGroupsUpdateOptions) (ProximityPlacementGroupResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, parameters, options)
	if err != nil {
		return ProximityPlacementGroupResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ProximityPlacementGroupResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ProximityPlacementGroupResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ProximityPlacementGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroupUpdate, options *ProximityPlacementGroupsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *ProximityPlacementGroupsClient) updateHandleResponse(resp *azcore.Response) (ProximityPlacementGroupResponse, error) {
	var val *ProximityPlacementGroup
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ProximityPlacementGroupResponse{}, err
	}
	return ProximityPlacementGroupResponse{RawResponse: resp.Response, ProximityPlacementGroup: val}, nil
}

// updateHandleError handles the Update error response.
func (client *ProximityPlacementGroupsClient) updateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
