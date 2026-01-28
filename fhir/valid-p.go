// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
)

func initValidP() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := validP{Function: slip.Function{Name: "valid-p", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "valid-p",
			Args: []*slip.DocArg{
				{
					Name: "value",
					Type: "object",
					Text: "An instance or primitive to validates against a type.",
				},
				{Name: "&key"},
				{
					Name: "type",
					Type: "fhir:type",
					Text: "A fhir type. Required for all but a fhir:Instance.",
				},
				{
					Name: "on-error",
					Type: "function",
					Text: "A function to call on validation failures.",
				},
			},
			Return: "boolean",
			Text: `__:validate__ validates an instance against the type validation rules.
If an _on-error_ function is provided it should return __t__ to abort validation or __nil__
to continue. The _on-error_ function is expected to take 3 arguments, a __bag-path__ which
identifies a property, value of the property, and a message describing a validation error.


If validation is successful then __nil__ is returned otherwise __t__ is returned,`,
		}, &Pkg)
}

type validP struct {
	slip.Function
}

// Call the the function with the arguments provvalidateed.
func (f *validP) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 5)
	var (
		data  any
		onErr slip.Caller
		typ   *Type
	)
	switch ta := args[0].(type) {
	case *Instance:
		data = ta.data
		typ = ta.class
	case *flavors.Instance: // bag
		if bag.Flavor() != ta.Type {
			slip.TypePanic(s, depth, "value", ta, "bag", "fhir:instance", "primitive type")
		}
		data = ta.Any
	default:
		data = slip.Simplify(ta)
	}
	args = args[1:]
	if value, has := slip.GetArgsKeyValue(args, slip.Symbol(":type")); has {
	ftype:
		switch tv := value.(type) {
		case *Type:
			typ = tv
		case slip.Symbol:
			value = slip.FindClass(string(tv))
			goto ftype
		default:
			slip.TypePanic(s, depth, ":type", value, "fhir:type", "symbol")
		}
	}
	if value, has := slip.GetArgsKeyValue(args, slip.Symbol(":on-error")); has {
		onErr = resolveToOnError(s, value, depth)
	}
	if typ == nil {
		slip.ErrorPanic(s, depth, "A :type is required if the value is not a fhir:instance.")
	}
	var onErrFn OnErrorFunc
	if onErr == nil {
		onErrFn = func(p jp.Expr, v any, message string) bool {
			panic(fmt.Sprintf("Value at %s, %s: %s.", p, v, message))
		}
	} else {
		fmt.Printf("*** %v\n", onErr)
		onErrFn = func(p jp.Expr, v any, message string) bool {

			// TBD
			panic(fmt.Sprintf("Value at %s, %s: %s.", p, v, message))
		}
	}
	if typ.Validate(data, onErrFn) {
		return slip.True
	}
	return nil
}

func resolveToOnError(s *slip.Scope, fn slip.Object, depth int) (caller slip.Caller) {
	d2 := depth + 1
	var doc *slip.FuncDoc
CallFunc:
	switch tf := fn.(type) {
	case *slip.Lambda:
		caller = tf
		doc = tf.Doc
	case *slip.FuncInfo:
		caller = tf.Create(nil).(slip.Funky).Caller()
		doc = tf.Doc
	case slip.Symbol:
		fn = slip.MustFindFunc(string(tf))
		goto CallFunc
	case slip.List:
		fn = s.Eval(tf, d2)
		goto CallFunc
	default:
		slip.TypePanic(s, depth, "function", tf, "function")
	}
	if doc == nil || len(doc.Args) != 3 {
		slip.TypePanic(s, depth, "function", fn, "function with 3 arguments")
	}
	return
}
