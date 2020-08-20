// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

// HTTPClientFailureOperations contains the methods for the HTTPClientFailure group.
type HTTPClientFailureOperations interface {
	// Delete400 - Return 400 status code - should be represented in the client as an error
	Delete400(ctx context.Context) (*http.Response, error)
	// Delete407 - Return 407 status code - should be represented in the client as an error
	Delete407(ctx context.Context) (*http.Response, error)
	// Delete417 - Return 417 status code - should be represented in the client as an error
	Delete417(ctx context.Context) (*http.Response, error)
	// Get400 - Return 400 status code - should be represented in the client as an error
	Get400(ctx context.Context) (*http.Response, error)
	// Get402 - Return 402 status code - should be represented in the client as an error
	Get402(ctx context.Context) (*http.Response, error)
	// Get403 - Return 403 status code - should be represented in the client as an error
	Get403(ctx context.Context) (*http.Response, error)
	// Get411 - Return 411 status code - should be represented in the client as an error
	Get411(ctx context.Context) (*http.Response, error)
	// Get412 - Return 412 status code - should be represented in the client as an error
	Get412(ctx context.Context) (*http.Response, error)
	// Get416 - Return 416 status code - should be represented in the client as an error
	Get416(ctx context.Context) (*http.Response, error)
	// Head400 - Return 400 status code - should be represented in the client as an error
	Head400(ctx context.Context) (*http.Response, error)
	// Head401 - Return 401 status code - should be represented in the client as an error
	Head401(ctx context.Context) (*http.Response, error)
	// Head410 - Return 410 status code - should be represented in the client as an error
	Head410(ctx context.Context) (*http.Response, error)
	// Head429 - Return 429 status code - should be represented in the client as an error
	Head429(ctx context.Context) (*http.Response, error)
	// Options400 - Return 400 status code - should be represented in the client as an error
	Options400(ctx context.Context) (*http.Response, error)
	// Options403 - Return 403 status code - should be represented in the client as an error
	Options403(ctx context.Context) (*http.Response, error)
	// Options412 - Return 412 status code - should be represented in the client as an error
	Options412(ctx context.Context) (*http.Response, error)
	// Patch400 - Return 400 status code - should be represented in the client as an error
	Patch400(ctx context.Context) (*http.Response, error)
	// Patch405 - Return 405 status code - should be represented in the client as an error
	Patch405(ctx context.Context) (*http.Response, error)
	// Patch414 - Return 414 status code - should be represented in the client as an error
	Patch414(ctx context.Context) (*http.Response, error)
	// Post400 - Return 400 status code - should be represented in the client as an error
	Post400(ctx context.Context) (*http.Response, error)
	// Post406 - Return 406 status code - should be represented in the client as an error
	Post406(ctx context.Context) (*http.Response, error)
	// Post415 - Return 415 status code - should be represented in the client as an error
	Post415(ctx context.Context) (*http.Response, error)
	// Put400 - Return 400 status code - should be represented in the client as an error
	Put400(ctx context.Context) (*http.Response, error)
	// Put404 - Return 404 status code - should be represented in the client as an error
	Put404(ctx context.Context) (*http.Response, error)
	// Put409 - Return 409 status code - should be represented in the client as an error
	Put409(ctx context.Context) (*http.Response, error)
	// Put413 - Return 413 status code - should be represented in the client as an error
	Put413(ctx context.Context) (*http.Response, error)
}

// httpClientFailureOperations implements the HTTPClientFailureOperations interface.
type httpClientFailureOperations struct {
	*Client
}

// Delete400 - Return 400 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Delete400(ctx context.Context) (*http.Response, error) {
	req, err := client.delete400CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.delete400HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// delete400CreateRequest creates the Delete400 request.
func (client *httpClientFailureOperations) delete400CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/400"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, req.MarshalAsJSON(true)
}

// delete400HandleResponse handles the Delete400 response.
func (client *httpClientFailureOperations) delete400HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.delete400HandleError(resp)
}

