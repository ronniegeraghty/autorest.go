package httpinfrastructuregroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// HTTPFailureClient is the test Infrastructure for AutoRest
type HTTPFailureClient struct {
	ManagementClient
}

// NewHTTPFailureClient creates an instance of the HTTPFailureClient client.
func NewHTTPFailureClient(p pipeline.Pipeline) HTTPFailureClient {
	return HTTPFailureClient{NewManagementClient(p)}
}

// GetEmptyError get empty error form server
func (client HTTPFailureClient) GetEmptyError(ctx context.Context) (*GetEmptyErrorResponse, error) {
	req, err := client.getEmptyErrorPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getEmptyErrorResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetEmptyErrorResponse), err
}

// getEmptyErrorPreparer prepares the GetEmptyError request.
func (client HTTPFailureClient) getEmptyErrorPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/failure/emptybody/error"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getEmptyErrorResponder handles the response to the GetEmptyError request.
func (client HTTPFailureClient) getEmptyErrorResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetEmptyErrorResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetNoModelEmpty get empty response from server
func (client HTTPFailureClient) GetNoModelEmpty(ctx context.Context) (*GetNoModelEmptyResponse, error) {
	req, err := client.getNoModelEmptyPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getNoModelEmptyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetNoModelEmptyResponse), err
}

// getNoModelEmptyPreparer prepares the GetNoModelEmpty request.
func (client HTTPFailureClient) getNoModelEmptyPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/failure/nomodel/empty"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getNoModelEmptyResponder handles the response to the GetNoModelEmpty request.
func (client HTTPFailureClient) getNoModelEmptyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetNoModelEmptyResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetNoModelError get empty error form server
func (client HTTPFailureClient) GetNoModelError(ctx context.Context) (*GetNoModelErrorResponse, error) {
	req, err := client.getNoModelErrorPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getNoModelErrorResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetNoModelErrorResponse), err
}

// getNoModelErrorPreparer prepares the GetNoModelError request.
func (client HTTPFailureClient) getNoModelErrorPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/failure/nomodel/error"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getNoModelErrorResponder handles the response to the GetNoModelError request.
func (client HTTPFailureClient) getNoModelErrorResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetNoModelErrorResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
