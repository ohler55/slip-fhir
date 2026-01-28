// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
)

func onErr(path jp.Expr, value any, message string) bool {
	panic(message)
}

func TestTypeInteger(t *testing.T) {
	pt, ok := slip.FindClass("fhir:integer").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate(0, onErr)
	pt.Validate(7, onErr)

	pt.Validate(-7, onErr)
	pt.Validate(math.MaxInt32, onErr)
	pt.Validate(math.MinInt32, onErr)

	pt.Validate(int64(7), onErr)
	pt.Validate(int32(7), onErr)
	pt.Validate(int16(7), onErr)
	pt.Validate(int8(7), onErr)

	pt.Validate(uint(7), onErr)
	pt.Validate(uint64(7), onErr)
	pt.Validate(uint32(7), onErr)
	pt.Validate(uint16(7), onErr)
	pt.Validate(uint8(7), onErr)

	pt.Validate(2.0, onErr)
	pt.Validate(float32(2.0), onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
	tt.Panic(t, func() { pt.Validate(math.MaxInt32+1, onErr) })
	tt.Panic(t, func() { pt.Validate(math.MinInt32-1, onErr) })
	tt.Panic(t, func() { pt.Validate(2.5, onErr) })
	tt.Panic(t, func() { pt.Validate(float32(2.5), onErr) })
}

func TestTypeInteger64(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate(0, onErr)
	pt.Validate(7, onErr)
	pt.Validate(-7, onErr)
	pt.Validate(math.MaxInt32+1, onErr)
	pt.Validate(math.MinInt32-1, onErr)

	pt.Validate(int64(7), onErr)
	pt.Validate(int32(7), onErr)
	pt.Validate(int16(7), onErr)
	pt.Validate(int8(7), onErr)

	pt.Validate(uint(7), onErr)
	pt.Validate(uint64(7), onErr)
	pt.Validate(uint32(7), onErr)
	pt.Validate(uint16(7), onErr)
	pt.Validate(uint8(7), onErr)

	pt.Validate(2.0, onErr)
	pt.Validate(float32(2.0), onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
	tt.Panic(t, func() { pt.Validate(2.5, onErr) })
}

func TestTypeUnsignedInt(t *testing.T) {
	pt, ok := slip.FindClass("unsignedInt").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate(0, onErr)
	pt.Validate(7, onErr)
	pt.Validate(math.MaxInt32, onErr)

	pt.Validate(2.0, onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
	tt.Panic(t, func() { pt.Validate(2.5, onErr) })
	tt.Panic(t, func() { pt.Validate(-1, onErr) })
}

func TestTypePositiveInt(t *testing.T) {
	pt, ok := slip.FindClass("positiveInt").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate(7, onErr)
	pt.Validate(math.MaxInt32, onErr)

	pt.Validate(2.0, onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
	tt.Panic(t, func() { pt.Validate(2.5, onErr) })
	tt.Panic(t, func() { pt.Validate(-1, onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })
}

func TestTypePositiveDecimal(t *testing.T) {
	pt, ok := slip.FindClass("decimal").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate(7, onErr)
	pt.Validate(7.5, onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
}

func TestTypeBoolean(t *testing.T) {
	pt, ok := slip.FindClass("boolean").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate(true, onErr)
	pt.Validate(false, onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
	tt.Panic(t, func() { pt.Validate(nil, onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })
}

func TestTypeTime(t *testing.T) {
	pt, ok := slip.FindClass("fhir:time").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("20:22:23", onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })
	tt.Panic(t, func() { pt.Validate(time.Now(), onErr) })
}

func TestTypeDate(t *testing.T) {
	pt, ok := slip.FindClass("date").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("2026-01-20", onErr)

	tt.Panic(t, func() { pt.Validate("20:21:22", onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })
	tt.Panic(t, func() { pt.Validate(time.Now(), onErr) })
}

func TestTypeInstant(t *testing.T) {
	pt, ok := slip.FindClass("instant").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("2026-01-20T20:21:22.123Z", onErr)
	pt.Validate(time.Now(), onErr)

	tt.Panic(t, func() { pt.Validate("20:21:22", onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })
}

func TestTypeDateTime(t *testing.T) {
	pt, ok := slip.FindClass("dateTime").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("2026-01-20T20:21:22.123Z", onErr)
	pt.Validate(time.Now(), onErr)

	tt.Panic(t, func() { pt.Validate("20:21:22", onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })
}

func TestTypeXHTML(t *testing.T) {
	pt, ok := slip.FindClass("xhtml").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("<x>y</x>", onErr)

	tt.Panic(t, func() { pt.Validate(0, onErr) })
}

func TestTypeSimplify(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
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

func TestTypeDescribe(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	desc := string(pt.Describe(nil, 0, 60, false))
	tt.Equal(t, true, strings.Contains(desc, "fhir:integer64 is a FHIR PrimitiveType"))
	tt.Equal(t, true, strings.Contains(desc, "Direct Ancestor: fixnum"))

	desc = string(pt.Describe(nil, 0, 60, true))
	tt.Equal(t, true, strings.Contains(desc, "is a FHIR PrimitiveType"))
	tt.Equal(t, true, strings.Contains(desc, "Direct Ancestor: fixnum"))
}

func TestTypeLoadForm(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Nil(t, pt.LoadForm())
}

func TestTypeVarNames(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, 0, len(pt.VarNames()))
}

func TestTypeEval(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, pt, pt.Eval(nil, 0))
}

func TestTypeEqual(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, true, pt.Equal(pt))
	var pt2 *fhir.Type
	pt2, ok = slip.FindClass("fhir:integer").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, false, pt.Equal(pt2))
}

func TestTypeHierachy(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, `[
  "fhir:integer64"
  "common-lisp:fixnum"
  "common-lisp:integer"
  "common-lisp:rational"
  "common-lisp:real"
  "common-lisp:number"
  t
]`, pretty.SEN(pt.Hierarchy()))
}

func TestTypeMetaclass(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, slip.Symbol("fhir-type"), pt.Metaclass())
}

func TestTypeInherits(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, true, pt.Inherits(slip.FindClass("real")))
	tt.Equal(t, false, pt.Inherits(slip.FindClass("float")))
}

func TestTypeDocumentation(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	orig := pt.Documentation()
	defer pt.SetDocumentation(orig)
	newDoc := "temporary"
	pt.SetDocumentation(newDoc)
	tt.Equal(t, newDoc, pt.Documentation())
}

func TestTypeMakeInstance(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Panic(t, func() { _ = pt.MakeInstance() })
}

func TestTypeBadInit(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Panic(t, func() { _ = pt.MakeInstance() })
}