// delete400HandleError handles the Delete400 error response.
func (client *httpClientFailureOperations) delete400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete407 - Return 407 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Delete407(ctx context.Context) (*http.Response, error) {
	req, err := client.delete407CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.delete407HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// delete407CreateRequest creates the Delete407 request.
func (client *httpClientFailureOperations) delete407CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/407"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, req.MarshalAsJSON(true)
}

// delete407HandleResponse handles the Delete407 response.
func (client *httpClientFailureOperations) delete407HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.delete407HandleError(resp)
}

// delete407HandleError handles the Delete407 error response.
func (client *httpClientFailureOperations) delete407HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete417 - Return 417 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Delete417(ctx context.Context) (*http.Response, error) {
	req, err := client.delete417CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.delete417HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// delete417CreateRequest creates the Delete417 request.
func (client *httpClientFailureOperations) delete417CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/417"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, req.MarshalAsJSON(true)
}

// delete417HandleResponse handles the Delete417 response.
func (client *httpClientFailureOperations) delete417HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.delete417HandleError(resp)
}

// delete417HandleError handles the Delete417 error response.
func (client *httpClientFailureOperations) delete417HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get400 - Return 400 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Get400(ctx context.Context) (*http.Response, error) {
	req, err := client.get400CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.get400HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// get400CreateRequest creates the Get400 request.
func (client *httpClientFailureOperations) get400CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/400"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// get400HandleResponse handles the Get400 response.
func (client *httpClientFailureOperations) get400HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.get400HandleError(resp)
}

// get400HandleError handles the Get400 error response.
func (client *httpClientFailureOperations) get400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get402 - Return 402 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Get402(ctx context.Context) (*http.Response, error) {
	req, err := client.get402CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.get402HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// get402CreateRequest creates the Get402 request.
func (client *httpClientFailureOperations) get402CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/402"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// get402HandleResponse handles the Get402 response.
func (client *httpClientFailureOperations) get402HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.get402HandleError(resp)
}

// get402HandleError handles the Get402 error response.
func (client *httpClientFailureOperations) get402HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get403 - Return 403 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Get403(ctx context.Context) (*http.Response, error) {
	req, err := client.get403CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.get403HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// get403CreateRequest creates the Get403 request.
func (client *httpClientFailureOperations) get403CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/403"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// get403HandleResponse handles the Get403 response.
func (client *httpClientFailureOperations) get403HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.get403HandleError(resp)
}

// get403HandleError handles the Get403 error response.
func (client *httpClientFailureOperations) get403HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get411 - Return 411 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Get411(ctx context.Context) (*http.Response, error) {
	req, err := client.get411CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.get411HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// get411CreateRequest creates the Get411 request.
func (client *httpClientFailureOperations) get411CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/411"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// get411HandleResponse handles the Get411 response.
func (client *httpClientFailureOperations) get411HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.get411HandleError(resp)
}

// get411HandleError handles the Get411 error response.
func (client *httpClientFailureOperations) get411HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get412 - Return 412 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Get412(ctx context.Context) (*http.Response, error) {
	req, err := client.get412CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.get412HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// get412CreateRequest creates the Get412 request.
func (client *httpClientFailureOperations) get412CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/412"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// get412HandleResponse handles the Get412 response.
func (client *httpClientFailureOperations) get412HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.get412HandleError(resp)
}

// get412HandleError handles the Get412 error response.
func (client *httpClientFailureOperations) get412HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get416 - Return 416 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Get416(ctx context.Context) (*http.Response, error) {
	req, err := client.get416CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.get416HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// get416CreateRequest creates the Get416 request.
func (client *httpClientFailureOperations) get416CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/416"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// get416HandleResponse handles the Get416 response.
func (client *httpClientFailureOperations) get416HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.get416HandleError(resp)
}

// get416HandleError handles the Get416 error response.
func (client *httpClientFailureOperations) get416HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Head400 - Return 400 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Head400(ctx context.Context) (*http.Response, error) {
	req, err := client.head400CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.head400HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// head400CreateRequest creates the Head400 request.
func (client *httpClientFailureOperations) head400CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/400"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodHead, *u)
	return req, nil
}

