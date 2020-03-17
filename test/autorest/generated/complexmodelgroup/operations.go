// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexmodelgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Operations contains the methods for the Operations group.
type Operations interface {
	// Create - Resets products.
	Create(ctx context.Context, subscriptionId string, resourceGroupName string, bodyParameter CatalogDictionaryOfArray) (*CreateResponse, error)
	// List - The Products endpoint returns information about the Uber products offered at a given location. The response includes the display name and other details about each product, and lists the products in the proper display order.
	List(ctx context.Context, resourceGroupName string) (*ListResponse, error)
	// Update - Resets products.
	Update(ctx context.Context, subscriptionId string, resourceGroupName string, bodyParameter CatalogArrayOfDictionary) (*UpdateResponse, error)
}

// operations implements the Operations interface.
type operations struct {
	*Client
}

// Create - Resets products.
func (client *operations) Create(ctx context.Context, subscriptionId string, resourceGroupName string, bodyParameter CatalogDictionaryOfArray) (*CreateResponse, error) {
	req, err := client.createCreateRequest(*client.u, subscriptionId, resourceGroupName, bodyParameter)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createCreateRequest creates the Create request.
func (client *operations) createCreateRequest(u url.URL, subscriptionId string, resourceGroupName string, bodyParameter CatalogDictionaryOfArray) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/Microsoft.Cache/Redis"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	u.Path = path.Join(u.Path, urlPath)
	query := u.Query()
	query.Set("apiVersion", "2014-04-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, u)
	err := req.MarshalAsJSON(bodyParameter)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *operations) createHandleResponse(resp *azcore.Response) (*CreateResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := CreateResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CatalogDictionary)
}

// List - The Products endpoint returns information about the Uber products offered at a given location. The response includes the display name and other details about each product, and lists the products in the proper display order.
func (client *operations) List(ctx context.Context, resourceGroupName string) (*ListResponse, error) {
	req, err := client.listCreateRequest(*client.u, resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listCreateRequest creates the List request.
func (client *operations) listCreateRequest(u url.URL, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/Microsoft.Cache/Redis"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape("123456"))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	u.Path = path.Join(u.Path, urlPath)
	query := u.Query()
	query.Set("apiVersion", "2014-04-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *operations) listHandleResponse(resp *azcore.Response) (*ListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := ListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CatalogArray)
}

// Update - Resets products.
func (client *operations) Update(ctx context.Context, subscriptionId string, resourceGroupName string, bodyParameter CatalogArrayOfDictionary) (*UpdateResponse, error) {
	req, err := client.updateCreateRequest(*client.u, subscriptionId, resourceGroupName, bodyParameter)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateCreateRequest creates the Update request.
func (client *operations) updateCreateRequest(u url.URL, subscriptionId string, resourceGroupName string, bodyParameter CatalogArrayOfDictionary) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/Microsoft.Cache/Redis"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	u.Path = path.Join(u.Path, urlPath)
	query := u.Query()
	query.Set("apiVersion", "2014-04-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(bodyParameter)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *operations) updateHandleResponse(resp *azcore.Response) (*UpdateResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := UpdateResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CatalogArray)
}
