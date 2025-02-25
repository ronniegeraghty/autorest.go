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

// ResourceNavigationLinksClient contains the methods for the ResourceNavigationLinks group.
// Don't use this type directly, use NewResourceNavigationLinksClient() instead.
type ResourceNavigationLinksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceNavigationLinksClient creates a new instance of ResourceNavigationLinksClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceNavigationLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceNavigationLinksClient, error) {
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
	client := &ResourceNavigationLinksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Gets a list of resource navigation links for a subnet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// virtualNetworkName - The name of the virtual network.
// subnetName - The name of the subnet.
// options - ResourceNavigationLinksClientListOptions contains the optional parameters for the ResourceNavigationLinksClient.List
// method.
func (client *ResourceNavigationLinksClient) List(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *ResourceNavigationLinksClientListOptions) (ResourceNavigationLinksClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
	if err != nil {
		return ResourceNavigationLinksClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceNavigationLinksClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceNavigationLinksClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ResourceNavigationLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *ResourceNavigationLinksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ResourceNavigationLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
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
func (client *ResourceNavigationLinksClient) listHandleResponse(resp *http.Response) (ResourceNavigationLinksClientListResponse, error) {
	result := ResourceNavigationLinksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceNavigationLinksListResult); err != nil {
		return ResourceNavigationLinksClientListResponse{}, err
	}
	return result, nil
}
