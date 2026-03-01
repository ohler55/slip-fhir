// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
	"github.com/ohler55/slip/sliptest"
)

func TestHTTPBatchString(t *testing.T) {
	su, hs := startMockServer(batchTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `
(http-batch '("{request: {method: GET url: '/Patient/id-001'}}"
              "{request: {method: GET url: '/Patient/id-002'}}"
              "{request: {method: GET url: '/Patient/id-003'}}")
            base-url
            :fhir-package 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)

			// fmt.Printf("*** %s\n", pretty.SEN(resource.Simplify()))
			tt.Equal(t, `{
  entry: [
    {
      extension: [{valueString: "{request: {method: GET url: \"/Patient/id-001\"}}"}]
      id: id-001
      resourceType: Patient
      response: {etag: "W/1" lastModified: "2026-03-01T10:11:01" status: "200"}
    }
    {
      extension: [{valueString: "{request: {method: GET url: \"/Patient/id-002\"}}"}]
      id: id-002
      resourceType: Patient
      response: {etag: "W/1" lastModified: "2026-03-01T10:11:02" status: "200"}
    }
    {
      extension: [{valueString: "{request: {method: GET url: \"/Patient/id-003\"}}"}]
      id: id-003
      resourceType: Patient
      response: {etag: "W/1" lastModified: "2026-03-01T10:11:03" status: "200"}
    }
  ]
  resourceType: Bundle
  type: batch-response
}`, pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
}

func TestHTTPBatchInstance(t *testing.T) {
	su, hs := startMockServer(batchTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `
(http-batch (list
             (make-instance 'fhir5:Bundle_Entry :data "{request: {method: GET url: '/Patient/id-001'}}")
             (make-instance 'fhir5:Bundle_Entry :data "{request: {method: GET url: '/Patient/id-002'}}"))
            base-url
            :transaction t
            :fhir-package 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)

			// fmt.Printf("*** %s\n", pretty.SEN(resource.Simplify()))
			tt.Equal(t, `{
  entry: [
    {
      extension: [{valueString: "{request: {method: GET url: \"/Patient/id-001\"}}"}]
      id: id-001
      resourceType: Patient
      response: {etag: "W/1" lastModified: "2026-03-01T10:11:01" status: "200"}
    }
    {
      extension: [{valueString: "{request: {method: GET url: \"/Patient/id-002\"}}"}]
      id: id-002
      resourceType: Patient
      response: {etag: "W/1" lastModified: "2026-03-01T10:11:02" status: "200"}
    }
  ]
  resourceType: Bundle
  type: transaction-response
}`, pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
}

func TestHTTPBatchBag(t *testing.T) {
	su, hs := startMockServer(batchTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `
(http-batch (list
             (make-bag "{request: {method: GET url: '/Patient/id-001'}}")
             (make-bag "{request: {method: GET url: '/Patient/id-002'}}"))
            base-url
            :fhir-package 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)

			// fmt.Printf("*** %s\n", pretty.SEN(resource.Simplify()))
			tt.Equal(t, `{
  entry: [
    {
      extension: [{valueString: "{request: {method: GET url: \"/Patient/id-001\"}}"}]
      id: id-001
      resourceType: Patient
      response: {etag: "W/1" lastModified: "2026-03-01T10:11:01" status: "200"}
    }
    {
      extension: [{valueString: "{request: {method: GET url: \"/Patient/id-002\"}}"}]
      id: id-002
      resourceType: Patient
      response: {etag: "W/1" lastModified: "2026-03-01T10:11:02" status: "200"}
    }
  ]
  resourceType: Bundle
  type: batch-response
}`, pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
}

func TestHTTPBatchNotList(t *testing.T) {
	(&sliptest.Function{
		Source:    `(http-batch t "http://localhost:8888" :fhir-package 'fhir5)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestHTTPBatchNotEntry(t *testing.T) {
	(&sliptest.Function{
		Source:    `(http-batch (list (make-instance 'vanilla-flavor)) "http://localhost:8888" :fhir-package 'fhir5)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func batchTestHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var res any
	if 0 < len(body) {
		res = oj.MustParse(body)
	}
	var entries []any
	// It's a mock so assume all entries are GETs for a Patient.
	for i, entry := range jp.C("entry").W().Get(res) {
		entries = append(entries, map[string]any{
			"resourceType": "Patient",
			"id":           fmt.Sprintf("id-%03d", i+1),
			"extension": []any{
				map[string]any{
					"valueString": pretty.SEN(entry),
				},
			},
			"response": map[string]any{
				"status":       "200",
				"etag":         "W/1",
				"lastModified": fmt.Sprintf("2026-03-01T10:11:%02d", i+1),
			},
		})
	}
	resp := map[string]any{
		"resourceType": "Bundle",
		"type":         "batch-response",
		"entry":        entries,
	}
	if jp.C("type").First(res) == "transaction" {
		resp["type"] = "transaction-response"
	}
	w.Header().Set("Content-Type", "application/fhir+json")
	w.Header().Set("Location", r.Host+r.URL.String())
	_ = oj.Write(w, resp)
}
