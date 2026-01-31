// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

func initPropertyCardinality() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := PropertyCardinality{Function: slip.Function{Name: "property-cardinality", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "property-cardinality",
			Args: []*slip.DocArg{
				{
					Name: "property",
					Type: "fhir:property",
					Text: `The property to return the cardinality of.`,
				},
			},
			Return: "fixnum,fixnum|nil",
			Text: `__property-cardinality__ returns the cardinality of _property_ as a minimum
and maximum pair of values. If the maximum is unlimited then the returned maximum will be __nil__.`,
		}, &Pkg)
}

// PropertyCardinality represents the property-cardinality function.
type PropertyCardinality struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *PropertyCardinality) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)

	prop, ok := args[0].(*Prop)
	if !ok {
		slip.TypePanic(s, depth, "property", args[0], "fhir:property")
	}
	card := slip.Values{slip.Fixnum(0), slip.Fixnum(1)}
	if prop.required {
		card[0] = slip.Fixnum(1)
	}
	if prop.array {
		card[1] = nil
	}
	return card
}
