// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

func initPropertyType() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := PropertyType{Function: slip.Function{Name: "property-type", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "property-type",
			Args: []*slip.DocArg{
				{
					Name: "property",
					Type: "fhir:property",
					Text: "The property to return the type of.",
				},
			},
			Return: "fhir:type",
			Text:   `__property-type__ returns the _property_ type.`,
		}, &Pkg)
}

// PropertyType represents the property-type function.
type PropertyType struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *PropertyType) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)

	prop, ok := args[0].(*Prop)
	if !ok {
		slip.TypePanic(s, depth, "property", args[0], "fhir:property")
	}
	return prop.ftype
}
