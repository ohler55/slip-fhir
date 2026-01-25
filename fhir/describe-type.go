// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"io"

	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/cl"
)

func initDescribeType() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := DescribeType{Function: slip.Function{Name: "describe-type", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "describe-type",
			Args: []*slip.DocArg{
				{
					Name: "type",
					Type: "fhir-type",
					Text: "The type to describe.",
				},
				{Name: "&optional"},
				{
					Name: "output-stream",
					Type: "output-stream",
					Text: "The stream to write to.",
				},
				{Name: "&key"},
				{
					Name: "full",
					Type: "bool",
					Text: "If true inherited properties are displayed as well as property extensions",
				},
				{
					Name: "stripe-color",
					Type: "string",
					Text: `If not empty, set the background color of every other property to the color provided.`,
				},
			},
			Return: "nil",
			Text:   `__describe-type__ is similar to __describe__ but with a more options for viewing FHIR types.`,
		}, &Pkg)
}

// DescribeType represents the describe-type function.
type DescribeType struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *DescribeType) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 6)
	obj := args[0]
	args = args[1:]
	so := s.Get("*standard-output*")
	ss, _ := so.(slip.Stream)
	w := so.(io.Writer)
	if 0 < len(args) {
		switch ta := args[0].(type) {
		case io.Writer:
			w = ta
			ss, _ = args[0].(slip.Stream)
			args = args[1:]
		case slip.Symbol:
			if len(ta) <= 1 && ta[0] != ':' {
				slip.TypePanic(s, depth, "describe-type output-stream", ta, "output-stream")
			}
			// otherwise a keyword
		default:
			slip.TypePanic(s, depth, "describe-type output-stream", ta, "output-stream")
		}
	}
	var (
		full bool
		bg   string
		b    []byte
	)
	if v, has := slip.GetArgsKeyValue(args, slip.Symbol(":full")); has && v != nil {
		full = true
	}
	if v, has := slip.GetArgsKeyValue(args, slip.Symbol(":stripe-color")); has {
		bg = slip.MustBeString(v, ":stripe-color")
	}
	ansi := s.Get("*print-ansi*") != nil
	right := int(s.Get("*print-right-margin*").(slip.Fixnum))
top:
	switch to := obj.(type) {
	case *Type:
		b = to.describe(b, 0, right, ansi, full, bg)
	case slip.Symbol:
		if c := slip.FindClass(string(to)); c != nil {
			obj = c
			goto top
		}
	default:
		b = cl.AppendDescribe(nil, obj, s, 0, right, ansi)
	}
	if _, err := w.Write(b); err != nil {
		slip.StreamPanic(s, depth, ss, "write failed: %s", err)
	}
	return nil
}
