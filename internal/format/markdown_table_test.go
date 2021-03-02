/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
	"github.com/terraform-docs/terraform-docs/internal/testutil"
)

func TestMarkdownTable(t *testing.T) {
	tests := map[string]struct {
		settings print.Settings
		options  terraform.Options
	}{
		// Base
		"Base": {
			settings: testutil.WithSections(),
			options:  terraform.Options{},
		},
		"Empty": {
			settings: testutil.WithSections(),
			options: terraform.Options{
				Path: "empty",
			},
		},
		"HideAll": {
			settings: print.Settings{},
			options: terraform.Options{
				ShowHeader:     false, // Since we don't show the header, the file won't be loaded at all
				HeaderFromFile: "bad.tf",
			},
		},

		// Settings
		"WithRequired": {
			settings: testutil.WithSections(
				print.Settings{
					ShowRequired: true,
				},
			),
			options: terraform.Options{},
		},
		"WithAnchor": {
			settings: testutil.WithSections(
				print.Settings{
					ShowAnchor: true,
				},
			),
			options: terraform.Options{},
		},
		"EscapeCharacters": {
			settings: testutil.WithSections(
				print.Settings{
					EscapeCharacters: true,
				},
			),
			options: terraform.Options{},
		},
		"IndentationOfFour": {
			settings: testutil.WithSections(
				print.Settings{
					IndentLevel: 4,
				},
			),
			options: terraform.Options{},
		},
		"OutputValues": {
			settings: print.Settings{
				ShowOutputs:     true,
				OutputValues:    true,
				ShowSensitivity: true,
			},
			options: terraform.Options{
				OutputValues:     true,
				OutputValuesPath: "output_values.json",
			},
		},
		"OutputValuesNoSensitivity": {
			settings: print.Settings{
				ShowOutputs:     true,
				OutputValues:    true,
				ShowSensitivity: false,
			},
			options: terraform.Options{
				OutputValues:     true,
				OutputValuesPath: "output_values.json",
			},
		},

		// Only section
		"OnlyHeader": {
			settings: print.Settings{ShowHeader: true},
			options:  terraform.Options{},
		},
		"OnlyInputs": {
			settings: print.Settings{ShowInputs: true},
			options:  terraform.Options{},
		},
		"OnlyOutputs": {
			settings: print.Settings{ShowOutputs: true},
			options:  terraform.Options{},
		},
		"OnlyModulecalls": {
			settings: print.Settings{ShowModuleCalls: true},
			options:  terraform.Options{},
		},
		"OnlyProviders": {
			settings: print.Settings{ShowProviders: true},
			options:  terraform.Options{},
		},
		"OnlyRequirements": {
			settings: print.Settings{ShowRequirements: true},
			options:  terraform.Options{},
		},
		"OnlyResources": {
			settings: print.Settings{ShowResources: true},
			options:  terraform.Options{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			expected, err := testutil.GetExpected("markdown", "table-"+name)
			assert.Nil(err)

			options, err := terraform.NewOptions().With(&tt.options)
			assert.Nil(err)

			module, err := testutil.GetModule(options)
			assert.Nil(err)

			printer := NewMarkdownTable(&tt.settings)
			actual, err := printer.Print(module, &tt.settings)

			assert.Nil(err)
			assert.Equal(expected, actual)
		})
	}
}
