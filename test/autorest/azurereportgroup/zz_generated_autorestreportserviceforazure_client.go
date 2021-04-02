// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurereportgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// AutoRestReportServiceForAzureClient contains the methods for the AutoRestReportServiceForAzure group.
// Don't use this type directly, use NewAutoRestReportServiceForAzureClient() instead.
type AutoRestReportServiceForAzureClient struct {
	con *Connection
}

// NewAutoRestReportServiceForAzureClient creates a new instance of AutoRestReportServiceForAzureClient with the specified values.
func NewAutoRestReportServiceForAzureClient(con *Connection) *AutoRestReportServiceForAzureClient {
	return &AutoRestReportServiceForAzureClient{con: con}
}

// GetReport - Get test coverage report
func (client *AutoRestReportServiceForAzureClient) GetReport(ctx context.Context, options *AutoRestReportServiceForAzureGetReportOptions) (MapOfInt32Response, error) {
	req, err := client.getReportCreateRequest(ctx, options)
	if err != nil {
		return MapOfInt32Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MapOfInt32Response{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MapOfInt32Response{}, client.getReportHandleError(resp)
	}
	return client.getReportHandleResponse(resp)
}

// getReportCreateRequest creates the GetReport request.
func (client *AutoRestReportServiceForAzureClient) getReportCreateRequest(ctx context.Context, options *AutoRestReportServiceForAzureGetReportOptions) (*azcore.Request, error) {
	urlPath := "/report/azure"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Qualifier != nil {
		reqQP.Set("qualifier", *options.Qualifier)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getReportHandleResponse handles the GetReport response.
func (client *AutoRestReportServiceForAzureClient) getReportHandleResponse(resp *azcore.Response) (MapOfInt32Response, error) {
	var val map[string]*int32
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return MapOfInt32Response{}, err
	}
	return MapOfInt32Response{RawResponse: resp.Response, Value: val}, nil
}

// getReportHandleError handles the GetReport error response.
func (client *AutoRestReportServiceForAzureClient) getReportHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
