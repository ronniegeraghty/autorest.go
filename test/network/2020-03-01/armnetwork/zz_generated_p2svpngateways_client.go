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

// P2SVPNGatewaysClient contains the methods for the P2SVPNGateways group.
// Don't use this type directly, use NewP2SVPNGatewaysClient() instead.
type P2SVPNGatewaysClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewP2SVPNGatewaysClient creates a new instance of P2SVPNGatewaysClient with the specified values.
func NewP2SVPNGatewaysClient(con *armcore.Connection, subscriptionID string) *P2SVPNGatewaysClient {
	return &P2SVPNGatewaysClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a virtual wan p2s vpn gateway if it doesn't exist else updates the existing gateway.
func (client *P2SVPNGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters P2SVPNGateway, options *P2SVPNGatewaysBeginCreateOrUpdateOptions) (P2SVPNGatewayPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, p2SVPNGatewayParameters, options)
	if err != nil {
		return P2SVPNGatewayPollerResponse{}, err
	}
	result := P2SVPNGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("P2SVPNGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return P2SVPNGatewayPollerResponse{}, err
	}
	poller := &p2SVPNGatewayPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (P2SVPNGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new P2SVPNGatewayPoller from the specified resume token.
// token - The value must come from a previous call to P2SVPNGatewayPoller.ResumeToken().
func (client *P2SVPNGatewaysClient) ResumeCreateOrUpdate(token string) (P2SVPNGatewayPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("P2SVPNGatewaysClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &p2SVPNGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates a virtual wan p2s vpn gateway if it doesn't exist else updates the existing gateway.
func (client *P2SVPNGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters P2SVPNGateway, options *P2SVPNGatewaysBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, p2SVPNGatewayParameters, options)
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
func (client *P2SVPNGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters P2SVPNGateway, options *P2SVPNGatewaysBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
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
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(p2SVPNGatewayParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *P2SVPNGatewaysClient) createOrUpdateHandleResponse(resp *azcore.Response) (P2SVPNGatewayResponse, error) {
	var val *P2SVPNGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return P2SVPNGatewayResponse{}, err
	}
	return P2SVPNGatewayResponse{RawResponse: resp.Response, P2SVPNGateway: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *P2SVPNGatewaysClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes a virtual wan p2s vpn gateway.
func (client *P2SVPNGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("P2SVPNGatewaysClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *P2SVPNGatewaysClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("P2SVPNGatewaysClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes a virtual wan p2s vpn gateway.
func (client *P2SVPNGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, options)
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
func (client *P2SVPNGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
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
func (client *P2SVPNGatewaysClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDisconnectP2SVPNConnections - Disconnect P2S vpn connections of the virtual wan P2SVpnGateway in the specified resource group.
func (client *P2SVPNGatewaysClient) BeginDisconnectP2SVPNConnections(ctx context.Context, resourceGroupName string, p2SVPNGatewayName string, request P2SVPNConnectionRequest, options *P2SVPNGatewaysBeginDisconnectP2SVPNConnectionsOptions) (HTTPPollerResponse, error) {
	resp, err := client.disconnectP2SVPNConnections(ctx, resourceGroupName, p2SVPNGatewayName, request, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("P2SVPNGatewaysClient.DisconnectP2SVPNConnections", "location", resp, client.disconnectP2SVPNConnectionsHandleError)
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

// ResumeDisconnectP2SVPNConnections creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *P2SVPNGatewaysClient) ResumeDisconnectP2SVPNConnections(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("P2SVPNGatewaysClient.DisconnectP2SVPNConnections", token, client.disconnectP2SVPNConnectionsHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// DisconnectP2SVPNConnections - Disconnect P2S vpn connections of the virtual wan P2SVpnGateway in the specified resource group.
func (client *P2SVPNGatewaysClient) disconnectP2SVPNConnections(ctx context.Context, resourceGroupName string, p2SVPNGatewayName string, request P2SVPNConnectionRequest, options *P2SVPNGatewaysBeginDisconnectP2SVPNConnectionsOptions) (*azcore.Response, error) {
	req, err := client.disconnectP2SVPNConnectionsCreateRequest(ctx, resourceGroupName, p2SVPNGatewayName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.disconnectP2SVPNConnectionsHandleError(resp)
	}
	return resp, nil
}

// disconnectP2SVPNConnectionsCreateRequest creates the DisconnectP2SVPNConnections request.
func (client *P2SVPNGatewaysClient) disconnectP2SVPNConnectionsCreateRequest(ctx context.Context, resourceGroupName string, p2SVPNGatewayName string, request P2SVPNConnectionRequest, options *P2SVPNGatewaysBeginDisconnectP2SVPNConnectionsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{p2sVpnGatewayName}/disconnectP2sVpnConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if p2SVPNGatewayName == "" {
		return nil, errors.New("parameter p2SVPNGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{p2sVpnGatewayName}", url.PathEscape(p2SVPNGatewayName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// disconnectP2SVPNConnectionsHandleError handles the DisconnectP2SVPNConnections error response.
func (client *P2SVPNGatewaysClient) disconnectP2SVPNConnectionsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginGenerateVPNProfile - Generates VPN profile for P2S client of the P2SVpnGateway in the specified resource group.
func (client *P2SVPNGatewaysClient) BeginGenerateVPNProfile(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVPNProfileParameters, options *P2SVPNGatewaysBeginGenerateVPNProfileOptions) (VPNProfileResponsePollerResponse, error) {
	resp, err := client.generateVPNProfile(ctx, resourceGroupName, gatewayName, parameters, options)
	if err != nil {
		return VPNProfileResponsePollerResponse{}, err
	}
	result := VPNProfileResponsePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("P2SVPNGatewaysClient.GenerateVPNProfile", "location", resp, client.generateVPNProfileHandleError)
	if err != nil {
		return VPNProfileResponsePollerResponse{}, err
	}
	poller := &vpnProfileResponsePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VPNProfileResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeGenerateVPNProfile creates a new VPNProfileResponsePoller from the specified resume token.
// token - The value must come from a previous call to VPNProfileResponsePoller.ResumeToken().
func (client *P2SVPNGatewaysClient) ResumeGenerateVPNProfile(token string) (VPNProfileResponsePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("P2SVPNGatewaysClient.GenerateVPNProfile", token, client.generateVPNProfileHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnProfileResponsePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// GenerateVPNProfile - Generates VPN profile for P2S client of the P2SVpnGateway in the specified resource group.
func (client *P2SVPNGatewaysClient) generateVPNProfile(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVPNProfileParameters, options *P2SVPNGatewaysBeginGenerateVPNProfileOptions) (*azcore.Response, error) {
	req, err := client.generateVPNProfileCreateRequest(ctx, resourceGroupName, gatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.generateVPNProfileHandleError(resp)
	}
	return resp, nil
}

// generateVPNProfileCreateRequest creates the GenerateVPNProfile request.
func (client *P2SVPNGatewaysClient) generateVPNProfileCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVPNProfileParameters, options *P2SVPNGatewaysBeginGenerateVPNProfileOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/generatevpnprofile"
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
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// generateVPNProfileHandleResponse handles the GenerateVPNProfile response.
func (client *P2SVPNGatewaysClient) generateVPNProfileHandleResponse(resp *azcore.Response) (VPNProfileResponseResponse, error) {
	var val *VPNProfileResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNProfileResponseResponse{}, err
	}
	return VPNProfileResponseResponse{RawResponse: resp.Response, VPNProfileResponse: val}, nil
}

// generateVPNProfileHandleError handles the GenerateVPNProfile error response.
func (client *P2SVPNGatewaysClient) generateVPNProfileHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a virtual wan p2s vpn gateway.
func (client *P2SVPNGatewaysClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysGetOptions) (P2SVPNGatewayResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return P2SVPNGatewayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return P2SVPNGatewayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return P2SVPNGatewayResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *P2SVPNGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
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
func (client *P2SVPNGatewaysClient) getHandleResponse(resp *azcore.Response) (P2SVPNGatewayResponse, error) {
	var val *P2SVPNGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return P2SVPNGatewayResponse{}, err
	}
	return P2SVPNGatewayResponse{RawResponse: resp.Response, P2SVPNGateway: val}, nil
}

// getHandleError handles the Get error response.
func (client *P2SVPNGatewaysClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginGetP2SVPNConnectionHealth - Gets the connection health of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.
func (client *P2SVPNGatewaysClient) BeginGetP2SVPNConnectionHealth(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthOptions) (P2SVPNGatewayPollerResponse, error) {
	resp, err := client.getP2SVPNConnectionHealth(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return P2SVPNGatewayPollerResponse{}, err
	}
	result := P2SVPNGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("P2SVPNGatewaysClient.GetP2SVPNConnectionHealth", "location", resp, client.getP2SVPNConnectionHealthHandleError)
	if err != nil {
		return P2SVPNGatewayPollerResponse{}, err
	}
	poller := &p2SVPNGatewayPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (P2SVPNGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeGetP2SVPNConnectionHealth creates a new P2SVPNGatewayPoller from the specified resume token.
// token - The value must come from a previous call to P2SVPNGatewayPoller.ResumeToken().
func (client *P2SVPNGatewaysClient) ResumeGetP2SVPNConnectionHealth(token string) (P2SVPNGatewayPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("P2SVPNGatewaysClient.GetP2SVPNConnectionHealth", token, client.getP2SVPNConnectionHealthHandleError)
	if err != nil {
		return nil, err
	}
	return &p2SVPNGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// GetP2SVPNConnectionHealth - Gets the connection health of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealth(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthOptions) (*azcore.Response, error) {
	req, err := client.getP2SVPNConnectionHealthCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.getP2SVPNConnectionHealthHandleError(resp)
	}
	return resp, nil
}

// getP2SVPNConnectionHealthCreateRequest creates the GetP2SVPNConnectionHealth request.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/getP2sVpnConnectionHealth"
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
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// getP2SVPNConnectionHealthHandleResponse handles the GetP2SVPNConnectionHealth response.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthHandleResponse(resp *azcore.Response) (P2SVPNGatewayResponse, error) {
	var val *P2SVPNGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return P2SVPNGatewayResponse{}, err
	}
	return P2SVPNGatewayResponse{RawResponse: resp.Response, P2SVPNGateway: val}, nil
}

// getP2SVPNConnectionHealthHandleError handles the GetP2SVPNConnectionHealth error response.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginGetP2SVPNConnectionHealthDetailed - Gets the sas url to get the connection health detail of P2S clients of the virtual wan P2SVpnGateway in the
// specified resource group.
func (client *P2SVPNGatewaysClient) BeginGetP2SVPNConnectionHealthDetailed(ctx context.Context, resourceGroupName string, gatewayName string, request P2SVPNConnectionHealthRequest, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthDetailedOptions) (P2SVPNConnectionHealthPollerResponse, error) {
	resp, err := client.getP2SVPNConnectionHealthDetailed(ctx, resourceGroupName, gatewayName, request, options)
	if err != nil {
		return P2SVPNConnectionHealthPollerResponse{}, err
	}
	result := P2SVPNConnectionHealthPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("P2SVPNGatewaysClient.GetP2SVPNConnectionHealthDetailed", "location", resp, client.getP2SVPNConnectionHealthDetailedHandleError)
	if err != nil {
		return P2SVPNConnectionHealthPollerResponse{}, err
	}
	poller := &p2SVPNConnectionHealthPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (P2SVPNConnectionHealthResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeGetP2SVPNConnectionHealthDetailed creates a new P2SVPNConnectionHealthPoller from the specified resume token.
// token - The value must come from a previous call to P2SVPNConnectionHealthPoller.ResumeToken().
func (client *P2SVPNGatewaysClient) ResumeGetP2SVPNConnectionHealthDetailed(token string) (P2SVPNConnectionHealthPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("P2SVPNGatewaysClient.GetP2SVPNConnectionHealthDetailed", token, client.getP2SVPNConnectionHealthDetailedHandleError)
	if err != nil {
		return nil, err
	}
	return &p2SVPNConnectionHealthPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// GetP2SVPNConnectionHealthDetailed - Gets the sas url to get the connection health detail of P2S clients of the virtual wan P2SVpnGateway in the specified
// resource group.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthDetailed(ctx context.Context, resourceGroupName string, gatewayName string, request P2SVPNConnectionHealthRequest, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthDetailedOptions) (*azcore.Response, error) {
	req, err := client.getP2SVPNConnectionHealthDetailedCreateRequest(ctx, resourceGroupName, gatewayName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.getP2SVPNConnectionHealthDetailedHandleError(resp)
	}
	return resp, nil
}

// getP2SVPNConnectionHealthDetailedCreateRequest creates the GetP2SVPNConnectionHealthDetailed request.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthDetailedCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, request P2SVPNConnectionHealthRequest, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthDetailedOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/getP2sVpnConnectionHealthDetailed"
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
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// getP2SVPNConnectionHealthDetailedHandleResponse handles the GetP2SVPNConnectionHealthDetailed response.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthDetailedHandleResponse(resp *azcore.Response) (P2SVPNConnectionHealthResponse, error) {
	var val *P2SVPNConnectionHealth
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return P2SVPNConnectionHealthResponse{}, err
	}
	return P2SVPNConnectionHealthResponse{RawResponse: resp.Response, P2SVPNConnectionHealth: val}, nil
}

// getP2SVPNConnectionHealthDetailedHandleError handles the GetP2SVPNConnectionHealthDetailed error response.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthDetailedHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all the P2SVpnGateways in a subscription.
func (client *P2SVPNGatewaysClient) List(options *P2SVPNGatewaysListOptions) ListP2SVPNGatewaysResultPager {
	return &listP2SVPNGatewaysResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ListP2SVPNGatewaysResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListP2SVPNGatewaysResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *P2SVPNGatewaysClient) listCreateRequest(ctx context.Context, options *P2SVPNGatewaysListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/p2svpnGateways"
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
func (client *P2SVPNGatewaysClient) listHandleResponse(resp *azcore.Response) (ListP2SVPNGatewaysResultResponse, error) {
	var val *ListP2SVPNGatewaysResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListP2SVPNGatewaysResultResponse{}, err
	}
	return ListP2SVPNGatewaysResultResponse{RawResponse: resp.Response, ListP2SVPNGatewaysResult: val}, nil
}

// listHandleError handles the List error response.
func (client *P2SVPNGatewaysClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all the P2SVpnGateways in a resource group.
func (client *P2SVPNGatewaysClient) ListByResourceGroup(resourceGroupName string, options *P2SVPNGatewaysListByResourceGroupOptions) ListP2SVPNGatewaysResultPager {
	return &listP2SVPNGatewaysResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ListP2SVPNGatewaysResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListP2SVPNGatewaysResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *P2SVPNGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *P2SVPNGatewaysListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *P2SVPNGatewaysClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ListP2SVPNGatewaysResultResponse, error) {
	var val *ListP2SVPNGatewaysResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListP2SVPNGatewaysResultResponse{}, err
	}
	return ListP2SVPNGatewaysResultResponse{RawResponse: resp.Response, ListP2SVPNGatewaysResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *P2SVPNGatewaysClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates virtual wan p2s vpn gateway tags.
func (client *P2SVPNGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters TagsObject, options *P2SVPNGatewaysUpdateTagsOptions) (P2SVPNGatewayResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, gatewayName, p2SVPNGatewayParameters, options)
	if err != nil {
		return P2SVPNGatewayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return P2SVPNGatewayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return P2SVPNGatewayResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *P2SVPNGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters TagsObject, options *P2SVPNGatewaysUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
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
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(p2SVPNGatewayParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *P2SVPNGatewaysClient) updateTagsHandleResponse(resp *azcore.Response) (P2SVPNGatewayResponse, error) {
	var val *P2SVPNGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return P2SVPNGatewayResponse{}, err
	}
	return P2SVPNGatewayResponse{RawResponse: resp.Response, P2SVPNGateway: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *P2SVPNGatewaysClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
