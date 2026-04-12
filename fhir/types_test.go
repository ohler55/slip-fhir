// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"testing"

	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestTypes(t *testing.T) {
	(&sliptest.Function{
		Source: `(fhir-types 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			list, _ := v.(slip.List)
			tt.Equal(t, 858, len(list))
			tt.Equal(t, true, classMemberP(list, "code"))
			tt.Equal(t, true, classMemberP(list, "quantity"))
			tt.Equal(t, true, classMemberP(list, "encounter_reason"))
			tt.Equal(t, true, classMemberP(list, "Patient"))
		},
	}).Test(t)
}
