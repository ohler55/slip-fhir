// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
)

func TestPrimitiveTypeInteger32(t *testing.T) {
	pt, ok := slip.FindClass("integer32").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate(0)
	pt.Validate(7)
	pt.Validate(-7)
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

func TestPrimitiveTypeInteger64(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate(0)
	pt.Validate(7)
	pt.Validate(-7)
	pt.Validate(math.MaxInt32 + 1)
	pt.Validate(math.MinInt32 - 1)

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
	tt.Panic(t, func() { pt.Validate(2.5) })
}

func TestPrimitiveTypeUnsignedInt(t *testing.T) {
	pt, ok := slip.FindClass("unsignedInt").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate(0)
	pt.Validate(7)
	pt.Validate(math.MaxInt32)

	pt.Validate(2.0)

	tt.Panic(t, func() { pt.Validate("string") })
	tt.Panic(t, func() { pt.Validate(2.5) })
	tt.Panic(t, func() { pt.Validate(-1) })
}

func TestPrimitiveTypePositiveInt(t *testing.T) {
	pt, ok := slip.FindClass("positiveInt").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate(7)
	pt.Validate(math.MaxInt32)

	pt.Validate(2.0)

	tt.Panic(t, func() { pt.Validate("string") })
	tt.Panic(t, func() { pt.Validate(2.5) })
	tt.Panic(t, func() { pt.Validate(-1) })
	tt.Panic(t, func() { pt.Validate(0) })
}

func TestPrimitiveTypePositiveDecimal(t *testing.T) {
	pt, ok := slip.FindClass("decimal").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate(7)
	pt.Validate(7.5)

	tt.Panic(t, func() { pt.Validate("string") })
}

func TestPrimitiveTypeBoolean(t *testing.T) {
	pt, ok := slip.FindClass("boolean").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate(true)
	pt.Validate(false)

	tt.Panic(t, func() { pt.Validate("string") })
	tt.Panic(t, func() { pt.Validate(nil) })
	tt.Panic(t, func() { pt.Validate(0) })
}

func TestPrimitiveTypeTime(t *testing.T) {
	pt, ok := slip.FindClass("ftime").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate("20:22:23")

	tt.Panic(t, func() { pt.Validate("string") })
	tt.Panic(t, func() { pt.Validate(0) })
	tt.Panic(t, func() { pt.Validate(time.Now()) })
}

func TestPrimitiveTypeDate(t *testing.T) {
	pt, ok := slip.FindClass("date").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate("2026-01-20")

	tt.Panic(t, func() { pt.Validate("20:21:22") })
	tt.Panic(t, func() { pt.Validate(0) })
	tt.Panic(t, func() { pt.Validate(time.Now()) })
}

func TestPrimitiveTypeInstant(t *testing.T) {
	pt, ok := slip.FindClass("instant").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate("2026-01-20T20:21:22.123Z")
	pt.Validate(time.Now())

	tt.Panic(t, func() { pt.Validate("20:21:22") })
	tt.Panic(t, func() { pt.Validate(0) })
}

func TestPrimitiveTypeDateTime(t *testing.T) {
	pt, ok := slip.FindClass("dateTime").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate("2026-01-20T20:21:22.123Z")
	pt.Validate(time.Now())

	tt.Panic(t, func() { pt.Validate("20:21:22") })
	tt.Panic(t, func() { pt.Validate(0) })
}

func TestPrimitiveTypeXHTML(t *testing.T) {
	pt, ok := slip.FindClass("xhtml").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)

	pt.Validate("<x>y</x>")

	tt.Panic(t, func() { pt.Validate(0) })
}

func TestPrimitiveTypeSimplify(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Equal(t, `{
  description: "A very large whole number"
  inherit: fixnum
  name: integer64
  package: fhir
  parent: fixnum
  pattern: "^[0]|[-+]?[1-9][0-9]*$"
}`, pretty.SEN(pt.Simplify()))
}

func TestPrimitiveTypeDescribe(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	desc := string(pt.Describe(nil, 0, 60, false))
	tt.Equal(t, true, strings.Contains(desc, "fhir:integer64 is a FHIR PrimitiveType"))
	tt.Equal(t, true, strings.Contains(desc, "Direct Ancestor: fixnum"))

	desc = string(pt.Describe(nil, 0, 60, true))
	tt.Equal(t, true, strings.Contains(desc, "is a FHIR PrimitiveType"))
	tt.Equal(t, true, strings.Contains(desc, "Direct Ancestor: fixnum"))
}

func TestPrimitiveTypeLoadForm(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Nil(t, pt.LoadForm())
}

func TestPrimitiveTypeVarNames(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Equal(t, 0, len(pt.VarNames()))
}

func TestPrimitiveTypeEval(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Equal(t, pt, pt.Eval(nil, 0))
}

func TestPrimitiveTypeEqual(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Equal(t, true, pt.Equal(pt))
	var pt2 *fhir.PrimitiveType
	pt2, ok = slip.FindClass("integer32").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Equal(t, false, pt.Equal(pt2))
}

func TestPrimitiveTypeHierachy(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Equal(t, "[integer64 fixnum integer rational real number t]", pretty.SEN(pt.Hierarchy()))
}

func TestPrimitiveTypeMetaclass(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Equal(t, slip.Symbol("primitive-type"), pt.Metaclass())
}

func TestPrimitiveTypeInherits(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Equal(t, true, pt.Inherits(slip.FindClass("real")))
	tt.Equal(t, false, pt.Inherits(slip.FindClass("float")))
}

func TestPrimitiveTypeDocumentation(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	orig := pt.Documentation()
	defer pt.SetDocumentation(orig)
	newDoc := "temporary"
	pt.SetDocumentation(newDoc)
	tt.Equal(t, newDoc, pt.Documentation())
}

func TestPrimitiveTypeMakeInstance(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Panic(t, func() { _ = pt.MakeInstance() })
}

func TestPrimitiveTypeBadInit(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.PrimitiveType)
	tt.Equal(t, true, ok)
	tt.Panic(t, func() { _ = pt.MakeInstance() })
}
