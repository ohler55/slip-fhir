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
)

func init() {
	Pkg.Initialize(map[string]*slip.VarVal{
		"*fhir*": {
			Val:    &Pkg,
			Const:  true,
			Export: true,
			Doc:    `The fhir package.`,
		},
	})

	initTypes()

	Pkg.Initialize(nil, &PrimitiveType{}) // lock
	slip.AddPackage(&Pkg)
	slip.UserPkg.Use(&Pkg)
}

func initTypes() {
	// slip/pkg/clos is needed so make sure it gets inited first with an
	// import.
	f5 := sen.MustParse(fhir5JSON)
	var primitives []*PrimitiveType

	for _, pa := range jp.C("primitives").W().Get(f5) {
		pt := PrimitiveType{
			name:        alt.String(jp.C("name").First(pa)),
			description: alt.String(jp.C("description").First(pa)),
			pattern:     alt.String(jp.C("pattern").First(pa)),
			parent:      alt.String(jp.C("parent").First(pa)),
			pkg:         &Pkg,
		}
		slip.RegisterClass(strings.ToLower(pt.name), &pt)
		primitives = append(primitives, &pt)
	}
	for _, pt := range primitives {
		pt.init()
	}
}
