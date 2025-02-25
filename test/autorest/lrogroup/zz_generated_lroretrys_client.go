//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package lrogroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// LRORetrysClient contains the methods for the LRORetrys group.
// Don't use this type directly, use NewLRORetrysClient() instead.
type LRORetrysClient struct {
	pl runtime.Pipeline
}

// NewLRORetrysClient creates a new instance of LRORetrysClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewLRORetrysClient(pl runtime.Pipeline) *LRORetrysClient {
	client := &LRORetrysClient{
		pl: pl,
	}
	return client
}

// BeginDelete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return
// this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - LRORetrysClientBeginDelete202Retry200Options contains the optional parameters for the LRORetrysClient.BeginDelete202Retry200
// method.
func (client *LRORetrysClient) BeginDelete202Retry200(ctx context.Context, options *LRORetrysClientBeginDelete202Retry200Options) (*runtime.Poller[LRORetrysClientDelete202Retry200Response], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.delete202Retry200(ctx, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LRORetrysClientDelete202Retry200Response](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LRORetrysClientDelete202Retry200Response](options.ResumeToken, client.pl, nil)
	}
}

// Delete202Retry200 - Long running delete request, service returns a 500, then a 202 to the initial request. Polls return
// this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
func (client *LRORetrysClient) delete202Retry200(ctx context.Context, options *LRORetrysClientBeginDelete202Retry200Options) (*http.Response, error) {
	req, err := client.delete202Retry200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// delete202Retry200CreateRequest creates the Delete202Retry200 request.
func (client *LRORetrysClient) delete202Retry200CreateRequest(ctx context.Context, options *LRORetrysClientBeginDelete202Retry200Options) (*policy.Request, error) {
	urlPath := "/lro/retryerror/delete/202/retry/200"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial
// request. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - LRORetrysClientBeginDeleteAsyncRelativeRetrySucceededOptions contains the optional parameters for the LRORetrysClient.BeginDeleteAsyncRelativeRetrySucceeded
// method.
func (client *LRORetrysClient) BeginDeleteAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysClientBeginDeleteAsyncRelativeRetrySucceededOptions) (*runtime.Poller[LRORetrysClientDeleteAsyncRelativeRetrySucceededResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteAsyncRelativeRetrySucceeded(ctx, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LRORetrysClientDeleteAsyncRelativeRetrySucceededResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LRORetrysClientDeleteAsyncRelativeRetrySucceededResponse](options.ResumeToken, client.pl, nil)
	}
}

// DeleteAsyncRelativeRetrySucceeded - Long running delete request, service returns a 500, then a 202 to the initial request.
// Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
func (client *LRORetrysClient) deleteAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysClientBeginDeleteAsyncRelativeRetrySucceededOptions) (*http.Response, error) {
	req, err := client.deleteAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteAsyncRelativeRetrySucceededCreateRequest creates the DeleteAsyncRelativeRetrySucceeded request.
func (client *LRORetrysClient) deleteAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LRORetrysClientBeginDeleteAsyncRelativeRetrySucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/retryerror/deleteasync/retry/succeeded"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a 202 to the
// initial request, with an entity that contains ProvisioningState=’Accepted’. Polls return this value until the last poll
// returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - LRORetrysClientBeginDeleteProvisioning202Accepted200SucceededOptions contains the optional parameters for the
// LRORetrysClient.BeginDeleteProvisioning202Accepted200Succeeded method.
func (client *LRORetrysClient) BeginDeleteProvisioning202Accepted200Succeeded(ctx context.Context, options *LRORetrysClientBeginDeleteProvisioning202Accepted200SucceededOptions) (*runtime.Poller[LRORetrysClientDeleteProvisioning202Accepted200SucceededResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteProvisioning202Accepted200Succeeded(ctx, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LRORetrysClientDeleteProvisioning202Accepted200SucceededResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LRORetrysClientDeleteProvisioning202Accepted200SucceededResponse](options.ResumeToken, client.pl, nil)
	}
}

// DeleteProvisioning202Accepted200Succeeded - Long running delete request, service returns a 500, then a 202 to the initial
// request, with an entity that contains ProvisioningState=’Accepted’. Polls return this value until the last poll returns
// a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
func (client *LRORetrysClient) deleteProvisioning202Accepted200Succeeded(ctx context.Context, options *LRORetrysClientBeginDeleteProvisioning202Accepted200SucceededOptions) (*http.Response, error) {
	req, err := client.deleteProvisioning202Accepted200SucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteProvisioning202Accepted200SucceededCreateRequest creates the DeleteProvisioning202Accepted200Succeeded request.
func (client *LRORetrysClient) deleteProvisioning202Accepted200SucceededCreateRequest(ctx context.Context, options *LRORetrysClientBeginDeleteProvisioning202Accepted200SucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/retryerror/delete/provisioning/202/accepted/200/succeeded"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginPost202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location'
// and 'Retry-After' headers, Polls return a 200 with a response body after success
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - LRORetrysClientBeginPost202Retry200Options contains the optional parameters for the LRORetrysClient.BeginPost202Retry200
// method.
func (client *LRORetrysClient) BeginPost202Retry200(ctx context.Context, options *LRORetrysClientBeginPost202Retry200Options) (*runtime.Poller[LRORetrysClientPost202Retry200Response], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.post202Retry200(ctx, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LRORetrysClientPost202Retry200Response](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LRORetrysClientPost202Retry200Response](options.ResumeToken, client.pl, nil)
	}
}

// Post202Retry200 - Long running post request, service returns a 500, then a 202 to the initial request, with 'Location'
// and 'Retry-After' headers, Polls return a 200 with a response body after success
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
func (client *LRORetrysClient) post202Retry200(ctx context.Context, options *LRORetrysClientBeginPost202Retry200Options) (*http.Response, error) {
	req, err := client.post202Retry200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// post202Retry200CreateRequest creates the Post202Retry200 request.
func (client *LRORetrysClient) post202Retry200CreateRequest(ctx context.Context, options *LRORetrysClientBeginPost202Retry200Options) (*policy.Request, error) {
	urlPath := "/lro/retryerror/post/202/retry/200"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Product != nil {
		return req, runtime.MarshalAsJSON(req, *options.Product)
	}
	return req, nil
}

// BeginPostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request,
// with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions contains the optional parameters for the LRORetrysClient.BeginPostAsyncRelativeRetrySucceeded
// method.
func (client *LRORetrysClient) BeginPostAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions) (*runtime.Poller[LRORetrysClientPostAsyncRelativeRetrySucceededResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.postAsyncRelativeRetrySucceeded(ctx, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LRORetrysClientPostAsyncRelativeRetrySucceededResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LRORetrysClientPostAsyncRelativeRetrySucceededResponse](options.ResumeToken, client.pl, nil)
	}
}

// PostAsyncRelativeRetrySucceeded - Long running post request, service returns a 500, then a 202 to the initial request,
// with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
func (client *LRORetrysClient) postAsyncRelativeRetrySucceeded(ctx context.Context, options *LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions) (*http.Response, error) {
	req, err := client.postAsyncRelativeRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// postAsyncRelativeRetrySucceededCreateRequest creates the PostAsyncRelativeRetrySucceeded request.
func (client *LRORetrysClient) postAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, options *LRORetrysClientBeginPostAsyncRelativeRetrySucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/retryerror/postasync/retry/succeeded"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Product != nil {
		return req, runtime.MarshalAsJSON(req, *options.Product)
	}
	return req, nil
}

// BeginPut201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with
// an entity that contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// product - Product to put
// options - LRORetrysClientBeginPut201CreatingSucceeded200Options contains the optional parameters for the LRORetrysClient.BeginPut201CreatingSucceeded200
// method.
func (client *LRORetrysClient) BeginPut201CreatingSucceeded200(ctx context.Context, product Product, options *LRORetrysClientBeginPut201CreatingSucceeded200Options) (*runtime.Poller[LRORetrysClientPut201CreatingSucceeded200Response], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.put201CreatingSucceeded200(ctx, product, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LRORetrysClientPut201CreatingSucceeded200Response](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LRORetrysClientPut201CreatingSucceeded200Response](options.ResumeToken, client.pl, nil)
	}
}

