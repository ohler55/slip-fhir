// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"io"

	"github.com/ohler55/slip"
)

var (
	propPrintSelfMethod = slip.Method{
		Name: ":print-self",
		Doc: &slip.FuncDoc{
			Name: ":print-self",
			Args: []*slip.DocArg{
				{Name: "&optional"},
				{
					Name:    "output-stream",
					Type:    "output-stream",
					Text:    "output-stream to write to.",
					Default: slip.Symbol("*standard-output*"),
				},
			},
			Return: "nil",
			Text:   `__:print-self__ writes the prop string representation to the _stream_.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &propPrintSelfCaller{}}},
	}
)

type propPrintSelfCaller struct{}

func (caller propPrintSelfCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &propPrintSelfMethod, args, 1, 2)
	p := args[0].(*Prop)

	so := s.Get("*standard-output*")
	ss, _ := so.(slip.Stream)
	w := so.(io.Writer)
	if 1 < len(args) {
		var ok bool
		ss, _ = args[1].(slip.Stream)
		if w, ok = args[1].(io.Writer); !ok {
			slip.TypePanic(s, depth, ":print-self output-stream", args[1], "output-stream")
		}
	}
	if _, err := w.Write(p.Append(nil)); err != nil {
		slip.StreamPanic(s, depth, ss, "%s", err)
	}
	return nil
}
