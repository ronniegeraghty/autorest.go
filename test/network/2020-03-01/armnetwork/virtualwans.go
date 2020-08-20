// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// VirtualWansOperations contains the methods for the VirtualWans group.
type VirtualWansOperations interface {
	// BeginCreateOrUpdate - Creates a VirtualWAN resource if it doesn't exist else updates the existing VirtualWAN.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualWanName string, wanParameters VirtualWan) (*VirtualWanPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualWanPoller, error)
	// BeginDelete - Deletes a VirtualWAN.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualWanName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves the details of a VirtualWAN.
	Get(ctx context.Context, resourceGroupName string, virtualWanName string) (*VirtualWanResponse, error)
	// List - Lists all the VirtualWANs in a subscription.
	List() (ListVirtualWaNsResultPager, error)
	// ListByResourceGroup - Lists all the VirtualWANs in a resource group.
	ListByResourceGroup(resourceGroupName string) (ListVirtualWaNsResultPager, error)
	// UpdateTags - Updates a VirtualWAN tags.
	UpdateTags(ctx context.Context, resourceGroupName string, virtualWanName string, wanParameters TagsObject) (*VirtualWanResponse, error)
}

// virtualWansOperations implements the VirtualWansOperations interface.
type virtualWansOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates a VirtualWAN resource if it doesn't exist else updates the existing VirtualWAN.
func (client *virtualWansOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualWanName string, wanParameters VirtualWan) (*VirtualWanPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, virtualWanName, wanParameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("virtualWansOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualWanPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualWanResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualWansOperations) ResumeCreateOrUpdate(token string) (VirtualWanPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualWansOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualWanPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *virtualWansOperations) createOrUpdateCreateRequest(resourceGroupName string, virtualWanName string, wanParameters VirtualWan) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWanName))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(wanParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *virtualWansOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*VirtualWanPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &VirtualWanPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *virtualWansOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes a VirtualWAN.
func (client *virtualWansOperations) BeginDelete(ctx context.Context, resourceGroupName string, virtualWanName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, virtualWanName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("virtualWansOperations.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *virtualWansOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("virtualWansOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *virtualWansOperations) deleteCreateRequest(resourceGroupName string, virtualWanName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWanName))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *virtualWansOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *virtualWansOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Retrieves the details of a VirtualWAN.
func (client *virtualWansOperations) Get(ctx context.Context, resourceGroupName string, virtualWanName string) (*VirtualWanResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, virtualWanName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *virtualWansOperations) getCreateRequest(resourceGroupName string, virtualWanName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWanName))
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

// getHandleResponse handles the Get response.
func (client *virtualWansOperations) getHandleResponse(resp *azcore.Response) (*VirtualWanResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := VirtualWanResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualWan)
}

// getHandleError handles the Get error response.
func (client *virtualWansOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all the VirtualWANs in a subscription.
func (client *virtualWansOperations) List() (ListVirtualWaNsResultPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &listVirtualWaNsResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *ListVirtualWaNsResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ListVirtualWaNsResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ListVirtualWaNsResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *virtualWansOperations) listCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualWans"
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
func (client *virtualWansOperations) listHandleResponse(resp *azcore.Response) (*ListVirtualWaNsResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ListVirtualWaNsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVirtualWaNsResult)
}

// listHandleError handles the List error response.
func (client *virtualWansOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Lists all the VirtualWANs in a resource group.
func (client *virtualWansOperations) ListByResourceGroup(resourceGroupName string) (ListVirtualWaNsResultPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &listVirtualWaNsResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *ListVirtualWaNsResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ListVirtualWaNsResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ListVirtualWaNsResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *virtualWansOperations) listByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *virtualWansOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*ListVirtualWaNsResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := ListVirtualWaNsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVirtualWaNsResult)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *virtualWansOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates a VirtualWAN tags.
func (client *virtualWansOperations) UpdateTags(ctx context.Context, resourceGroupName string, virtualWanName string, wanParameters TagsObject) (*VirtualWanResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, virtualWanName, wanParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *virtualWansOperations) updateTagsCreateRequest(resourceGroupName string, virtualWanName string, wanParameters TagsObject) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{VirtualWANName}", url.PathEscape(virtualWanName))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(wanParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *virtualWansOperations) updateTagsHandleResponse(resp *azcore.Response) (*VirtualWanResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := VirtualWanResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualWan)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *virtualWansOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
