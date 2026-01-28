// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"

	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/cl"
)

var (
	instanceValidateMethod = slip.Method{
		Name: ":validate",
		Doc: &slip.FuncDoc{
			Name: ":validate",
			Args: []*slip.DocArg{
				{Name: "&optional"},
				{
					Name: "on-error",
					Type: "function",
					Text: "A function to call on validation failures.",
				},
			},
			Return: "nil",
			Text: `__:validate__ validates an instance against the type validation rules.
If an _on-error_ function is provided it should return one of __:continue__, __:reject__,
or __:raise__.
  __:continue__ indicates validation should continue and the error ignored.
  __:reject__ indicates validation should continue but validation should fail after completion.
  __:raise__ indicates validation should stop and an error raised.
  __nil__ indicates any validation failure should raise an error.


If validation is successful then __t__ is returned otherwise __nil__ is returned,`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceValidate{}}},
	}
)

func initInstanceValidate() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := instanceValidate{Function: slip.Function{Name: "instance-validate", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "instance-validate",
			Args: []*slip.DocArg{
				{
					Name: "instance",
					Type: "instance",
					Text: "An instance.",
				},
				{Name: "&optional"},
				{
					Name: "on-error",
					Type: "function",
					Text: "A function to call on validation failures.",
				},
			},
			Return: "boolean",
			Text: `__instance-validate__ validates an instance against the type validation rules.
If an _on-error_ function is provided it should return one of __:continue__, __:reject__,
or __:raise__.
  __:continue__ indicates validation should continue and the error ignored.
  __:reject__ indicates validation should continue but validation should fail after completion.
  __:raise__ indicates validation should stop and an error raised.
  __nil__ indicates any validation failure should raise an error.


If validation is successful then __t__ is returned otherwise __nil__ is returned,`,
		}, &Pkg)
}

type instanceValidate struct {
	slip.Function
}

// Call the the function with the arguments provvalidateed.
func (f *instanceValidate) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 2)
	inst, ok := args[0].(*Instance)
	if !ok {
		slip.TypePanic(s, depth, "instance", args[0], "fhir:instance")
	}
	var onErr slip.Caller
	if 1 < len(args) {
		onErr = cl.ResolveToCaller(s, args[1], depth)
	}

	fmt.Printf("*** %v %s\n", onErr, inst)

	// TBD run valdation on data

	return slip.True
}
