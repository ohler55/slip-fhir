// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"io"

	"github.com/ohler55/slip"
)

var (
	instanceDescribeMethod = slip.Method{
		Name: ":describe",
		Doc: &slip.FuncDoc{
			Name: ":describe",
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
			Text:   `__:describe__ writes a description of the instance to _output-stream_.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceDescribeCaller{}}},
	}
)

type instanceDescribeCaller struct{}

func (caller instanceDescribeCaller) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, &instanceDescribeMethod, args, 1, 2)
	inst := args[0].(*Instance)

	ansi := s.Get("*print-ansi*") != nil
	right := int(s.Get("*print-right-margin*").(slip.Fixnum))
	b := inst.Describe(nil, 0, right, ansi)
	so := s.Get("*standard-output*")
	ss, _ := so.(slip.Stream)
	w := so.(io.Writer)
	if 1 < len(args) {
		var ok bool
		ss, _ = args[1].(slip.Stream)
		if w, ok = args[1].(io.Writer); !ok {
			slip.TypePanic(s, depth, ":describe output-stream", args[1], "output-stream")
		}
	}
	if _, err := w.Write(b); err != nil {
		slip.StreamPanic(s, depth, ss, "%s", err)
	}

	return nil
}
