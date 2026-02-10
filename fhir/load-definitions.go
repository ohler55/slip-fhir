// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"os"
	"strings"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/slip"
)

func initLoadDefinitions() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := LoadDefinitions{Function: slip.Function{Name: "load-definitions", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "load-definitions",
			Args: []*slip.DocArg{
				{
					Name: "filename",
					Type: "string",
					Text: "The path to a FHIR definitions file.",
				},
				{
					Name: "package",
					Type: "symbol|string|package",
					Text: "The name of the package to define the FHIR types in.",
				},
			},
			Return: "list",
			Text: `__load-definitions__ defines fhir:Types for the definitions in the _filename_.
Typically the defines types are placed in a separate package from the __fhire__ package. This
allows multiple FHIR versions to be loaded at the same time. If the named _package_ does not
exist it is created. The __use-package__ function for __common-lisp-user__ will not be called
to avoid conflicts but the user can make that call or add that to the Slip _config.lisp_ file.
A list of all the newly defined types is returned.`,
		}, &Pkg)
}

// LoadDefinitions represents the load-definitions function.
type LoadDefinitions struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *LoadDefinitions) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 2, 2)

	filename := slip.MustBeString(args[0], "filename")
	var p *slip.Package
	arg1 := args[1]
pkgArg:
	switch ta := arg1.(type) {
	case *slip.Package:
		p = ta
	case slip.String:
		arg1 = slip.Symbol(ta)
		goto pkgArg
	case slip.Symbol:
		if p = slip.FindPackage(string(ta)); p == nil {
			p = slip.DefPackage(string(ta), []string{}, "")
		}
	default:
		slip.TypePanic(s, depth, "package", ta, "package", "string", "symbol")
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		slip.FilePanic(s, depth, slip.String(filename), "read failed. %s", err)
	}
	defs := sen.MustParse(content)

	types := defineTypes(defs, p)

	lst := make(slip.List, len(types))
	for i, t := range types {
		lst[i] = t
	}
	return lst
}

func defineTypes(schema any, p *slip.Package) []*Type {
	if version := alt.String(jp.C("version").First(schema)); 0 < len(version) {
		p.DefConst("*version*", slip.String(version),
			"The FHIR version in the package.").Export = true
	}
	var types []*Type

	types = definePrimitives(types, schema, p)
	types = defineComplexTypes(types, jp.C("hierarchy").W().Get(schema), p)
	types = defineComplexTypes(types, jp.C("datatypes").W().Get(schema), p)
	types = defineComplexTypes(types, jp.C("backbones").W().Get(schema), p)
	types = defineComplexTypes(types, jp.C("resources").W().Get(schema), p)
	initTypeParents(types)

	return types
}

func definePrimitives(types []*Type, schema any, p *slip.Package) []*Type {
	var primitives []*Type
	for _, pa := range jp.C("primitives").W().Get(schema) {
		pt := Type{
			name:        alt.String(jp.C("name").First(pa)),
			description: alt.String(jp.C("description").First(pa)),
			pkg:         p,
			parent:      alt.String(jp.C("parent").First(pa)),
			pattern:     alt.String(jp.C("pattern").First(pa)),
		}
		p.RegisterClass(strings.ToLower(pt.name), &pt)
		primitives = append(primitives, &pt)
		types = append(types, &pt)
	}
	for _, pt := range primitives {
		pt.init()
	}
	return types
}

func defineComplexTypes(types []*Type, defs []any, p *slip.Package) []*Type {
	for _, ts := range defs {
		ft := Type{
			name:        alt.String(jp.C("name").First(ts)),
			description: alt.String(jp.C("description").First(ts)),
			pkg:         p,
			parent:      alt.String(jp.C("parent").First(ts)),
		}
		for _, ps := range jp.C("properties").W().Get(ts) {
			ft.props = append(ft.props, NewProp(ps))
		}
		p.RegisterClass(strings.ToLower(ft.name), &ft)
		types = append(types, &ft)
	}
	return types
}

func initTypeParents(types []*Type) {
	for _, ft := range types {
		ft.init()
	}
}
