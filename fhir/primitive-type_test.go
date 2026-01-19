// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"math"
	"testing"

	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
)

func TestPrimitiveTypeInteger32(t *testing.T) {
	//scope := slip.NewScope()
	pt, ok := slip.FindClass("integer32").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate(7)
	pt.Validate(math.MaxInt32)
	pt.Validate(math.MinInt32)

	pt.Validate(int64(7))
	pt.Validate(int32(7))
	pt.Validate(int16(7))
	pt.Validate(int8(7))

	pt.Validate(uint(7))
	pt.Validate(uint64(7))
	pt.Validate(uint32(7))
	pt.Validate(uint16(7))
	pt.Validate(uint8(7))

	pt.Validate(2.0)
	pt.Validate(float32(2.0))

	tt.Panic(t, func() { pt.Validate("string") })
	tt.Panic(t, func() { pt.Validate(math.MaxInt32 + 1) })
	tt.Panic(t, func() { pt.Validate(math.MinInt32 - 1) })
	tt.Panic(t, func() { pt.Validate(2.5) })
	tt.Panic(t, func() { pt.Validate(float32(2.5)) })
}
