//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataboxedge

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

// AvailableSKUsClient contains the methods for the AvailableSKUs group.
// Don't use this type directly, use NewAvailableSKUsClient() instead.
type AvailableSKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailableSKUsClient creates a new instance of AvailableSKUsClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailableSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailableSKUsClient, error) {
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
	client := &AvailableSKUsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - List all the available Skus and information related to them.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01
// options - AvailableSKUsClientListOptions contains the optional parameters for the AvailableSKUsClient.List method.
func (client *AvailableSKUsClient) NewListPager(options *AvailableSKUsClientListOptions) *runtime.Pager[AvailableSKUsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailableSKUsClientListResponse]{
		More: func(page AvailableSKUsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailableSKUsClientListResponse) (AvailableSKUsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AvailableSKUsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AvailableSKUsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AvailableSKUsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AvailableSKUsClient) listCreateRequest(ctx context.Context, options *AvailableSKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataBoxEdge/availableSkus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailableSKUsClient) listHandleResponse(resp *http.Response) (AvailableSKUsClientListResponse, error) {
	result := AvailableSKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUList); err != nil {
		return AvailableSKUsClientListResponse{}, err
	}
	return result, nil
}
