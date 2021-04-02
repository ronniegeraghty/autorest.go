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

// InboundNatRulesClient contains the methods for the InboundNatRules group.
// Don't use this type directly, use NewInboundNatRulesClient() instead.
type InboundNatRulesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewInboundNatRulesClient creates a new instance of InboundNatRulesClient with the specified values.
func NewInboundNatRulesClient(con *armcore.Connection, subscriptionID string) *InboundNatRulesClient {
	return &InboundNatRulesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a load balancer inbound nat rule.
func (client *InboundNatRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, options *InboundNatRulesBeginCreateOrUpdateOptions) (InboundNatRulePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, inboundNatRuleParameters, options)
	if err != nil {
		return InboundNatRulePollerResponse{}, err
	}
	result := InboundNatRulePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("InboundNatRulesClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return InboundNatRulePollerResponse{}, err
	}
	poller := &inboundNatRulePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (InboundNatRuleResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new InboundNatRulePoller from the specified resume token.
// token - The value must come from a previous call to InboundNatRulePoller.ResumeToken().
func (client *InboundNatRulesClient) ResumeCreateOrUpdate(token string) (InboundNatRulePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("InboundNatRulesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &inboundNatRulePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a load balancer inbound nat rule.
func (client *InboundNatRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, options *InboundNatRulesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, inboundNatRuleParameters, options)
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
func (client *InboundNatRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, options *InboundNatRulesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if inboundNatRuleName == "" {
		return nil, errors.New("parameter inboundNatRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inboundNatRuleName}", url.PathEscape(inboundNatRuleName))
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
	return req, req.MarshalAsJSON(inboundNatRuleParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *InboundNatRulesClient) createOrUpdateHandleResponse(resp *azcore.Response) (InboundNatRuleResponse, error) {
	var val *InboundNatRule
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return InboundNatRuleResponse{}, err
	}
	return InboundNatRuleResponse{RawResponse: resp.Response, InboundNatRule: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *InboundNatRulesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified load balancer inbound nat rule.
func (client *InboundNatRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("InboundNatRulesClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *InboundNatRulesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("InboundNatRulesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified load balancer inbound nat rule.
func (client *InboundNatRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, options)
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
func (client *InboundNatRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if inboundNatRuleName == "" {
		return nil, errors.New("parameter inboundNatRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inboundNatRuleName}", url.PathEscape(inboundNatRuleName))
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
func (client *InboundNatRulesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified load balancer inbound nat rule.
func (client *InboundNatRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesGetOptions) (InboundNatRuleResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, options)
	if err != nil {
		return InboundNatRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return InboundNatRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return InboundNatRuleResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InboundNatRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if inboundNatRuleName == "" {
		return nil, errors.New("parameter inboundNatRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inboundNatRuleName}", url.PathEscape(inboundNatRuleName))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InboundNatRulesClient) getHandleResponse(resp *azcore.Response) (InboundNatRuleResponse, error) {
	var val *InboundNatRule
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return InboundNatRuleResponse{}, err
	}
	return InboundNatRuleResponse{RawResponse: resp.Response, InboundNatRule: val}, nil
}

// getHandleError handles the Get error response.
func (client *InboundNatRulesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the inbound nat rules in a load balancer.
func (client *InboundNatRulesClient) List(resourceGroupName string, loadBalancerName string, options *InboundNatRulesListOptions) InboundNatRuleListResultPager {
	return &inboundNatRuleListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp InboundNatRuleListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.InboundNatRuleListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *InboundNatRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *InboundNatRulesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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

// listHandleResponse handles the List response.
func (client *InboundNatRulesClient) listHandleResponse(resp *azcore.Response) (InboundNatRuleListResultResponse, error) {
	var val *InboundNatRuleListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return InboundNatRuleListResultResponse{}, err
	}
	return InboundNatRuleListResultResponse{RawResponse: resp.Response, InboundNatRuleListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *InboundNatRulesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
