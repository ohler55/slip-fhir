// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
)

// SearchParam contains information about resource search parameters.
type SearchParam struct {
	name     string
	docs     string
	typeName string
	expr     string
}

// NewSearchParam creates a new SearchParam from a simple map (JSON).
func NewSearchParam(simple any) *SearchParam {
	return &SearchParam{
		name:     alt.String(jp.C("name").First(simple)),
		docs:     alt.String(jp.C("description").First(simple)),
		typeName: alt.String(jp.C("type").First(simple)),
		expr:     alt.String(jp.C("expr").First(simple)),
	}
}
