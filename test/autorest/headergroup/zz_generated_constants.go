//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package headergroup

const host = "http://localhost:3000"

type GreyscaleColors string

const (
	GreyscaleColorsWhite GreyscaleColors = "White"
	GreyscaleColorsBlack GreyscaleColors = "black"
	GreyscaleColorsGREY  GreyscaleColors = "GREY"
)

// PossibleGreyscaleColorsValues returns the possible values for the GreyscaleColors const type.
func PossibleGreyscaleColorsValues() []GreyscaleColors {
	return []GreyscaleColors{
		GreyscaleColorsWhite,
		GreyscaleColorsBlack,
		GreyscaleColorsGREY,
	}
}
