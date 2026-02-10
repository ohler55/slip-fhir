// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"testing"

	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestTypeProperty(t *testing.T) {
	(&sliptest.Function{
		Source: `(type-property (find-class 'fhir5:patient) 'gender)`,
		Expect: "/#<fhir5:Property gender [0-9a-f]+>/",
	}).Test(t)
}

func TestTypePropertySymbol(t *testing.T) {
	(&sliptest.Function{
		Source: `(type-property 'fhir5:patient "gender")`,
		Expect: "/#<fhir5:Property gender [0-9a-f]+>/",
	}).Test(t)
}

func TestTypePropertyInherit(t *testing.T) {
	(&sliptest.Function{
		Source: `(type-property 'fhir5:patient 'id)`,
		Expect: "/#<fhir5:Property id [0-9a-f]+>/",
	}).Test(t)
}

func TestTypePropertyNotFound(t *testing.T) {
	(&sliptest.Function{
		Source: `(type-property 'fhir5:patient 'quux)`,
		Expect: "nil",
	}).Test(t)
}

func TestTypePropertyNotType(t *testing.T) {
	(&sliptest.Function{
		Source:    `(type-property 7 "gender")`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}
