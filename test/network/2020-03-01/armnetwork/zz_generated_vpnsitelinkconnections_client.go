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

// VPNSiteLinkConnectionsClient contains the methods for the VPNSiteLinkConnections group.
// Don't use this type directly, use NewVPNSiteLinkConnectionsClient() instead.
type VPNSiteLinkConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVPNSiteLinkConnectionsClient creates a new instance of VPNSiteLinkConnectionsClient with the specified values.
func NewVPNSiteLinkConnectionsClient(con *armcore.Connection, subscriptionID string) *VPNSiteLinkConnectionsClient {
	return &VPNSiteLinkConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Get - Retrieves the details of a vpn site link connection.
func (client *VPNSiteLinkConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNSiteLinkConnectionsGetOptions) (VPNSiteLinkConnectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName, options)
	if err != nil {
		return VPNSiteLinkConnectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VPNSiteLinkConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VPNSiteLinkConnectionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNSiteLinkConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNSiteLinkConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}"
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
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
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
func (client *VPNSiteLinkConnectionsClient) getHandleResponse(resp *azcore.Response) (VPNSiteLinkConnectionResponse, error) {
	var val *VPNSiteLinkConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNSiteLinkConnectionResponse{}, err
	}
	return VPNSiteLinkConnectionResponse{RawResponse: resp.Response, VPNSiteLinkConnection: val}, nil
}

// getHandleError handles the Get error response.
func (client *VPNSiteLinkConnectionsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
