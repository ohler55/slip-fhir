// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"sort"
	"strings"

	"github.com/ohler55/slip"
)

func initDatatypes() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := Datatypes{Function: slip.Function{Name: "fhir-datatypes", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "fhir-datatypes",
			Args: []*slip.DocArg{
				{
					Name: "package",
					Type: "package|string|symbol",
					Text: `Package designator.`,
				},
			},
			Return: "list",
			Text:   `__fhir-datatypes__ returns a list of all the FHIR complex datatypes in the _package_.`,
		}, &Pkg)
}

// Datatypes represents the fhir-datatypes function.
type Datatypes struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *Datatypes) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 1)
	p := slip.PackageFromArg(args[0])

	classes := p.AllClasses()
	sort.Slice(classes, func(i, j int) bool {
		return classes[i].Name() < classes[j].Name()
	})
	var types slip.List
	resClass := p.FindClass("Resource")
	elementClass := p.FindClass("Element")
	for _, class := range classes {
		if class.Inherits(elementClass) && !class.Inherits(resClass) && !strings.Contains(class.Name(), "_") {
			types = append(types, class)
		}
	}
	return types
}
