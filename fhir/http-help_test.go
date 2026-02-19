// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ohler55/ojg/tt"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/sliptest"
)

func testHTTPHelp(t *testing.T, topic string, includes ...string) {
	var out strings.Builder
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: &out})
	(&sliptest.Function{
		Scope:  scope,
		Source: fmt.Sprintf(`(let ((*standard-output* out)) (http-help %s))`, topic),
		Expect: "nil",
	}).Test(t)
	help := out.String()
	for _, inc := range includes {
		tt.Equal(t, true, strings.Contains(help, inc), "%s", inc)
	}
}

func TestHTTPHelp(t *testing.T) {
	// Check enough of the content of each to verify it at least displays
	// something other than empty space.
	testHTTPHelp(t, "", "functions", "resources", "datatypes", "primitives")
	testHTTPHelp(t, "functions", "http-read", "http-create", "http-update")
	testHTTPHelp(t, "resources", "Account", "Patient", "Encounter")
	testHTTPHelp(t, "datatypes", "Address", "CodeableConcept", "Range")
	testHTTPHelp(t, "backbones", "Patient_Contact", "Patient_Link")
	testHTTPHelp(t, "primitives", "fhir5:integer", "fhir5:string", "fhir5:code")
	testHTTPHelp(t, "explore", "describe ", "describe-type", "http-help")
	testHTTPHelp(t, "summary", "http-read", "/[type]/[id]")
	testHTTPHelp(t, "headers", "ETag", "Accept", "Last-Modified")
	testHTTPHelp(t, "parameters", " summary", " filter", " pretty")
	// testHTTPHelp(t, "search", "")
	// testHTTPHelp(t, "history", "")
	// testHTTPHelp(t, "compartment", "")
	testHTTPHelp(t, "read-example", "(defvar resp", `(cadr (assoc "Location" (nth 2 resp)))`)
	// testHTTPHelp(t, "create-example", "")
	// testHTTPHelp(t, "update-example", "")
	// testHTTPHelp(t, "delete-example", "")
	// testHTTPHelp(t, "patch-example", "")
	// testHTTPHelp(t, "batch-example", "")
}

func TestHTTPHelpWriteError(t *testing.T) {
	scope := slip.NewScope()
	scope.Let("out", &slip.OutputStream{Writer: badWriter(0)})
	(&sliptest.Function{
		Scope:     scope,
		Source:    `(let ((*standard-output* out)) (http-help))`,
		PanicType: slip.StreamErrorSymbol,
	}).Test(t)
}
