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

func TestHTTPPatchRFC6902(t *testing.T) {
	su, hs := startMockServer(patchTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(http-patch '((:op add :path "/birthDate" :value "1956-02-09"))
                             base-url
                            :type "Patient"
                            :id "id-123"
                            :fhir-package 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)

			tt.Equal(t, `{
  birthDate: "1956-02-09"
  extension: [
    {
      valueString: "/Patient/id-123 - [{op: add path: \"/birthDate\" value: \"1956-02-09\"}]"
    }
  ]
  id: id-123
  resourceType: Patient
}`, pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
}

func TestHTTPPatchParameters(t *testing.T) {
	su, hs := startMockServer(patchTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `
(http-patch (make-instance 'fhir5:Parameters
                           :data "{
                                    resourceType: Parameters,
                                    parameter:[
                                      {
                                        name: operation
                                        part: [
                                          {name: type valueString: add}
                                          {name: path valueString: Patient}
                                          {name: name valueString: birthDate}
                                          {name: value valueString: '1956-02-09'}
                                        ]
                                      }
                                    ]
                                  }")
            base-url
            :type "Patient"
            :id "id-123"
            :fhir-package 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)

			expect := `{
  birthDate: "1956-02-09"
  extension: [
    {
      valueString: "/Patient/id-123 - {
  parameter: [
    {
      name: operation
      part: [
        {name: type valueString: add}
        {name: path valueString: Patient}
        {name: name valueString: birthDate}
        {name: value valueString: \"1956-02-09\"}
      ]
    }
  ]
  resourceType: Parameters
}"
    }
  ]
  id: id-123
  resourceType: Patient
}`
			tt.Equal(t, expect, pretty.SEN(resource.Simplify()))
		},
	}).Test(t)
}

func TestHTTPPatchBadPatch(t *testing.T) {
	(&sliptest.Function{
		Source: `(http-patch t
                             "http://localhost:8888"
                            :type "Patient"
                            :id "id-123"
                            :fhir-package 'fhir5)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source: `(http-patch '((::op add :path "/birthDate" value "1956-02-09"))
                             "http://localhost:8888"
                            :type "Patient"
                            :id "id-123"
                            :fhir-package 'fhir5)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source: `(http-patch '(7 t)
                             "http://localhost:8888"
                            :type "Patient"
                            :id "id-123"
                            :fhir-package 'fhir5)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source: `(http-patch (make-instance 'fhir5:Patient)
                             "http://localhost:8888"
                            :type "Patient"
                            :id "id-123"
                            :fhir-package 'fhir5)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func patchTestHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	if 0 < len(body) {
		if j, err := oj.Parse(body); err == nil {
			body = []byte(pretty.SEN(j, 100.5))
		}
	}
	resp := map[string]any{
		"resourceType": "Patient",
		"id":           "id-123",
		"birthDate":    "1956-02-09",
		"extension": []any{
			map[string]any{
				"valueString": r.URL.String() + " - " + string(body),
			},
		},
	}
	w.Header().Set("Content-Type", "application/fhir+json")
	w.Header().Set("Location", r.Host+r.URL.String())
	_ = oj.Write(w, resp)
}
