//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package headergroup

import "time"

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// HeaderClientCustomRequestIDOptions contains the optional parameters for the HeaderClient.CustomRequestID method.
type HeaderClientCustomRequestIDOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamBoolOptions contains the optional parameters for the HeaderClient.ParamBool method.
type HeaderClientParamBoolOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamByteOptions contains the optional parameters for the HeaderClient.ParamByte method.
type HeaderClientParamByteOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamDateOptions contains the optional parameters for the HeaderClient.ParamDate method.
type HeaderClientParamDateOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamDatetimeOptions contains the optional parameters for the HeaderClient.ParamDatetime method.
type HeaderClientParamDatetimeOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamDatetimeRFC1123Options contains the optional parameters for the HeaderClient.ParamDatetimeRFC1123 method.
type HeaderClientParamDatetimeRFC1123Options struct {
	// Send a post request with header values "Wed, 01 Jan 2010 12:34:56 GMT" or "Mon, 01 Jan 0001 00:00:00 GMT"
	Value *time.Time
}

// HeaderClientParamDoubleOptions contains the optional parameters for the HeaderClient.ParamDouble method.
type HeaderClientParamDoubleOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamDurationOptions contains the optional parameters for the HeaderClient.ParamDuration method.
type HeaderClientParamDurationOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamEnumOptions contains the optional parameters for the HeaderClient.ParamEnum method.
type HeaderClientParamEnumOptions struct {
	// Send a post request with header values 'GREY'
	Value *GreyscaleColors
}

// HeaderClientParamExistingKeyOptions contains the optional parameters for the HeaderClient.ParamExistingKey method.
type HeaderClientParamExistingKeyOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamFloatOptions contains the optional parameters for the HeaderClient.ParamFloat method.
type HeaderClientParamFloatOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamIntegerOptions contains the optional parameters for the HeaderClient.ParamInteger method.
type HeaderClientParamIntegerOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamLongOptions contains the optional parameters for the HeaderClient.ParamLong method.
type HeaderClientParamLongOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamProtectedKeyOptions contains the optional parameters for the HeaderClient.ParamProtectedKey method.
type HeaderClientParamProtectedKeyOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientParamStringOptions contains the optional parameters for the HeaderClient.ParamString method.
type HeaderClientParamStringOptions struct {
	// Send a post request with header values "The quick brown fox jumps over the lazy dog" or null or ""
	Value *string
}

// HeaderClientResponseBoolOptions contains the optional parameters for the HeaderClient.ResponseBool method.
type HeaderClientResponseBoolOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseByteOptions contains the optional parameters for the HeaderClient.ResponseByte method.
type HeaderClientResponseByteOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseDateOptions contains the optional parameters for the HeaderClient.ResponseDate method.
type HeaderClientResponseDateOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseDatetimeOptions contains the optional parameters for the HeaderClient.ResponseDatetime method.
type HeaderClientResponseDatetimeOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseDatetimeRFC1123Options contains the optional parameters for the HeaderClient.ResponseDatetimeRFC1123
// method.
type HeaderClientResponseDatetimeRFC1123Options struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseDoubleOptions contains the optional parameters for the HeaderClient.ResponseDouble method.
type HeaderClientResponseDoubleOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseDurationOptions contains the optional parameters for the HeaderClient.ResponseDuration method.
type HeaderClientResponseDurationOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseEnumOptions contains the optional parameters for the HeaderClient.ResponseEnum method.
type HeaderClientResponseEnumOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseExistingKeyOptions contains the optional parameters for the HeaderClient.ResponseExistingKey method.
type HeaderClientResponseExistingKeyOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseFloatOptions contains the optional parameters for the HeaderClient.ResponseFloat method.
type HeaderClientResponseFloatOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseIntegerOptions contains the optional parameters for the HeaderClient.ResponseInteger method.
type HeaderClientResponseIntegerOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseLongOptions contains the optional parameters for the HeaderClient.ResponseLong method.
type HeaderClientResponseLongOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseProtectedKeyOptions contains the optional parameters for the HeaderClient.ResponseProtectedKey method.
type HeaderClientResponseProtectedKeyOptions struct {
	// placeholder for future optional parameters
}

// HeaderClientResponseStringOptions contains the optional parameters for the HeaderClient.ResponseString method.
type HeaderClientResponseStringOptions struct {
	// placeholder for future optional parameters
}