// Put201CreatingSucceeded200 - Long running put request, service returns a 500, then a 201 to the initial request, with an
// entity that contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a
// ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
func (client *LRORetrysClient) put201CreatingSucceeded200(ctx context.Context, product Product, options *LRORetrysClientBeginPut201CreatingSucceeded200Options) (*http.Response, error) {
	req, err := client.put201CreatingSucceeded200CreateRequest(ctx, product, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client *LRORetrysClient) put201CreatingSucceeded200CreateRequest(ctx context.Context, product Product, options *LRORetrysClientBeginPut201CreatingSucceeded200Options) (*policy.Request, error) {
	urlPath := "/lro/retryerror/put/201/creating/succeeded/200"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, product)
}

// BeginPutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request,
// with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// product - Product to put
// options - LRORetrysClientBeginPutAsyncRelativeRetrySucceededOptions contains the optional parameters for the LRORetrysClient.BeginPutAsyncRelativeRetrySucceeded
// method.
func (client *LRORetrysClient) BeginPutAsyncRelativeRetrySucceeded(ctx context.Context, product Product, options *LRORetrysClientBeginPutAsyncRelativeRetrySucceededOptions) (*runtime.Poller[LRORetrysClientPutAsyncRelativeRetrySucceededResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.putAsyncRelativeRetrySucceeded(ctx, product, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LRORetrysClientPutAsyncRelativeRetrySucceededResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LRORetrysClientPutAsyncRelativeRetrySucceededResponse](options.ResumeToken, client.pl, nil)
	}
}

// PutAsyncRelativeRetrySucceeded - Long running put request, service returns a 500, then a 200 to the initial request, with
// an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation
// header for operation status
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
func (client *LRORetrysClient) putAsyncRelativeRetrySucceeded(ctx context.Context, product Product, options *LRORetrysClientBeginPutAsyncRelativeRetrySucceededOptions) (*http.Response, error) {
	req, err := client.putAsyncRelativeRetrySucceededCreateRequest(ctx, product, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// putAsyncRelativeRetrySucceededCreateRequest creates the PutAsyncRelativeRetrySucceeded request.
func (client *LRORetrysClient) putAsyncRelativeRetrySucceededCreateRequest(ctx context.Context, product Product, options *LRORetrysClientBeginPutAsyncRelativeRetrySucceededOptions) (*policy.Request, error) {
	urlPath := "/lro/retryerror/putasync/retry/succeeded"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, product)
}
