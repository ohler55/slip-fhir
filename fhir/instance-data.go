// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
)

var (
	instanceDataMethod = slip.Method{
		Name: ":data",
		Doc: &slip.FuncDoc{
			Name:   ":data",
			Return: "bag",
			Text: `__:data__ returns the instance __data__ as a __bag__. Modifying the
returned bag will modify the instance's data.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceData{}}},
	}
)

func initInstanceData() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := instanceData{Function: slip.Function{Name: "instance-data", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "instance-data",
			Args: []*slip.DocArg{
				{
					Name: "instance",
					Type: "instance",
					Text: "An instance.",
				},
			},
			Return: "fixnum",
			Text: `__instance-data__ returns the instance __data__ as a __bag__. Modifying the
returned bag will modify the instance's data.`,
		}, &Pkg)
}

type instanceData struct {
	slip.Function
}

// Call the the function with the arguments provdataed.
func (f *instanceData) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)
	inst, ok := args[0].(*Instance)
	if !ok {
		slip.TypePanic(s, depth, "instance", args[0], "fhir:instance")
	}
	bg := bag.Flavor().MakeInstance().(*flavors.Instance)
	bg.Any = inst.data

	return bg
}
