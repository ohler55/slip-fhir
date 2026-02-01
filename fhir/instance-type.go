// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var (
	instanceTypeMethod = slip.Method{
		Name: ":type",
		Doc: &slip.FuncDoc{
			Name:   ":type",
			Return: "class",
			Text:   `__:type__ returns the FHIR type of the instance.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceTypeCaller{}}},
	}
	instanceClassMethod = slip.Method{
		Name: ":class",
		Doc: &slip.FuncDoc{
			Name:   ":class",
			Return: "class",
			Text:   `__:class__ returns the FHIR type of the instance.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceTypeCaller{}}},
	}
)

type instanceTypeCaller struct{}

func (caller instanceTypeCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &instanceTypeMethod, args, 1, 1)
	inst := args[0].(slip.Instance)

	return inst.Class()
}