// head400HandleResponse handles the Head400 response.
func (client *httpClientFailureOperations) head400HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.head400HandleError(resp)
}

// head400HandleError handles the Head400 error response.
func (client *httpClientFailureOperations) head400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Head401 - Return 401 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Head401(ctx context.Context) (*http.Response, error) {
	req, err := client.head401CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.head401HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// head401CreateRequest creates the Head401 request.
func (client *httpClientFailureOperations) head401CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/401"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodHead, *u)
	return req, nil
}

// head401HandleResponse handles the Head401 response.
func (client *httpClientFailureOperations) head401HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.head401HandleError(resp)
}

// head401HandleError handles the Head401 error response.
func (client *httpClientFailureOperations) head401HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Head410 - Return 410 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Head410(ctx context.Context) (*http.Response, error) {
	req, err := client.head410CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.head410HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// head410CreateRequest creates the Head410 request.
func (client *httpClientFailureOperations) head410CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/410"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodHead, *u)
	return req, nil
}

// head410HandleResponse handles the Head410 response.
func (client *httpClientFailureOperations) head410HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.head410HandleError(resp)
}

// head410HandleError handles the Head410 error response.
func (client *httpClientFailureOperations) head410HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Head429 - Return 429 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Head429(ctx context.Context) (*http.Response, error) {
	req, err := client.head429CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.head429HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// head429CreateRequest creates the Head429 request.
func (client *httpClientFailureOperations) head429CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/429"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodHead, *u)
	return req, nil
}

// head429HandleResponse handles the Head429 response.
func (client *httpClientFailureOperations) head429HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.head429HandleError(resp)
}

// head429HandleError handles the Head429 error response.
func (client *httpClientFailureOperations) head429HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Options400 - Return 400 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Options400(ctx context.Context) (*http.Response, error) {
	req, err := client.options400CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.options400HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// options400CreateRequest creates the Options400 request.
func (client *httpClientFailureOperations) options400CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/400"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodOptions, *u)
	return req, nil
}

// options400HandleResponse handles the Options400 response.
func (client *httpClientFailureOperations) options400HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.options400HandleError(resp)
}

// options400HandleError handles the Options400 error response.
func (client *httpClientFailureOperations) options400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Options403 - Return 403 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Options403(ctx context.Context) (*http.Response, error) {
	req, err := client.options403CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.options403HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// options403CreateRequest creates the Options403 request.
func (client *httpClientFailureOperations) options403CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/403"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodOptions, *u)
	return req, nil
}

// options403HandleResponse handles the Options403 response.
func (client *httpClientFailureOperations) options403HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.options403HandleError(resp)
}

// options403HandleError handles the Options403 error response.
func (client *httpClientFailureOperations) options403HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Options412 - Return 412 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Options412(ctx context.Context) (*http.Response, error) {
	req, err := client.options412CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.options412HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// options412CreateRequest creates the Options412 request.
func (client *httpClientFailureOperations) options412CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/412"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodOptions, *u)
	return req, nil
}

// options412HandleResponse handles the Options412 response.
func (client *httpClientFailureOperations) options412HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.options412HandleError(resp)
}

// options412HandleError handles the Options412 error response.
func (client *httpClientFailureOperations) options412HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Patch400 - Return 400 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Patch400(ctx context.Context) (*http.Response, error) {
	req, err := client.patch400CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.patch400HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// patch400CreateRequest creates the Patch400 request.
func (client *httpClientFailureOperations) patch400CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/400"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(true)
}

// patch400HandleResponse handles the Patch400 response.
func (client *httpClientFailureOperations) patch400HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.patch400HandleError(resp)
}

// patch400HandleError handles the Patch400 error response.
func (client *httpClientFailureOperations) patch400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Patch405 - Return 405 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Patch405(ctx context.Context) (*http.Response, error) {
	req, err := client.patch405CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.patch405HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// patch405CreateRequest creates the Patch405 request.
func (client *httpClientFailureOperations) patch405CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/405"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(true)
}

// patch405HandleResponse handles the Patch405 response.
func (client *httpClientFailureOperations) patch405HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.patch405HandleError(resp)
}

