// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var propNameMethod = slip.Method{
	Name: ":name",
	Doc: &slip.FuncDoc{
		Name:   ":name",
		Args:   []*slip.DocArg{},
		Return: "string",
		Text:   `__:name__ returns the property name.`,
	},
	Combinations: []*slip.Combination{{From: &blankType, Primary: &propertyName{}}},
}

func initPropertyName() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := propertyName{Function: slip.Function{Name: "property-name", Args: args}}
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

type propertyName struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *propertyName) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)

	prop, ok := args[0].(*Prop)
	if !ok {
		slip.TypePanic(s, depth, "property", args[0], "fhir:property")
	}
	return slip.String(prop.name)
}
