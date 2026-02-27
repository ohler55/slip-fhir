// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
)

var (
	httpURLKeys = []slip.Symbol{
		slip.Symbol(":type"),
		slip.Symbol(":id"),
		slip.Symbol(":version"),
		slip.Symbol(":history"),
	}
	compartmentURLKeys = []slip.Symbol{
		slip.Symbol(":id"),
		slip.Symbol(":type"),
	}
)

func httpRequest(
	s *slip.Scope,
	args slip.List,
	depth int,
	reqMod func(req *http.Request),
	body any) (uu *url.URL, data any, fhirPkg string, res *http.Response, timeout time.Duration) {

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

	fhirPkg = "fhir5"
	if v, has := slip.GetArgsKeyValue(base, slip.Symbol(":fhir-package")); has {
		fhirPkg = slip.MustBeString(v, ":fhir-package")
	}
	if v, has := slip.GetArgsKeyValue(args, slip.Symbol(":fhir-package")); has {
		fhirPkg = slip.MustBeString(v, ":fhir-package")
	}

	uu = httpKeysParser(s, depth, base, args)
	ctx := context.Background()
	if timeout = timeoutFromArgs(base, args); 0 < timeout {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	var bodyReader io.Reader
	switch tb := body.(type) {
	case nil:
		// no body
	case io.Reader:
		bodyReader = tb
	default:
		bodyReader = strings.NewReader(oj.JSON(body))
	}
	if req, err := http.NewRequestWithContext(ctx, http.MethodGet, uu.String(), bodyReader); err == nil {
		httpKeysHeader(s, depth, base, args, req)
		if reqMod != nil {
			reqMod(req)
		}
		if res, err = (&http.Client{}).Do(req); err != nil {
			panic(err)
		}
		var body []byte
		if body, err = io.ReadAll(res.Body); err == nil {
			_ = res.Body.Close()
			data = oj.MustParse(body)
		}
	}
	return
}

func httpKeysParser(s *slip.Scope, depth int, base, args slip.List) *url.URL {
	uv, _ := slip.GetArgsKeyValue(base, slip.Symbol(":url"))
	uu, err := url.Parse(slip.MustBeString(uv, ":url"))
	if err != nil {
		panic(err)
	}
	bargs := make(slip.List, len(base)+len(args))
	copy(bargs, args)
	copy(bargs[len(args):], base)
	pb := []byte(uu.Path)
	keys := httpURLKeys
	var (
		isComp  bool
		hasType bool
	)
	if v, has := slip.GetArgsKeyValue(bargs, slip.Symbol(":compartment")); has {
		pb = append(pb, '/')
		pb = append(pb, slip.MustBeString(v, ":compartment")...)
		keys = compartmentURLKeys
		isComp = true
	}
	for _, key := range keys {
		if v, has := slip.GetArgsKeyValue(bargs, key); has {
			switch key {
			case slip.Symbol(":type"):
				pb = append(pb, '/')
				pb = append(pb, slip.MustBeString(v, string(key))...)
				hasType = true
			case slip.Symbol(":id"):
				pb = append(pb, '/')
				pb = append(pb, slip.MustBeString(v, string(key))...)
			case slip.Symbol(":version"):
				pb = append(pb, "/_history/"...)
				pb = append(pb, slip.MustBeString(v, string(key))...)
			case slip.Symbol(":history"):
				if v != nil {
					pb = append(pb, "/_history"...)
				}
			}
		}
	}
	if isComp && !hasType {
		// A search so must end in _search which is handled in http-search. If
		// not a search then a * is needed at the end.
		if _, has := slip.GetArgsKeyValue(bargs, slip.Symbol(":query")); !has {
			pb = append(pb, '/', '*')
		}
	}
	uu.Path = string(pb)

	var params slip.List
	// Merge args and base with the higher precedence at the start.
	if v, has := slip.GetArgsKeyValue(args, slip.Symbol(":params")); has {
		if sv, ok := v.(slip.List); ok {
			params = append(params, sv...)
		}
	}
	if v, has := slip.GetArgsKeyValue(base, slip.Symbol(":params")); has {
		if sv, ok := v.(slip.List); ok {
			params = append(params, sv...)
		}
	}
	qb := encodeParams([]byte(uu.RawQuery), s, params, depth)

	uu.RawQuery = string(qb)

	return uu
}

func encodeParams(qb []byte, s *slip.Scope, params slip.List, depth int) []byte {
	for pos := 0; pos < len(params); pos += 2 {
		ks, ok := params[pos].(slip.String)
		if !ok {
			slip.TypePanic(s, depth, "params key", params[pos], "string")
		}
		if len(params)-1 <= pos {
			panic(fmt.Sprintf("%s missing an argument", ks))
		}
		var vs slip.String
		if vs, ok = params[pos+1].(slip.String); !ok {
			slip.TypePanic(s, depth, "params value", params[pos+1], "string")
		}
		if 0 < len(qb) {
			qb = append(qb, '&')
		}
		qb = append(qb, url.QueryEscape(string(ks))...)
		qb = append(qb, '=')
		qb = append(qb, url.QueryEscape(string(vs))...)
	}
	return qb
}

func httpKeysHeader(s *slip.Scope, depth int, base slip.List, args slip.List, req *http.Request) {
	// Header values are pulled from base first followed by args with values
	// overwritten if they have the same key.
	var headers slip.List
	if v, has := slip.GetArgsKeyValue(base, slip.Symbol(":headers")); has {
		if sv, ok := v.(slip.List); ok {
			headers = append(headers, sv...)
		}
	}
	if v, has := slip.GetArgsKeyValue(args, slip.Symbol(":headers")); has {
		if sv, ok := v.(slip.List); ok {
			headers = append(headers, sv...)
		}
	}
	for _, fv := range headers {
		field, ok := fv.(slip.List)
		if !ok || len(field) < 2 {
			slip.TypePanic(s, depth, "header field", fv, "list")
		}
		key, _ := field[0].(slip.String)
		if len(key) == 0 {
			slip.TypePanic(s, depth, "header field key", field[0], "string")
		}
		for i, v := range field[1:] {
			if ss, _ := v.(slip.String); 0 < len(ss) {
				if 0 < i {
					req.Header.Add(string(key), string(ss))
				} else {
					req.Header.Set(string(key), string(ss))
				}
			} else {
				slip.TypePanic(s, depth, "header field value", v, "string")
			}
		}
	}
	// Overrides
	req.Header.Set("Content-Type", "application/fhir+json")
	// Ideally an Accept field should be included but some servers don't
	// understand that header field and return a warning.
}

func respHeaders(res *http.Response) slip.List {
	var header slip.List

	for k, va := range res.Header {
		// Go incorrectly changes the capitalization on ETag to Etag. This is
		// a known issue that will not befixed since receiving clients are
		// supposed to ignore case. Since not all users will adhere to that
		// specification, Etag is converted to ETag.
		if k == "Etag" {
			k = "ETag"
		}
		value := slip.List{slip.String(k)}
		for _, v := range va {
			value = append(value, slip.String(v))
		}
		header = append(header, value)
	}
	return header
}

func timeoutFromArgs(base, args slip.List) time.Duration {
	var timeout time.Duration

	if v, has := slip.GetArgsKeyValue(base, slip.Symbol(":timeout")); has {
		if r, ok := v.(slip.Real); ok {
			timeout = time.Duration(r.RealValue() * float64(time.Second))
		}
	}
	if v, has := slip.GetArgsKeyValue(args, slip.Symbol(":timeout")); has {
		if r, ok := v.(slip.Real); ok {
			timeout = time.Duration(r.RealValue() * float64(time.Second))
		}
	}
	return timeout
}

func makeAnyResource(data any, fhirPkg string) (resource slip.Object) {
	resType := alt.String(jp.C("resourceType").First(data))
	if class := slip.FindClass(fhirPkg + ":" + resType); class != nil {
		if inst, ok := class.MakeInstance().(*Instance); ok {
			inst.data, _ = data.(map[string]any)
			resource = inst
		}
	} else {
		bg := bag.Flavor().MakeInstance().(*flavors.Instance)
		bg.Any = data
		resource = bg
	}
	return
}

func loadPage(uu string, timeout time.Duration) (data any, res *http.Response) {
	ctx := context.Background()
	if 0 < timeout {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	if req, err := http.NewRequestWithContext(ctx, http.MethodGet, uu, nil); err == nil {
		if res, err = (&http.Client{}).Do(req); err != nil {
			panic(err)
		}
		var body []byte
		if body, err = io.ReadAll(res.Body); err == nil {
			_ = res.Body.Close()
			data = oj.MustParse(body)
		}
	}
	return
}
