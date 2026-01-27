// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var (
	instanceEqualMethod = slip.Method{
		Name: ":equal",
		Doc: &slip.FuncDoc{
			Name: ":equal",
			Args: []*slip.DocArg{
				{
					Name: "other",
					Type: "object",
					Text: "Other object to compare to _self_.",
				},
			},
			Return: "boolean",
			Text:   `__:equal__ returns _t_ if the instance is of the same flavor as _other_ and has the same content.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceEqualCaller{}}},
	}
)

type instanceEqualCaller struct{}

func (caller instanceEqualCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &instanceEqualMethod, args, 2, 2)
	inst := args[0].(slip.Instance)
	if inst.Equal(args[1]) {
		return slip.True
	}
	return nil
}

func (caller instanceEqualCaller) FuncDocs() *slip.FuncDoc {
	return instanceEqualMethod.Doc
}
