// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	_ "github.com/ohler55/slip/pkg/clos"
	"github.com/ohler55/slip/pkg/flavors"
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
the fhir classes can be used to build as well as access fields of FHIR datatypes. The
__fhir__ package contains the functions and metaclasses for dynamically loaded FHIR
types. For more details invoke the __http-help__ funcion.


The default FHIR version is v5.0.0 and is in package fhir5.`,
		PreSet: slip.DefaultPreSet,
	}
	blankType = Type{
		name:        "Type",
		pkg:         &Pkg,
		description: "The meta-class for all FHIR types. All FHIR instance methods are defined by this class.",
	}
	blankProp = Property{
		name: "Property",
		pkg:  &Pkg,
		docs: "The meta-class for all FHIR properties.",
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
	slip.RegisterClass("type", &blankType)
	slip.RegisterClass("property", &blankProp)

	initDescribeType()
	initInstanceID()
	initInstanceData()
	initInstanceGet()
	initInstanceSet()
	initInstanceReplace()
	initValidP()
	initTypeProperties()
	initTypeProperty()
	initPropertyName()
	initPropertyType()
	initPropertyEnum()
	initPropertyGroup()
	initPropertyCardinality()
	initPropertyValidP()
	initLoadDefinitions()

	initHTTPHelp()
	initHTTPRead()
	initHTTPEach()
	initHTTPCapabilities()
	initHTTPCreate()

	p5 := slip.DefPackage("fhir5", []string{}, "FHIR version 5.0.0")
	defineTypes(sen.MustParse(fhir5JSON), p5)

	Pkg.Initialize(nil, &Type{}) // lock
	slip.AddPackage(&Pkg)
	slip.UserPkg.Use(&Pkg)
}

func objectify(v any) (obj slip.Object) {
	switch v.(type) {
	case map[string]any, []any:
		bg := bag.Flavor().MakeInstance().(*flavors.Instance)
		bg.Any = v
		obj = bg
	default:
		obj = slip.SimpleObject(v)
	}
	return
}

func primitiveName(v any) (name string) {
	switch v.(type) {
	case string:
		name = "string"
	case int64:
		name = "integer"
	case float64, float32:
		name = "decimal"
	case time.Time:
		name = "dateTime"
	default:
		name = fmt.Sprintf("%T", v)
	}
	return
}
