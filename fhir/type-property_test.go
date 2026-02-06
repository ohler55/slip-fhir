// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"testing"

	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestTypeProperty(t *testing.T) {
	(&sliptest.Function{
		Source: `(type-property (find-class 'patient) 'gender)`,
		Expect: "/#<fhir:property gender [0-9a-f]+>/",
	}).Test(t)
}

func TestTypePropertySymbol(t *testing.T) {
	(&sliptest.Function{
		Source: `(type-property 'patient "gender")`,
		Expect: "/#<fhir:property gender [0-9a-f]+>/",
	}).Test(t)
}

func TestTypePropertyInherit(t *testing.T) {
	(&sliptest.Function{
		Source: `(type-property 'patient 'id)`,
		Expect: "/#<fhir:property id [0-9a-f]+>/",
	}).Test(t)
}

func TestTypePropertyNotFound(t *testing.T) {
	(&sliptest.Function{
		Source: `(type-property 'patient 'quux)`,
		Expect: "nil",
	}).Test(t)
}

func TestTypePropertyNotType(t *testing.T) {
	(&sliptest.Function{
		Source:    `(type-property 7 "gender")`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}
