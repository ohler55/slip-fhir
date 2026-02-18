// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
	"github.com/ohler55/slip/sliptest"
)

type badWriter int

func (w badWriter) Write([]byte) (int, error) {
	return 0, fmt.Errorf("oops")
}

func TestInstanceIDMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (make-instance 'fhir5:patient) :id)`,
		Expect: "/[0-9]+/",
	}).Test(t)
}

func TestInstanceIDFunction(t *testing.T) {
	(&sliptest.Function{
		Source: `(instance-id (make-instance 'fhir5:patient))`,
		Expect: "/[0-9]+/",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(instance-id 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceTypeFunction(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (make-instance 'fhir5:patient) :type)`,
		Expect: "#<fhir5:Type Patient>",
	}).Test(t)
}

func TestInstanceWhichOperationsMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(sort (send (make-instance 'fhir5:patient) :which-operations))`,
		Expect: `(:class :data :describe :equal :get :id :init :operation-handled-p :print-self
        :replace :set :type :valid-p :which-operations)`,
	}).Test(t)
}

func TestInstanceOperationHandledPMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (make-instance 'fhir5:patient) :operation-handled-p :id)`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(send (make-instance 'fhir5:patient) :operation-handled-p :quux)`,
		Expect: "nil",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(send (make-instance 'fhir5:patient) :operation-handled-p 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceEqualMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(let ((pat (make-instance 'fhir5:patient)))
                   (send pat :equal pat))`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(let ((pat (make-instance 'fhir5:range :data (make-bag "{low:{value:1} high:{value:2}}")))
                       (pat2 (make-instance 'fhir5:range :data (make-bag "{low:{value:1} high:{value:2}}"))))
                   (send pat :equal pat2))`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(send (make-instance 'fhir5:patient) :equal 7)`,
		Expect: "nil",
	}).Test(t)
}

func TestInstanceDataMethod(t *testing.T) {
	(&sliptest.Function{
		Source:    `(instance-data 'quux)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceDataFunction(t *testing.T) {
	(&sliptest.Function{
		Source:    `(instance-data 'quux)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceGetMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (make-instance 'fhir5:range :data (make-bag "{low:{value:1} high:{value:2}}")) :get "low.value")`,
		Expect: "1",
	}).Test(t)
}

func TestInstanceGetFunction(t *testing.T) {
	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'fhir5:range :data (make-bag "{low:{value:1} high:{value:2}}")) "low.value")`,
		Expect: "1",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(instance-get 7 "low.value")`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'fhir5:HumanName :data (make-bag "{given:[bill bob]}")) 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceGetPathList(t *testing.T) {
	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'fhir5:range :data (make-bag "{low:{value:1} high:{value:2}}")) '(low "value"))`,
		Expect: "1",
	}).Test(t)
	(&sliptest.Function{
		Source: `(send
                   (instance-get
                     (make-instance 'fhir5:range :data (make-bag "{low:{value:1} high:{value:2}}")) "low")
                   :write nil)`,
		Expect: `"{value: 1}"`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'fhir5:HumanName :data (make-bag "{given:[bill bob]}")) '(given 1))`,
		Expect: `"bob"`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'fhir5:HumanName :data (make-bag "{given:[bill bob]}")) '(given nil))`,
		Expect: `"bill"`,
	}).Test(t)

	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'fhir5:HumanName :data (make-bag "{given:[bill bob]}")) '(given t))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceSetMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(let ((inst (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}"))))
                   (send inst :set "high" (make-bag "{value:2}"))
                   (send inst :get "high.value"))`,
		Expect: "2",
	}).Test(t)
}

func TestInstanceSetFunction(t *testing.T) {
	(&sliptest.Function{
		Source: `(let ((inst (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}"))))
                   (instance-set inst "high" (make-bag "{value:2}"))
                   (send inst :get "high.value"))`,
		Expect: "2",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(instance-set 7 "high" 8)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceInit(t *testing.T) {
	(&sliptest.Function{
		Source:    `(send (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}")) :init)`,
		PanicType: slip.InvalidMethodErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(make-instance 'fhir5:range :data nil)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(make-instance 'fhir5:range :data (make-bag "{low:{value:1}}") :on-error 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(make-instance 'fhir5:range :data (make-bag "{quux:{value:1}}"))`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source: `(let ((inst (make-instance 'fhir5:range
                                            :data (make-bag "{quux:{value:1}}") :on-error (lambda (p v m) t))))
                   (send (send inst :data) :write nil))`,
		Expect: `"{}"`,
	}).Test(t)
}

func TestInstancePrintSelfMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(with-output-to-string (s)
                   (send (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}")) :print-self s))`,
		Expect: "/#<fhir5:Range [0-9a-f]+>/",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(send (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}")) :print-self 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstancePrintSelfBadWrite(t *testing.T) {
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: badWriter(0)})
	(&sliptest.Function{
		Scope:     scope,
		Source:    `(send (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}")) :print-self out)`,
		PanicType: slip.StreamErrorSymbol,
	}).Test(t)
}

func TestInstanceDescribeMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(with-output-to-string (s)
                   (send (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}")) :describe s))`,
		Expect: "/an instance of .+fhir5:Range/",
	}).Test(t)
	(&sliptest.Function{
		Source: `(let ((*print-ansi* nil))
                   (with-output-to-string (s)
                     (send (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}")) :describe s)))`,
		Expect: "/an instance of fhir5:Range/",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(send (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}")) :describe 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceDescribeBadWrite(t *testing.T) {
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: badWriter(0)})
	(&sliptest.Function{
		Scope:     scope,
		Source:    `(send (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}")) :describe out)`,
		PanicType: slip.StreamErrorSymbol,
	}).Test(t)
}

func TestInstanceReplaceMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(let ((inst (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}"))))
                   (send inst :replace (make-bag "{low:{value:2}}"))
                   (send inst :get "low.value"))`,
		Expect: "2",
	}).Test(t)
}

func TestInstanceReplaceFunction(t *testing.T) {
	(&sliptest.Function{
		Source: `(let ((inst (make-instance 'fhir5:range :data (make-bag "{low:{value:1}}"))))
                   (instance-replace inst (make-bag "{low:{value:2}}"))
                   (send inst :get "low.value"))`,
		Expect: "2",
	}).Test(t)
}

func TestInstanceReplaceNotInstance(t *testing.T) {
	(&sliptest.Function{
		Source:    `(instance-replace 7 (make-bag "{low:{value:2}}"))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceReplaceBadValue(t *testing.T) {
	(&sliptest.Function{
		Source:    `(instance-replace (make-instance 'fhir5:range) 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceReplaceBadBagValue(t *testing.T) {
	(&sliptest.Function{
		Source:    `(instance-replace (make-instance 'fhir5:range) (make-bag "7"))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(instance-replace (make-instance 'fhir5:range) (make-bag "{avg:{value:3}}"))`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
}

func TestInstanceMisc(t *testing.T) {
	rt, ok := slip.FindClass("fhir5:range").(*fhir.Type)
	tt.Equal(t, true, ok)

	bg := bag.Flavor().MakeInstance().(*flavors.Instance)
	bg.Init(slip.NewScope(), slip.List{}, 0)
	bg.Any = map[string]any{"value": int64(3)}

	inst := rt.MakeInstance()
	inst.SetSlotValue(slip.Symbol("low"), bg)
	tt.Equal(t, "{low: {value: 3}}", pretty.SEN(inst.Simplify()))

	names := inst.SlotNames()
	sort.Strings(names)
	tt.Equal(t, "[extension high id low]", pretty.SEN(names))

	tt.Equal(t, true, inst.IsA("fhir5:Range"))
	tt.Equal(t, true, inst.IsA("fhir5:Element"))
	tt.Equal(t, false, inst.IsA("fhir5:patient"))

	fi, ok := inst.(*fhir.Instance)
	tt.Equal(t, true, ok)
	tt.Equal(t, true, fi.HasMethod(":ID"))
	tt.Equal(t, false, fi.HasMethod(":xyz"))

	tt.Equal(t, `{
  combinations: [{from: Type primary: true}]
  name: ":id"
}`, pretty.SEN(fi.GetMethod(":Id")))

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
]`, pretty.SEN(fi.MethodNames()))

	dup := fi.Dup()
	tt.Equal(t, true, fi.Equal(dup))
	fi.SetSynchronized(true)
	tt.Equal(t, true, fi.Synchronized())
	dup = fi.Dup()
	tt.Equal(t, true, fi.Equal(dup))
	fi.SetSynchronized(false)
	tt.Equal(t, true, fi.Equal(dup.Eval(nil, 0)))

	tt.Equal(t, `(let ((inst (make-instance (quote range)))) (setf (slot-value inst (quote extension)) nil)
     (setf (slot-value inst (quote high)) nil) (setf (slot-value inst (quote id)) nil)
     (setf (slot-value inst (quote low)) (("value" . 3))) inst)`, fi.LoadForm().String())

	(&sliptest.Function{
		Source:    `(send (make-instance 'fhir5:range) :quux)`,
		PanicType: slip.InvalidMethodErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(make-instance 'fhir5:range :data (make-instance 'vanilla-flavor))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceSlots(t *testing.T) {
	rt, ok := slip.FindClass("fhir5:range").(*fhir.Type)
	tt.Equal(t, true, ok)

	bg := bag.Flavor().MakeInstance().(*flavors.Instance)
	bg.Init(slip.NewScope(), slip.List{}, 0)
	bg.Any = map[string]any{"value": int64(3)}

	inst := rt.MakeInstance()
	inst.SetSlotValue(slip.Symbol("low"), bg)
	tt.Equal(t, "{low: {value: 3}}", pretty.SEN(inst.Simplify()))

	value, has := inst.SlotValue(slip.Symbol("low"))
	tt.Equal(t, true, has)
	tt.Equal(t, "[[value 3]]", pretty.SEN(value))

	scope := slip.NewScope()
	vanilla := slip.ReadString("(make-instance 'vanilla-flavor)", scope).Eval(scope, nil)
	tt.Panic(t, func() { _ = inst.SetSlotValue(slip.Symbol("low"), vanilla) })

	var qt *fhir.Type
	qt, ok = slip.FindClass("fhir5:quantity").(*fhir.Type)
	tt.Equal(t, true, ok)
	q := qt.MakeInstance()
	q.SetSlotValue(slip.Symbol("value"), slip.Fixnum(5))
	q.SetSlotValue(slip.Symbol("unit"), slip.String("mg"))

	inst.SetSlotValue(slip.Symbol("low"), q)
	tt.Equal(t, "{low: {unit: mg value: 5}}", pretty.SEN(inst.Simplify()))

	inst.SetSlotValue(slip.Symbol("low"), slip.String("{value: 6}"))
	tt.Equal(t, "{low: {value: 6}}", pretty.SEN(inst.Simplify()))

	inst.SlotValue(slip.Symbol("low"))
	tt.Equal(t, "{low: {value: 6}}", pretty.SEN(inst.Simplify()))

	tt.Panic(t, func() { _ = inst.SetSlotValue(slip.Symbol("quux"), q) })
	tt.Panic(t, func() { _ = inst.SetSlotValue(slip.Symbol("low"), slip.Fixnum(3)) })
}

func TestInstanceValidPtrue(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (make-instance 'fhir5:range :data (make-bag "{low:{value:2}}")) :valid-p)`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(send (make-instance 'fhir5:range :data "{low:{value:2}}") :valid-p
                       (lambda (p v m) nil))`,
		Expect: "t",
	}).Test(t)
}

func TestInstanceValidPfalse(t *testing.T) {
	(&sliptest.Function{
		Source: `(let ((inst (make-instance 'fhir5:range :data (make-bag "{quux:{value:2}}") :no-validation t))
                       errors)
                   (send inst :valid-p (lambda (p v m) (setq errors (add errors (list p v m)))))
                   errors)`,
		Expect: `((#<bag-path $.quux> nil "quux is not a property of Range"))`,
	}).Test(t)

	(&sliptest.Function{
		Source: `(let ((inst (make-instance 'fhir5:range :data (make-bag "{quux:{value:2}}") :no-validation t)))
                   (send inst :valid-p))`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)

	(&sliptest.Function{
		Source: `(let ((inst (make-instance 'fhir5:range :data (make-bag "{quux:{value:2}}") :no-validation t)))
                   (send inst :valid-p (lambda (p v m) t)))`,
		Expect: "nil",
	}).Test(t)
}
