// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package main

import (
	"context"
	"fmt"
	"generatortests/autorest/generated/azurereportgroup"
	"generatortests/autorest/generated/reportgroup"
	"sort"
)

// generate autorest test server coverage report
func main() {
	vanillaClient := reportgroup.NewDefaultClient(nil)
	vanillaReport, err := vanillaClient.AutoRestReportServiceOperations().GetReport(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	azureClient := azurereportgroup.NewDefaultClient(nil)
	azureReport, err := azureClient.AutoRestReportServiceForAzureOperations().GetReport(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	notRun := []string{}
	for key, val := range *vanillaReport.Value {
		if val <= 0 {
			notRun = append(notRun, fmt.Sprintf("GENERAL: %s", key))
		}
	}
	for key, val := range *azureReport.Value {
		if val <= 0 {
			notRun = append(notRun, fmt.Sprintf("AZURE: %s", key))
		}
	}
	sort.Strings(notRun)
	for _, s := range notRun {
		fmt.Printf("Not Run: %s\n", s)
	}
	passed := len(*vanillaReport.Value) + len(*azureReport.Value) - len(notRun)
	fmt.Printf("\nReport:	Passed(%d)  Not Run(%d)\n", passed, len(notRun))
}
