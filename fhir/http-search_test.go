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

func TestHTTPSearchGET(t *testing.T) {
	su, hs := startMockServer(searchTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(let (resources)
                   (http-search (lambda (r) (addf resources r))
                              base-url
                              :type "Patient"
                              :params '("given" "Pete")
                              :fhir-package 'fhir5)
                   resources)`,
		Validate: func(t *testing.T, v slip.Object) {
			resources, _ := v.(slip.List)
			tt.Equal(t, 2, len(resources))
			resource, _ := resources[0].(*fhir.Instance)
			tt.NotNil(t, resource)
			tt.Equal(t, `{
  id: p01
  name: [
    {family: Parrot given: [Pete]}
  ]
  resourceType: Patient
}`, pretty.SEN(resource.Simplify()))
			resource, _ = resources[1].(*fhir.Instance)
			tt.NotNil(t, resource)
			tt.Equal(t, `{
  id: p02
  name: [
    {family: Porcupine given: [Pete]}
  ]
  resourceType: Patient
}`, pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
	(&sliptest.Function{
		Scope: scope,
		Source: `(let (resources)
                   (http-search (lambda (r) (addf resources r))
                              base-url
                              :type "Quux"
                              :fhir-package 'fhir5)
                   resources)`,
		Validate: func(t *testing.T, v slip.Object) {
			resources, _ := v.(slip.List)
			tt.Equal(t, 1, len(resources))
			resource, _ := resources[0].(*fhir.Instance)
			tt.NotNil(t, resource)
			tt.Equal(t, `{
  issue: [{code: security severity: error}]
  resourceType: OperationOutcome
}`, pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
}

func TestHTTPSearchPOST(t *testing.T) {
	su, hs := startMockServer(searchTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(let (resources)
                   (http-search (lambda (r) (addf resources r))
                              base-url
                              :type "Patient"
                              :query '("given" "Pete")
                              :limit 1
                              :fhir-package 'fhir5)
                   resources)`,
		Validate: func(t *testing.T, v slip.Object) {
			resources, _ := v.(slip.List)
			tt.Equal(t, 1, len(resources))
			resource, _ := resources[0].(*fhir.Instance)
			tt.NotNil(t, resource)
			tt.Equal(t, `{
  id: p01
  name: [
    {family: Parrot given: [Pete]}
  ]
  resourceType: Patient
}`, pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
}

func TestHTTPSearchEmpty(t *testing.T) {
	su, hs := startMockServer(eachTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope:  scope,
		Source: `(http-search (lambda (r) nil) base-url :type "Empty")`,
		Expect: "nil",
	}).Test(t)
}

func TestHTTPSearchBadQuery(t *testing.T) {
	su, hs := startMockServer(eachTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope:     scope,
		Source:    `(http-search (lambda (r) nil) base-url :type "Patient" :query t)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestHTTPSearchBadLimit(t *testing.T) {
	su, hs := startMockServer(searchTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(http-search (lambda (r) nil)
                              base-url
                              :type "Patient"
                              :query '("given" "Pete")
                              :limit t
                              :fhir-package 'fhir5)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func searchTestHandler(w http.ResponseWriter, r *http.Request) {
	var resp any
	switch r.URL.String() {
	case "/Patient?given=Pete":
		resp = map[string]any{
			"resourceType": "Bundle",
			"entry": []any{
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Patient",
						"id":           "p01",
						"name":         []any{map[string]any{"given": []any{"Pete"}, "family": "Parrot"}},
					},
				},
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Patient",
						"id":           "p02",
						"name":         []any{map[string]any{"given": []any{"Pete"}, "family": "Porcupine"}},
					},
				},
			},
		}
	case "/Patient/_search":
		if r.Method != "POST" {
			_, _ = w.Write([]byte("Must be a POST."))
			return
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) != "given=Pete" {
			_, _ = w.Write([]byte("Missing or wrong query."))
			return
		}
		resp = map[string]any{
			"resourceType": "Bundle",
			"entry": []any{
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Patient",
						"id":           "p01",
						"name":         []any{map[string]any{"given": []any{"Pete"}, "family": "Parrot"}},
					},
				},
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Patient",
						"id":           "p02",
						"name":         []any{map[string]any{"given": []any{"Pete"}, "family": "Porcupine"}},
					},
				},
			},
		}
	case "/Quux":
		resp = map[string]any{
			"resourceType": "OperationOutcome",
			"issue": []any{
				map[string]any{
					"severity": "error",
					"code":     "security",
				},
			},
		}
		w.WriteHeader(404)
	case "/Empty":
		// nil resp
	default:
		resp = map[string]any{}
	}
	w.Header().Set("Content-Type", "application/fhir+json")
	_ = oj.Write(w, resp)
}
