//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azalias

import "time"

// AliasesCreateResponse - The response model for the Alias Create API for the case when the alias was successfully created.
type AliasesCreateResponse struct {
	// READ-ONLY; The id for the alias.
	AliasID *string `json:"aliasId,omitempty" azure:"ro"`

	// READ-ONLY; The created timestamp for the alias.
	CreatedTimestamp *string `json:"createdTimestamp,omitempty" azure:"ro"`

	// READ-ONLY; The id for the creator data item that this alias references (could be null if the alias has not been assigned).
	CreatorDataItemID *string `json:"creatorDataItemId,omitempty" azure:"ro"`

	// READ-ONLY; The timestamp of the last time the alias was assigned.
	LastUpdatedTimestamp *string `json:"lastUpdatedTimestamp,omitempty" azure:"ro"`
}

// ErrorResponse - An error happened.
type ErrorResponse struct {
	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// GeoJSONFeature - A valid GeoJSON Feature object type. Please refer to RFC 7946 [https://tools.ietf.org/html/rfc7946#section-3.2]
// for details.
type GeoJSONFeature struct {
	// REQUIRED; Specifies the GeoJSON type. Must be one of the nine valid GeoJSON object types - Point, MultiPoint, LineString,
	// MultiLineString, Polygon, MultiPolygon, GeometryCollection, Feature and
	// FeatureCollection.
	Type *GeoJSONObjectType `json:"type,omitempty"`

	// The type of the feature. The value depends on the data model the current feature is part of. Some data models may have
	// an empty value.
	FeatureType *string `json:"featureType,omitempty"`

	// Identifier for the feature.
	ID *string `json:"id,omitempty"`

	// Properties can contain any additional metadata about the Feature. Value can be any JSON object or a JSON null value
	Properties interface{} `json:"properties,omitempty"`
}

// GetGeoJSONObject implements the GeoJSONObjectClassification interface for type GeoJSONFeature.
func (g *GeoJSONFeature) GetGeoJSONObject() *GeoJSONObject {
	return &GeoJSONObject{
		Type: g.Type,
		ID:   g.ID,
	}
}

type GeoJSONFeatureData struct {
	// The type of the feature. The value depends on the data model the current feature is part of. Some data models may have
	// an empty value.
	FeatureType *string `json:"featureType,omitempty"`

	// Identifier for the feature.
	ID *string `json:"id,omitempty"`

	// Properties can contain any additional metadata about the Feature. Value can be any JSON object or a JSON null value
	Properties interface{} `json:"properties,omitempty"`
}

// GeoJSONObjectClassification provides polymorphic access to related types.
// Call the interface's GetGeoJSONObject() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *GeoJSONFeature, *GeoJSONObject
type GeoJSONObjectClassification interface {
	// GetGeoJSONObject returns the GeoJSONObject content of the underlying type.
	GetGeoJSONObject() *GeoJSONObject
}

// GeoJSONObject - A valid GeoJSON object. Please refer to RFC 7946 [https://tools.ietf.org/html/rfc7946#section-3] for details.
type GeoJSONObject struct {
	// REQUIRED; Specifies the GeoJSON type. Must be one of the nine valid GeoJSON object types - Point, MultiPoint, LineString,
	// MultiLineString, Polygon, MultiPolygon, GeometryCollection, Feature and
	// FeatureCollection.
	Type *GeoJSONObjectType `json:"type,omitempty"`

	// Identifier for the feature.
	ID *string `json:"id,omitempty"`
}

// GetGeoJSONObject implements the GeoJSONObjectClassification interface for type GeoJSONObject.
func (g *GeoJSONObject) GetGeoJSONObject() *GeoJSONObject { return g }

// GeoJSONObjectNamedCollection - A named collection of GeoJSON object
type GeoJSONObjectNamedCollection struct {
	// Name of the collection
	CollectionName *string `json:"collectionName,omitempty"`

	// Dictionary of
	Objects map[string]GeoJSONObjectClassification `json:"objects,omitempty"`
}

// ListItem - Detailed information for the alias.
type ListItem struct {
	// READ-ONLY; The id for the alias.
	AliasID *string `json:"aliasId,omitempty" azure:"ro"`

	// READ-ONLY; The created timestamp for the alias.
	CreatedTimestamp *string `json:"createdTimestamp,omitempty" azure:"ro"`

	// READ-ONLY; The id for the creator data item that this alias references (could be null if the alias has not been assigned).
	CreatorDataItemID *string `json:"creatorDataItemId,omitempty" azure:"ro"`

	// READ-ONLY; The timestamp of the last time the alias was assigned.
	LastUpdatedTimestamp *string `json:"lastUpdatedTimestamp,omitempty" azure:"ro"`
}

// ListResponse - The response model for the List API. Returns a list of all the previously created aliases.
type ListResponse struct {
	// READ-ONLY; A list of all the previously created aliases.
	Aliases []*ListItem `json:"aliases,omitempty" azure:"ro"`

	// READ-ONLY; If present, the location of the next page of data.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

type ParameterMetadataValue struct {
	// a JSON object
	Value interface{} `json:"value,omitempty"`
}

// ParameterValuesValue - The value of a parameter.
type ParameterValuesValue struct {
	// The value of the parameter.
	Value interface{} `json:"value,omitempty"`
}

type PolicyAssignmentProperties struct {
	// The display name of the policy assignment.
	DisplayName *string `json:"displayName,omitempty"`

	// Key-value pairs of extra info.
	Metadata map[string]*ParameterMetadataValue `json:"metadata,omitempty"`

	// The parameter values for the assigned policy rule. The keys are the parameter names.
	Parameters map[string]*ParameterValuesValue `json:"parameters,omitempty"`
}

// ScheduleCreateOrUpdateProperties - The parameters supplied to the create or update schedule operation.
type ScheduleCreateOrUpdateProperties struct {
	// A list of all the previously created aliases.
	Aliases []*string `json:"aliases,omitempty"`

	// Gets or sets the description of the schedule.
	Description *string `json:"description,omitempty"`

	// Gets or sets the interval of the schedule.
	Interval interface{} `json:"interval,omitempty"`

	// Gets or sets the start time of the schedule.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// clientCreateOptions contains the optional parameters for the client.Create method.
type clientCreateOptions struct {
	GroupBy []SomethingCount
	// The unique id that references the assigned data item to be aliased.
	AssignedID *float32
	// The unique id that references a creator data item to be aliased.
	CreatorID *int32
}

// clientGetScriptOptions contains the optional parameters for the client.GetScript method.
type clientGetScriptOptions struct {
	// placeholder for future optional parameters
}

// clientListOptions contains the optional parameters for the client.List method.
type clientListOptions struct {
	GroupBy []LogMetricsGroupBy
}

// clientPolicyAssignmentOptions contains the optional parameters for the client.PolicyAssignment method.
type clientPolicyAssignmentOptions struct {
	// placeholder for future optional parameters
}
