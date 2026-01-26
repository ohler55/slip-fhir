// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/ohler55/slip"
)

const (
	// TypeSymbol is the symbol with a value of "fhir-type".
	TypeSymbol = slip.Symbol("fhir-type")
)

// Type is the meta class for FHIR types.
type Type struct {
	name        string
	description string
	pkg         *slip.Package
	parent      string
	inherit     slip.Class // direct super
	supers      []slip.Class
	valid       func(v any) bool // called on bag (simple) elements
	// primitive types may have a pattern and regexp
	pattern string
	rx      *regexp.Regexp
	// complex types have properties
	props []*Prop
}

// String representation of the Object.
func (t *Type) String() string {
	return string(t.Append([]byte{}))
}

// Append a buffer with a representation of the Object.
func (t *Type) Append(b []byte) []byte {
	b = append(b, "#<fhir:Type "...)
	b = append(b, t.name...)
	return append(b, '>')
}

// Name of the class.
func (t *Type) Name() string {
	return t.name
}

// Pkg returns the package the flavor was defined in.
func (t *Type) Pkg() *slip.Package {
	return t.pkg
}

// Documentation of the class.
func (t *Type) Documentation() string {
	return t.description
}

// SetDocumentation of the class.
func (t *Type) SetDocumentation(doc string) {
	t.description = doc
}

// LoadForm returns a list that can be evaluated to create the class or nil if
// the class is a built in class.
func (t *Type) LoadForm() slip.Object {
	return nil
}

// Validate return without panicing if the value is acceptable for the
// instance and panics otherwise.
func (t *Type) Validate(value any) {
	if !t.valid(value) {
		panic(fmt.Sprintf("%s, a %T is not a valid value for a %s.", value, value, t))
	}
}

// Simplify by returning the string representation of the class.
func (t *Type) Simplify() any {
	simple := map[string]any{
		"name":        t.name,
		"description": t.description,
		"parent":      t.parent,
		"package":     t.pkg.Name,
	}
	if t.inherit != nil {
		simple["inherit"] = t.inherit.Name()
	}
	if 0 < len(t.pattern) {
		simple["pattern"] = t.pattern
	}
	if 0 < len(t.props) {
		props := make([]any, len(t.props))
		for i, p := range t.props {
			props[i] = p.Simplify()
		}
		simple["properties"] = props
	}
	return simple
}

// Equal returns true if this Object and the other are equal in value.
func (t *Type) Equal(other slip.Object) (eq bool) {
	return t == other
}

// Hierarchy returns the class hierarchy as symbols for the instance.
func (t *Type) Hierarchy() []slip.Symbol {
	names := make([]slip.Symbol, len(t.supers)+2)
	names[0] = slip.Symbol("fhir:" + t.name)
	for i, sc := range t.supers {
		names[i+1] = slip.Symbol(fmt.Sprintf("%s:%s", sc.Pkg().Name, sc.Name()))
	}
	names[len(names)-1] = slip.TrueSymbol

	return names
}

// Inherits returns true if this Class inherits from a specified Class.
func (t *Type) Inherits(sc slip.Class) bool {
	for _, s := range t.supers {
		if sc == s {
			return true
		}
	}
	return false
}

// InheritsList returns a list of all inherited classes.
func (t *Type) InheritsList() (supers []slip.Class) {
	if t.inherit != nil {
		supers = append(supers, t.inherit)
		supers = append(supers, t.inherit.InheritsList()...)
	}
	return
}

// Metaclass returns the symbol built-in-class.
func (t *Type) Metaclass() slip.Symbol {
	return TypeSymbol
}

// Eval returns self.
func (t *Type) Eval(s *slip.Scope, depth int) slip.Object {
	return t
}

// VarNames for DefMethod, requiredVars and defaultVars combined.
func (t *Type) VarNames() (names []string) {
	if t.inherit != nil {
		names = t.inherit.VarNames()
	}
	for _, p := range t.props {
		names = append(names, p.name)
	}
	return names
}

