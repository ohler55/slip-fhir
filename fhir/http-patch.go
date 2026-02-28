// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

func initHTTPPatch() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := HTTPPatch{Function: slip.Function{Name: "http-patch", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "http-patch",
			Args: []*slip.DocArg{
				{
					Name: "patch",
					Type: "list",
					Text: `A list of property lists that specify the operations to perform on the resource.
Each property list must have an _:op_ and a _:path_. Depending on the operation additional properties
are needed. Note that the path syntax is different for RFC 6902 and FHIR Path. RFC 6902 uses __/__ as
separator and can include indexes while FHIR Path uses __.__ separators and does not include indexes.
Operations differ as well. For example, RFC 6902 includes a remove operation while FHIR Path has
a delete operation instead. Examples of the operation property lists using the RFC 6902 syntax are:

 (:op add :path "/name/0/given" :value "Pete")
 (:op replace :path "/name/0/given/0" :value "Pete")
 (:op remove :path "/name/0/given/0")
 (:op test :path "/name/0/given/0" :value "Pete")
 (:op copy :from "/name/0/family" :path "/name/1/family")
 (:op move :from "/name/0/given/1" :path "/name/1/given/0")


The syntax of the _:path_ in the first property list will determine which patch format to use. A leading
__/__ indicates RFC 6902 should be used and any other character indicates FHIR Path should be used.
`,
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
			Text: `__http-patch__ forms a URL from the provided parameters and sends a PATCH request to
the host and port provided in the _base_ which can either be the _base_ itself if the _base_ is
a string or if _base_ is a property list then the _:url_ in the property list.

TBD

The return value should include a resource of either the expected resource, nil, an OperationOutcome,
or if an _ elements_ parameter is specified, a __bag__. The return value from the call will be a
list of three members. The first is the HTTP status as a __fixnum__. The second is the resource
retrieved, nil, __bag__, or OperationOutcome. The last element in the list are the headers.


For additional information about the FHIR HTTP patch refer to https://www.hl7.org/fhir//http.html#patch.
`,
		}, &Pkg)
}

// HTTPPatch represents the http-patch function.
type HTTPPatch struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPPatch) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 1, 15)

	// TBD mod and content

	_, data, fhirPkg, res, _ := httpRequest(s, args, depth, nil, nil)

	resource := makeAnyResource(data, fhirPkg)

	return slip.List{
		slip.Fixnum(res.StatusCode),
		resource,
		respHeaders(res),
	}
}
