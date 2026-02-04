// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var (
	propEqualMethod = slip.Method{
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
			Text:   `__:equal__ returns _t_ if the prop is of the same flavor as _other_ and has the same content.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &propEqualCaller{}}},
	}
)

type propEqualCaller struct{}

func (caller propEqualCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &propEqualMethod, args, 2, 2)
	p := args[0].(*Prop)
	if p.Equal(args[1]) {
		return slip.True
	}
	return nil
}
