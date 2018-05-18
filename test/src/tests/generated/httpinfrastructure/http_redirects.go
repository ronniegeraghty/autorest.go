package httpinfrastructuregroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io"
	"io/ioutil"
	"net/http"
)

// HTTPRedirectsClient is the test Infrastructure for AutoRest
type HTTPRedirectsClient struct {
	ManagementClient
}

// NewHTTPRedirectsClient creates an instance of the HTTPRedirectsClient client.
func NewHTTPRedirectsClient(p pipeline.Pipeline) HTTPRedirectsClient {
	return HTTPRedirectsClient{NewManagementClient(p)}
}

// Delete307 delete redirected with 307, resulting in a 200 after redirect
//
// booleanValue is simple boolean value true
func (client HTTPRedirectsClient) Delete307(ctx context.Context, booleanValue *bool) (*HTTPRedirectsDelete307Response, error) {
	req, err := client.delete307Preparer(booleanValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.delete307Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsDelete307Response), err
}

// delete307Preparer prepares the Delete307 request.
func (client HTTPRedirectsClient) delete307Preparer(booleanValue *bool) (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/307"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(booleanValue)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// delete307Responder handles the response to the Delete307 request.
func (client HTTPRedirectsClient) delete307Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusTemporaryRedirect)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsDelete307Response{rawResponse: resp.Response()}, err
}

// Get300 return 300 status code and redirect to /http/success/200
func (client HTTPRedirectsClient) Get300(ctx context.Context) (*Get300Response, error) {
	req, err := client.get300Preparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.get300Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Get300Response), err
}

// get300Preparer prepares the Get300 request.
func (client HTTPRedirectsClient) get300Preparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/300"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// get300Responder handles the response to the Get300 request.
func (client HTTPRedirectsClient) get300Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusMultipleChoices)
	if resp == nil {
		return nil, err
	}
	result := &Get300Response{rawResponse: resp.Response()}
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

// Get301 return 301 status code and redirect to /http/success/200
func (client HTTPRedirectsClient) Get301(ctx context.Context) (*HTTPRedirectsGet301Response, error) {
	req, err := client.get301Preparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.get301Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsGet301Response), err
}

// get301Preparer prepares the Get301 request.
func (client HTTPRedirectsClient) get301Preparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/301"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// get301Responder handles the response to the Get301 request.
func (client HTTPRedirectsClient) get301Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusMovedPermanently)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsGet301Response{rawResponse: resp.Response()}, err
}

// Get302 return 302 status code and redirect to /http/success/200
func (client HTTPRedirectsClient) Get302(ctx context.Context) (*HTTPRedirectsGet302Response, error) {
	req, err := client.get302Preparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.get302Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsGet302Response), err
}

// get302Preparer prepares the Get302 request.
func (client HTTPRedirectsClient) get302Preparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/302"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// get302Responder handles the response to the Get302 request.
func (client HTTPRedirectsClient) get302Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusFound)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsGet302Response{rawResponse: resp.Response()}, err
}

// Get307 redirect get with 307, resulting in a 200 success
func (client HTTPRedirectsClient) Get307(ctx context.Context) (*HTTPRedirectsGet307Response, error) {
	req, err := client.get307Preparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.get307Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsGet307Response), err
}

// get307Preparer prepares the Get307 request.
func (client HTTPRedirectsClient) get307Preparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/307"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// get307Responder handles the response to the Get307 request.
func (client HTTPRedirectsClient) get307Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusTemporaryRedirect)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsGet307Response{rawResponse: resp.Response()}, err
}

// Head300 return 300 status code and redirect to /http/success/200
func (client HTTPRedirectsClient) Head300(ctx context.Context) (*HTTPRedirectsHead300Response, error) {
	req, err := client.head300Preparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.head300Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsHead300Response), err
}

// head300Preparer prepares the Head300 request.
func (client HTTPRedirectsClient) head300Preparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/300"
	req, err := pipeline.NewRequest("HEAD", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// head300Responder handles the response to the Head300 request.
func (client HTTPRedirectsClient) head300Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusMultipleChoices)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsHead300Response{rawResponse: resp.Response()}, err
}

// Head301 return 301 status code and redirect to /http/success/200
func (client HTTPRedirectsClient) Head301(ctx context.Context) (*HTTPRedirectsHead301Response, error) {
	req, err := client.head301Preparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.head301Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsHead301Response), err
}

// head301Preparer prepares the Head301 request.
func (client HTTPRedirectsClient) head301Preparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/301"
	req, err := pipeline.NewRequest("HEAD", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// head301Responder handles the response to the Head301 request.
func (client HTTPRedirectsClient) head301Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusMovedPermanently)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsHead301Response{rawResponse: resp.Response()}, err
}

// Head302 return 302 status code and redirect to /http/success/200
func (client HTTPRedirectsClient) Head302(ctx context.Context) (*HTTPRedirectsHead302Response, error) {
	req, err := client.head302Preparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.head302Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsHead302Response), err
}

// head302Preparer prepares the Head302 request.
func (client HTTPRedirectsClient) head302Preparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/302"
	req, err := pipeline.NewRequest("HEAD", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// head302Responder handles the response to the Head302 request.
func (client HTTPRedirectsClient) head302Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusFound)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsHead302Response{rawResponse: resp.Response()}, err
}

