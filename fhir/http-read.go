// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
)

func initHTTPRead() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := HTTPRead{Function: slip.Function{Name: "http-read", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "http-read",
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
					Text: "The resource type if needed and not already in the _base_ as _:url_.",
				},
				{
					Name: "id",
					Type: "string",
					Text: "The resource id if needed and not already in the _base_.",
				},
				{
					Name: "version",
					Type: "string",
					Text: `The resource version if needed and not already in the _base_. If set
then a vread request is sent.`,
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
			Text: `__http-read__ forms a URL from the provided parameters and sends a GET request to
the host and port provided in the _base_ which can either be the _base_ itself if the _base_ is
a string or if _base_ is a property list then the _:url_ in the property list. Only the
_application/fhir+json_ format is currently supported.


The return value should include a resource of either the expected resource, nil, an OperationOutcome,
or if an _ elements_ parameter is specified, a __bag__. The return value from the call will be a
list of three members. The first is the HTTP status as a __fixnum__. The second is the resource
retrieved, nil, __bag__, or OperationOutcome. The last element in the list are the headers.


For additional information about the FHIR HTTP read refer to https://www.hl7.org/fhir//http.html#read.
`,
		}, &Pkg)
}

// HTTPRead represents the http-read function.
type HTTPRead struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPRead) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 15)

	uu, data, fhirPkg, res, _ := httpRequest(s, args, depth, nil, nil)

	var resource slip.Object

	if _, has := uu.Query()["_elements"]; data != nil && (has || res.StatusCode != 200) {
		bg := bag.Flavor().MakeInstance().(*flavors.Instance)
		bg.Any = data
		resource = bg
	} else {
		resource = makeAnyResource(data, fhirPkg)
	}
	return slip.List{
		slip.Fixnum(res.StatusCode),
		resource,
		respHeaders(res),
	}
}
