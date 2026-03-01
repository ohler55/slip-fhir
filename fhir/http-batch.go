// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"net/http"

	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
)

func initHTTPBatch() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := HTTPBatch{Function: slip.Function{Name: "http-batch", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "http-batch",
			Args: []*slip.DocArg{
				{
					Name: "entries",
					Type: "list",
					Text: `A list of __Bundle_Entry__ instances. The request property must be present and
in the case of a PUT or POST a resource property is also needed.`,
				},
				{
					Name: "base",
					Type: "string|property-list",
					Text: `Identifies the FHIR server to connect to. It may also include default
or base values if a property list. Any of the _&key_ arguments can be included in the property
list and will serve as a base or defaults for the _&key_ arguments.`,
				},
				{Name: "&key"},
				{
					Name: "headers",
					Type: "assoc-list",
					Text: `If present, the values in the association list are merged and supersede any
_:headers_ in the _base_. The __car__ of each element of the list is header field key and the remaining
values in the list element are the values for header field. An example is
  (("Content-Type" "application/fhir+json") ("ETag" "W/"))


The headers FHIR servers should handle are describe at https://www.hl7.org/fhir//http.html#Http-Headers.`,
				},
				{
					Name: "params",
					Type: "property-list",
					Text: `If present, the values in the property are merged and supersede any
_:params_ in the _base_. The property list indicators should be strings that will be used as
values in the request URL query. Multiple values with the same key are allowed.`,
				},
				{
					Name: "timeout",
					Type: "real",
					Text: `The number of seconds to wait before giving and returning a 408,
Request Timeout code in the response.`,
				},
				{
					Name: "fhir-package",
					Type: "string|symbol",
					Text: `The FHIR package to use when creating FHIR types from responses.
Default: fhir5.`,
				},
				{
					Name: "transaction",
					Type: "boolean",
					Text: `If true then the Bundle type property is set to "transaction" otherwise the
type property is set to "batch".`,
				},
			},
			Return: "list",
			Text: `__http-batch__ forms a URL from the provided parameters and sends a POST request to
the host and port provided in the _base_ which can either be the _base_ itself if the _base_ is
a string or if _base_ is a property list then the _:url_ in the property list. An HTTP POST request
containing a __Bundle__ resource made up of the _entries_ provided.


The return value should include a Bundle resource or an OperationOutcome in the event of an failure.
The return value from the call will be a list of three members. The first is the HTTP status as a
__fixnum__. The second is the resource retrieved. The last element in the list are the headers.


For additional information about the FHIR HTTP batch refer to https://www.hl7.org/fhir//http.html#transaction.
`,
		}, &Pkg)
}

// HTTPBatch represents the http-batch function.
type HTTPBatch struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPBatch) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 2, 12)
	entryList, ok := args[0].(slip.List)
	if !ok {
		slip.TypePanic(s, depth, "entries", args[0], "list")
	}
	args = args[1:]
	entries := make([]any, len(entryList))
	for i, entry := range entryList {
		entries[i] = simplifyEntry(s, depth, entry)
	}
	bundle := map[string]any{
		"type":  "batch",
		"entry": entries,
	}
	if v, has := slip.GetArgsKeyValue(args[1:], slip.Symbol(":transaction")); has && v != nil {
		bundle["type"] = "transaction"
	}
	bmod := func(req *http.Request) {
		req.Method = http.MethodPost
	}
	_, data, fhirPkg, res, _ := httpRequest(s, args, depth, bmod, bundle)

	resource := makeAnyResource(data, fhirPkg)

	return slip.List{
		slip.Fixnum(res.StatusCode),
		resource,
		respHeaders(res),
	}
}

func simplifyEntry(s *slip.Scope, depth int, entry slip.Object) (simple any) {
	switch te := entry.(type) {
	case slip.String:
		simple = sen.MustParse([]byte(te))
	case *Instance:
		if te.Class().Name() == "Bundle_Entry" {
			simple = te.data
		}
	case *flavors.Instance:
		if te.Class() == bag.Flavor() {
			simple = te.Any
		}
	}
	if simple == nil {
		slip.TypePanic(s, depth, "entry", entry, "Bundle_Entry", "bag", "json/sen string")
	}
	return
}
