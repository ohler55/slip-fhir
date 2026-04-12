// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"testing"

	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestPrimitives(t *testing.T) {
	(&sliptest.Function{
		Source: `(fhir-primitives 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			list, _ := v.(slip.List)
			tt.Equal(t, 21, len(list))
			tt.Equal(t, true, classMemberP(list, "code"))
			tt.Equal(t, true, classMemberP(list, "oid"))
			tt.Equal(t, false, classMemberP(list, "quantity"))
			tt.Equal(t, false, classMemberP(list, "Patient"))
		},
	}).Test(t)
}
