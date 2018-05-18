package report

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/xml"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
)

type responder func(resp pipeline.Response) (result pipeline.Response, err error)

// ResponderPolicyFactory is a Factory capable of creating a responder pipeline.
type responderPolicyFactory struct {
	responder responder
}

// New creates a responder policy factory.
func (arpf responderPolicyFactory) New(next pipeline.Policy, po *pipeline.PolicyOptions) pipeline.Policy {
	return responderPolicy{next: next, responder: arpf.responder}
}

type responderPolicy struct {
	next      pipeline.Policy
	responder responder
}

// Do sends the request to the service and validates/deserializes the HTTP response.
func (arp responderPolicy) Do(ctx context.Context, request pipeline.Request) (pipeline.Response, error) {
	resp, err := arp.next.Do(ctx, request)
	if err != nil {
		return resp, err
	}
	return arp.responder(resp)
}

// validateResponse checks an HTTP response's status code against a legal set of codes.
// If the response code is not legal, then validateResponse reads all of the response's body
// (containing error information) and returns a response error.
func validateResponse(resp pipeline.Response, successStatusCodes ...int) error {
	if resp == nil {
		return NewResponseError(nil, nil, "nil response")
	}
	responseCode := resp.Response().StatusCode
	for _, i := range successStatusCodes {
		if i == responseCode {
			return nil
		}
	}
	// only close the body in the failure case. in the
	// success case responders will close the body as required.
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return NewResponseError(err, resp.Response(), "failed to read response body")
	}
	// the service code, description and details will be populated during unmarshalling
	responseError := NewResponseError(nil, resp.Response(), resp.Response().Status)
	if len(b) > 0 {
		if err = xml.Unmarshal(b, &responseError); err != nil {
			return NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return responseError
}

// removes any BOM from the byte slice
func removeBOM(b []byte) []byte {
	// UTF8
	return bytes.TrimPrefix(b, []byte("\xef\xbb\xbf"))
}
