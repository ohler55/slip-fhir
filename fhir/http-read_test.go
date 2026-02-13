// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip-fhir/fhir"
	"github.com/ohler55/slip/sliptest"
)

func TestHTTPReadWithServer(t *testing.T) {
	port := availablePort()
	hs := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(readTestHandler),
	}
	go func() { _ = hs.ListenAndServe() }()
	defer func() { _ = hs.Close() }()

	su := fmt.Sprintf("http://localhost:%d", port)
	start := time.Now()
	for time.Since(start) < time.Second*2 {
		time.Sleep(time.Millisecond * 50)
		if resp, err := http.Get(su); err == nil {
			_ = resp.Body.Close()
			break
		}
	}
	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope:  scope,
		Source: `(http-read base-url :type "Patient" :id "id-123")`, // TBD add some other params and headers
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			resource, _ := resp[1].(*fhir.Instance)
			tt.NotNil(t, resource)
			tt.Equal(t, "Patient", resource.Class().Name())

			sv, _ := resource.SlotValue(slip.Symbol("id"))
			tt.Equal(t, `"id-123"`, sv.String())
			ex, _ := jp.C("extension").N(0).C("valueString").First(resource.Simplify()).(string)
			headers := sen.MustParse([]byte(ex))
			tt.Equal(t,
				[]any{"application/fhir+json", "application/json+fhir", "application/json"},
				jp.C("Accept").First(headers))
			tt.Equal(t, []any{"application/fhir+json"}, jp.C("Content-Type").First(headers))
			// Verify headers were received. Check for Content-Type.
			tt.Equal(t, `/Content-Type\" \"application\/fhir\+json/`, resp[2].String())
		},
	}).Test(t)
}

func readTestHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"resourceType": "Patient",
		"id":           "id-123",
		"extension": []any{
			map[string]any{
				"valueString": pretty.SEN(r.Header),
			},
		},
	}
	w.Header().Set("Content-Type", "application/fhir+json")
	w.Header().Set("Location", r.Host+r.URL.String())
	_ = oj.Write(w, resp)
}
