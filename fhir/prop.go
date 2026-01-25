// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"
	"sort"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
)

// Prop contains information about the properties of a type.
type Prop struct {
	name     string
	docs     string
	typeName string
	ftype    Validator
	enum     []string
	group    []*Prop
	required bool
	array    bool
}

// Simplify the Object into simple go types of nil, bool, int64, float64,
// string, []any, map[string]any, or time.Time.
func (p *Prop) Simplify() any {
	simple := map[string]any{
		"name":        p.name,
		"description": p.docs,
		"type":        p.typeName,
		"required":    p.required,
		"array":       p.array,
		"enum":        p.enum,
	}
	if 0 < len(p.group) {
		group := make([]any, len(p.group))
		for i, g := range p.group {
			group[i] = g.Simplify()
		}
		simple["group"] = group
	}
	return simple
}

// NewProp creates a new Prop from a simple map (JSON).
func NewProp(simple any) *Prop {
	p := Prop{
		name:     alt.String(jp.C("name").First(simple)),
		docs:     alt.String(jp.C("description").First(simple)),
		typeName: alt.String(jp.C("type").First(simple)),
		required: alt.Bool(jp.C("required").First(simple)),
		array:    alt.Bool(jp.C("array").First(simple)),
	}
	for _, e := range jp.C("enum").W().Get(simple) {
		p.enum = append(p.enum, alt.String(e))
	}
	for _, gp := range jp.C("group").W().Get(simple) {
		p.group = append(p.group, NewProp(gp))
	}
	return &p
}

func (p *Prop) init(t *Type) {
	if 0 < len(p.group) {
		for _, gp := range p.group {
			gp.init(t)
		}
		sortProps(p.group)
		return
	}
	pt := Pkg.FindClass(p.typeName)
	if pt == nil {
		panic(fmt.Sprintf("FHIR type %s property %s specifies an undefined parent of %s",
			t.name, p.name, p.typeName))
	}
	p.ftype = pt.(Validator)
}

func sortProps(props []*Prop) {
	sort.Slice(props, func(i, j int) bool {
		ni := props[i].name
		if ni == "resourceType" {
			return true
		}
		if ni[0] == '_' {
			ni = ni[1:]
		}
		nj := props[j].name
		if nj == "resourceType" {
			return false
		}
		if nj[0] == '_' {
			nj = nj[1:]
		}
		if ni == nj { // one is an _
			return props[j].name < props[i].name
		}
		return ni < nj
	})
}