// Describe the class in detail.
func (t *Type) Describe(b []byte, indent, right int, ansi bool) []byte {
	return t.describe(b, indent, right, ansi, false, "")
}

func (t *Type) describe(b []byte, indent, right int, ansi, full bool, bg string) []byte {
	b = append(b, indentSpaces[:indent]...)
	if ansi {
		b = append(b, bold...)
		b = append(b, "fhir:"...)
		b = append(b, t.name...)
		b = append(b, colorOff...)
	} else {
		b = append(b, "fhir:"...)
		b = append(b, t.name...)
	}
	b = append(b, " is a FHIR "...)
	switch {
	case 'a' <= t.name[0] && t.name[0] <= 'z':
		b = append(b, "PrimitiveType"...)
	case strings.ContainsRune(t.name, '_'):
		b = append(b, "BackboneType"...)
	case t.inherit != nil && t.inherit.Name() == "Element":
		b = append(b, "DataType"...)
	default:
		b = append(b, "Resource"...)
	}
	b = append(b, ":\n"...)
	i2 := indent + 2
	i3 := indent + 4
	if 0 < len(t.description) {
		b = append(b, indentSpaces[:i2]...)
		b = append(b, "Documentation:\n"...)
		b = slip.AppendDoc(b, t.description, i3, right, ansi)
		b = append(b, '\n')
	}
	b = append(b, indentSpaces[:i2]...)
	if t.inherit != nil {
		b = append(b, "Direct Ancestor:"...)
		b = append(b, ' ')
		b = append(b, t.inherit.Name()...)
		b = append(b, '\n')
	}

	b = append(b, indentSpaces[:i2]...)
	b = append(b, "Class precedence list: fhir:"...)
	b = append(b, t.name...)
	for _, sc := range t.supers {
		b = append(b, ' ')
		b = append(b, sc.Pkg().Name...)
		b = append(b, ':')
		b = append(b, sc.Name()...)
	}
	b = append(b, " t\n"...)
	if 0 < len(t.pattern) {
		b = append(b, indentSpaces[:i2]...)
		b = append(b, "Pattern: "...)
		b = append(b, t.pattern...)
		b = append(b, '\n')
	}
	if 0 < len(t.props) {
		b = t.describeProps(b, indent, right, ansi, full, bg)
	}
	return b
}

func (t *Type) describeProps(b []byte, indent, right int, ansi, full bool, bg string) []byte {
	i2 := indent + 2
	pspace := indentSpaces[:indent+4]
	var (
		nameWidth int
		typeWidth int
		props     []*Prop
	)
	for _, p := range t.props {
		if full || p.name[0] != '_' {
			props = append(props, p)
		}
	}
	if full {
		it, _ := t.inherit.(*Type)
		for it != nil {
			for _, ip := range it.props {
				if ip.name != "resourceType" {
					props = append(props, ip)
				}
			}
			it, _ = it.inherit.(*Type)
		}
	}
	sortProps(props)
	for _, p := range props {
		w := len(p.name)
		if nameWidth < w {
			nameWidth = w
		}
		w = len(p.typeName)
		if typeWidth < w {
			typeWidth = w
		}
	}
	docEdge := indent + nameWidth + typeWidth + 14
	b = append(b, indentSpaces[:i2]...)
	b = append(b, "Properties:\n"...)
	if ansi {
		b = append(b, bold...)
	}
	b = fmt.Appendf(b, "%s%-*s  Card. %-*s  Description\n", pspace, nameWidth, "Name", typeWidth, "Type")
	if ansi {
		b = append(b, colorOff...)
	}
	for i, p := range props {
		if i%2 == 0 && ansi {
			b = append(b, bg...)
		}
		b = fmt.Appendf(b, "%s%-*s  ", pspace, nameWidth, p.name)
		if p.required {
			b = append(b, '1')
		} else {
			b = append(b, '0')
		}
		b = append(b, '.', '.')
		if p.array {
			b = append(b, '*')
		} else {
			b = append(b, '1')
		}
		b = append(b, ' ', ' ')
		b = fmt.Appendf(b, "%-*s  ", typeWidth, p.typeName)
		b = slip.AppendDoc(b, p.docs, docEdge, right, ansi, 0)
		if i%2 == 0 && ansi && 0 < len(bg) {
			b = append(b, colorOff...)
		}
		b = append(b, '\n')

		if 0 < len(p.group) {
			var group []*Prop
			for _, gp := range p.group {
				if full || gp.name[0] != '_' {
					group = append(group, gp)
				}
			}
			left := '┣'
			for i, gp := range group {
				if i == len(group)-1 {
					left = '┗'
				}
				b = fmt.Appendf(b, "%s%c %-*s      %-*s", pspace, left, nameWidth, gp.name, typeWidth, gp.typeName)
				if 0 < len(gp.docs) {
					b = append(b, ' ', ' ')
					b = slip.AppendDoc(b, gp.docs, docEdge, right, ansi, 0)
				}
				b = append(b, '\n')
			}
		}
	}
	return b
}

