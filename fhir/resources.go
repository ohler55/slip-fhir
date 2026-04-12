// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"sort"

	"github.com/ohler55/slip"
)

func initResources() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := Resources{Function: slip.Function{Name: "fhir-resources", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "fhir-resources",
			Args: []*slip.DocArg{
				{
					Name: "package",
					Type: "package|string|symbol",
					Text: `Package designator.`,
				},
			},
			Return: "list",
			Text:   `__fhir-resources__ returns a list of all the FHIR types in the _package_.`,
		}, &Pkg)
}

// Resources represents the fhir-resources function.
type Resources struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *Resources) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)
	p := slip.PackageFromArg(args[0])

	classes := p.AllClasses()
	sort.Slice(classes, func(i, j int) bool {
		return classes[i].Name() < classes[j].Name()
	})
	var types slip.List
	resClass := p.FindClass("Resource")
	for _, class := range classes {
		if class.Inherits(resClass) {
			types = append(types, class)
		}
	}
	return types
}
