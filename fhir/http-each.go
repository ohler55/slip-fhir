// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"net/http"
	"time"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/cl"
)

func initHTTPEach() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := HTTPEach{Function: slip.Function{Name: "http-each", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "http-each",
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
					Text: "The resource type if needed and not already in the _base_ as _:url_.",
				},
				{
					Name: "id",
					Type: "string",
					Text: "The resource id if needed and not already in the _base_.",
				},
				{
					Name: "history",
					Type: "boolean",
					Text: `If true "_history" is appended to the URL path.`,
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
			Return: "nil|resource|bag",
			Text: `__http-each__ forms a URL from the provided parameters and sents a GET request to
the host and port provided in the _base_ which can either be the _base_ itself if the _base_ is
a string or if _base_ is a property list then the _:url_ in the property list. Only the
_application/fhir+json_ format is currently supported. The __http-each__ function is generally
used for processing multiple returns in a FHIR Bundle and then continuing to load pages until
the _limit_ is reached or until no more resource are pending processing.


The return value will be __nil__ on success. If an error occurrs then either a FHIR OperationOutcome
is returned or a __bag__ with what ever the JSON content of the last response was. This does not
match any specific FHIR HTTP request but can be used for any FHIR GET request.
`,
		}, &Pkg)
}

// HTTPEach represents the http-each function.
type HTTPEach struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPEach) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 2, 18)
	d2 := depth + 1
	caller := cl.ResolveToCaller(s, args[0], d2)
	_, data, fhirPkg, res, timeout := httpRequest(s, args[1:], depth, nil, nil)

	resType := alt.String(jp.C("resourceType").First(data))

	switch resType {
	case "":
		// No content so an error with no OperationOutcome.
		return nil
	case "Bundle":
		args = args[2:]
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

var linkNextPath = jp.MustParseString("link[?@.relation == 'next'].url")

func eachInBundle(
	s *slip.Scope,
	data any,
	caller slip.Caller,
	fhirPkg string,
	limit, d2 int,
	res *http.Response,
	timeout time.Duration) slip.Object {

	result := slip.List{slip.Fixnum(200), nil, respHeaders(res)}

	for data != nil {
		for _, dv := range jp.C("entry").W().C("resource").Get(data) {
			resource := makeAnyResource(dv, fhirPkg)
			_ = caller.Call(s, slip.List{resource}, d2)
			if 0 < limit {
				limit--
				if limit == 0 {
					return result
				}
			}
		}
		if limit != 0 { // covers negative and remaining
			next := alt.String(linkNextPath.First(data))
			data = nil
			if 0 < len(next) {
				data, res = loadPage(next, timeout)
				if res.StatusCode != 200 {
					return makeAnyResource(data, fhirPkg)
				}
			}
		}
	}
	return nil
}
