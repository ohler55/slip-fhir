// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/pretty"
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
					Type: "property-list",
					Text: `If present, the values in the property are merged and supersede any
_:headers_ in the _base_. The property list indicators should be strings that will be used as
the key in a HTTP header. The headers FHIR servers should handle are describe at
https://www.hl7.org/fhir//http.html#Http-Headers.`,
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
					Name: "xml",
					Type: "boolean",
					Text: `If true the mime type of the request will be set to "application/fhir+xml"
otherwise the mime type is "application/fhir+json".`,
				},
			},
			Return: "property-list",
			Text: `__http-read__ forms a URL from the provided parameters and sents a GET request to
the host and port provided in the _base_ which can either be the _base_ itself if the _base_ is
a string or if _base_ is a property list then the _:url_ in the property list.


The return value should include a resource of either the expected resource, nil, an OperationOutcome,
or if an _ elements_ parameter is specified, a __bag__. The return value from the call will be a
property list with three indicators:
  __:resource__ retrieved, nil, __bag__, or OperationOutcome.
  __:status__ of the response such as 200.
  __:headers__ in the response.


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

	var base slip.List
	switch ta := args[0].(type) {
	case slip.String:
		base = slip.List{slip.Symbol(":url"), ta}
	case slip.List:
		base = ta
	default:
		slip.TypePanic(s, depth, "base", ta, "string", "property-list")
	}
	args = args[1:]

	uu := httpKeysParser(base, args, []slip.Symbol{slip.Symbol(":type")}) // more keys

	res, err := http.Get(uu.String())
	if err != nil {
		panic(err)
	}
	var body []byte
	if body, err = io.ReadAll(res.Body); err != nil {
		panic(err)
	}
	data := oj.MustParse(body)
	fmt.Printf("*** %s\n", pretty.SEN(data))
	// resType := alt.String(jp.C("resourceType").First(data))
	// TBD find class else bag, if _elements was present in uu params then bag

	resource := bag.Flavor().MakeInstance().(*flavors.Instance)
	resource.Any = data

	return slip.List{
		slip.Symbol(":resource"), resource,
		slip.Symbol(":status"), slip.Fixnum(res.StatusCode),
		slip.Symbol(":headers"), respHeaders(res),
	}
}

func httpKeysParser(base slip.List, args slip.List, keys []slip.Symbol) *url.URL {
	uv, _ := slip.GetArgsKeyValue(base, slip.Symbol(":url"))
	uu, err := url.Parse(slip.MustBeString(uv, ":url"))
	if err != nil {
		panic(err)
	}
	// TBD update uu according to base keys and args

	return uu
}

func respHeaders(res *http.Response) slip.List {
	// TBD
	return nil
}
