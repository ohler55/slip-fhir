// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"math"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
	"github.com/ohler55/slip/sliptest"
)

func onErr(path jp.Expr, value any, message string) bool {
	panic(message)
}

func onErrStop(path jp.Expr, value any, message string) bool {
	return true
}

func TestTypeValidateInteger(t *testing.T) {
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

	tt.Equal(t, true, pt.Validate("string", onErrStop))
	tt.Equal(t, true, pt.Validate(math.MaxInt32+1, onErrStop))
}

func TestTypeValidateInteger64(t *testing.T) {
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

	tt.Equal(t, true, pt.Validate("string", onErrStop))
}

func TestTypeValidateUnsignedInt(t *testing.T) {
	pt, ok := slip.FindClass("unsignedInt").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate(0, onErr)
	pt.Validate(7, onErr)
	pt.Validate(math.MaxInt32, onErr)

	pt.Validate(2.0, onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
	tt.Panic(t, func() { pt.Validate(2.5, onErr) })
	tt.Panic(t, func() { pt.Validate(-1, onErr) })

	tt.Equal(t, true, pt.Validate("string", onErrStop))
	tt.Equal(t, true, pt.Validate(-1, onErrStop))
}

func TestTypeValidatePositiveInt(t *testing.T) {
	pt, ok := slip.FindClass("positiveInt").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate(7, onErr)
	pt.Validate(math.MaxInt32, onErr)

	pt.Validate(2.0, onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
	tt.Panic(t, func() { pt.Validate(2.5, onErr) })
	tt.Panic(t, func() { pt.Validate(-1, onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })

	tt.Equal(t, true, pt.Validate("string", onErrStop))
	tt.Equal(t, true, pt.Validate(0, onErrStop))
}

func TestTypeValidatePositiveDecimal(t *testing.T) {
	pt, ok := slip.FindClass("decimal").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate(7, onErr)
	pt.Validate(7.5, onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })

	tt.Equal(t, true, pt.Validate("string", onErrStop))
}

func TestTypeValidateBoolean(t *testing.T) {
	pt, ok := slip.FindClass("boolean").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate(true, onErr)
	pt.Validate(false, onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
	tt.Panic(t, func() { pt.Validate(nil, onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })

	tt.Equal(t, true, pt.Validate("string", onErrStop))
}

func TestTypeValidateTime(t *testing.T) {
	pt, ok := slip.FindClass("fhir:time").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("20:22:23", onErr)

	tt.Panic(t, func() { pt.Validate("string", onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })
	tt.Panic(t, func() { pt.Validate(time.Now(), onErr) })

	tt.Equal(t, true, pt.Validate("string", onErrStop))
}

func TestTypeValidateDate(t *testing.T) {
	pt, ok := slip.FindClass("date").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("2026-01-20", onErr)

	tt.Panic(t, func() { pt.Validate("20:21:22", onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })
	tt.Panic(t, func() { pt.Validate(time.Now(), onErr) })

	tt.Equal(t, true, pt.Validate("string", onErrStop))
}

func TestTypeValidateInstant(t *testing.T) {
	pt, ok := slip.FindClass("instant").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("2026-01-20T20:21:22.123Z", onErr)
	pt.Validate(time.Now(), onErr)

	tt.Panic(t, func() { pt.Validate("20:21:22", onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })

	tt.Equal(t, true, pt.Validate("string", onErrStop))
}

func TestTypeValidateDateTime(t *testing.T) {
	pt, ok := slip.FindClass("dateTime").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("2026-01-20T20:21:22.123Z", onErr)
	pt.Validate(time.Now(), onErr)

	tt.Panic(t, func() { pt.Validate("20:21:22", onErr) })
	tt.Panic(t, func() { pt.Validate(0, onErr) })

	tt.Equal(t, true, pt.Validate("string", onErrStop))
}

func TestTypeValidateCode(t *testing.T) {
	pt, ok := slip.FindClass("code").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("abc", onErr)

	tt.Panic(t, func() { pt.Validate(0, onErr) })

	tt.Equal(t, true, pt.Validate(0, onErrStop))
}

func TestTypeValidateXHTML(t *testing.T) {
	pt, ok := slip.FindClass("xhtml").(*fhir.Type)
	tt.Equal(t, true, ok)

	pt.Validate("<x>y</x>", onErr)

	tt.Panic(t, func() { pt.Validate(0, onErr) })

	tt.Equal(t, true, pt.Validate(0, onErrStop))
}

func TestTypeValidateBase(t *testing.T) {
	bt, ok := slip.FindClass("base").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, false, bt.Validate(nil, onErrStop))
}

func TestTypeValidateComplex(t *testing.T) {
	ct, ok := slip.FindClass("Range").(*fhir.Type)
	tt.Equal(t, true, ok)

	ct.Validate(map[string]any{
		"low":  map[string]any{"value": 30, "unit": "mL"},
		"high": map[string]any{"value": 50, "unit": "mL"},
	}, onErr)

	tt.Panic(t, func() { ct.Validate(0, onErr) })
	tt.Panic(t, func() { ct.Validate(map[string]any{"average": 5}, onErr) })
	tt.Panic(t, func() { ct.Validate(map[string]any{"low": 5}, onErr) })

	tt.Equal(t, true, ct.Validate(7, onErrStop))
	tt.Equal(t, true, ct.Validate(map[string]any{"average": 5}, onErrStop))
	tt.Equal(t, true, ct.Validate(map[string]any{"low": 5}, onErrStop))
	tt.Equal(t, true, ct.Validate(map[string]any{
		"LOW": map[string]any{"value": 30, "unit": "mL"},
	}, onErrStop))

}

func TestTypeName(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, "integer64", pt.Name())
}

func TestTypeSimplifyPrimitive(t *testing.T) {
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

func TestTypeSimplifyComplex(t *testing.T) {
	rt, ok := slip.FindClass("Range").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, `{
  description: "A set of ordered Quantities defined by a low and high limit."
  inherit: Element
  name: Range
  package: fhir
  parent: Element
  properties: [
    {
      array: false
      description: "The low limit. The boundary is inclusive."
      enum: []
      name: low
      required: false
      type: Quantity
    }
    {
      array: false
      description: "The high limit. The boundary is inclusive."
      enum: []
      name: high
      required: false
      type: Quantity
    }
  ]
}`, pretty.SEN(rt.Simplify()))
}

func TestTypeVarNamesComplex(t *testing.T) {
	rt, ok := slip.FindClass("Range").(*fhir.Type)
	tt.Equal(t, true, ok)
	names := rt.VarNames()
	sort.Strings(names)
	tt.Equal(t, `[extension high id low]`, pretty.SEN(names))
}

func TestTypeMethods(t *testing.T) {
	rt, ok := slip.FindClass("Range").(*fhir.Type)
	tt.Equal(t, true, ok)
	var names []string
	for _, m := range rt.Methods() {
		names = append(names, m.Name)
	}
	sort.Strings(names)
	tt.Equal(t, `[
  ":class"
  ":data"
  ":describe"
  ":equal"
  ":get"
  ":id"
  ":init"
  ":operation-handled-p"
  ":print-self"
  ":replace"
  ":set"
  ":type"
  ":valid-p"
  ":which-operations"
]`, pretty.SEN(names))
}

func TestTypeGetMethod(t *testing.T) {
	rt, ok := slip.FindClass("Range").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Equal(t, `{
  combinations: [{from: Type primary: true}]
  name: ":set"
}`, pretty.SEN(rt.GetMethod(":set")))
}

func TestTypeDescribe(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	desc := string(pt.Describe(nil, 0, 80, false))
	tt.Equal(t, true, strings.Contains(desc, "fhir:integer64 is a FHIR PrimitiveType"))
	tt.Equal(t, true, strings.Contains(desc, "Direct Ancestor: fixnum"))

	desc = string(pt.Describe(nil, 0, 80, true))
	tt.Equal(t, true, strings.Contains(desc, "is a FHIR PrimitiveType"))
	tt.Equal(t, true, strings.Contains(desc, "Direct Ancestor: fixnum"))

	pt, ok = slip.FindClass("Range").(*fhir.Type)
	tt.Equal(t, true, ok)
	desc = string(pt.Describe(nil, 0, 80, false))
	tt.Equal(t, true, strings.Contains(desc, "is a FHIR DataType"))

	pt, ok = slip.FindClass("Account").(*fhir.Type)
	tt.Equal(t, true, ok)
	desc = string(pt.Describe(nil, 0, 80, false))
	tt.Equal(t, true, strings.Contains(desc, "is a FHIR Resource"))

	pt, ok = slip.FindClass("Account_coverage").(*fhir.Type)
	tt.Equal(t, true, ok)
	desc = string(pt.Describe(nil, 0, 80, false))
	tt.Equal(t, true, strings.Contains(desc, "is a FHIR BackboneType"))
}

func TestTypeDescribeSelf(t *testing.T) {
	pt, ok := slip.FindClass("type").(*fhir.Type)
	tt.Equal(t, true, ok)
	desc := string(pt.Describe(nil, 0, 80, false))

	tt.Equal(t, true, strings.Contains(desc, "fhir:Type is the FHIR meta-class"))
	tt.Equal(t, true, strings.Contains(desc, "Methods:"))
	tt.Equal(t, true, strings.Contains(desc, ":which-operations"))

	desc = string(pt.Describe(nil, 0, 80, true))
	tt.Equal(t, true, strings.Contains(desc, "is the FHIR meta-class"))
}

func TestTypeLoadForm(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Nil(t, pt.LoadForm())
}

func TestTypeVarNamesPrimitive(t *testing.T) {
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

func TestTypeMakeInstancePrimitive(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Panic(t, func() { _ = pt.MakeInstance() })
}

func TestTypeMakeInstanceComplex(t *testing.T) {
	(&sliptest.Function{
		Source: `(make-instance 'range :data (make-bag "{low:{value:30 unit:mL} high:{value:50 unit:mL}}"))`,
		Expect: "/#<fhir:Range [0-9a-f]+>/",
	}).Test(t)
	(&sliptest.Function{
		Source: `(send
                   (send
                     (make-instance 'range :data (make-bag "{low:{value:30 unit:mL} high:{value:50 unit:mL}}"))
                     :data)
                   :write nil)`,
		Expect: `"{high: {unit: mL value: 50} low: {unit: mL value: 30}}"`,
	}).Test(t)
}

func TestTypeBadInit(t *testing.T) {
	pt, ok := slip.FindClass("integer64").(*fhir.Type)
	tt.Equal(t, true, ok)
	tt.Panic(t, func() { _ = pt.MakeInstance() })
}
