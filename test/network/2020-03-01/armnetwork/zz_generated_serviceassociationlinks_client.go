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
)

// ServiceAssociationLinksClient contains the methods for the ServiceAssociationLinks group.
// Don't use this type directly, use NewServiceAssociationLinksClient() instead.
type ServiceAssociationLinksClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewServiceAssociationLinksClient creates a new instance of ServiceAssociationLinksClient with the specified values.
func NewServiceAssociationLinksClient(con *armcore.Connection, subscriptionID string) *ServiceAssociationLinksClient {
	return &ServiceAssociationLinksClient{con: con, subscriptionID: subscriptionID}
}

// List - Gets a list of service association links for a subnet.
func (client *ServiceAssociationLinksClient) List(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *ServiceAssociationLinksListOptions) (ServiceAssociationLinksListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
	if err != nil {
		return ServiceAssociationLinksListResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceAssociationLinksListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceAssociationLinksListResultResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ServiceAssociationLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *ServiceAssociationLinksListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ServiceAssociationLinks"
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
func (client *ServiceAssociationLinksClient) listHandleResponse(resp *azcore.Response) (ServiceAssociationLinksListResultResponse, error) {
	var val *ServiceAssociationLinksListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ServiceAssociationLinksListResultResponse{}, err
	}
	return ServiceAssociationLinksListResultResponse{RawResponse: resp.Response, ServiceAssociationLinksListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ServiceAssociationLinksClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
