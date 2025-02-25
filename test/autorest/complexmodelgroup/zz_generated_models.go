//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package complexmodelgroup

type CatalogArray struct {
	// Array of products
	ProductArray []*Product `json:"productArray,omitempty"`
}

type CatalogArrayOfDictionary struct {
	// Array of dictionary of products
	ProductArrayOfDictionary []map[string]*Product `json:"productArrayOfDictionary,omitempty"`
}

type CatalogDictionary struct {
	// Dictionary of products
	ProductDictionary map[string]*Product `json:"productDictionary,omitempty"`
}

type CatalogDictionaryOfArray struct {
	// Dictionary of Array of product
	ProductDictionaryOfArray map[string][]*Product `json:"productDictionaryOfArray,omitempty"`
}

// ComplexModelClientCreateOptions contains the optional parameters for the ComplexModelClient.Create method.
type ComplexModelClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ComplexModelClientListOptions contains the optional parameters for the ComplexModelClient.List method.
type ComplexModelClientListOptions struct {
	// placeholder for future optional parameters
}

// ComplexModelClientUpdateOptions contains the optional parameters for the ComplexModelClient.Update method.
type ComplexModelClientUpdateOptions struct {
	// placeholder for future optional parameters
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Product - The product documentation.
type Product struct {
	// Capacity of product. For example, 4 people.
	Capacity *string `json:"capacity,omitempty"`

	// Description of product.
	Description *string `json:"description,omitempty"`

	// Display name of product.
	DisplayName *string `json:"display_name,omitempty"`

	// Image URL representing the product.
	Image *string `json:"image,omitempty"`

	// Unique identifier representing a specific product for a given latitude & longitude. For example, uberX in San Francisco
	// will have a different product_id than uberX in Los Angeles.
	ProductID *string `json:"product_id,omitempty"`
}
