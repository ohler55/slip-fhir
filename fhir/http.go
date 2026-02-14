// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ohler55/slip"
)

var httpURLKeys = []slip.Symbol{
	slip.Symbol(":type"),
	slip.Symbol(":id"),
	slip.Symbol(":version"),
	slip.Symbol(":history"),
	slip.Symbol(":search"),
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
	for _, key := range httpURLKeys {
		if v, has := slip.GetArgsKeyValue(bargs, key); has {
			if ss, _ := v.(slip.String); 0 < len(ss) {
				switch key {
				case slip.Symbol(":type"), slip.Symbol(":id"):
					pb = append(pb, '/')
					pb = append(pb, string(ss)...)
				case slip.Symbol(":version"):
					pb = append(pb, "/_history/"...)
					pb = append(pb, string(ss)...)
				case slip.Symbol(":history"):
					pb = append(pb, "/_history"...)
				case slip.Symbol(":search"):
					pb = append(pb, "/_search"...)
				}
			}
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
	qb := []byte(uu.RawQuery)
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
	uu.RawQuery = string(qb)

	return uu
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
	req.Header.Set("Accept", "application/fhir+json")
	req.Header.Add("Accept", "application/json+fhir")
	req.Header.Add("Accept", "application/json")

	// TBD set Content-Type overriding any other value
}

func respHeaders(res *http.Response) slip.List {
	var header slip.List

	for k, va := range res.Header {
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
