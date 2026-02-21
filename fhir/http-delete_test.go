// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"net/http"
	"testing"

	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
	"github.com/ohler55/slip/sliptest"
)

func TestHTTPDeleteEmpty(t *testing.T) {
	su, hs := startMockServer(deleteTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope:  scope,
		Source: `(http-delete base-url :type "Patient" :id "P001" :fhir-package 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(204), resp[0])
			tt.Nil(t, resp[1])
		},
	}).Test(t)
}

func TestHTTPDeleteOpOut(t *testing.T) {
	su, hs := startMockServer(deleteOpOutTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope:  scope,
		Source: `(http-delete base-url :type "Patient" :id "P001" :fhir-package 'fhir5)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)
			tt.Equal(t, "OperationOutcome", resource.Class().Name())
		},
	}).Test(t)
}

func deleteTestHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()

	w.WriteHeader(204)
}

func deleteOpOutTestHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()

	resp := map[string]any{
		"resourceType": "OperationOutcome",
		"issue": []any{
			map[string]any{
				"severity": "success",
				"code":     "deleted",
			},
		},
	}
	_ = oj.Write(w, resp)
}
