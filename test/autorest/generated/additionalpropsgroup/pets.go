// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package additionalpropsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

// PetsOperations contains the methods for the Pets group.
type PetsOperations interface {
	// CreateApInProperties - Create a Pet which contains more properties than what is defined.
	CreateApInProperties(ctx context.Context, createParameters PetApInProperties) (*PetApInPropertiesResponse, error)
	// CreateApInPropertiesWithApstring - Create a Pet which contains more properties than what is defined.
	CreateApInPropertiesWithApstring(ctx context.Context, createParameters PetApInPropertiesWithApstring) (*PetApInPropertiesWithApstringResponse, error)
	// CreateApObject - Create a Pet which contains more properties than what is defined.
	CreateApObject(ctx context.Context, createParameters PetApObject) (*PetApObjectResponse, error)
	// CreateApString - Create a Pet which contains more properties than what is defined.
	CreateApString(ctx context.Context, createParameters PetApString) (*PetApStringResponse, error)
	// CreateApTrue - Create a Pet which contains more properties than what is defined.
	CreateApTrue(ctx context.Context, createParameters PetApTrue) (*PetApTrueResponse, error)
	// CreateCatApTrue - Create a CatAPTrue which contains more properties than what is defined.
	CreateCatApTrue(ctx context.Context, createParameters CatApTrue) (*CatApTrueResponse, error)
}

// petsOperations implements the PetsOperations interface.
type petsOperations struct {
	*Client
}

// CreateApInProperties - Create a Pet which contains more properties than what is defined.
func (client *petsOperations) CreateApInProperties(ctx context.Context, createParameters PetApInProperties) (*PetApInPropertiesResponse, error) {
	req, err := client.createApInPropertiesCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createApInPropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createApInPropertiesCreateRequest creates the CreateApInProperties request.
func (client *petsOperations) createApInPropertiesCreateRequest(createParameters PetApInProperties) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/in/properties"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// createApInPropertiesHandleResponse handles the CreateApInProperties response.
func (client *petsOperations) createApInPropertiesHandleResponse(resp *azcore.Response) (*PetApInPropertiesResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.createApInPropertiesHandleError(resp)
	}
	result := PetApInPropertiesResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApInProperties)
}

// createApInPropertiesHandleError handles the CreateApInProperties error response.
func (client *petsOperations) createApInPropertiesHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateApInPropertiesWithApstring - Create a Pet which contains more properties than what is defined.
func (client *petsOperations) CreateApInPropertiesWithApstring(ctx context.Context, createParameters PetApInPropertiesWithApstring) (*PetApInPropertiesWithApstringResponse, error) {
	req, err := client.createApInPropertiesWithApstringCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createApInPropertiesWithApstringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createApInPropertiesWithApstringCreateRequest creates the CreateApInPropertiesWithApstring request.
func (client *petsOperations) createApInPropertiesWithApstringCreateRequest(createParameters PetApInPropertiesWithApstring) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/in/properties/with/additionalProperties/string"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// createApInPropertiesWithApstringHandleResponse handles the CreateApInPropertiesWithApstring response.
func (client *petsOperations) createApInPropertiesWithApstringHandleResponse(resp *azcore.Response) (*PetApInPropertiesWithApstringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.createApInPropertiesWithApstringHandleError(resp)
	}
	result := PetApInPropertiesWithApstringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApInPropertiesWithApstring)
}

// createApInPropertiesWithApstringHandleError handles the CreateApInPropertiesWithApstring error response.
func (client *petsOperations) createApInPropertiesWithApstringHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateApObject - Create a Pet which contains more properties than what is defined.
func (client *petsOperations) CreateApObject(ctx context.Context, createParameters PetApObject) (*PetApObjectResponse, error) {
	req, err := client.createApObjectCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createApObjectHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createApObjectCreateRequest creates the CreateApObject request.
func (client *petsOperations) createApObjectCreateRequest(createParameters PetApObject) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/type/object"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// createApObjectHandleResponse handles the CreateApObject response.
func (client *petsOperations) createApObjectHandleResponse(resp *azcore.Response) (*PetApObjectResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.createApObjectHandleError(resp)
	}
	result := PetApObjectResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApObject)
}

// createApObjectHandleError handles the CreateApObject error response.
func (client *petsOperations) createApObjectHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateApString - Create a Pet which contains more properties than what is defined.
func (client *petsOperations) CreateApString(ctx context.Context, createParameters PetApString) (*PetApStringResponse, error) {
	req, err := client.createApStringCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createApStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createApStringCreateRequest creates the CreateApString request.
func (client *petsOperations) createApStringCreateRequest(createParameters PetApString) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/type/string"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// createApStringHandleResponse handles the CreateApString response.
func (client *petsOperations) createApStringHandleResponse(resp *azcore.Response) (*PetApStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.createApStringHandleError(resp)
	}
	result := PetApStringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApString)
}

// createApStringHandleError handles the CreateApString error response.
func (client *petsOperations) createApStringHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateApTrue - Create a Pet which contains more properties than what is defined.
func (client *petsOperations) CreateApTrue(ctx context.Context, createParameters PetApTrue) (*PetApTrueResponse, error) {
	req, err := client.createApTrueCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createApTrueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createApTrueCreateRequest creates the CreateApTrue request.
func (client *petsOperations) createApTrueCreateRequest(createParameters PetApTrue) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/true"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// createApTrueHandleResponse handles the CreateApTrue response.
func (client *petsOperations) createApTrueHandleResponse(resp *azcore.Response) (*PetApTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.createApTrueHandleError(resp)
	}
	result := PetApTrueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApTrue)
}

// createApTrueHandleError handles the CreateApTrue error response.
func (client *petsOperations) createApTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateCatApTrue - Create a CatAPTrue which contains more properties than what is defined.
func (client *petsOperations) CreateCatApTrue(ctx context.Context, createParameters CatApTrue) (*CatApTrueResponse, error) {
	req, err := client.createCatApTrueCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createCatApTrueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createCatApTrueCreateRequest creates the CreateCatApTrue request.
func (client *petsOperations) createCatApTrueCreateRequest(createParameters CatApTrue) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/true-subclass"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// createCatApTrueHandleResponse handles the CreateCatApTrue response.
func (client *petsOperations) createCatApTrueHandleResponse(resp *azcore.Response) (*CatApTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.createCatApTrueHandleError(resp)
	}
	result := CatApTrueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CatApTrue)
}

// createCatApTrueHandleError handles the CreateCatApTrue error response.
func (client *petsOperations) createCatApTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
