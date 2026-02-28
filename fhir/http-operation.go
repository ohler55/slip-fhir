// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/ohler55/slip"
)

func initHTTPOperation() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := HTTPOperation{Function: slip.Function{Name: "http-operation", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "http-operation",
			Args: []*slip.DocArg{
				{
					Name: "operation",
					Type: "string",
					Text: `The operation to request.`,
				},
				{
					Name: "args",
					Type: "property-list|Parameters",
					Text: `Arguments to the operation. If __nil__ then a GTE request is sent and arguments
are expected to be in the _params_ if there are any. If the args is a __Paramters__ instance then the request
will be a POST. If the args is a __list__ then it is assumed to be properties encoded using using
application/x-www-form-urlencoded.`,
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
					Text: "The resource type if needed and not aloperationy in the _base_ as _:url_.",
				},
				{
					Name: "id",
					Type: "string",
					Text: "The resource id if needed and not aloperationy in the _base_.",
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
			Text: `__http-operation__ forms a URL from the provided arguments. There are three forms
that can be generated. They are:
   /$[name]
   /[type]/$[name]
   /[type]/[id]/$[name]


The return value from the call will be a list of three members. The first is the HTTP status as
a __fixnum__. The second is the resource retrieved. The last element in the list are the headers.


For additional information about some example FHIR HTTP operation refer to
https://www.hl7.org/fhir//capabilitystatement-operation-subset.html and
https://www.hl7.org/fhir//capabilitystatement-operation-implements.html.

`,
		}, &Pkg)
}

// HTTPOperation represents the http-operation function.
type HTTPOperation struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPOperation) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 3, 15)

	op := slip.MustBeString(args[0], ":operation")

	var (
		omod func(req *http.Request)
		body any
	)
	switch ta := args[1].(type) {
	case nil:
		omod = func(req *http.Request) {
			req.Method = http.MethodGet
			// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.URL.Path = fmt.Sprintf("%s/$%s", req.URL.Path, op)
		}
	case slip.List:
		omod = func(req *http.Request) {
			req.Method = http.MethodPost
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.URL.Path = fmt.Sprintf("%s/$%s", req.URL.Path, op)
			qb := encodeParams(nil, s, ta, depth)
			body = bytes.NewReader(qb)
		}
	case *Instance:
		omod = func(req *http.Request) {
			req.Method = http.MethodPost
			req.Header.Set("Content-Type", "application/fhir+json")
			req.URL.Path = fmt.Sprintf("%s/$%s", req.URL.Path, op)
			// TBD check class, must be Parameters
			body = ta
		}
	default:
		// TBD type error
	}
	_, data, fhirPkg, res, _ := httpRequest(s, args[2:], depth, omod, body)

	resource := makeAnyResource(data, fhirPkg)

	return slip.List{
		slip.Fixnum(res.StatusCode),
		resource,
		respHeaders(res),
	}
}
