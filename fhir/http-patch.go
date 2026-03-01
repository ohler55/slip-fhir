// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"net/http"
	"strings"

	"github.com/ohler55/ojg/oj"
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
					Type: "list|Parameters",
					Text: `Either an instance of Paramters or a list of property lists that specify
the operations to perform on the resource. If a __Parameters__ instance then it should comply with the
FHIR Path specification at https://www.hl7.org/fhir//fhirpatch.html. If a list then it should be
consistent with RFC 6902 and be a list of property lists. Each property list must have an _:op_ and
a _:path_. Depending on the operation additional properties are needed. Note that the path syntax is
different for RFC 6902 and FHIR Path. RFC 6902 uses __/__ as a separator and FHIR Path uses __.__
separators and does not include indexes. Operations differ as well. For example, RFC 6902 includes
a remove operation while FHIR Path has a delete operation instead. Examples of the operation property
lists using the RFC 6902 syntax are:

  (:op add :path "/name/0/given" :value "Pete")
  (:op replace :path "/name/0/given/0" :value "Pete")
  (:op remove :path "/name/0/given/0")
  (:op test :path "/name/0/given/0" :value "Pete")
  (:op copy :from "/name/0/family" :path "/name/1/family")
  (:op move :from "/name/0/given/1" :path "/name/1/given/0")
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
a string or if _base_ is a property list then the _:url_ in the property list. An HTTP PATCH request is
sent to the target server with either RFC 6902 or FHIR Path content depending on the syntax of the _patch_
argument.


The return value should include a Bundle resource or an OperationOutcome in the event of an failure.
The return value from the call will be a list of three members. The first is the HTTP status as a
__fixnum__. The second is the resource retrieved. The last element in the list are the headers.


For additional information about the FHIR HTTP patch refer to https://www.hl7.org/fhir//http.html#patch,.
https://www.hl7.org/fhir//fhirpatch.html, and https://datatracker.ietf.org/doc/html/rfc6902.
`,
		}, &Pkg)
}

// HTTPPatch represents the http-patch function.
type HTTPPatch struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPPatch) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 2, 14)

	var (
		body any
		mime string
	)
	switch ta := args[0].(type) {
	case slip.List:
		mime = "application/json-patch+json"
		patches := make([]any, len(ta))
		for i, v := range ta {
			pl, ok := v.(slip.List)
			if !ok {
				slip.TypePanic(s, depth, "patch", v, "property list")
			}
			patch := map[string]any{}
			for j := 0; j < len(pl); j += 2 {
				if sym, _ := pl[j].(slip.Symbol); 0 < len(sym) && sym[0] == ':' {
					patch[string(sym[1:])] = slip.Simplify(pl[j+1])
				} else {
					slip.TypePanic(s, depth, "indicator", pl[j], "keyword")
				}
			}
			patches[i] = patch
		}
		body = strings.NewReader(oj.JSON(patches))
	case *Instance:
		mime = "application/fhir+json"
		if ta.Class().Name() != "Parameters" {
			slip.TypePanic(s, depth, "patch", ta, "list of property lists", "Parameters")
		}
		body = ta
	default:
		slip.TypePanic(s, depth, "patch", ta, "list of property lists", "Parameters")
	}
	pmod := func(req *http.Request) {
		req.Method = http.MethodPatch
		req.Header.Set("Content-Type", mime)
	}
	_, data, fhirPkg, res, _ := httpRequest(s, args[1:], depth, pmod, body)

	resource := makeAnyResource(data, fhirPkg)

	return slip.List{
		slip.Fixnum(res.StatusCode),
		resource,
		respHeaders(res),
	}
}
