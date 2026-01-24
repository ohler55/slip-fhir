// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	_ "embed"
	"strings"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/slip"
	_ "github.com/ohler55/slip/pkg/clos"
)

//go:embed "fhir5.json"
var fhir5JSON []byte

const (
	bold         = "\x1b[1m"
	colorOff     = "\x1b[m"
	indentSpaces = "                                                                                "
)

var (
	// Pkg is the jet package.
	Pkg = slip.Package{
		Name:      "fhir",
		Nicknames: []string{"fhir"},
		Doc: `The __fhir__ package provides a FHIR client to aid in making requests
to a FHIR server and in handling responses. Classes for each of the FHIR resource and
non-primitive types are defined from a specified FHIR JSON Schema file. Instances of
the fhir classes can be used to build as well as access fields of FHIR datatypes.


The default FHIR version is v5.0.0. The _set-fhir-version_ function can be used to
change the FHIR version which also redefines the fhir classes.`,
		PreSet: slip.DefaultPreSet,
	}
	types []Type
)

func init() {
	Pkg.Initialize(map[string]*slip.VarVal{
		"*fhir*": {
			Val:    &Pkg,
			Const:  true,
			Export: true,
			Doc:    `The fhir package.`,
		},
		"*fhir-version*": {
			Val:    slip.String("unknown"),
			Const:  true,
			Export: true,
			Doc:    `The FHIR version.`,
		},
	})
	// slip/pkg/clos is needed so make sure it gets inited first with an
	// import.
	initTypes(sen.MustParse(fhir5JSON))

	Pkg.Initialize(nil, &PrimitiveType{}) // lock
	slip.AddPackage(&Pkg)
	slip.UserPkg.Use(&Pkg)
}

func initTypes(schema any) {
	if version := alt.String(jp.C("version").First(schema)); 0 < len(version) {
		vv := Pkg.GetVarVal("*fhir-version*")
		vv.Val = slip.String(version)
	}
	initPrimitives(schema)
	loadTypes(jp.C("hierarchy").W().Get(schema))
	loadTypes(jp.C("datatypes").W().Get(schema))
	loadTypes(jp.C("backbones").W().Get(schema))
	loadTypes(jp.C("resources").W().Get(schema))
	initTypeParents()
}

func initPrimitives(schema any) {
	var primitives []*PrimitiveType
	for _, pa := range jp.C("primitives").W().Get(schema) {
		pt := PrimitiveType{
			name:        alt.String(jp.C("name").First(pa)),
			description: alt.String(jp.C("description").First(pa)),
			pattern:     alt.String(jp.C("pattern").First(pa)),
			parent:      alt.String(jp.C("parent").First(pa)),
			pkg:         &Pkg,
		}
		slip.RegisterClass(strings.ToLower(pt.name), &pt)
		primitives = append(primitives, &pt)
		types = append(types, &pt)
	}
	for _, pt := range primitives {
		pt.init()
	}
}

func loadTypes(defs []any) {
	for _, ts := range defs {
		ft := Base{
			name:   alt.String(jp.C("name").First(ts)),
			docs:   alt.String(jp.C("description").First(ts)),
			parent: alt.String(jp.C("parent").First(ts)),
			pkg:    &Pkg,
		}
		for _, ps := range jp.C("properties").W().Get(ts) {
			p := Prop{
				name:     alt.String(jp.C("name").First(ps)),
				docs:     alt.String(jp.C("description").First(ps)),
				typeName: alt.String(jp.C("type").First(ps)),
				required: alt.Bool(jp.C("required").First(ps)),
				array:    alt.Bool(jp.C("array").First(ps)),
			}
			for _, e := range jp.C("enum").W().Get(ps) {
				p.enum = append(p.enum, alt.String(e))
			}
			ft.props = append(ft.props, &p)
		}
		slip.RegisterClass(strings.ToLower(ft.name), &ft)
		types = append(types, &ft)
	}
}

func initTypeParents() {
	for _, ft := range types {
		if base, ok := ft.(*Base); ok {
			base.init()
		}
	}
}
