// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

// BoolOperations contains the methods for the Bool group.
type BoolOperations interface {
	// GetFalse - Get false Boolean value
	GetFalse(ctx context.Context) (*BoolGetFalseResponse, error)
	// GetInvalid - Get invalid Boolean value
	GetInvalid(ctx context.Context) (*BoolGetInvalidResponse, error)
	// GetNull - Get null Boolean value
	GetNull(ctx context.Context) (*BoolGetNullResponse, error)
	// GetTrue - Get true Boolean value
	GetTrue(ctx context.Context) (*BoolGetTrueResponse, error)
	// PutFalse - Set Boolean value false
	PutFalse(ctx context.Context) (*BoolPutFalseResponse, error)
	// PutTrue - Set Boolean value true
	PutTrue(ctx context.Context) (*BoolPutTrueResponse, error)
}

// boolOperations implements the BoolOperations interface.
type boolOperations struct {
	*Client
}

// GetFalse - Get false Boolean value
func (client *boolOperations) GetFalse(ctx context.Context) (*BoolGetFalseResponse, error) {
	req, err := client.getFalseCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getFalseHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getFalseCreateRequest creates the GetFalse request.
func (client *boolOperations) getFalseCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/bool/false"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getFalseHandleResponse handles the GetFalse response.
func (client *boolOperations) getFalseHandleResponse(resp *azcore.Response) (*BoolGetFalseResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := BoolGetFalseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalid - Get invalid Boolean value
func (client *boolOperations) GetInvalid(ctx context.Context) (*BoolGetInvalidResponse, error) {
	req, err := client.getInvalidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *boolOperations) getInvalidCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/bool/invalid"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *boolOperations) getInvalidHandleResponse(resp *azcore.Response) (*BoolGetInvalidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := BoolGetInvalidResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNull - Get null Boolean value
func (client *boolOperations) GetNull(ctx context.Context) (*BoolGetNullResponse, error) {
	req, err := client.getNullCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNullCreateRequest creates the GetNull request.
func (client *boolOperations) getNullCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/bool/null"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *boolOperations) getNullHandleResponse(resp *azcore.Response) (*BoolGetNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := BoolGetNullResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetTrue - Get true Boolean value
func (client *boolOperations) GetTrue(ctx context.Context) (*BoolGetTrueResponse, error) {
	req, err := client.getTrueCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getTrueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getTrueCreateRequest creates the GetTrue request.
func (client *boolOperations) getTrueCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/bool/true"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getTrueHandleResponse handles the GetTrue response.
func (client *boolOperations) getTrueHandleResponse(resp *azcore.Response) (*BoolGetTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := BoolGetTrueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// PutFalse - Set Boolean value false
func (client *boolOperations) PutFalse(ctx context.Context) (*BoolPutFalseResponse, error) {
	req, err := client.putFalseCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putFalseHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putFalseCreateRequest creates the PutFalse request.
func (client *boolOperations) putFalseCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/bool/false"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(false)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putFalseHandleResponse handles the PutFalse response.
func (client *boolOperations) putFalseHandleResponse(resp *azcore.Response) (*BoolPutFalseResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &BoolPutFalseResponse{RawResponse: resp.Response}, nil
}

// PutTrue - Set Boolean value true
func (client *boolOperations) PutTrue(ctx context.Context) (*BoolPutTrueResponse, error) {
	req, err := client.putTrueCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putTrueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putTrueCreateRequest creates the PutTrue request.
func (client *boolOperations) putTrueCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/bool/true"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(true)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putTrueHandleResponse handles the PutTrue response.
func (client *boolOperations) putTrueHandleResponse(resp *azcore.Response) (*BoolPutTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &BoolPutTrueResponse{RawResponse: resp.Response}, nil
}
