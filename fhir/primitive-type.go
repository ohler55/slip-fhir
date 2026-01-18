// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"runtime/debug"

	"github.com/ohler55/slip"
)

const (
	// PrimitiveTypeSymbol is the symbol with a value of "primitive-type".
	PrimitiveTypeSymbol = slip.Symbol("primitive-type")
)

// PrimitiveType is the class for FHIR PrimitiveTypes.
type PrimitiveType struct {
	name        string
	description string
	pattern     string
	parent      string
	valid       func(v any) bool
	inherit     slip.Class // direct super
	pkg         *slip.Package
	precedence  []slip.Symbol
}

func (pt *PrimitiveType) init() {
	defer func() {
		if rec := recover(); rec != nil {
			debug.PrintStack()
		}
	}()
	pt.inherit = slip.FindClass(pt.parent)
	if pt.inherit == nil {
		// try fhir package
		pt.inherit = Pkg.FindClass(pt.parent)
	}
	if ipt, ok := pt.inherit.(*PrimitiveType); ok && ipt.valid == nil {
		ipt.init()
	}
	pt.precedence = []slip.Symbol{slip.Symbol(pt.name), slip.Symbol(pt.inherit.Name())}
	for _, ic := range pt.inherit.InheritsList() {
		pt.precedence = append(pt.precedence, slip.Symbol(ic.Name()))
	}

	pt.precedence = append(pt.precedence, slip.TrueSymbol)

	// TBD also an indicator of inited
	pt.valid = func(v any) bool {
		return true
	}
	// TBD switch on type
}

// String representation of the Object.
func (pt *PrimitiveType) String() string {
	return string(pt.Append([]byte{}))
}

// Append a buffer with a representation of the Object.
func (pt *PrimitiveType) Append(b []byte) []byte {
	b = append(b, "#<PrimitiveType "...)
	b = append(b, pt.name...)
	return append(b, '>')
}

// Simplify by returning the string representation of the class.
func (pt *PrimitiveType) Simplify() any {
	simple := map[string]any{
		"name":        pt.name,
		"description": pt.description,
		"pattern":     pt.pattern,
		"package":     pt.pkg.Name,
	}
	if pt.inherit != nil {
		simple["inherit"] = pt.inherit.Name()
	}
	return simple
}

// Equal returns true if this Object and the other are equal in value.
func (pt *PrimitiveType) Equal(other slip.Object) (eq bool) {
	return pt == other
}

// Hierarchy returns the class hierarchy as symbols for the instance.
func (pt *PrimitiveType) Hierarchy() []slip.Symbol {
	return pt.precedence
}

// Inherits returns true if this Class inherits from a specified Class.
func (pt *PrimitiveType) Inherits(sc slip.Class) bool {
	name := slip.Symbol(sc.Name())
	for _, sym := range pt.precedence[1:] {
		if name == sym {
			return true
		}
	}
	return false
}

// InheritsList returns a list of all inherited classes.
func (pt *PrimitiveType) InheritsList() []slip.Class {
	return append([]slip.Class{pt.inherit}, pt.inherit.InheritsList()...)
}

// Metaclass returns the symbol built-in-class.
func (pt *PrimitiveType) Metaclass() slip.Symbol {
	return PrimitiveTypeSymbol
}

// Eval returns self.
func (pt *PrimitiveType) Eval(s *slip.Scope, depth int) slip.Object {
	return pt
}

// Name of the class.
func (pt *PrimitiveType) Name() string {
	return pt.name
}

// Pkg returns the package the class was defined in.
func (pt *PrimitiveType) Pkg() *slip.Package {
	return pt.pkg
}

// Documentation of the class.
func (pt *PrimitiveType) Documentation() string {
	return pt.description
}

// SetDocumentation of the class.
func (pt *PrimitiveType) SetDocumentation(doc string) {
	pt.description = doc
}

// VarNames for DefMethod, requiredVars and defaultVars combined.
func (pt *PrimitiveType) VarNames() []string {
	return nil
}

// Describe the class in detail.
func (pt *PrimitiveType) Describe(b []byte, indent, right int, ansi bool) []byte {
	b = append(b, indentSpaces[:indent]...)
	if ansi {
		b = append(b, bold...)
		b = append(b, pt.name...)
		b = append(b, colorOff...)
	} else {
		b = append(b, pt.name...)
	}
	b = append(b, " is a FHIR PrimitiveType:\n"...)
	i2 := indent + 2
	i3 := indent + 4
	if 0 < len(pt.description) {
		b = append(b, indentSpaces[:i2]...)
		b = append(b, "Documentation:\n"...)
		b = slip.AppendDoc(b, pt.description, i3, right, ansi)
		b = append(b, '\n')
	}
	if 0 < len(pt.pattern) {
		b = append(b, indentSpaces[:i2]...)
		b = append(b, "Pattern: "...)
		b = append(b, pt.pattern...)
		b = append(b, '\n')
	}
	b = append(b, indentSpaces[:i2]...)
	if pt.inherit != nil {
		b = append(b, "Direct Ancestor:"...)
		b = append(b, ' ')
		b = append(b, pt.inherit.Name()...)
		b = append(b, '\n')
	}

	b = append(b, indentSpaces[:i2]...)
	b = append(b, "Class precedence list:"...)
	for _, sym := range pt.precedence {
		b = append(b, ' ')
		b = append(b, sym...)
	}
	b = append(b, '\n')

	return b
}

// MakeInstance creates a new instance but does not call the :init method.
func (pt *PrimitiveType) MakeInstance() slip.Instance {
	panic(slip.ErrorNew(slip.NewScope(), 0, "Can not allocate an instance of %s.", pt))
}

// LoadForm returns a list that can be evaluated to create the class or nil if
// the class is a built in class.
func (pt *PrimitiveType) LoadForm() slip.Object {
	return nil
}
