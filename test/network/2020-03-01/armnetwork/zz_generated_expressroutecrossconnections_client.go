//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRouteCrossConnectionsClient contains the methods for the ExpressRouteCrossConnections group.
// Don't use this type directly, use NewExpressRouteCrossConnectionsClient() instead.
type ExpressRouteCrossConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteCrossConnectionsClient creates a new instance of ExpressRouteCrossConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteCrossConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRouteCrossConnectionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ExpressRouteCrossConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Update the specified ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// parameters - Parameters supplied to the update express route crossConnection operation.
// options - ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.BeginCreateOrUpdate
// method.
func (client *ExpressRouteCrossConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ExpressRouteCrossConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, crossConnectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ExpressRouteCrossConnectionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExpressRouteCrossConnectionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Update the specified ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *ExpressRouteCrossConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteCrossConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Gets details about the specified ExpressRouteCrossConnection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group (peering location of the circuit).
// crossConnectionName - The name of the ExpressRouteCrossConnection (service key of the circuit).
// options - ExpressRouteCrossConnectionsClientGetOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.Get
// method.
func (client *ExpressRouteCrossConnectionsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsClientGetOptions) (ExpressRouteCrossConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
	if err != nil {
		return ExpressRouteCrossConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCrossConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCrossConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRouteCrossConnectionsClient) getHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientGetResponse, error) {
	result := ExpressRouteCrossConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnection); err != nil {
		return ExpressRouteCrossConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves all the ExpressRouteCrossConnections in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// options - ExpressRouteCrossConnectionsClientListOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.List
// method.
func (client *ExpressRouteCrossConnectionsClient) NewListPager(options *ExpressRouteCrossConnectionsClientListOptions) *runtime.Pager[ExpressRouteCrossConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExpressRouteCrossConnectionsClientListResponse]{
		More: func(page ExpressRouteCrossConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRouteCrossConnectionsClientListResponse) (ExpressRouteCrossConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExpressRouteCrossConnectionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExpressRouteCrossConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExpressRouteCrossConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCrossConnectionsClient) listCreateRequest(ctx context.Context, options *ExpressRouteCrossConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteCrossConnectionsClient) listHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientListResponse, error) {
	result := ExpressRouteCrossConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnectionListResult); err != nil {
		return ExpressRouteCrossConnectionsClientListResponse{}, err
	}
	return result, nil
}

// BeginListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// devicePath - The path of the device.
// options - ExpressRouteCrossConnectionsClientBeginListArpTableOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.BeginListArpTable
// method.
func (client *ExpressRouteCrossConnectionsClient) BeginListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListArpTableOptions) (*runtime.Poller[ExpressRouteCrossConnectionsClientListArpTableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listArpTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ExpressRouteCrossConnectionsClientListArpTableResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExpressRouteCrossConnectionsClientListArpTableResponse](options.ResumeToken, client.pl, nil)
	}
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *ExpressRouteCrossConnectionsClient) listArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListArpTableOptions) (*http.Response, error) {
	req, err := client.listArpTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// listArpTableCreateRequest creates the ListArpTable request.
func (client *ExpressRouteCrossConnectionsClient) listArpTableCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListArpTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/arpTables/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByResourceGroupPager - Retrieves all the ExpressRouteCrossConnections in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// options - ExpressRouteCrossConnectionsClientListByResourceGroupOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.ListByResourceGroup
// method.
func (client *ExpressRouteCrossConnectionsClient) NewListByResourceGroupPager(resourceGroupName string, options *ExpressRouteCrossConnectionsClientListByResourceGroupOptions) *runtime.Pager[ExpressRouteCrossConnectionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExpressRouteCrossConnectionsClientListByResourceGroupResponse]{
		More: func(page ExpressRouteCrossConnectionsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRouteCrossConnectionsClientListByResourceGroupResponse) (ExpressRouteCrossConnectionsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExpressRouteCrossConnectionsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExpressRouteCrossConnectionsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExpressRouteCrossConnectionsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExpressRouteCrossConnectionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRouteCrossConnectionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ExpressRouteCrossConnectionsClient) listByResourceGroupHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientListByResourceGroupResponse, error) {
	result := ExpressRouteCrossConnectionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnectionListResult); err != nil {
		return ExpressRouteCrossConnectionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in
// a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// devicePath - The path of the device.
// options - ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.BeginListRoutesTable
// method.
func (client *ExpressRouteCrossConnectionsClient) BeginListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions) (*runtime.Poller[ExpressRouteCrossConnectionsClientListRoutesTableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listRoutesTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ExpressRouteCrossConnectionsClientListRoutesTableResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExpressRouteCrossConnectionsClientListRoutesTableResponse](options.ResumeToken, client.pl, nil)
	}
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *ExpressRouteCrossConnectionsClient) listRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions) (*http.Response, error) {
	req, err := client.listRoutesTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// listRoutesTableCreateRequest creates the ListRoutesTable request.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTables/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the ExpressRouteCrossConnection.
// peeringName - The name of the peering.
// devicePath - The path of the device.
// options - ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions contains the optional parameters for the
// ExpressRouteCrossConnectionsClient.BeginListRoutesTableSummary method.
func (client *ExpressRouteCrossConnectionsClient) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions) (*runtime.Poller[ExpressRouteCrossConnectionsClientListRoutesTableSummaryResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listRoutesTableSummary(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ExpressRouteCrossConnectionsClientListRoutesTableSummaryResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExpressRouteCrossConnectionsClientListRoutesTableSummaryResponse](options.ResumeToken, client.pl, nil)
	}
}

// ListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions) (*http.Response, error) {
	req, err := client.listRoutesTableSummaryCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// listRoutesTableSummaryCreateRequest creates the ListRoutesTableSummary request.
func (client *ExpressRouteCrossConnectionsClient) listRoutesTableSummaryCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if devicePath == "" {
		return nil, errors.New("parameter devicePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// UpdateTags - Updates an express route cross connection tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// crossConnectionName - The name of the cross connection.
// crossConnectionParameters - Parameters supplied to update express route cross connection tags.
// options - ExpressRouteCrossConnectionsClientUpdateTagsOptions contains the optional parameters for the ExpressRouteCrossConnectionsClient.UpdateTags
// method.
func (client *ExpressRouteCrossConnectionsClient) UpdateTags(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsClientUpdateTagsOptions) (ExpressRouteCrossConnectionsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, crossConnectionName, crossConnectionParameters, options)
	if err != nil {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ExpressRouteCrossConnectionsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if crossConnectionName == "" {
		return nil, errors.New("parameter crossConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, crossConnectionParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ExpressRouteCrossConnectionsClient) updateTagsHandleResponse(resp *http.Response) (ExpressRouteCrossConnectionsClientUpdateTagsResponse, error) {
	result := ExpressRouteCrossConnectionsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCrossConnection); err != nil {
		return ExpressRouteCrossConnectionsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
