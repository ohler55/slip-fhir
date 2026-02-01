// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"testing"

	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestInstanceIDMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (make-instance 'patient) :id)`,
		Expect: "/[0-9]+/",
	}).Test(t)
}

func TestInstanceIDFunction(t *testing.T) {
	(&sliptest.Function{
		Source: `(instance-id (make-instance 'patient))`,
		Expect: "/[0-9]+/",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(instance-id 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceTypeFunction(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (make-instance 'patient) :type)`,
		Expect: "#<fhir:Type Patient>",
	}).Test(t)
}

func TestInstanceWhichOperationsMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(sort (send (make-instance 'patient) :which-operations))`,
		Expect: `(:class :data :describe :equal :get :id :init :operation-handled-p :print-self
        :replace :set :type :validate :which-operations)`,
	}).Test(t)
}

func TestInstanceOperationHandledPMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(send (make-instance 'patient) :operation-handled-p :id)`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(send (make-instance 'patient) :operation-handled-p :quux)`,
		Expect: "nil",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(send (make-instance 'patient) :operation-handled-p 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceEqualMethod(t *testing.T) {
	(&sliptest.Function{
		Source: `(let ((pat (make-instance 'patient)))
                   (send pat :equal pat))`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(let ((pat (make-instance 'range :data (make-bag "{low:{value:1} high:{value:2}}")))
                       (pat2 (make-instance 'range :data (make-bag "{low:{value:1} high:{value:2}}"))))
                   (send pat :equal pat2))`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(send (make-instance 'patient) :equal 7)`,
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
		Source: `(send (make-instance 'range :data (make-bag "{low:{value:1} high:{value:2}}")) :get "low.value")`,
		Expect: "1",
	}).Test(t)
}

func TestInstanceGetFunction(t *testing.T) {
	(&sliptest.Function{
		Source: `(instance-get (make-instance 'range :data (make-bag "{low:{value:1} high:{value:2}}")) "low.value")`,
		Expect: "1",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(instance-get 7 "low.value")`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'HumanName :data (make-bag "{given:[bill bob]}")) 7)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestInstanceGetPathList(t *testing.T) {
	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'range :data (make-bag "{low:{value:1} high:{value:2}}")) '(low "value"))`,
		Expect: "1",
	}).Test(t)
	(&sliptest.Function{
		Source: `(send
                   (instance-get
                     (make-instance 'range :data (make-bag "{low:{value:1} high:{value:2}}")) "low")
                   :write nil)`,
		Expect: `"{value: 1}"`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'HumanName :data (make-bag "{given:[bill bob]}")) '(given 1))`,
		Expect: `"bob"`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'HumanName :data (make-bag "{given:[bill bob]}")) '(given nil))`,
		Expect: `"bill"`,
	}).Test(t)

	(&sliptest.Function{
		Source: `(instance-get
                   (make-instance 'HumanName :data (make-bag "{given:[bill bob]}")) '(given t))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}