// Head307 redirect with 307, resulting in a 200 success
func (client HTTPRedirectsClient) Head307(ctx context.Context) (*HTTPRedirectsHead307Response, error) {
	req, err := client.head307Preparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.head307Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsHead307Response), err
}

// head307Preparer prepares the Head307 request.
func (client HTTPRedirectsClient) head307Preparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/307"
	req, err := pipeline.NewRequest("HEAD", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// head307Responder handles the response to the Head307 request.
func (client HTTPRedirectsClient) head307Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusTemporaryRedirect)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsHead307Response{rawResponse: resp.Response()}, err
}

// Patch302 patch true Boolean value in request returns 302.  This request should not be automatically redirected, but
// should return the received 302 to the caller for evaluation
//
// booleanValue is simple boolean value true
func (client HTTPRedirectsClient) Patch302(ctx context.Context, booleanValue *bool) (*HTTPRedirectsPatch302Response, error) {
	req, err := client.patch302Preparer(booleanValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.patch302Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsPatch302Response), err
}

// patch302Preparer prepares the Patch302 request.
func (client HTTPRedirectsClient) patch302Preparer(booleanValue *bool) (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/302"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(booleanValue)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// patch302Responder handles the response to the Patch302 request.
func (client HTTPRedirectsClient) patch302Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusFound)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsPatch302Response{rawResponse: resp.Response()}, err
}

// Patch307 patch redirected with 307, resulting in a 200 after redirect
//
// booleanValue is simple boolean value true
func (client HTTPRedirectsClient) Patch307(ctx context.Context, booleanValue *bool) (*HTTPRedirectsPatch307Response, error) {
	req, err := client.patch307Preparer(booleanValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.patch307Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsPatch307Response), err
}

// patch307Preparer prepares the Patch307 request.
func (client HTTPRedirectsClient) patch307Preparer(booleanValue *bool) (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/307"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(booleanValue)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// patch307Responder handles the response to the Patch307 request.
func (client HTTPRedirectsClient) patch307Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusTemporaryRedirect)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsPatch307Response{rawResponse: resp.Response()}, err
}

// Post303 post true Boolean value in request returns 303.  This request should be automatically redirected usign a
// get, ultimately returning a 200 status code
//
// booleanValue is simple boolean value true
func (client HTTPRedirectsClient) Post303(ctx context.Context, booleanValue *bool) (*HTTPRedirectsPost303Response, error) {
	req, err := client.post303Preparer(booleanValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.post303Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsPost303Response), err
}

// post303Preparer prepares the Post303 request.
func (client HTTPRedirectsClient) post303Preparer(booleanValue *bool) (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/303"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(booleanValue)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// post303Responder handles the response to the Post303 request.
func (client HTTPRedirectsClient) post303Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusSeeOther)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsPost303Response{rawResponse: resp.Response()}, err
}

// Post307 post redirected with 307, resulting in a 200 after redirect
//
// booleanValue is simple boolean value true
func (client HTTPRedirectsClient) Post307(ctx context.Context, booleanValue *bool) (*HTTPRedirectsPost307Response, error) {
	req, err := client.post307Preparer(booleanValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.post307Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsPost307Response), err
}

// post307Preparer prepares the Post307 request.
func (client HTTPRedirectsClient) post307Preparer(booleanValue *bool) (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/307"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(booleanValue)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// post307Responder handles the response to the Post307 request.
func (client HTTPRedirectsClient) post307Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusTemporaryRedirect)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsPost307Response{rawResponse: resp.Response()}, err
}

// Put301 put true Boolean value in request returns 301.  This request should not be automatically redirected, but
// should return the received 301 to the caller for evaluation
//
// booleanValue is simple boolean value true
func (client HTTPRedirectsClient) Put301(ctx context.Context, booleanValue *bool) (*HTTPRedirectsPut301Response, error) {
	req, err := client.put301Preparer(booleanValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.put301Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsPut301Response), err
}

// put301Preparer prepares the Put301 request.
func (client HTTPRedirectsClient) put301Preparer(booleanValue *bool) (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/301"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(booleanValue)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// put301Responder handles the response to the Put301 request.
func (client HTTPRedirectsClient) put301Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusMovedPermanently)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsPut301Response{rawResponse: resp.Response()}, err
}

// Put307 put redirected with 307, resulting in a 200 after redirect
//
// booleanValue is simple boolean value true
func (client HTTPRedirectsClient) Put307(ctx context.Context, booleanValue *bool) (*HTTPRedirectsPut307Response, error) {
	req, err := client.put307Preparer(booleanValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.put307Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*HTTPRedirectsPut307Response), err
}

// put307Preparer prepares the Put307 request.
func (client HTTPRedirectsClient) put307Preparer(booleanValue *bool) (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/redirect/307"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(booleanValue)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// put307Responder handles the response to the Put307 request.
func (client HTTPRedirectsClient) put307Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusTemporaryRedirect)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &HTTPRedirectsPut307Response{rawResponse: resp.Response()}, err
}
