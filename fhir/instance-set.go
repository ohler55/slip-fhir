// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var (
	instanceSetMethod = slip.Method{
		Name: ":set",
		Doc: &slip.FuncDoc{
			Name: ":set",
			Args: []*slip.DocArg{
				{
					Name: "property",
					Type: "string",
					Text: "The name of the property to set.",
				},
				{
					Name: "value",
					Type: "object",
					Text: "A replacement value for the property.",
				},
			},
			Return: "object",
			Text:   `__:set__ sets the _property_ with _value_ if the _value_ passes validation.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceSet{}}},
	}
)

func initInstanceSet() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := instanceSet{Function: slip.Function{Name: "instance-set", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "instance-set",
			Args: []*slip.DocArg{
				{
					Name: "instance",
					Type: "instance",
					Text: "An instance.",
				},
				{
					Name: "property",
					Type: "string",
					Text: "The name of the property to set.",
				},
				{
					Name: "value",
					Type: "object",
					Text: "A replacement value for the property.",
				},
			},
			Return: "nil",
			Text:   `__instance-set__ sets the _property_ with _value_ if the _value_ passes validation.`,
		}, &Pkg)
}

type instanceSet struct {
	slip.Function
}

// Call the the function with the arguments provseted.
func (f *instanceSet) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 3, 3)
	inst, ok := args[0].(*Instance)
	if !ok {
		slip.TypePanic(s, depth, "instance", args[0], "fhir:instance")
	}
	pname := slip.MustBeString(args[1], "property")

	inst.SetSlotValue(slip.Symbol(pname), args[2])

	return nil
}
