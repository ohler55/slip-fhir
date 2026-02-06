// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
)

var propValidPMethod = slip.Method{
	Name: ":valid-p",
	Doc: &slip.FuncDoc{
		Name: ":valid-p",
		Args: []*slip.DocArg{
			{
				Name: "value",
				Type: "object",
				Text: "A value to validate.",
			},
			{Name: "&optional"},
			{
				Name: "on-error",
				Type: "function",
				Text: "A function to call on validation failures.",
			},
		},
		Return: "list",
		Text: `__:valid-p__ validates a _value_ for comformance with the property.
If an _on-error_ function is provided it should return __t__ to abort validation or __nil__
to continue. The _on-error_ function is expected to take 3 arguments, a __bag-path__ which
identifies a property, value of the property, and a message describing a validation error.


If validation is successful then __t__ is returned otherwise __nil__ is returned.`,
	},
	Combinations: []*slip.Combination{{From: &blankType, Primary: &propertyValidP{}}},
}

func initPropertyValidP() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := propertyValidP{Function: slip.Function{Name: "property-valid-p", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "property-valid-p",
			Args: []*slip.DocArg{
				{
					Name: "property",
					Type: "fhir:property",
					Text: "The property to return the use for validation.",
				},
				{
					Name: "value",
					Type: "object",
					Text: "A value to validate.",
				},
				{Name: "&optional"},
				{
					Name: "on-error",
					Type: "function",
					Text: "A function to call on validation failures.",
				},
			},
			Return: "boolean",
			Text: `__property-valid-p__ validates a _value_ for comformance with the property.
If an _on-error_ function is provided it should return __t__ to abort validation or __nil__
to continue. The _on-error_ function is expected to take 3 arguments, a __bag-path__ which
identifies a property, value of the property, and a message describing a validation error.


If validation is successful then __t__ is returned otherwise __nil__ is returned.`,
		}, &Pkg)
}

type propertyValidP struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *propertyValidP) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 2, 3)

	prop, ok := args[0].(*Property)
	if !ok {
		slip.TypePanic(s, depth, "property", args[0], "fhir:property")
	}
	var (
		data  any
		onErr slip.Caller
	)
	switch ta := args[1].(type) {
	case *Instance:
		data = ta.data
	case *flavors.Instance: // bag
		if bag.Flavor() != ta.Type {
			slip.TypePanic(s, depth, "value", ta, "bag", "fhir:instance", "primitive type")
		}
		data = ta.Any
	default:
		data = slip.Simplify(ta)
	}
	if 2 < len(args) {
		onErr = resolveToOnError(s, args[2], depth)
	}
	var onErrFn OnErrorFunc
	if onErr == nil {
		onErrFn = func(p jp.Expr, v any, message string) bool {
			return true
		}
	} else {
		onErrFn = func(p jp.Expr, v any, message string) bool {
			return onErr.Call(s, slip.List{bag.Path(p), objectify(v), slip.String(message)}, depth) == slip.True
		}
	}
	if prop.validateValue(data, onErrFn) {
		return nil
	}
	return slip.True
}
