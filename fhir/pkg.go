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
	bold          = "\x1b[1m"
	bgBlack       = "\x1b[0;40m"
	bgRed         = "\x1b[0;41m"
	bgGreen       = "\x1b[0;42m"
	bgYellow      = "\x1b[0;43m"
	bgBlue        = "\x1b[0;44m"
	bgPurple      = "\x1b[0;45m"
	bgCyan        = "\x1b[0;46m"
	bgGray        = "\x1b[0;47m"
	bgDarkGray    = "\x1b[0;100m"
	bgLightRed    = "\x1b[0;101m"
	bgLightGreen  = "\x1b[0;102m"
	bgLightYellow = "\x1b[0;103m"
	bgLightBlue   = "\x1b[0;104m"
	bgLightPurple = "\x1b[0;105m"
	bgLightCyan   = "\x1b[0;106m"
	bgWhite       = "\x1b[0;107m"

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
	types []Validator
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
		"bg-black": {
			Val:    slip.String("\x1b[0;40m"),
			Const:  true,
			Export: true,
			Doc:    "Black background",
		},
		"bg-red": {
			Val:    slip.String("\x1b[0;41m"),
			Const:  true,
			Export: true,
			Doc:    "Red background",
		},
		"bg-green": {
			Val:    slip.String("\x1b[0;42m"),
			Const:  true,
			Export: true,
			Doc:    "Green background",
		},
		"bg-yellow": {
			Val:    slip.String("\x1b[0;43m"),
			Const:  true,
			Export: true,
			Doc:    "Yellow background",
		},
		"bg-blue": {
			Val:    slip.String("\x1b[0;44m"),
			Const:  true,
			Export: true,
			Doc:    "Blue background",
		},
		"bg-purple": {
			Val:    slip.String("\x1b[0;45m"),
			Const:  true,
			Export: true,
			Doc:    "Purple background",
		},
		"bg-cyan": {
			Val:    slip.String("\x1b[0;46m"),
			Const:  true,
			Export: true,
			Doc:    "Cyan background",
		},
		"bg-gray": {
			Val:    slip.String("\x1b[0;47m"),
			Const:  true,
			Export: true,
			Doc:    "Gray background",
		},
		"bg-dark-gray": {
			Val:    slip.String("\x1b[0;100m"),
			Const:  true,
			Export: true,
			Doc:    "Dark gray background",
		},
		"bg-light-red": {
			Val:    slip.String("\x1b[0;101m"),
			Const:  true,
			Export: true,
			Doc:    "Light red background",
		},
		"bg-light-green": {
			Val:    slip.String("\x1b[0;102m"),
			Const:  true,
			Export: true,
			Doc:    "Light green background",
		},
		"bg-light-yellow": {
			Val:    slip.String("\x1b[0;103m"),
			Const:  true,
			Export: true,
			Doc:    "Light yellow background",
		},
		"bg-light-blue": {
			Val:    slip.String("\x1b[0;104m"),
			Const:  true,
			Export: true,
			Doc:    "Light blue background",
		},
		"bg-light-purple": {
			Val:    slip.String("\x1b[0;105m"),
			Const:  true,
			Export: true,
			Doc:    "Light purple background",
		},
		"bg-light-cyan": {
			Val:    slip.String("\x1b[0;106m"),
			Const:  true,
			Export: true,
			Doc:    "Light cyan background",
		},
		"bg-white": {
			Val:    slip.String("\x1b[0;107m"),
			Const:  true,
			Export: true,
			Doc:    "White background",
		},
	})
	// slip/pkg/clos is needed so make sure it gets inited first with an
	// import.
	initTypes(sen.MustParse(fhir5JSON))

	initDescribeType()

	Pkg.Initialize(nil, &Type{}) // lock
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
	var primitives []*Type
	for _, pa := range jp.C("primitives").W().Get(schema) {
		pt := Type{
			name:        alt.String(jp.C("name").First(pa)),
			description: alt.String(jp.C("description").First(pa)),
			pkg:         &Pkg,
			parent:      alt.String(jp.C("parent").First(pa)),
			pattern:     alt.String(jp.C("pattern").First(pa)),
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
		ft := Type{
			name:        alt.String(jp.C("name").First(ts)),
			description: alt.String(jp.C("description").First(ts)),
			pkg:         &Pkg,
			parent:      alt.String(jp.C("parent").First(ts)),
		}
		for _, ps := range jp.C("properties").W().Get(ts) {
			ft.props = append(ft.props, NewProp(ps))
		}
		slip.RegisterClass(strings.ToLower(ft.name), &ft)
		types = append(types, &ft)
	}
}

func initTypeParents() {
	for _, ft := range types {
		if base, ok := ft.(*Type); ok {
			base.init()
		}
	}
}
