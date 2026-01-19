// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"
	"math"
	"regexp"
	"time"

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
	rx          *regexp.Regexp
	valid       func(v any) bool // called on bag (simple) elements
	inherit     slip.Class       // direct super
	pkg         *slip.Package
	precedence  []slip.Symbol
}

func (pt *PrimitiveType) init() {
	if pt.valid != nil {
		return
	}
	if pt.name == pt.parent {
		panic(fmt.Sprintf("primitive type %s specifies a parent of itself", pt.name))
	}

	pt.inherit = Pkg.FindClass(pt.parent)
	if pt.inherit == nil {
		// try current package
		pt.inherit = slip.FindClass(pt.parent)
	}
	if pt.inherit == nil {
		panic(fmt.Sprintf("primitive type %s specifies an undefined parent of %s", pt.name, pt.parent))
	}
	if ipt, ok := pt.inherit.(*PrimitiveType); ok && ipt.valid == nil {
		ipt.init()
	}
	pt.precedence = []slip.Symbol{slip.Symbol(pt.name), slip.Symbol(pt.inherit.Name())}
	for _, ic := range pt.inherit.InheritsList() {
		pt.precedence = append(pt.precedence, slip.Symbol(ic.Name()))
	}

	pt.precedence = append(pt.precedence, slip.TrueSymbol)

	// The valid field is also an indicator of init() having been called.
	switch pt.name {
	case "boolean":
		pt.valid = func(v any) bool {
			_, ok := v.(bool)
			return ok
		}
	case "base64Binary", "canonical", "code", "id", "markdown", "oid", "fstring", "uri", "url", "uuid":
		pt.rx = regexp.MustCompile(pt.pattern)
		pt.valid = func(v any) bool {
			if s, ok := v.(string); ok && pt.rx.MatchString(s) {
				return true
			}
			return false
		}
	case "date", "dateTime", "instant", "time":
		pt.rx = regexp.MustCompile(pt.pattern)
		pt.valid = func(v any) bool {
			return primitiveTime(v, pt.rx)
		}
	case "decimal":
		pt.valid = func(v any) bool {
			if _, ok := v.(float64); ok {
				return true
			}
			return false
		}
	case "integer32":
		pt.valid = func(v any) bool {
			if i, ok := primitiveInt(v); ok && math.MinInt32 <= i && i <= math.MaxInt32 {
				return true
			}
			return false
		}
	case "integer64":
		pt.valid = func(v any) bool {
			if _, ok := primitiveInt(v); ok {
				return true
			}
			return false
		}
	case "positiveint":
		pt.valid = func(v any) bool {
			if i, ok := primitiveInt(v); ok && 0 < i && i <= math.MaxInt32 {
				return true
			}
			return false
		}
	case "unsignedint":
		pt.valid = func(v any) bool {
			if i, ok := primitiveInt(v); ok && 0 <= i && i <= math.MaxInt32 {
				return true
			}
			return false
		}
	case "xhtml":
		pt.valid = func(v any) bool {
			_, ok := v.(string)
			return ok
		}
	default:
		pt.valid = func(v any) bool {
			return true
		}
	}
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
		"parent":      pt.parent,
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

// Validate return without panicing if the value is acceptable for the
// instance and panics otherwise.
func (pt *PrimitiveType) Validate(value any) {
	if !pt.valid(value) {
		panic(fmt.Sprintf("%s, a %T is not a valid value for a %s.", value, value, pt))
	}
}

func primitiveInt(v any) (i int64, ok bool) {
	ok = true
	switch tv := v.(type) {
	case int64:
		i = tv
	case int:
		i = int64(tv)
	case int8:
		i = int64(tv)
	case int16:
		i = int64(tv)
	case int32:
		i = int64(tv)
	case uint:
		i = int64(tv)
	case uint8:
		i = int64(tv)
	case uint16:
		i = int64(tv)
	case uint32:
		i = int64(tv)
	case uint64:
		i = int64(tv)
	case float32:
		i = int64(tv)
		ok = float32(i) == tv
	case float64:
		i = int64(tv)
		ok = float64(i) == tv
	default:
		ok = false
	}
	return
}

func primitiveTime(v any, rx *regexp.Regexp) (ok bool) {
	switch tv := v.(type) {
	case time.Time:
		ok = true
	case string:
		ok = rx.MatchString(tv)
	}
	return
}
