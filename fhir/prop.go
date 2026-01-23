// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

// Prop contains information about the properties of a type.
type Prop struct {
	name        string
	description string
	typeName    string
	ftype       Type
	// enum     []string // also items.enum
	// pattern string
	// required bool
	// array bool
}

// Simplify the Object into simple go types of nil, bool, int64, float64,
// string, []any, map[string]any, or time.Time.
func (p *Prop) Simplify() any {
	simple := map[string]any{
		"name":        p.name,
		"description": p.description,
		"type":        p.typeName,
	}
	// TBD add choices
	return simple
}
