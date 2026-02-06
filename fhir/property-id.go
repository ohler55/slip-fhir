// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var (
	propIDMethod = slip.Method{
		Name: ":id",
		Doc: &slip.FuncDoc{
			Name:   ":id",
			Args:   []*slip.DocArg{},
			Return: "fixnum",
			Text:   `__:id__ returns the instance __id__.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &prodIDCaller{}}},
	}
)

type prodIDCaller struct{}

func (caller prodIDCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &propIDMethod, args, 1, 1)
	p := args[0].(*Property)

	return slip.Fixnum(p.ID())
}