// patch405HandleError handles the Patch405 error response.
func (client *httpClientFailureOperations) patch405HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Patch414 - Return 414 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Patch414(ctx context.Context) (*http.Response, error) {
	req, err := client.patch414CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.patch414HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// patch414CreateRequest creates the Patch414 request.
func (client *httpClientFailureOperations) patch414CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/414"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(true)
}

// patch414HandleResponse handles the Patch414 response.
func (client *httpClientFailureOperations) patch414HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.patch414HandleError(resp)
}

// patch414HandleError handles the Patch414 error response.
func (client *httpClientFailureOperations) patch414HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post400 - Return 400 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Post400(ctx context.Context) (*http.Response, error) {
	req, err := client.post400CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.post400HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// post400CreateRequest creates the Post400 request.
func (client *httpClientFailureOperations) post400CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/400"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, req.MarshalAsJSON(true)
}

// post400HandleResponse handles the Post400 response.
func (client *httpClientFailureOperations) post400HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.post400HandleError(resp)
}

// post400HandleError handles the Post400 error response.
func (client *httpClientFailureOperations) post400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post406 - Return 406 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Post406(ctx context.Context) (*http.Response, error) {
	req, err := client.post406CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.post406HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// post406CreateRequest creates the Post406 request.
func (client *httpClientFailureOperations) post406CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/406"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, req.MarshalAsJSON(true)
}

// post406HandleResponse handles the Post406 response.
func (client *httpClientFailureOperations) post406HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.post406HandleError(resp)
}

// post406HandleError handles the Post406 error response.
func (client *httpClientFailureOperations) post406HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post415 - Return 415 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Post415(ctx context.Context) (*http.Response, error) {
	req, err := client.post415CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.post415HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// post415CreateRequest creates the Post415 request.
func (client *httpClientFailureOperations) post415CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/415"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, req.MarshalAsJSON(true)
}

// post415HandleResponse handles the Post415 response.
func (client *httpClientFailureOperations) post415HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.post415HandleError(resp)
}

// post415HandleError handles the Post415 error response.
func (client *httpClientFailureOperations) post415HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put400 - Return 400 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Put400(ctx context.Context) (*http.Response, error) {
	req, err := client.put400CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.put400HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// put400CreateRequest creates the Put400 request.
func (client *httpClientFailureOperations) put400CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/400"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(true)
}

// put400HandleResponse handles the Put400 response.
func (client *httpClientFailureOperations) put400HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.put400HandleError(resp)
}

// put400HandleError handles the Put400 error response.
func (client *httpClientFailureOperations) put400HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put404 - Return 404 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Put404(ctx context.Context) (*http.Response, error) {
	req, err := client.put404CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.put404HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// put404CreateRequest creates the Put404 request.
func (client *httpClientFailureOperations) put404CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/404"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(true)
}

// put404HandleResponse handles the Put404 response.
func (client *httpClientFailureOperations) put404HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.put404HandleError(resp)
}

// put404HandleError handles the Put404 error response.
func (client *httpClientFailureOperations) put404HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put409 - Return 409 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Put409(ctx context.Context) (*http.Response, error) {
	req, err := client.put409CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.put409HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// put409CreateRequest creates the Put409 request.
func (client *httpClientFailureOperations) put409CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/409"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(true)
}

// put409HandleResponse handles the Put409 response.
func (client *httpClientFailureOperations) put409HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.put409HandleError(resp)
}

// put409HandleError handles the Put409 error response.
func (client *httpClientFailureOperations) put409HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put413 - Return 413 status code - should be represented in the client as an error
func (client *httpClientFailureOperations) Put413(ctx context.Context) (*http.Response, error) {
	req, err := client.put413CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.put413HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// put413CreateRequest creates the Put413 request.
func (client *httpClientFailureOperations) put413CreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/http/failure/client/413"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(true)
}

// put413HandleResponse handles the Put413 response.
func (client *httpClientFailureOperations) put413HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.put413HandleError(resp)
}

// put413HandleError handles the Put413 error response.
func (client *httpClientFailureOperations) put413HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
