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
	"github.com/ohler55/slip/pkg/flavors"
	"github.com/ohler55/slip/sliptest"
)

func startMockServer(handler func(w http.ResponseWriter, r *http.Request)) (string, *http.Server) {
	port := availablePort()
	hs := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}
	go func() { _ = hs.ListenAndServe() }()

	su := fmt.Sprintf("http://localhost:%d", port)
	start := time.Now()
	for time.Since(start) < time.Second*2 {
		time.Sleep(time.Millisecond * 50)
		if resp, err := http.Get(su); err == nil {
			_ = resp.Body.Close()
			break
		}
	}
	return su, &hs
}

func TestHTTPReadUrlBase(t *testing.T) {
	su, hs := startMockServer(readTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(http-read base-url
                            :type "Patient"
                            :id "id-123"
                            :version "v3"
                            :headers '(("ETag" "W/") ("Accept" "application/json" "palin/text"))
                            :params '("_pretty" "true")
                            :timeout 1.5
                            :fhir-package 'fhir5)`,
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
			tt.Equal(t, `/\/Patient\/id-123\/_history\/v3/`, resp[2].String())
		},
	}).Test(t)
}

func TestHTTPReadListBase(t *testing.T) {
	su, hs := startMockServer(readTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(http-read
                   (list :url base-url
                         :type "Patient"
                         :id "id-123"
                         :headers '(("ETag" "W/"))
                         :params '("_elements" "id,extension" "_pretty" "true")
                         :timeout 1.5
                         :fhir-package 'fhir5))`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			bg, _ := resp[1].(*flavors.Instance)
			tt.NotNil(t, bg)
			tt.Equal(t, "bag-flavor", bg.Class().Name())

			id := jp.C("id").First(bg.Any)
			tt.Equal(t, "id-123", id)
		},
	}).Test(t)
}

func TestHTTPReadNotResource(t *testing.T) {
	su, hs := startMockServer(readTestNotResourceHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope:  scope,
		Source: `(http-read base-url :type "Patient" :id "q123")`,
		Validate: func(t *testing.T, v slip.Object) {
			resp, _ := v.(slip.List)
			tt.Equal(t, 3, len(resp))
			tt.Equal(t, slip.Fixnum(200), resp[0])
			bg, _ := resp[1].(*flavors.Instance)
			tt.NotNil(t, bg)
			tt.Equal(t, "bag-flavor", bg.Class().Name())

			id := jp.C("id").First(bg.Any)
			tt.Equal(t, "q123", id)
		},
	}).Test(t)
}

func TestHTTPReadBadBase(t *testing.T) {
	(&sliptest.Function{
		Source:    `(http-read t)`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestHTTPReadBadUrl(t *testing.T) {
	(&sliptest.Function{
		Source:    `(http-read "hzzz://")`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(http-read "http://\t\n")`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
}

func TestHTTPReadBadParams(t *testing.T) {
	(&sliptest.Function{
		Source:    `(http-read "http://localhost:1234" :params '("xyz"))`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(http-read "http://localhost:1234" :params '(:xyz "xx"))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(http-read "http://localhost:1234" :params '("xyz" 7))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
}

func TestHTTPReadBadHeader(t *testing.T) {
	(&sliptest.Function{
		Source:    `(http-read "http://localhost:1234" :headers '(("Content-Type")))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(http-read "http://localhost:1234" :headers '(("Content-Type" 7)))`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(http-read "http://localhost:1234" :headers '((7 "xyz")))`,
		PanicType: slip.TypeErrorSymbol,
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

func readTestNotResourceHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"resourceType": "Quux",
		"id":           "q123",
	}
	w.Header().Set("Content-Type", "application/fhir+json")
	w.Header().Set("Location", r.Host+r.URL.String())
	_ = oj.Write(w, resp)
}
