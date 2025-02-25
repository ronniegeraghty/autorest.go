//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package errorsgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PetClient contains the methods for the Pet group.
// Don't use this type directly, use NewPetClient() instead.
type PetClient struct {
	pl runtime.Pipeline
}

// NewPetClient creates a new instance of PetClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewPetClient(pl runtime.Pipeline) *PetClient {
	client := &PetClient{
		pl: pl,
	}
	return client
}

// DoSomething - Asks pet to do something
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 0.0.0
// whatAction - what action the pet should do
// options - PetClientDoSomethingOptions contains the optional parameters for the PetClient.DoSomething method.
func (client *PetClient) DoSomething(ctx context.Context, whatAction string, options *PetClientDoSomethingOptions) (PetClientDoSomethingResponse, error) {
	req, err := client.doSomethingCreateRequest(ctx, whatAction, options)
	if err != nil {
		return PetClientDoSomethingResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PetClientDoSomethingResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PetClientDoSomethingResponse{}, runtime.NewResponseError(resp)
	}
	return client.doSomethingHandleResponse(resp)
}

// doSomethingCreateRequest creates the DoSomething request.
func (client *PetClient) doSomethingCreateRequest(ctx context.Context, whatAction string, options *PetClientDoSomethingOptions) (*policy.Request, error) {
	urlPath := "/errorStatusCodes/Pets/doSomething/{whatAction}"
	if whatAction == "" {
		return nil, errors.New("parameter whatAction cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{whatAction}", url.PathEscape(whatAction))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// doSomethingHandleResponse handles the DoSomething response.
func (client *PetClient) doSomethingHandleResponse(resp *http.Response) (PetClientDoSomethingResponse, error) {
	result := PetClientDoSomethingResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PetAction); err != nil {
		return PetClientDoSomethingResponse{}, err
	}
	return result, nil
}

// GetPetByID - Gets pets by id.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 0.0.0
// petID - pet id
// options - PetClientGetPetByIDOptions contains the optional parameters for the PetClient.GetPetByID method.
func (client *PetClient) GetPetByID(ctx context.Context, petID string, options *PetClientGetPetByIDOptions) (PetClientGetPetByIDResponse, error) {
	req, err := client.getPetByIDCreateRequest(ctx, petID, options)
	if err != nil {
		return PetClientGetPetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PetClientGetPetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return PetClientGetPetByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPetByIDHandleResponse(resp)
}

// getPetByIDCreateRequest creates the GetPetByID request.
func (client *PetClient) getPetByIDCreateRequest(ctx context.Context, petID string, options *PetClientGetPetByIDOptions) (*policy.Request, error) {
	urlPath := "/errorStatusCodes/Pets/{petId}/GetPet"
	if petID == "" {
		return nil, errors.New("parameter petID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{petId}", url.PathEscape(petID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPetByIDHandleResponse handles the GetPetByID response.
func (client *PetClient) getPetByIDHandleResponse(resp *http.Response) (PetClientGetPetByIDResponse, error) {
	result := PetClientGetPetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pet); err != nil {
		return PetClientGetPetByIDResponse{}, err
	}
	return result, nil
}

// HasModelsParam - Ensure you can correctly deserialize the returned PetActionError and deserialization doesn't conflict
// with the input param name 'models'
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 0.0.0
// options - PetClientHasModelsParamOptions contains the optional parameters for the PetClient.HasModelsParam method.
func (client *PetClient) HasModelsParam(ctx context.Context, options *PetClientHasModelsParamOptions) (PetClientHasModelsParamResponse, error) {
	req, err := client.hasModelsParamCreateRequest(ctx, options)
	if err != nil {
		return PetClientHasModelsParamResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PetClientHasModelsParamResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PetClientHasModelsParamResponse{}, runtime.NewResponseError(resp)
	}
	return PetClientHasModelsParamResponse{}, nil
}

// hasModelsParamCreateRequest creates the HasModelsParam request.
func (client *PetClient) hasModelsParamCreateRequest(ctx context.Context, options *PetClientHasModelsParamOptions) (*policy.Request, error) {
	urlPath := "/errorStatusCodes/Pets/hasModelsParam"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	modelsDefault := "value1"
	if options != nil && options.Models != nil {
		modelsDefault = *options.Models
	}
	reqQP.Set("models", modelsDefault)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
