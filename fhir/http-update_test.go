// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"fmt"
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

func TestHTTPUpdateOk(t *testing.T) {
	su, hs := startMockServer(updateTestHandler)
	defer func() { _ = hs.Close() }()

	scope := slip.NewScope()
	scope.Let("base-url", slip.String(su))
	(&sliptest.Function{
		Scope: scope,
		Source: `(http-update "{resourceType: Patient, id:P002, name:[{given:[Rocky] family:Racoon}]}"
                              base-url :fhir-package 'fhir5)`,
		Validate: validateUpdateResponse,
	}).Test(t)
	(&sliptest.Function{
		Scope: scope,
		Source: `(http-update (make-instance 'fhir5:Patient
                                             :data "{resourceType: Patient
                                                     id:P002
                                                     name:[{given:[Rocky] family:Racoon}]}")
                              base-url :fhir-package 'fhir5)`,
		Validate: validateUpdateResponse,
	}).Test(t)
	(&sliptest.Function{
		Scope: scope,
		Source: `(http-update (make-bag "{resourceType: Patient, id:P002, name:[{given:[Rocky] family:Racoon}]}")
                              base-url :fhir-package 'fhir5)`,
		Validate: validateUpdateResponse,
	}).Test(t)
}

func TestHTTPUpdateBadResource(t *testing.T) {
	(&sliptest.Function{
		Source:    `(http-update t "http://localhost:1234")`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(http-update (make-instance 'vanilla-flavor) "http://localhost:1234")`,
		PanicType: slip.TypeErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source:    `(http-update "{name:[{given:[Rocky] family:Racoon}]}" "http://localhost:1234")`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
	(&sliptest.Function{
		Source: `(http-update "{resourceType:Patient name:[{given:[Rocky] family:Racoon}]}"
                              "http://localhost:1234")`,
		PanicType: slip.ErrorSymbol,
	}).Test(t)
}

func updateTestHandler(w http.ResponseWriter, r *http.Request) {
	version := "v02"
	defer func() { _ = r.Body.Close() }()

	resp := oj.MustLoad(r.Body)

	_ = jp.C("meta").C("versionID").Set(resp, version)
	_ = jp.C("meta").C("lastUpdated").Set(resp, "2026-02-17T21:36:27Z")

	w.Header().Set("Content-Type", "application/fhir+json")
	w.Header().Set("Location", fmt.Sprintf("%s%s/_history/%s", r.Host, r.URL.String(), version))
	w.Header().Set("Last-Modified", "Tue, 17 Feb 2026 21:36:27 GMT")
	w.Header().Set("ETag", fmt.Sprintf("W/%q", version))
	w.WriteHeader(201)
	_ = oj.Write(w, resp)
}

func validateUpdateResponse(t *testing.T, v slip.Object) {
	resp, _ := v.(slip.List)
	tt.Equal(t, 3, len(resp))
	tt.Equal(t, slip.Fixnum(201), resp[0])
	inst, _ := resp[1].(*fhir.Instance)
	tt.NotNil(t, inst)
	tt.Equal(t, "Patient", inst.Class().Name())

	id, has := inst.SlotValue(slip.Symbol("id"))
	tt.Equal(t, true, has)
	tt.NotNil(t, id)
	tt.Equal(t, `{
  id: P002
  meta: {lastUpdated: "2026-02-17T21:36:27Z" versionID: v02}
  name: [
    {family: Racoon given: [Rocky]}
  ]
  resourceType: Patient
}`, pretty.SEN(inst))

	// The case of tags should not matter but check anyway.
	tt.Equal(t, `/ETag\" \"W\/\"v02\"/`, resp[2].String())
	tt.Equal(t, `/\/Patient\/P002\/_history\/v02/`, resp[2].String())
	tt.Equal(t, `/Last-Modified\" \"Tue, 17 Feb 2026 21:36:27 GMT\"/`, resp[2].String())
}