// MakeInstance creates a new instance but does not call the :init method.
func (t *Type) MakeInstance() slip.Instance {
	if 0 < len(t.pattern) { // primitive type
		panic(slip.ErrorNew(slip.NewScope(), 0, "Can not allocate an instance of %s.", t))
	}
	return &Instance{class: t, data: map[string]any{}}
}

func (t *Type) init() {
	if t.valid != nil {
		return
	}
	if t.name == t.parent {
		panic(fmt.Sprintf("primitive type %s specifies a parent of itself", t.name))
	}
	if 0 < len(t.parent) {
		if t.inherit = Pkg.FindClass(t.parent); t.inherit == nil {
			// try current package
			if t.inherit = slip.FindClass(t.parent); t.inherit == nil {
				panic(fmt.Sprintf("FHIR type %s specifies an undefined parent of %s", t.name, t.parent))
			}
		}
	}
	if t.inherit != nil {
		if it, ok := t.inherit.(*Type); ok && it.valid == nil {
			it.init()
		}
		t.supers = append(t.supers, t.inherit)
		t.supers = append(t.supers, t.inherit.InheritsList()...)
	}

	for _, p := range t.props {
		p.init(t)
	}

	// The valid field is also an indicator of init() having been called.
	switch t.name {
	case "boolean":
		t.valid = func(v any) bool {
			_, ok := v.(bool)
			return ok
		}
	case "base64Binary", "canonical", "code", "id", "markdown", "oid", "string", "uri", "url", "uuid",
		"time", "date":
		t.rx = regexp.MustCompile(t.pattern)
		t.valid = func(v any) bool {
			if s, ok := v.(string); ok && t.rx.MatchString(s) {
				return true
			}
			return false
		}
	case "dateTime", "instant":
		t.rx = regexp.MustCompile(t.pattern)
		t.valid = func(v any) bool {
			return primitiveTime(v, t.rx)
		}
	case "decimal":
		t.valid = func(v any) (ok bool) {
			switch v.(type) {
			case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32:
				ok = true
			}
			return
		}
	case "integer":
		t.valid = func(v any) bool {
			if i, ok := primitiveInt(v); ok && math.MinInt32 <= i && i <= math.MaxInt32 {
				return true
			}
			return false
		}
	case "integer64":
		t.valid = func(v any) bool {
			if _, ok := primitiveInt(v); ok {
				return true
			}
			return false
		}
	case "positiveInt":
		t.valid = func(v any) bool {
			if i, ok := primitiveInt(v); ok && 0 < i && i <= math.MaxInt32 {
				return true
			}
			return false
		}
	case "unsignedInt":
		t.valid = func(v any) bool {
			if i, ok := primitiveInt(v); ok && 0 <= i && i <= math.MaxInt32 {
				return true
			}
			return false
		}
	case "xhtml":
		t.valid = func(v any) bool {
			_, ok := v.(string)
			return ok
		}
	default:
		// TBD set correct for other types
		t.valid = func(v any) bool {
			return true
		}
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
