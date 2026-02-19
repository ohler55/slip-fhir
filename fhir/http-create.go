// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"net/http"
	"strings"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
)

func initHTTPCreate() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := HTTPCreate{Function: slip.Function{Name: "http-create", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "http-create",
			Args: []*slip.DocArg{
				{
					Name: "resource",
					Type: "Resource",
					Text: `The resource to include as the body of the request which will be used to
create the new resource. The resource must be a valid resource. The resource can be an instance of
FHIR type, a __bag__ with a valid structure for the type to create, or a JSON or SEN string that
can be parsed to a FHIR instance.`,
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
			},
			Return: "list",
			Text: `__http-create__ forms a URL from the provided parameters and sents a POST request to
the host and port provided in the _base_ which can either be the _base_ itself if the _base_ is
a string or if _base_ is a property list then the _:url_ in the property list. Only the
_application/fhir+json_ format is currently supported.


The return value should include a resource of the created resource or an OperationOutcome. The return
value from the call will be a list of three members. The first is the HTTP status as a __fixnum__.
The second is the resource created or an OperationOutcome. The last element in the list are the headers.


For additional information about the FHIR HTTP create refer to https://www.hl7.org/fhir//http.html#create.
`,
		}, &Pkg)
}

// HTTPCreate represents the http-create function.
type HTTPCreate struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPCreate) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 2, 10)

	var body any
	switch ta := args[0].(type) {
	case *Instance:
		body = ta.data
	case *flavors.Instance:
		if bag.Flavor() == ta.Type {
			body = ta.Any
		} else {
			slip.TypePanic(s, depth, "resource", ta, "bag", "instance", "string")
		}
	case slip.String:
		body = sen.MustParse([]byte(ta))
	default:
		slip.TypePanic(s, depth, "resource", ta, "bag", "instance", "string")
	}
	resType, _ := jp.C("resourceType").First(body).(string)
	if len(resType) == 0 {
		panic("Resource is missing the resourceType field.")
	}
	rmod := func(req *http.Request) {
		req.Method = http.MethodPost
		suffix := "/" + resType
		if !strings.HasSuffix(req.URL.Path, suffix) {
			req.URL.Path = string(append([]byte(req.URL.Path), suffix...))
		}
	}
	_, data, fhirPkg, res, _ := httpRequest(s, args[1:], depth, rmod, body)

	resource := makeAnyResource(data, fhirPkg)

	return slip.List{
		slip.Fixnum(res.StatusCode),
		resource,
		respHeaders(res),
	}
}
