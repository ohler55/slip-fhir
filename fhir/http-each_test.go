// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
	"github.com/ohler55/slip/sliptest"
)

func TestHTTPEachOne(t *testing.T) {
	su, hs := startMockServer(eachTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(let (resources)
                   (http-each (lambda (r) (addf resources r))
                              base-url
                              :type "Patient"
                              :id "id-123"
                              :fhir-package 'fhir5)
                   resources)`,
		Validate: func(t *testing.T, v slip.Object) {
			resources, _ := v.(slip.List)
			tt.Equal(t, 1, len(resources))
			resource, _ := resources[0].(*fhir.Instance)
			tt.NotNil(t, resource)
			tt.Equal(t, "Patient", resource.Class().Name())
		},
	}).Test(t)
}

func TestHTTPEachMulti(t *testing.T) {
	su, hs := startMockServer(eachTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(let (resources)
                   (http-each (lambda (r) (addf resources r))
                              base-url
                              :type "Patient"
                              :timeout 1.5
                              :fhir-package 'fhir5)
                   resources)`,
		Validate: func(t *testing.T, v slip.Object) {
			resources, _ := v.(slip.List)
			tt.Equal(t, 5, len(resources))
			for _, element := range resources {
				resource, _ := element.(*fhir.Instance)
				tt.NotNil(t, resource)
				tt.Equal(t, "Patient", resource.Class().Name())
			}
		},
	}).Test(t)
}

func TestHTTPEachLimit(t *testing.T) {
	su, hs := startMockServer(eachTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(let (resources)
                   (http-each (lambda (r) (addf resources r))
                              base-url
                              :type "Patient"
                              :limit 2
                              :fhir-package 'fhir5)
                   resources)`,
		Validate: func(t *testing.T, v slip.Object) {
			resources, _ := v.(slip.List)
			tt.Equal(t, 2, len(resources))
			for _, element := range resources {
				resource, _ := element.(*fhir.Instance)
				tt.NotNil(t, resource)
				tt.Equal(t, "Patient", resource.Class().Name())
			}
		},
	}).Test(t)
	(&sliptest.Function{
		Scope:     scope,
		Source:    `(http-each (lambda (r) nil) base-url :type "Patient" :limit t)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestHTTPEachCompartment(t *testing.T) {
	su, hs := startMockServer(eachTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(let (resources)
                   (http-each (lambda (r) (addf resources r))
                              base-url
                              :compartment "Patient"
                              :id "P001"
                              :limit 2
                              :fhir-package 'fhir5)
                   resources)`,
		Validate: func(t *testing.T, v slip.Object) {
			resources, _ := v.(slip.List)
			tt.Equal(t, 2, len(resources))
			var (
				hasEncounter      bool
				hasServiceRequest bool
			)
			for _, element := range resources {
				resource, _ := element.(*fhir.Instance)
				tt.NotNil(t, resource)
				switch resource.Class().Name() {
				case "Encounter":
					hasEncounter = true
				case "ServiceRequest":
					hasServiceRequest = true
				}
			}
			tt.Equal(t, true, hasEncounter && hasServiceRequest)
		},
	}).Test(t)
}

func TestHTTPEachBadNextPage(t *testing.T) {
	su, hs := startMockServer(eachTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(let (resources)
                   (http-each (lambda (r) (addf resources r))
                              base-url
                              :compartment "Patient"
                              :id "quux"
                              :fhir-package 'fhir5)
                   resources)`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
}

func TestHTTPEachEmpty(t *testing.T) {
	su, hs := startMockServer(eachTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope:  scope,
		Source: `(http-each (lambda (r) nil) base-url :type "Empty")`,
		Expect: "nil",
	}).Test(t)
}

func TestHTTPEachError(t *testing.T) {
	su, hs := startMockServer(eachTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope:  scope,
		Source: `(http-each (lambda (r) nil) base-url :type "Status404" :history "true")`,
		Expect: "nil",
		Validate: func(t *testing.T, v slip.Object) {
			resource, _ := v.(*fhir.Instance)
			tt.NotNil(t, resource)
			tt.Equal(t, "OperationOutcome", resource.Class().Name())
		},
	}).Test(t)
}

func eachTestHandler(w http.ResponseWriter, r *http.Request) {
	var resp any
	switch r.URL.String() {
	case "/Patient/id-123":
		resp = map[string]any{
			"resourceType": "Patient",
			"id":           "id-123",
		}
	case "/Patient":
		resp = map[string]any{
			"resourceType": "Bundle",
			"entry": []any{
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Patient",
						"id":           "p01",
					},
				},
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Patient",
						"id":           "p02",
					},
				},
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Patient",
						"id":           "p03",
					},
				},
			},
			"link": []any{
				map[string]any{
					"relation": "next",
					"url":      fmt.Sprintf("http://%s/next-page", r.Host),
				},
			},
		}
	case "/Patient/P001/%2A":
		resp = map[string]any{
			"resourceType": "Bundle",
			"entry": []any{
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Encounter",
						"id":           "e01",
					},
				},
				map[string]any{
					"resource": map[string]any{
						"resourceType": "ServiceRequest",
						"id":           "sr01",
					},
				},
			},
			"link": []any{
				map[string]any{
					"relation": "next",
					"url":      fmt.Sprintf("http://%s/next-page", r.Host),
				},
			},
		}
	case "/Patient/quux/%2A":
		resp = map[string]any{
			"resourceType": "Bundle",
			"entry": []any{
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Encounter",
						"id":           "e01",
					},
				},
			},
			"link": []any{
				map[string]any{
					"relation": "next",
					"url":      "...",
				},
			},
		}
	case "/Empty":
		// nil resp
	case "/Status404/_history":
		resp = map[string]any{
			"resourceType": "Bundle",
			"entry": []any{
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Patient",
						"id":           "p01",
					},
				},
			},
			"link": []any{
				map[string]any{
					"relation": "next",
					"url":      fmt.Sprintf("http://%s/Status404Next", r.Host),
				},
			},
		}
	case "/Status404Next":
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
	default:
		resp = map[string]any{
			"resourceType": "Bundle",
			"entry": []any{
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Patient",
						"id":           "p04",
					},
				},
				map[string]any{
					"resource": map[string]any{
						"resourceType": "Patient",
						"id":           "p05",
					},
				},
			},
		}
	}
	w.Header().Set("Content-Type", "application/fhir+json")
	_ = oj.Write(w, resp)
}
