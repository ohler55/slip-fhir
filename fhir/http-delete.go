// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"net/http"

	"github.com/ohler55/slip"
)

func initHTTPDelete() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := HTTPDelete{Function: slip.Function{Name: "http-delete", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "http-delete",
			Args: []*slip.DocArg{
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
					Text: "The resource type if needed and not alredy in the _base_ as _:url_.",
				},
				{
					Name: "id",
					Type: "string",
					Text: "The resource id if needed and not alredy in the _base_.",
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
					Name: "fhir-package",
					Type: "string|symbol",
					Text: `The FHIR package to use when creating FHIR types from responses.
Default: fhir5.`,
				},
			},
			Return: "list",
			Text: `__http-delete__ forms a URL from the provided parameters and sends a DELETE request to
the host and port provided in the _base_ which can either be the _base_ itself if the _base_ is
a string or if _base_ is a property list then the _:url_ in the property list. Only the
_application/fhir+json_ format is currently supported. Both a _:type_ and a _:id_ are required. They can
be specified in either the _base_ or as a key value.


The return value should include a resource of either the expected resource, nil, an OperationOutcome,
or if an _ elements_ parameter is specified, a __bag__. The return value from the call will be a
list of three members. The first is the HTTP status as a __fixnum__. The second is the resource
either nil or an OperationOutcome. The last element in the list are the headers.


For additional information about the FHIR HTTP delete refer to https://www.hl7.org/fhir//http.html#delete.
`,
		}, &Pkg)
}

// HTTPDelete represents the http-delete function.
type HTTPDelete struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPDelete) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 13)

	dmod := func(req *http.Request) {
		req.Method = http.MethodDelete
	}
	_, data, fhirPkg, res, _ := httpRequest(s, args, depth, dmod, nil)

	var resource slip.Object
	if data != nil {
		resource = makeAnyResource(data, fhirPkg)
	}
	return slip.List{
		slip.Fixnum(res.StatusCode),
		resource,
		respHeaders(res),
	}
}
