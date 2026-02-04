// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"strings"

	"github.com/ohler55/slip"
)

var (
	propOperationHandledPMethod = slip.Method{
		Name: ":operation-handled-p",
		Doc: &slip.FuncDoc{
			Name: ":operation-handled-p",
			Args: []*slip.DocArg{
				{
					Name: "method",
					Type: "keyword",
					Text: "Symbol to check.",
				},
			},
			Return: "boolean",
			Text:   `__:operation-handled-p__ returns _t_ if the prop handles the method and _nil_ otherwise.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &propOperationHandledPCaller{}}},
	}
)

type propOperationHandledPCaller struct{}

func (caller propOperationHandledPCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &propOperationHandledPMethod, args, 2, 2)
	method, ok := args[1].(slip.Symbol)
	if !ok {
		slip.TypePanic(s, depth, ":operation-handled-p method", args[1], "symbol")
	}
	if propMethods[strings.ToLower(string(method))] != nil {
		return slip.True
	}
	return nil
}
