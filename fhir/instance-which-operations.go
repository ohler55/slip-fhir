// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var (
	instanceWhichOperationsMethod = slip.Method{
		Name: ":which-operations",
		Doc: &slip.FuncDoc{
			Name:   ":which-operations",
			Return: "list",
			Text:   `__:which-operations__ returns a list of all the methods the instance handles.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceWhichOperationsCaller{}}},
	}
)

type instanceWhichOperationsCaller struct{}

func (caller instanceWhichOperationsCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &instanceWhichOperationsMethod, args, 1, 1)

	return typeMethodNames()
}

func (caller instanceWhichOperationsCaller) FuncDocs() *slip.FuncDoc {
	return instanceWhichOperationsMethod.Doc
}
