// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var (
	propWhichOperationsMethod = slip.Method{
		Name: ":which-operations",
		Doc: &slip.FuncDoc{
			Name:   ":which-operations",
			Return: "list",
			Text:   `__:which-operations__ returns a list of all the methods the prop handles.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &propWhichOperationsCaller{}}},
	}
)

type propWhichOperationsCaller struct{}

func (caller propWhichOperationsCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &propWhichOperationsMethod, args, 1, 1)

	return propMethodNames()
}
