// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

var (
	instanceIDMethod = slip.Method{
		Name: ":id",
		Doc: &slip.FuncDoc{
			Name:   ":id",
			Return: "fixnum",
			Text: `__:id__ returns the instance __id__ which should not be confused with
the FHIR _id_ property.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceID{}}},
	}
)

func initInstanceID() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := instanceID{Function: slip.Function{Name: "instance-id", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "instance-id",
			Args: []*slip.DocArg{
				{
					Name: "instance",
					Type: "instance",
					Text: "An instance.",
				},
			},
			Return: "fixnum",
			Text: `__instance-id__ returns the instance __id__ which should not be confused with
the FHIR _id_ property.`,
		}, &Pkg)
}

type instanceID struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *instanceID) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)
	inst, ok := args[0].(slip.Instance)
	if !ok {
		slip.TypePanic(s, depth, "instance", args[0], "instance")
	}
	return slip.Fixnum(inst.ID())
}
