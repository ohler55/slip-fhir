// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
	"github.com/ohler55/slip/sliptest"
)

func TestHTTPOperationGET(t *testing.T) {
	su, hs := startMockServer(operationTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(http-operation "reflect" nil
                                 base-url
                                 :type "Patient"
                                 :id "id-123"
                                 :params '("arg1" "1" "arg2" "2")
                                 :fhir-package 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)

			tt.Equal(t, `{
  details: {text: "/Patient/id-123/$reflect?arg1=1&arg2=2 - "}
  issue: {code: success severity: success}
  resourceType: OperationOutcome
}`, pretty.SEN(resource.Simplify()))
			// fmt.Printf("*** %s\n", pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
}

func TestHTTPOperationPOSTURL(t *testing.T) {
	su, hs := startMockServer(operationTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(http-operation "reflect" '("arg1" "1" "arg2" "2")
                                 base-url
                                 :type "Patient"
                                 :fhir-package 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)

			tt.Equal(t, `{
  details: {text: "/Patient/$reflect - arg1=1&arg2=2"}
  issue: {code: success severity: success}
  resourceType: OperationOutcome
}`, pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
}

func TestHTTPOperationPOSTParameters(t *testing.T) {
	su, hs := startMockServer(operationTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(http-operation "reflect"
                                 (make-instance 'fhir5:Parameters
                                                :data "{resourceType:Parameters parameter:[{name:arg1}]}")
                                 base-url
                                 :fhir-package 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)

			tt.Equal(t, `{
  details: {text: "/$reflect - {parameter: [{name: arg1}] resourceType: Parameters}"}
  issue: {code: success severity: success}
  resourceType: OperationOutcome
}`, pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
}

func TestHTTPOperationNotParameters(t *testing.T) {
	(&sliptest.Function{
		Source: `(http-operation "reflect" (make-instance 'fhir5:Patient)
                                 "http://localhost:8888"
                                 :fhir-package 'fhir5)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestHTTPOperationBadArgs(t *testing.T) {
	(&sliptest.Function{
		Source:    `(http-operation "reflect" t "http://localhost:8888" :fhir-package 'fhir5)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func operationTestHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	if 0 < len(body) {
		if j, err := oj.Parse(body); err == nil {
			body = []byte(pretty.SEN(j, 100.5))
		}
	}
	resp := map[string]any{
		"resourceType": "OperationOutcome",
		"issue": map[string]any{
			"severity": "success",
			"code":     "success",
		},
		"details": map[string]any{
			"text": r.URL.String() + " - " + string(body),
		},
	}
	w.Header().Set("Content-Type", "application/fhir+json")
	w.Header().Set("Location", r.Host+r.URL.String())
	_ = oj.Write(w, resp)
}
