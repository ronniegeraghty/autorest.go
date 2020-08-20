// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// NetworkInterfaceLoadBalancersOperations contains the methods for the NetworkInterfaceLoadBalancers group.
type NetworkInterfaceLoadBalancersOperations interface {
	// List - List all load balancers in a network interface.
	List(resourceGroupName string, networkInterfaceName string) (NetworkInterfaceLoadBalancerListResultPager, error)
}

// networkInterfaceLoadBalancersOperations implements the NetworkInterfaceLoadBalancersOperations interface.
type networkInterfaceLoadBalancersOperations struct {
	*Client
	subscriptionID string
}

// List - List all load balancers in a network interface.
func (client *networkInterfaceLoadBalancersOperations) List(resourceGroupName string, networkInterfaceName string) (NetworkInterfaceLoadBalancerListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, networkInterfaceName)
	if err != nil {
		return nil, err
	}
	return &networkInterfaceLoadBalancerListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *NetworkInterfaceLoadBalancerListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.NetworkInterfaceLoadBalancerListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.NetworkInterfaceLoadBalancerListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *networkInterfaceLoadBalancersOperations) listCreateRequest(resourceGroupName string, networkInterfaceName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
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
func (client *networkInterfaceLoadBalancersOperations) listHandleResponse(resp *azcore.Response) (*NetworkInterfaceLoadBalancerListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := NetworkInterfaceLoadBalancerListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkInterfaceLoadBalancerListResult)
}

// listHandleError handles the List error response.
func (client *networkInterfaceLoadBalancersOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
