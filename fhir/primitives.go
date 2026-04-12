// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"sort"

	"github.com/ohler55/slip"
)

func initPrimitives() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := Primitives{Function: slip.Function{Name: "fhir-primitives", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "fhir-primitives",
			Args: []*slip.DocArg{
				{
					Name: "package",
					Type: "package|string|symbol",
					Text: `Package designator.`,
				},
			},
			Return: "list",
			Text:   `__fhir-primitives__ returns a list of all the FHIR primitive types in the _package_.`,
		}, &Pkg)
}

// Primitives represents the fhir-primitives function.
type Primitives struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *Primitives) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)
	p := slip.PackageFromArg(args[0])

	classes := p.AllClasses()
	sort.Slice(classes, func(i, j int) bool {
		return classes[i].Name() < classes[j].Name()
	})
	var types slip.List
	baseClass := p.FindClass("Base")
	for _, class := range classes {
		if !class.Inherits(baseClass) && baseClass != class {
			types = append(types, class)
		}
	}
	return types
}
