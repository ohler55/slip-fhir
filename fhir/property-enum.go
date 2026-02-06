// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var propEnumMethod = slip.Method{
	Name: ":enum",
	Doc: &slip.FuncDoc{
		Name:   ":enum",
		Args:   []*slip.DocArg{},
		Return: "list",
		Text:   `__:enum__ returns the property enum or nil is there are none.`,
	},
	Combinations: []*slip.Combination{{From: &blankType, Primary: &propertyEnum{}}},
}

func initPropertyEnum() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := propertyEnum{Function: slip.Function{Name: "property-enum", Args: args}}
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

type propertyEnum struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *propertyEnum) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)

	prop, ok := args[0].(*Property)
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
