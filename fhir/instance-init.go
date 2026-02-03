// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var (
	instanceInitMethod = slip.Method{
		Name: ":init",
		Doc: &slip.FuncDoc{
			Name: ":init",
			Args: []*slip.DocArg{
				{Name: "&key"},
				{
					Name: "data",
					Type: "bag",
					Text: "The initial data for and instance.",
				},
				{
					Name: "on-error",
					Type: "function",
					Text: "The function to call on validation errors.",
				},
				{
					Name: "no-validation",
					Type: "boolean",
					Text: "If non-nil then no validation is performed in the _data_.",
				},
			},
			Return: "nil",
			Text: `__:init__ is called after an instance is created with __make-instance__.
It can not be called directly.`,
		},
		// This method is never called so any caller is fine for the After
		// combination is fine. The method exists for documentation only.
		Combinations: []*slip.Combination{{From: &blankType, After: &instanceTypeCaller{}}},
	}
)
