// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
)

var (
	instanceValidPMethod = slip.Method{
		Name: ":valid-p",
		Doc: &slip.FuncDoc{
			Name: ":valid-p",
			Args: []*slip.DocArg{
				{Name: "&optional"},
				{
					Name: "on-error",
					Type: "function",
					Text: "A function to call on validation failures.",
				},
			},
			Return: "nil",
			Text: `__:valid-p__ validates an instance against the type validation rules.
If an _on-error_ function is provided it should return __t__ to abort validation or __nil__
to continue. The _on-error_ function is expected to take 3 arguments, a __bag-path__ which
identifies a property, value of the property, and a message describing a validation error.


If validation is successful then __t__ is returned otherwise __nil__ is returned,`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceValidPCaller{}}},
	}
)

type instanceValidPCaller struct{}

func (caller instanceValidPCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &instanceValidPMethod, args, 1, 2)
	inst := args[0].(*Instance)
	var onErrFn OnErrorFunc
	if 1 < len(args) {
		onErr := resolveToOnError(s, args[1], depth)
		onErrFn = func(p jp.Expr, v any, message string) bool {
			return onErr.Call(s, slip.List{bag.Path(p), objectify(v), slip.String(message)}, depth) == slip.True
		}
	} else {
		onErrFn = func(p jp.Expr, v any, message string) bool {
			panic(fmt.Sprintf("Value at %s, %s: %s.", p, pretty.SEN(v), message))
		}
	}
	if inst.class.Validate(inst.data, onErrFn) {
		return nil
	}
	return slip.True
}
