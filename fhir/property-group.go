// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

func initPropertyGroup() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := PropertyGroup{Function: slip.Function{Name: "property-group", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "property-group",
			Args: []*slip.DocArg{
				{
					Name: "property",
					Type: "fhir:property",
					Text: `The property to return the group properties if the _property_ has
group members. An example of a property with a groups is 'value[x]'.`,
				},
			},
			Return: "list",
			Text:   `__property-group__ returns the _property_ group or nil if there are none.`,
		}, &Pkg)
}

// PropertyGroup represents the property-group function.
type PropertyGroup struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *PropertyGroup) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)

	prop, ok := args[0].(*Prop)
	if !ok {
		slip.TypePanic(s, depth, "property", args[0], "fhir:property")
	}
	if 0 < len(prop.group) {
		lst := make(slip.List, len(prop.group))
		for i, gp := range prop.group {
			lst[i] = gp
		}
		return lst
	}
	return nil
}
