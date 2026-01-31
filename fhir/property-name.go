// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

func initPropertyName() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := PropertyName{Function: slip.Function{Name: "property-name", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "property-name",
			Args: []*slip.DocArg{
				{
					Name: "property",
					Type: "fhir:property",
					Text: "The property to return the name of.",
				},
			},
			Return: "string",
			Text:   `__property-name__ returns the _property_ name.`,
		}, &Pkg)
}

// PropertyName represents the property-name function.
type PropertyName struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *PropertyName) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)

	prop, ok := args[0].(*Prop)
	if !ok {
		slip.TypePanic(s, depth, "property", args[0], "fhir:property")
	}
	return slip.String(prop.name)
}
