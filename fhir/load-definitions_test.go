// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"testing"

	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestLoadDefinitionsOk(t *testing.T) {
	p5 := slip.FindPackage("fhir5")
	defer func() {
		_ = p5.Remove("Dummy")
		_ = p5.Remove("Dummy0")
	}()

	(&sliptest.Function{
		Source: `(load-definitions "testdata/dummy.json" "fhir5")`,
		Expect: "(#<fhir5:Type Dummy> #<fhir5:Type Dummy0>)",
	}).Test(t)
	tt.NotNil(t, slip.FindClass("fhir5:Dummy"))
}

func TestLoadDefinitionsPkg(t *testing.T) {
	p5 := slip.FindPackage("fhir5")
	defer func() {
		_ = p5.Remove("Dummy")
		_ = p5.Remove("Dummy0")
	}()

	(&sliptest.Function{
		Source: `(load-definitions "testdata/dummy.json" (find-package 'fhir5))`,
		Expect: "(#<fhir5:Type Dummy> #<fhir5:Type Dummy0>)",
	}).Test(t)
	tt.NotNil(t, slip.FindClass("fhir5:Dummy"))
}

func TestLoadDefinitionsBadPropertyType(t *testing.T) {
	p5 := slip.FindPackage("fhir5")
	defer func() { _ = p5.Remove("Dummy") }()

	(&sliptest.Function{
		Source:    `(load-definitions "testdata/bad-property-type.json" 'fhir5)`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
}

func TestLoadDefinitionsSelfParent(t *testing.T) {
	p5 := slip.FindPackage("fhir5")
	defer func() { _ = p5.Remove("Dummy") }()

	(&sliptest.Function{
		Source:    `(load-definitions "testdata/bad-dummy.json" 'fhir5)`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
}

func TestLoadDefinitionsBadParent(t *testing.T) {
	p5 := slip.FindPackage("fhir5")
	defer func() { _ = p5.Remove("Dummy") }()

	(&sliptest.Function{
		Source:    `(load-definitions "testdata/bad-parent.json" 'fhir5)`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
}

func TestLoadDefinitionsNewPkg(t *testing.T) {
	fire := slip.FindPackage("fire")
	defer func() {
		slip.RemovePackage(fire)
	}()

	(&sliptest.Function{
		Source: `(load-definitions "testdata/standalone.json" 'fire)`,
		Expect: "(#<fire:Type integer>)",
	}).Test(t)
	tt.NotNil(t, slip.FindClass("fire:integer"))
}

func TestLoadDefinitionsBadPackage(t *testing.T) {
	(&sliptest.Function{
		Source:    `(load-definitions "testdata/bad-parent.json" t)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestLoadDefinitionsBadFilename(t *testing.T) {
	(&sliptest.Function{
		Source:    `(load-definitions "testdata/not-found" 'fire)`,
		PanicType: slip.FileErrorSymbol,
	}).Test(t)
}
