// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"sort"

	"github.com/ohler55/slip"
)

func initTypes() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := Types{Function: slip.Function{Name: "fhir-types", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "fhir-types",
			Args: []*slip.DocArg{
				{
					Name: "package",
					Type: "package|string|symbol",
					Text: `Package designator.`,
				},
			},
			Return: "list",
			Text:   `__fhir-types__ returns a list of all the FHIR types in the _package_.`,
		}, &Pkg)
}

// Types represents the fhir-types function.
type Types struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *Types) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)
	p := slip.PackageFromArg(args[0])

	classes := p.AllClasses()
	sort.Slice(classes, func(i, j int) bool {
		return classes[i].Name() < classes[j].Name()
	})
	types := make(slip.List, len(classes))
	for i, class := range classes {
		types[i] = class
	}
	return types
}
