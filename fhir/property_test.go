// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
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

func TestPropertyBadMethods(t *testing.T) {
	(&sliptest.Function{
		Source:    `(send (type-property 'patient "gender") :init)`,
		PanicType: slip.InvalidMethodErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(send (type-property 'patient "gender") :quux)`,
		PanicType: slip.InvalidMethodErrorSymbol,
	}).Test(t)
}

func TestPropertyString(t *testing.T) {
	rt, ok := slip.FindClass("Range").(*fhir.Type)
	tt.Equal(t, true, ok)
	p := rt.FindProperty("low")
	tt.NotNil(t, p)
	tt.Equal(t, "/#<fhir:property low [0-9a-f]+>/", p.String())
}

func TestPropertySimplify(t *testing.T) {
	pt, ok := slip.FindClass("Patient").(*fhir.Type)
	tt.Equal(t, true, ok)
	p := pt.FindProperty("deceased[x]")
	tt.NotNil(t, p)
	names := jp.C("group").W().C("name").Get(p.Simplify())
	sort.Slice(names, func(i, j int) bool { return names[i].(string) < names[j].(string) })
	tt.Equal(t, "[_deceasedBoolean _deceasedDateTime deceasedBoolean deceasedDateTime]", pretty.SEN(names))
}

func TestPropertyMisc(t *testing.T) {
	rt, ok := slip.FindClass("Range").(*fhir.Type)
	tt.Equal(t, true, ok)
	p := rt.FindProperty("low")
	tt.NotNil(t, p)

	tt.Equal(t, true, p.Equal(p))
	tt.Equal(t, p, p.Eval(nil, 0))
	tt.Equal(t, "[property t]", pretty.SEN(p.Hierarchy()))

	tt.Equal(t, `[
  ":cardinality"
  ":class"
  ":describe"
  ":enum"
  ":equal"
  ":group"
  ":id"
  ":name"
  ":operation-handled-p"
  ":print-self"
  ":type"
  ":valid-p"
  ":which-operations"
]`, pretty.SEN(p.MethodNames()))

	tt.Equal(t, "low", p.Name())
	tt.Equal(t, "[]", pretty.SEN(p.VarNames()))
	tt.Equal(t, "[]", pretty.SEN(p.InheritsList()))
	tt.Equal(t, "property", string(p.Metaclass()))
	tt.Nil(t, p.LoadForm())
	tt.Panic(t, func() { _ = p.MakeInstance() })
	tt.Equal(t, `{
  combinations: [{from: Type primary: true}]
  name: ":enum"
}`, pretty.SEN(p.GetMethod(":enum")))
	tt.Equal(t, 13, len(p.Methods()))

	docs := p.Documentation()
	defer p.SetDocumentation(docs)
	tt.Equal(t, "The low limit. The boundary is inclusive.", p.Documentation())
	p.SetDocumentation("testing")
	tt.Equal(t, "testing", p.Documentation())
}

func TestPropertyDescribeAnsi(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(describe (type-property 'patient 'gender) out)`,
		Expect: "",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, "/, an instance of .*fhir:Property/", desc)
	tt.Equal(t, "/the gender that the patient is considered to have/", desc)
	tt.Equal(t, "/Type: code/", desc)
	tt.Equal(t, "/Cardinality: 0..1/", desc)
	tt.Equal(t, "/Enum: male female other unknown/", desc)
}

func TestPropertyDescribePlain(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(let ((*print-ansi* nil)) (describe (type-property 'patient "deceased[x]") out))`,
		Expect: "",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, "/, an instance of .*fhir:Property/", desc)
	tt.Equal(t, "/Group:/", desc)
	tt.Equal(t, "/deceasedBoolean/", desc)
	tt.Equal(t, "/deceasedDateTime/", desc)
}

func TestPropertyDescribeCardinalityArray(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(describe (type-property 'patient 'name) out)`,
		Expect: "",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, `/Cardinality: 0\.\.\*/`, desc)
}

func TestPropertyDescribeRequired(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(describe (type-property 'patient_link 'other) out)`,
		Expect: "",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, `/Cardinality: 1\.\.1/`, desc)
}

func TestPropertyDescribeSelf(t *testing.T) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: `(describe 'fhir:property out)`,
		Expect: "",
	}).Test(t)
	desc := out.String()
	tt.Equal(t, "/ is the FHIR property meta-class/", desc)
	tt.Equal(t, "/The meta-class for all FHIR properties./", desc)
	tt.Equal(t, "/Methods:/", desc)
	tt.Equal(t, "/:cardinality/", desc)

	out.Reset()
	(&sliptest.Function{
		Scope:  scope,
		Source: `(let ((*print-ansi* nil)) (describe 'fhir:property out))`,
		Expect: "",
	}).Test(t)
	desc = out.String()
	tt.Equal(t, "/fhir:Property is the FHIR property meta-class/", desc)
}
