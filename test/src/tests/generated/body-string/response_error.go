package stringgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"net"
	"net/http"
)

// if you want to provide custom error handling set this variable to your constructor function
var responseErrorFactory func(cause error, response *http.Response, description string) error

// NewResponseError creates an error object that implements the error interface.
func NewResponseError(cause error, response *http.Response, description string) error {
	if responseErrorFactory != nil {
		return responseErrorFactory(cause, response, description)
	}
	return &responseError{
		ErrorNode:   pipeline.ErrorNode{}.Initialize(cause, 3),
		response:    response,
		description: description,
	}
}

// responseError is the internal struct that implements the public ResponseError interface.
type responseError struct {
	pipeline.ErrorNode // This is embedded so that responseError "inherits" Error, Temporary, Timeout, and Cause
	response           *http.Response
	description        string
}

// Error implements the error interface's Error method to return a string representation of the error.
func (e *responseError) Error() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "===== RESPONSE ERROR (Code=%v) =====\n", e.response.StatusCode)
	fmt.Fprintf(b, "Status=%s, Description: %s\n", e.response.Status, e.description)
	s := b.String()
	return e.ErrorNode.Error(s)
}

// Response implements the ResponseError interface's method to return the HTTP response.
func (e *responseError) Response() *http.Response {
	return e.response
}

// RFC7807 PROBLEM ------------------------------------------------------------------------------------
// RFC7807Problem ... This type can be publicly embedded in another type that wants to add additional members.
/*type RFC7807Problem struct {
	// Mandatory: A (relative) URI reference identifying the problem type (it MAY refer to human-readable documentation).
	typeURI string // Should default to "about:blank"
	// Optional: Short, human-readable summary (maybe localized).
	title string
	// Optional: HTTP status code generated by the origin server
	status int
	// Optional: Human-readable explanation for this problem occurance.
	// Should help client correct the problem. Clients should NOT parse this string.
	detail string
	// Optional: A (relative) URI identifying this specific problem occurence (it may or may not be dereferenced).
	instance string
}
// NewRFC7807Problem ...
func NewRFC7807Problem(typeURI string, status int, titleFormat string, a ...interface{}) error {
	return &RFC7807Problem{
		typeURI: typeURI,
		status:  status,
		title:   fmt.Sprintf(titleFormat, a...),
	}
}
// Error returns the error information as a string.
func (e *RFC7807Problem) Error() string {
	return e.title
}
// TypeURI ...
func (e *RFC7807Problem) TypeURI() string {
	if e.typeURI == "" {
		e.typeURI = "about:blank"
	}
	return e.typeURI
}
// Members ...
func (e *RFC7807Problem) Members() (status int, title, detail, instance string) {
	return e.status, e.title, e.detail, e.instance
}*/
