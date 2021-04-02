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

// ExpressRouteServiceProvidersClient contains the methods for the ExpressRouteServiceProviders group.
// Don't use this type directly, use NewExpressRouteServiceProvidersClient() instead.
type ExpressRouteServiceProvidersClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteServiceProvidersClient creates a new instance of ExpressRouteServiceProvidersClient with the specified values.
func NewExpressRouteServiceProvidersClient(con *armcore.Connection, subscriptionID string) *ExpressRouteServiceProvidersClient {
	return &ExpressRouteServiceProvidersClient{con: con, subscriptionID: subscriptionID}
}

// List - Gets all the available express route service providers.
func (client *ExpressRouteServiceProvidersClient) List(options *ExpressRouteServiceProvidersListOptions) ExpressRouteServiceProviderListResultPager {
	return &expressRouteServiceProviderListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ExpressRouteServiceProviderListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteServiceProviderListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRouteServiceProvidersClient) listCreateRequest(ctx context.Context, options *ExpressRouteServiceProvidersListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders"
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
func (client *ExpressRouteServiceProvidersClient) listHandleResponse(resp *azcore.Response) (ExpressRouteServiceProviderListResultResponse, error) {
	var val *ExpressRouteServiceProviderListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteServiceProviderListResultResponse{}, err
	}
	return ExpressRouteServiceProviderListResultResponse{RawResponse: resp.Response, ExpressRouteServiceProviderListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ExpressRouteServiceProvidersClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
