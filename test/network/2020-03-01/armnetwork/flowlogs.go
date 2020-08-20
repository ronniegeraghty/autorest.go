// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// FlowLogsOperations contains the methods for the FlowLogs group.
type FlowLogsOperations interface {
	// BeginCreateOrUpdate - Create or update a flow log for the specified network security group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog) (*FlowLogPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (FlowLogPoller, error)
	// BeginDelete - Deletes the specified flow log resource.
	BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets a flow log resource by name.
	Get(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string) (*FlowLogResponse, error)
	// List - Lists all flow log resources for the specified Network Watcher.
	List(resourceGroupName string, networkWatcherName string) (FlowLogListResultPager, error)
}

// flowLogsOperations implements the FlowLogsOperations interface.
type flowLogsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Create or update a flow log for the specified network security group.
func (client *flowLogsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog) (*FlowLogPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, networkWatcherName, flowLogName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("flowLogsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &flowLogPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*FlowLogResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *flowLogsOperations) ResumeCreateOrUpdate(token string) (FlowLogPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("flowLogsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &flowLogPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *flowLogsOperations) createOrUpdateCreateRequest(resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *flowLogsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*FlowLogPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &FlowLogPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *flowLogsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified flow log resource.
func (client *flowLogsOperations) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, networkWatcherName, flowLogName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("flowLogsOperations.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *flowLogsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("flowLogsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *flowLogsOperations) deleteCreateRequest(resourceGroupName string, networkWatcherName string, flowLogName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *flowLogsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *flowLogsOperations) deleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets a flow log resource by name.
func (client *flowLogsOperations) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string) (*FlowLogResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, networkWatcherName, flowLogName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *flowLogsOperations) getCreateRequest(resourceGroupName string, networkWatcherName string, flowLogName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *flowLogsOperations) getHandleResponse(resp *azcore.Response) (*FlowLogResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := FlowLogResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FlowLog)
}

// getHandleError handles the Get error response.
func (client *flowLogsOperations) getHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all flow log resources for the specified Network Watcher.
func (client *flowLogsOperations) List(resourceGroupName string, networkWatcherName string) (FlowLogListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, networkWatcherName)
	if err != nil {
		return nil, err
	}
	return &flowLogListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *FlowLogListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.FlowLogListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.FlowLogListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *flowLogsOperations) listCreateRequest(resourceGroupName string, networkWatcherName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *flowLogsOperations) listHandleResponse(resp *azcore.Response) (*FlowLogListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := FlowLogListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FlowLogListResult)
}

// listHandleError handles the List error response.
func (client *flowLogsOperations) listHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
