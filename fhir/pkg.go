// Copyright (c) 2026, Peter Ohler, All rights reserved.

package jet

import (
	"github.com/ohler55/slip"
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

	// TBD

	// TBD Pkg.Initialize(nil, &PubMsg{}) // lock
	slip.AddPackage(&Pkg)
	slip.UserPkg.Use(&Pkg)
}
