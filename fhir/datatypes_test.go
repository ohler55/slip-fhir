// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"testing"

	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestDatatypes(t *testing.T) {
	(&sliptest.Function{
		Source: `(fhir-datatypes 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			list, _ := v.(slip.List)
			tt.Equal(t, 45, len(list))
			tt.Equal(t, true, classMemberP(list, "address"))
			tt.Equal(t, true, classMemberP(list, "quantity"))
			tt.Equal(t, false, classMemberP(list, "Patient"))
			tt.Equal(t, false, classMemberP(list, "code"))
		},
	}).Test(t)
}
