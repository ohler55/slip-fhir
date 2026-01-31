// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

func initPropertyEnum() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := PropertyEnum{Function: slip.Function{Name: "property-enum", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "property-enum",
			Args: []*slip.DocArg{
				{
					Name: "property",
					Type: "fhir:property",
					Text: "The property to return the enum of.",
				},
			},
			Return: "list",
			Text:   `__property-enum__ returns the _property_ enum or nil if there are none.`,
		}, &Pkg)
}

// PropertyEnum represents the property-enum function.
type PropertyEnum struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *PropertyEnum) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)

	prop, ok := args[0].(*Prop)
	if !ok {
		slip.TypePanic(s, depth, "property", args[0], "fhir:property")
	}
	if 0 < len(prop.enum) {
		lst := make(slip.List, len(prop.enum))
		for i, e := range prop.enum {
			lst[i] = slip.String(e)
		}
		return lst
	}
	return nil
}
