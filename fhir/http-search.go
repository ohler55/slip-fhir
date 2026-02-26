// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/cl"
)

func initHTTPSearch() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := HTTPSearch{Function: slip.Function{Name: "http-search", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "http-search",
			Args: []*slip.DocArg{
				{
					Name: "function",
					Type: "function",
					Text: `Function to call with each resource in the bundle returned and subsequent
linked page's bundles.`,
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
					Name: "type",
					Type: "string",
					Text: "The resource type if needed and not alsearchy in the _base_ as _:url_.",
				},
				{
					Name: "id",
					Type: "string",
					Text: "The compartment id. The inclusion of this argument indicates a compartment search.",
				},
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
					Name: "query",
					Type: "property-list",
					Text: `If present, the values in the property are merged and supersede any
_:query_ in the _base_. The property list indicators should be strings that will be used as
values in the request POST body. Multiple values with the same key are allowed.`,
				},
				{
					Name: "timeout",
					Type: "real",
					Text: `The number of seconds to wait before giving and returning a 408,
Request Timeout code in the response.`,
				},
				{
					Name: "limit",
					Type: "fixnum",
					Text: `The limit of the number of resources fetched when paging.`,
				},
				{
					Name: "fhir-package",
					Type: "string|symbol",
					Text: `The FHIR package to use when creating FHIR types from responses.
Default: fhir5.`,
				},
			},
			Return: "nil",
			Text: `__http-search__ forms a URL from the provided parameters and sends a GET request to
the host and port provided in the _base_ which can either be the _base_ itself if the _base_ is
a string or if _base_ is a property list then the _:url_ in the property list. Only the
_application/fhir+json_ format is currently supported. Search parameters are specified in the
_params_ argument of if _:query_ is specified an HTTP POST is made with the _:query_ as the content
encoded as x-www-form-urlencoded.


The return value should include a resource of either a Bundle, nil, or an OperationOutcome.
The return value from the call will be a list of three members. The first is the HTTP status as
a __fixnum__. The second is the resource or nil. The last element in the list are the headers.


For additional information about the FHIR HTTP search refer to https://www.hl7.org/fhir//http.html#search.
`,
		}, &Pkg)
}

// HTTPSearch represents the http-search function.
type HTTPSearch struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPSearch) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 2, 18)
	d2 := depth + 1
	caller := cl.ResolveToCaller(s, args[0], d2)

	var (
		data    any
		fhirPkg string
		res     *http.Response
		timeout time.Duration
	)
	args = args[1:]
	if v, has := slip.GetArgsKeyValue(args[1:], slip.Symbol(":query")); has {
		if query, ok := v.(slip.List); ok {
			qb := encodeParams(nil, s, query, depth)
			smod := func(req *http.Request) {
				req.Method = http.MethodPost
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				suffix := "/_search"
				if !strings.Contains(req.URL.Path, suffix) {
					req.URL.Path = string(append([]byte(req.URL.Path), suffix...))
				}
			}
			_, data, fhirPkg, res, timeout = httpRequest(s, args, depth, smod, bytes.NewReader(qb))
		} else {
			slip.TypePanic(s, depth, ":query", v, "list")
		}
	} else {
		_, data, fhirPkg, res, _ = httpRequest(s, args, depth, nil, nil)
	}

	resType := alt.String(jp.C("resourceType").First(data))

	switch resType {
	case "":
		// No content so an error with no OperationOutcome.
		return nil
	case "Bundle":
		args = args[1:]
		limit := -1
		if v, has := slip.GetArgsKeyValue(args, slip.Symbol(":limit")); has {
			if num, ok := v.(slip.Fixnum); ok {
				limit = int(num)
			} else {
				slip.TypePanic(s, depth, ":limit", v, "fixnum")
			}
		}
		return eachInBundle(s, data, caller, fhirPkg, limit, d2, res, timeout)
	}
	resource := makeAnyResource(data, fhirPkg)

	_ = caller.Call(s, slip.List{resource}, d2)

	return nil
}
