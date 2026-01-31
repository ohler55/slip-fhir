// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

func initTypeProperty() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := TypeProperty{Function: slip.Function{Name: "type-property", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "type-property",
			Args: []*slip.DocArg{
				{
					Name: "type",
					Type: "fhir:type",
					Text: "The type to get the property of.",
				},
				{
					Name: "name",
					Type: "string",
					Text: "The name of the property to return.",
				},
			},
			Return: "fhir:property",
			Text: `__type-property__ returns the property with the _name_ or nil
if the type does not include that property.`,
		}, &Pkg)
}

// TypeProperty represents the type-property function.
type TypeProperty struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *TypeProperty) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 2, 2)
	obj := args[0]
	var prop *Prop
top:
	switch to := obj.(type) {
	case *Type:
		prop = to.findProp(slip.MustBeString(args[1], "name"))
	case slip.Symbol:
		if c := slip.FindClass(string(to)); c != nil {
			obj = c
			goto top
		}
	default:
		slip.TypePanic(s, depth, "type", to, "fhir type")
	}
	return prop
}
