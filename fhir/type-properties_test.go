// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"testing"

	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestTypeProperties(t *testing.T) {
	(&sliptest.Function{
		Source: `(sort (mapcar (lambda (p) (send p :name)) (type-properties (find-class 'range))))`,
		Expect: `("extension" "high" "id" "low")`,
	}).Test(t)
}

func TestTypePropertiesSymbol(t *testing.T) {
	(&sliptest.Function{
		Source: `(sort (mapcar (lambda (p) (send p :name)) (type-properties 'range)))`,
		Expect: `("extension" "high" "id" "low")`,
	}).Test(t)
}

func TestTypePropertiesNotFound(t *testing.T) {
	(&sliptest.Function{
		Source: `(type-properties 'quux-type)`,
		Expect: "nil",
	}).Test(t)
}

func TestTypePropertiesNotType(t *testing.T) {
	(&sliptest.Function{
		Source:    `(type-properties 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}
