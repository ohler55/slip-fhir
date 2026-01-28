// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
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
If an _on-error_ function is provided it should return __t__ to abort validation or __nil__
to continue. The _on-error_ function is expected to take 3 arguments, a __bag-path__ which
identifies a property, value of the property, and a message describing a validation error.


If validation is successful then __nil__ is returned otherwise __t__ is returned,`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceValidateCaller{}}},
	}
)

type instanceValidateCaller struct{}

func (caller instanceValidateCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &instanceValidateMethod, args, 1, 1)
	inst := args[0].(slip.Instance)

	return inst.Class()
}

func (caller instanceValidateCaller) FuncDocs() *slip.FuncDoc {
	return instanceValidateMethod.Doc
}
