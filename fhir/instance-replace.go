// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
)

var (
	instanceReplaceMethod = slip.Method{
		Name: ":replace",
		Doc: &slip.FuncDoc{
			Name: ":replace",
			Args: []*slip.DocArg{
				{
					Name: "data",
					Type: "bag",
					Text: "The replacement data.",
				},
			},
			Return: "object",
			Text: `__:replace__ replaces all the data in the instance with the content of the _data_.
The data is not duplicated leaving the bag and and instance both referring to the same data.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceReplace{}}},
	}
)

func initInstanceReplace() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := instanceReplace{Function: slip.Function{Name: "instance-replace", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "instance-replace",
			Args: []*slip.DocArg{
				{
					Name: "instance",
					Type: "instance",
					Text: "An instance.",
				},
				{
					Name: "data",
					Type: "bag",
					Text: "The replacement data.",
				},
			},
			Return: "nil",
			Text: `__instance-replace__ replaces all the data in the instance with the content of the _data_.
The data is not duplicated leaving the bag and and instance both referring to the same data.`,
		}, &Pkg)
}

type instanceReplace struct {
	slip.Function
}

// Call the the function with the arguments provreplaceed.
func (f *instanceReplace) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 2, 2)
	inst, ok := args[0].(*Instance)
	if !ok {
		slip.TypePanic(s, depth, "instance", args[0], "fhir:instance")
	}
	var bg *flavors.Instance
	if bg, ok = args[1].(*flavors.Instance); !ok || bag.Flavor() != bg.Type {
		slip.TypePanic(s, depth, "data", args[1], "bag")
	}
	var data map[string]any
	data, ok = bg.Any.(map[string]any)
	if !ok {
		slip.TypePanic(s, depth, "data", bg, "complex type")
	}

	// TBD run valdation on data

	inst.data = data

	return nil
}
