// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"strings"
	"testing"

	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestDescribeTypeMinimum(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(describe-type (find-class 'fhir5:patient) out)`,
		Expect: "nil",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, "/ is a FHIR Resource:/", desc)
	tt.Equal(t, `/ deceased\[x\] /`, desc)
	tt.Equal(t, "/ deceasedBoolean /", desc)
	tt.Equal(t, false, strings.Contains(desc, "_gender"))
}

func TestDescribeTypeFlags(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(describe-type (find-class 'fhir5:account) out)`,
		Expect: "nil",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, `/ ?! Σ/`, desc)
}

func TestDescribeTypeFull(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(describe-type 'fhir5:patient out :full t :stripe-color bg-light-blue)`,
		Expect: "nil",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, "/ is a FHIR Resource:/", desc)
	tt.Equal(t, `/ deceased\[x\] /`, desc)
	tt.Equal(t, "/ deceasedBoolean /", desc)
	tt.Equal(t, "/_gender /", desc)
	// Assure the color code is in the output.
	tt.Equal(t, `/\[0;104m/`, desc)
}

func TestDescribeTypeTight(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(let ((*print-right-margin* 40)) (describe-type (find-class 'fhir5:patient) out))`,
		Expect: "nil",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, "/ is a FHIR Resource:/", desc)
	tt.Equal(t, `/ deceased\[x\] /`, desc)
	tt.Equal(t, "/ deceasedBoolean /", desc)
	tt.Equal(t, false, strings.Contains(desc, "_gender"))
}

func TestDescribeTypeFullTight(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(let ((*print-right-margin* 40)) (describe-type (find-class 'fhir5:patient) out :full t))`,
		Expect: "nil",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, "/ is a FHIR Resource:/", desc)
	tt.Equal(t, `/ deceased\[x\] /`, desc)
	tt.Equal(t, "/ deceasedBoolean /", desc)
}

func TestDescribeTypeOther(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(describe-type (find-class 'bag-flavor) out)`,
		Expect: "nil",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, "/ is a flavor:/", desc)
}

func TestDescribeTypeBadOption(t *testing.T) {
	(&sliptest.Function{
		Source:    `(describe-type 'patient 'quux)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(describe-type 'patient 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestDescribeTypeWriteError(t *testing.T) {
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: badWriter(0)})
	(&sliptest.Function{
		Scope:     scope,
		Source:    `(describe-type (find-class 'patient) out)`,
		PanicType: slip.StreamErrorSymbol,
	}).Test(t)
}
