// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"

	"github.com/ohler55/slip"
)

// BaseSymbol is the symbol with a value of "base".
const BaseSymbol = slip.Symbol("base")

// Base represents the FHIR Base type. It is the foundation of all complex
// types.
type Base struct {
	name    string
	docs    string
	props   []*Prop
	parent  string
	inherit Type
	pkg     *slip.Package
}

func (base *Base) init() {
	if base.inherit, _ = Pkg.FindClass(base.parent).(Type); base.inherit == nil {
		panic(fmt.Sprintf("FHIR type %s specifies an undefined parent of %s", base.name, base.parent))
	}
	for _, p := range base.props {
		p.ftype = Pkg.FindClass(p.typeName).(Type)
	}
	// TBD
}

// String representation of the Object.
func (base *Base) String() string {
	return string(base.Append([]byte{}))
}

// Append a buffer with a representation of the type.
func (base *Base) Append(b []byte) []byte {
	b = append(b, "#<fhir:"...)
	b = append(b, base.name...)
	return append(b, '>')
}

// Simplify the Object into simple go types of nil, bool, int64, float64,
// string, []any, map[string]any, or time.Time.
func (base *Base) Simplify() any {
	props := make([]any, len(base.props))
	for i, p := range base.props {
		props[i] = p.Simplify()
	}
	return map[string]any{
		"name":        base.name,
		"description": base.docs,
		"properties":  props,
	}
}

// Equal returns true if this Object and the other are equal in value.
func (base *Base) Equal(other slip.Object) (eq bool) {
	return base == other
}

// Hierarchy returns the class hierarchy as symbols for the instance.
func (base *Base) Hierarchy() []slip.Symbol {
	return []slip.Symbol{BaseSymbol, slip.TrueSymbol}
}

// Eval returns self.
func (base *Base) Eval(s *slip.Scope, depth int) slip.Object {
	return base
}

// Name of the class.
func (base *Base) Name() string {
	return base.name
}

// Pkg returns the package the flavor was defined in.
func (base *Base) Pkg() *slip.Package {
	return base.pkg
}

// Documentation of the class.
func (base *Base) Documentation() string {
	return base.docs
}

// SetDocumentation of the class.
func (base *Base) SetDocumentation(doc string) {
	base.docs = doc
}

// Describe the flavor in detail.
func (base *Base) Describe(b []byte, indent, right int, ansi bool) []byte {
	b = append(b, indentSpaces[:indent]...)
	if ansi {
		b = append(b, bold...)
		b = append(b, "fhir:"...)
		b = append(b, base.name...)
		b = append(b, colorOff...)
	} else {
		b = append(b, "fhir:"...)
		b = append(b, base.name...)
	}
	// TBD
	return b
}

// MakeInstance creates a new instance but does not call the :init method.
func (base *Base) MakeInstance() slip.Instance {
	if base.name == "Base" || base.name == "Element" || base.name == "Resource" || base.name == "DomainResource" {
		slip.ErrorPanic(slip.NewScope(), 0, "Can not create an instance of fhir %s.", base.name)
	}
	// inst := Instance{Type: obj}
	// inst.Vars = map[string]slip.Object{}
	// inst.SetSynchronized(true)
	// for k, v := range obj.defaultVars {
	// 	inst.Vars[k] = v
	// }
	// inst.Vars["self"] = &inst

	// TBD complete once an instance type is defined, maybe Data or Instance

	return nil
}

// Inherits returns true if this Class inherits from a specified Class.
func (base *Base) Inherits(sc slip.Class) bool {
	// TBD
	return false
}

// InheritsList returns a list of all inherited classes.
func (base *Base) InheritsList() []slip.Class {
	// TBD
	// ca := make([]slip.Class, len(obj.inherit))
	// for i, f := range obj.inherit {
	// 	ca[i] = f
	// }
	return nil
}

// LoadForm returns a list that can be evaluated to create the class or nil if
// the class is a built in class.
func (base *Base) LoadForm() slip.Object {
	// TBD
	return nil
}

// Metaclass returns the symbol flavor.
func (base *Base) Metaclass() slip.Symbol {
	return slip.Symbol("fhir-type")
}

// VarNames for DefMethod, requiredVars and defaultVars combined.
func (base *Base) VarNames() (names []string) {
	if base.inherit != nil {
		names = base.inherit.VarNames()
	}
	for _, p := range base.props {
		names = append(names, p.name)
	}
	return names
}

// Validate should return without panicing if the value is acceptable for
// the instance and panics otherwise.
func (base *Base) Validate(value any) {
	// TBD
}
