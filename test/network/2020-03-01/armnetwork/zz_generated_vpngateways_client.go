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

// VPNGatewaysClient contains the methods for the VPNGateways group.
// Don't use this type directly, use NewVPNGatewaysClient() instead.
type VPNGatewaysClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVPNGatewaysClient creates a new instance of VPNGatewaysClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVPNGatewaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VPNGatewaysClient, error) {
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
	client := &VPNGatewaysClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a virtual wan vpn gateway if it doesn't exist else updates the existing gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// vpnGatewayParameters - Parameters supplied to create or Update a virtual wan vpn gateway.
// options - VPNGatewaysClientBeginCreateOrUpdateOptions contains the optional parameters for the VPNGatewaysClient.BeginCreateOrUpdate
// method.
func (client *VPNGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysClientBeginCreateOrUpdateOptions) (*runtime.Poller[VPNGatewaysClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VPNGatewaysClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VPNGatewaysClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a virtual wan vpn gateway if it doesn't exist else updates the existing gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *VPNGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VPNGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters VPNGateway, options *VPNGatewaysClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, vpnGatewayParameters)
}

// BeginDelete - Deletes a virtual wan vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// options - VPNGatewaysClientBeginDeleteOptions contains the optional parameters for the VPNGatewaysClient.BeginDelete method.
func (client *VPNGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginDeleteOptions) (*runtime.Poller[VPNGatewaysClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VPNGatewaysClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VPNGatewaysClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a virtual wan vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *VPNGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VPNGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a virtual wan vpn gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// options - VPNGatewaysClientGetOptions contains the optional parameters for the VPNGatewaysClient.Get method.
func (client *VPNGatewaysClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientGetOptions) (VPNGatewaysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return VPNGatewaysClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNGatewaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNGatewaysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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
func (client *VPNGatewaysClient) getHandleResponse(resp *http.Response) (VPNGatewaysClientGetResponse, error) {
	result := VPNGatewaysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNGateway); err != nil {
		return VPNGatewaysClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the VpnGateways in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// options - VPNGatewaysClientListOptions contains the optional parameters for the VPNGatewaysClient.List method.
func (client *VPNGatewaysClient) NewListPager(options *VPNGatewaysClientListOptions) *runtime.Pager[VPNGatewaysClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VPNGatewaysClientListResponse]{
		More: func(page VPNGatewaysClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VPNGatewaysClientListResponse) (VPNGatewaysClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VPNGatewaysClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VPNGatewaysClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VPNGatewaysClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VPNGatewaysClient) listCreateRequest(ctx context.Context, options *VPNGatewaysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnGateways"
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
func (client *VPNGatewaysClient) listHandleResponse(resp *http.Response) (VPNGatewaysClientListResponse, error) {
	result := VPNGatewaysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNGatewaysResult); err != nil {
		return VPNGatewaysClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the VpnGateways in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The resource group name of the VpnGateway.
// options - VPNGatewaysClientListByResourceGroupOptions contains the optional parameters for the VPNGatewaysClient.ListByResourceGroup
// method.
func (client *VPNGatewaysClient) NewListByResourceGroupPager(resourceGroupName string, options *VPNGatewaysClientListByResourceGroupOptions) *runtime.Pager[VPNGatewaysClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VPNGatewaysClientListByResourceGroupResponse]{
		More: func(page VPNGatewaysClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VPNGatewaysClientListByResourceGroupResponse) (VPNGatewaysClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VPNGatewaysClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VPNGatewaysClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VPNGatewaysClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VPNGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VPNGatewaysClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *VPNGatewaysClient) listByResourceGroupHandleResponse(resp *http.Response) (VPNGatewaysClientListByResourceGroupResponse, error) {
	result := VPNGatewaysClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNGatewaysResult); err != nil {
		return VPNGatewaysClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginReset - Resets the primary of the vpn gateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// options - VPNGatewaysClientBeginResetOptions contains the optional parameters for the VPNGatewaysClient.BeginReset method.
func (client *VPNGatewaysClient) BeginReset(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginResetOptions) (*runtime.Poller[VPNGatewaysClientResetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reset(ctx, resourceGroupName, gatewayName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VPNGatewaysClientResetResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VPNGatewaysClientResetResponse](options.ResumeToken, client.pl, nil)
	}
}

// Reset - Resets the primary of the vpn gateway in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *VPNGatewaysClient) reset(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginResetOptions) (*http.Response, error) {
	req, err := client.resetCreateRequest(ctx, resourceGroupName, gatewayName, options)
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

// resetCreateRequest creates the Reset request.
func (client *VPNGatewaysClient) resetCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNGatewaysClientBeginResetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/reset"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// UpdateTags - Updates virtual wan vpn gateway tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// vpnGatewayParameters - Parameters supplied to update a virtual wan vpn gateway tags.
// options - VPNGatewaysClientUpdateTagsOptions contains the optional parameters for the VPNGatewaysClient.UpdateTags method.
func (client *VPNGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VPNGatewaysClientUpdateTagsOptions) (VPNGatewaysClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, gatewayName, vpnGatewayParameters, options)
	if err != nil {
		return VPNGatewaysClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNGatewaysClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNGatewaysClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VPNGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnGatewayParameters TagsObject, options *VPNGatewaysClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, vpnGatewayParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VPNGatewaysClient) updateTagsHandleResponse(resp *http.Response) (VPNGatewaysClientUpdateTagsResponse, error) {
	result := VPNGatewaysClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNGateway); err != nil {
		return VPNGatewaysClientUpdateTagsResponse{}, err
	}
	return result, nil
}
