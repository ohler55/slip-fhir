// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"testing"

	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestPropertyName(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (type-property 'patient "gender") :name)`,
		Expect: `"gender"`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(property-name (type-property 'patient "gender"))`,
		Expect: `"gender"`,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(property-name t)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestPropertyType(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (type-property 'patient "gender") :type)`,
		Expect: "#<fhir:Type code>",
	}).Test(t)
	(&sliptest.Function{
		Source: `(property-type (type-property 'patient "gender"))`,
		Expect: "#<fhir:Type code>",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(property-type t)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestPropertyEnum(t *testing.T) {
	(&sliptest.Function{
		Source: `(sort (send (type-property 'patient "gender") :enum))`,
		Expect: `("female" "male" "other" "unknown")`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(sort (property-enum (type-property 'patient "gender")))`,
		Expect: `("female" "male" "other" "unknown")`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(sort (property-enum (type-property 'patient "id")))`,
		Expect: "nil",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(property-enum t)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestPropertyGroup(t *testing.T) {
	(&sliptest.Function{
		Source: `(sort (mapcar (lambda (g) (send g :name)) (send (type-property 'patient "deceased[x]") :group)))`,
		Expect: `("_deceasedBoolean" "_deceasedDateTime" "deceasedBoolean" "deceasedDateTime")`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(sort (mapcar (lambda (g) (send g :name)) (property-group (type-property 'patient "deceased[x]"))))`,
		Expect: `("_deceasedBoolean" "_deceasedDateTime" "deceasedBoolean" "deceasedDateTime")`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(send (type-property 'patient "gender") :group)`,
		Expect: "nil",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(property-group t)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestPropertyCardinality(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (type-property 'patient "gender") :cardinality)`,
		Expect: `0, 1`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(property-cardinality (type-property 'patient "gender"))`,
		Expect: `0, 1`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(send (type-property 'patient "name") :cardinality)`,
		Expect: `0, nil`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(send (type-property 'patient_link "other") :cardinality)`,
		Expect: `1, 1`,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(property-cardinality t)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestPropertyValidPtrue(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (type-property 'patient "gender") :valid-p "female")`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(property-valid-p (type-property 'patient "gender") "male")`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(property-valid-p (type-property 'patient "name")
                   (make-bag "[{given:[Fred] family:Fox}]"))`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(let ((prop (type-property 'patient "maritalStatus")))
                   (property-valid-p prop
                                     (make-instance 'CodeableConcept :data (make-bag "{text:'married'}"))))`,
		Expect: "t",
	}).Test(t)
}

func TestPropertyValidPfalse(t *testing.T) {
	(&sliptest.Function{
		Source: `(property-valid-p (type-property 'patient "gender") "quux")`,
		Expect: "nil",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(property-valid-p t "quux")`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)

	(&sliptest.Function{
		Source: `(let ((prop (type-property 'patient "name"))
                       errors)
                   (property-valid-p prop
                                     (make-bag "[{given:[Fred] family:Fox quux:1}]")
                                     (lambda (p v m) (setq errors (add errors (list p v m)))))
                   errors)`,
		Expect: `((#<bag-path @.name[0].quux> nil "quux is not a property of HumanName"))`,
	}).Test(t)

	(&sliptest.Function{
		Source:    `(property-valid-p (type-property 'patient "name") (make-instance 'vanilla-flavor))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}
