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

func TestHTTPCapabilitiesUrlBase(t *testing.T) {
	su, hs := startMockServer(capabilitiesTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope:  scope,
		Source: `(http-capabilities base-url)`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)
			tt.Equal(t, "CapabilityStatement", resource.Class().Name())
		},
	}).Test(t)
}

func capabilitiesTestHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"resourceType": "CapabilityStatement",
		"status":       "active",
		"kind":         "instance",
		"fhirVersion":  "5.0.0",
		"format":       "application/fhir+json",
	}
	w.Header().Set("Content-Type", "application/fhir+json")
	w.Header().Set("Location", r.Host+r.URL.String())
	_ = oj.Write(w, resp)
}
