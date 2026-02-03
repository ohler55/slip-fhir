// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"testing"

	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func TestValidPBasic(t *testing.T) {
	(&sliptest.Function{
		Source: `(valid-p (make-instance 'range :data (make-bag "{low:{value:1} high:{value:2}}")))`,
		Expect: "t",
	}).Test(t)
}

func TestValidPWithType(t *testing.T) {
	(&sliptest.Function{
		Source: `(valid-p (make-bag "{low:{value:1} high:{value:2}}") :type 'range)`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(valid-p 7 :type 'fhir:integer)`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(valid-p 7 :type t)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(valid-p 7)`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(valid-p 7.5 :type 'fhir:integer)`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
}

func TestValidPWithOnError(t *testing.T) {
	(&sliptest.Function{
		Source: `(with-output-to-string (s)
                   (valid-p 7 :type 'fhir:integer :on-error (lambda (p v m) (format s "~A: ~A" p m))))`,
		Expect: `""`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(with-output-to-string (s)
                   (valid-p 7.5 :type 'fhir:integer :on-error (lambda (p v m) (format s "~A: ~A" p m))))`,
		Expect: `"#<bag-path $>: a decimal is not a valid type for a integer"`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(with-output-to-string (s)
                   (valid-p (make-bag "{low:{value:{a:1}}}")
                            :type 'fhir:range
                            :on-error (lambda (p v m) (format s "~A: ~A ~A" p v m))))`,
		Expect: `/#<bag-flavor .*is not a valid type for a decimal/`,
	}).Test(t)
	(&sliptest.Function{
		Source: `(valid-p 7.5 :type 'fhir:integer :on-error (lambda (p v m) t))`,
		Expect: "nil",
	}).Test(t)
	(&sliptest.Function{
		Source:    `(valid-p 7.5 :type 'fhir:integer :on-error (lambda (p v) t))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestValidPWithOnErrorFunc(t *testing.T) {
	scope := slip.NewScope()
	defer func() {
		_ = slip.ReadString("(makunbound 'on-err-func)", scope).Eval(scope, nil)
	}()
	_ = slip.ReadString("(defun on-err-func (p v m) t)", scope).Eval(scope, nil)
	(&sliptest.Function{
		Source: `(valid-p 7 :type 'fhir:integer :on-error #'on-err-func)`,
		Expect: "t",
	}).Test(t)
	(&sliptest.Function{
		Source: `(valid-p 7.5 :type 'fhir:integer :on-error #'on-err-func)`,
		Expect: "nil",
	}).Test(t)
	(&sliptest.Function{
		Source: `(valid-p 7 :type 'fhir:string :on-error 'on-err-func)`,
		Expect: "nil",
	}).Test(t)
}

func TestValidPWithOnErrorList(t *testing.T) {
	(&sliptest.Function{
		Source: `(valid-p 7 :type 'fhir:integer :on-error '(lambda (p v m) t))`,
		Expect: "t",
	}).Test(t)
}

func TestValidPNotBag(t *testing.T) {
	(&sliptest.Function{
		Source:    `(valid-p (make-instance 'vanilla-flavor))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}
