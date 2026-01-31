// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

func initTypeProperties() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := TypeProperties{Function: slip.Function{Name: "type-properties", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "type-properties",
			Args: []*slip.DocArg{
				{
					Name: "type",
					Type: "fhir:type",
					Text: "The type to get the properties of.",
				},
			},
			Return: "list",
			Text:   `__type-properties__ returns a list of a types properties.`,
		}, &Pkg)
}

// TypeProperties represents the type-properties function.
type TypeProperties struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *TypeProperties) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)
	obj := args[0]
	var props slip.List
top:
	switch to := obj.(type) {
	case *Type:
		props = to.propList()
	case slip.Symbol:
		if c := slip.FindClass(string(to)); c != nil {
			obj = c
			goto top
		}
	default:
		slip.TypePanic(s, depth, "type", to, "fhir type")
	}
	return props
}
