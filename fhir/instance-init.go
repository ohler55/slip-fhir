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
					Name: "no-error",
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
		Combinations: []*slip.Combination{{From: &blankType, After: &instanceInitCaller{}}},
	}
)

type instanceInitCaller struct{}

func (caller instanceInitCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	return nil
}

func (caller instanceInitCaller) FuncDocs() *slip.FuncDoc {
	return instanceInitMethod.Doc
}
