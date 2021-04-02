// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ServiceEndpointPolicyDefinitionsClient contains the methods for the ServiceEndpointPolicyDefinitions group.
// Don't use this type directly, use NewServiceEndpointPolicyDefinitionsClient() instead.
type ServiceEndpointPolicyDefinitionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewServiceEndpointPolicyDefinitionsClient creates a new instance of ServiceEndpointPolicyDefinitionsClient with the specified values.
func NewServiceEndpointPolicyDefinitionsClient(con *armcore.Connection, subscriptionID string) *ServiceEndpointPolicyDefinitionsClient {
	return &ServiceEndpointPolicyDefinitionsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a service endpoint policy definition in the specified service endpoint policy.
func (client *ServiceEndpointPolicyDefinitionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition, options *ServiceEndpointPolicyDefinitionsBeginCreateOrUpdateOptions) (ServiceEndpointPolicyDefinitionPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, serviceEndpointPolicyDefinitions, options)
	if err != nil {
		return ServiceEndpointPolicyDefinitionPollerResponse{}, err
	}
	result := ServiceEndpointPolicyDefinitionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ServiceEndpointPolicyDefinitionsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return ServiceEndpointPolicyDefinitionPollerResponse{}, err
	}
	poller := &serviceEndpointPolicyDefinitionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServiceEndpointPolicyDefinitionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ServiceEndpointPolicyDefinitionPoller from the specified resume token.
// token - The value must come from a previous call to ServiceEndpointPolicyDefinitionPoller.ResumeToken().
func (client *ServiceEndpointPolicyDefinitionsClient) ResumeCreateOrUpdate(token string) (ServiceEndpointPolicyDefinitionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ServiceEndpointPolicyDefinitionsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &serviceEndpointPolicyDefinitionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a service endpoint policy definition in the specified service endpoint policy.
func (client *ServiceEndpointPolicyDefinitionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition, options *ServiceEndpointPolicyDefinitionsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, serviceEndpointPolicyDefinitions, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceEndpointPolicyDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition, options *ServiceEndpointPolicyDefinitionsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if serviceEndpointPolicyDefinitionName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyDefinitionName}", url.PathEscape(serviceEndpointPolicyDefinitionName))
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
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(serviceEndpointPolicyDefinitions)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServiceEndpointPolicyDefinitionsClient) createOrUpdateHandleResponse(resp *azcore.Response) (ServiceEndpointPolicyDefinitionResponse, error) {
	var val *ServiceEndpointPolicyDefinition
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ServiceEndpointPolicyDefinitionResponse{}, err
	}
	return ServiceEndpointPolicyDefinitionResponse{RawResponse: resp.Response, ServiceEndpointPolicyDefinition: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServiceEndpointPolicyDefinitionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified ServiceEndpoint policy definitions.
func (client *ServiceEndpointPolicyDefinitionsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, options *ServiceEndpointPolicyDefinitionsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ServiceEndpointPolicyDefinitionsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ServiceEndpointPolicyDefinitionsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ServiceEndpointPolicyDefinitionsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified ServiceEndpoint policy definitions.
func (client *ServiceEndpointPolicyDefinitionsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, options *ServiceEndpointPolicyDefinitionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceEndpointPolicyDefinitionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, options *ServiceEndpointPolicyDefinitionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if serviceEndpointPolicyDefinitionName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyDefinitionName}", url.PathEscape(serviceEndpointPolicyDefinitionName))
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
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ServiceEndpointPolicyDefinitionsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Get the specified service endpoint policy definitions from service endpoint policy.
func (client *ServiceEndpointPolicyDefinitionsClient) Get(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, options *ServiceEndpointPolicyDefinitionsGetOptions) (ServiceEndpointPolicyDefinitionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, options)
	if err != nil {
		return ServiceEndpointPolicyDefinitionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceEndpointPolicyDefinitionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceEndpointPolicyDefinitionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceEndpointPolicyDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, options *ServiceEndpointPolicyDefinitionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if serviceEndpointPolicyDefinitionName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyDefinitionName}", url.PathEscape(serviceEndpointPolicyDefinitionName))
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
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceEndpointPolicyDefinitionsClient) getHandleResponse(resp *azcore.Response) (ServiceEndpointPolicyDefinitionResponse, error) {
	var val *ServiceEndpointPolicyDefinition
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ServiceEndpointPolicyDefinitionResponse{}, err
	}
	return ServiceEndpointPolicyDefinitionResponse{RawResponse: resp.Response, ServiceEndpointPolicyDefinition: val}, nil
}

// getHandleError handles the Get error response.
func (client *ServiceEndpointPolicyDefinitionsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Gets all service endpoint policy definitions in a service end point policy.
func (client *ServiceEndpointPolicyDefinitionsClient) ListByResourceGroup(resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPolicyDefinitionsListByResourceGroupOptions) ServiceEndpointPolicyDefinitionListResultPager {
	return &serviceEndpointPolicyDefinitionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ServiceEndpointPolicyDefinitionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ServiceEndpointPolicyDefinitionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServiceEndpointPolicyDefinitionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPolicyDefinitionsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
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
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServiceEndpointPolicyDefinitionsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ServiceEndpointPolicyDefinitionListResultResponse, error) {
	var val *ServiceEndpointPolicyDefinitionListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ServiceEndpointPolicyDefinitionListResultResponse{}, err
	}
	return ServiceEndpointPolicyDefinitionListResultResponse{RawResponse: resp.Response, ServiceEndpointPolicyDefinitionListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ServiceEndpointPolicyDefinitionsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
