// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unsafe"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/slip"
)

const (
	// PropertySymbol is the symbol with a value of "fhir-property".
	PropertySymbol = slip.Symbol("property")
)

var propMethods = map[string]*slip.Method{
	propCardinalityMethod.Name:       &propCardinalityMethod,
	propClassMethod.Name:             &propClassMethod,
	propDescribeMethod.Name:          &propDescribeMethod,
	propEnumMethod.Name:              &propEnumMethod,
	propEqualMethod.Name:             &propEqualMethod,
	propGroupMethod.Name:             &propGroupMethod,
	propIDMethod.Name:                &propIDMethod,
	propNameMethod.Name:              &propNameMethod,
	propOperationHandledPMethod.Name: &propOperationHandledPMethod,
	propPrintSelfMethod.Name:         &propPrintSelfMethod,
	propTypeMethod.Name:              &propTypeMethod,
	propValidPMethod.Name:            &propValidPMethod,
	propWhichOperationsMethod.Name:   &propWhichOperationsMethod,
}

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
	pkg      *slip.Package
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

// String representation of the Object.
func (p *Prop) String() string {
	return string(p.Append([]byte{}))
}

// Append a buffer with a representation of the Object.
func (p *Prop) Append(b []byte) []byte {
	b = append(b, "#<fhir:property "...)
	b = append(b, p.name...)
	b = append(b, ' ')
	b = strconv.AppendUint(b, p.ID(), 16)
	return append(b, '>')
}

// ID returns unique ID for the instance.
func (p *Prop) ID() uint64 {
	return uint64(uintptr(unsafe.Pointer(p)))
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

// Equal returns true if this Object and the other are equal in value.
func (p *Prop) Equal(other slip.Object) bool {
	return p == other
}

// Hierarchy returns the class hierarchy as symbols for the instance.
func (p *Prop) Hierarchy() []slip.Symbol {
	return []slip.Symbol{slip.Symbol("property"), slip.TrueSymbol}
}

// Eval returns self.
func (p *Prop) Eval(s *slip.Scope, depth int) slip.Object {
	return p
}

// Name of the class.
func (p *Prop) Name() string {
	return p.name
}

// Pkg returns the package the flavor was defined in.
func (p *Prop) Pkg() *slip.Package {
	return p.pkg
}

// Documentation of the class.
func (p *Prop) Documentation() string {
	return p.docs
}

// SetDocumentation of the class.
func (p *Prop) SetDocumentation(doc string) {
	p.docs = doc
}

// MakeInstance creates a new instance but does not call the :init method.
func (p *Prop) MakeInstance() slip.Instance {
	panic(slip.ErrorNew(slip.NewScope(), 0, "Can not allocate an instance of %s.", p))
}

// Inherits returns true if this Class inherits from a specified Class.
func (p *Prop) Inherits(sc slip.Class) bool {
	return false
}

// InheritsList returns a list of all inherited classes.
func (p *Prop) InheritsList() (supers []slip.Class) {
	return
}

// Metaclass returns the symbol built-in-class.
func (p *Prop) Metaclass() slip.Symbol {
	return PropertySymbol
}

// VarNames for DefMethod, requiredVars and defaultVars combined.
func (p *Prop) VarNames() (names []string) {
	return
}

// MethodNames returns a sorted list of the methods of the instance.
func (p *Prop) MethodNames() slip.List {
	return propMethodNames()
}

// LoadForm returns a list that can be evaluated to create the class or nil if
// the class is a built in class.
func (p *Prop) LoadForm() slip.Object {
	return nil
}

// Receive a method invocation from the send function. Not intended to be
// called by any code other than the send function but is public to allow it
// to be over-ridden.
func (p *Prop) Receive(s *slip.Scope, message string, args slip.List, depth int) (result slip.Object) {
	method := propMethods[strings.ToLower(message)]
	if method == nil {
		slip.InvalidMethodPanic(s, depth,
			p, nil, slip.Symbol(message), "Property does not include the %s method.", message)
	}
	if method.Combinations[0].Primary != nil {
		result = method.Combinations[0].Primary.Call(s, append(slip.List{p}, args...), depth)
	}
	return
}

// GetMethod returns the method if it exists.
func (p *Prop) GetMethod(name string) *slip.Method {
	return propMethods[strings.ToLower(name)]
}

// Methods returns a map of the methods.
func (p *Prop) Methods() map[string]*slip.Method {
	return propMethods
}

// Describe the instance in detail.
func (p *Prop) Describe(b []byte, indent, right int, ansi bool) []byte {
	if strings.EqualFold(p.name, "Property") {
		return p.describeSelf(b, indent, right, ansi)
	}
	b = append(b, indentSpaces[:indent]...)
	if ansi {
		b = append(b, bold...)
		b = p.Append(b)
		b = append(b, colorOff...)
	} else {
		b = p.Append(b)
	}
	b = append(b, ", an instance of "...)
	if ansi {
		b = append(b, bold...)
		b = append(b, "fhir:Property "...)
		b = append(b, colorOff...)
	} else {
		b = append(b, "fhir:Property "...)
	}
	b = append(b, '\n')
	i2 := indent + 2
	i3 := indent + 4
	b = append(b, indentSpaces[:i2]...)
	b = append(b, "Documentation:\n"...)
	b = slip.AppendDoc(b, p.docs, i3, right, ansi)
	b = append(b, '\n')
	b = append(b, indentSpaces[:i2]...)
	b = append(b, "Type: "...)
	b = append(b, p.typeName...)
	b = append(b, '\n')
	b = append(b, indentSpaces[:i2]...)
	b = append(b, "Cardinality: "...)
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
	b = append(b, '\n')
	if 0 < len(p.enum) {
		b = append(b, indentSpaces[:i2]...)
		b = append(b, "Enum:"...)
		for _, e := range p.enum {
			b = append(b, ' ')
			b = append(b, e...)
		}
		b = append(b, '\n')
	}
	if 0 < len(p.group) {
		b = append(b, indentSpaces[:i2]...)
		b = append(b, "Group:\n"...)
		for _, gp := range p.group {
			b = append(b, indentSpaces[:i3]...)
			b = append(b, gp.name...)
			b = append(b, '\n')
		}
	}
	return b
}

func (p *Prop) describeSelf(b []byte, indent, right int, ansi bool) []byte {
	b = append(b, indentSpaces[:indent]...)
	if ansi {
		b = append(b, bold...)
		b = append(b, "fhir:Property"...)
		b = append(b, colorOff...)
	} else {
		b = append(b, "fhir:Property"...)
	}
	b = append(b, " is the FHIR property meta-class\n"...)
	i2 := indent + 2
	i3 := indent + 4
	b = append(b, indentSpaces[:i2]...)
	b = append(b, "Documentation:\n"...)
	b = slip.AppendDoc(b, p.docs, i3, right, ansi)
	b = append(b, '\n')
	b = append(b, indentSpaces[:i2]...)
	b = append(b, "Methods:\n"...)
	for _, name := range propMethodNames() {
		b = append(b, indentSpaces[:i3]...)
		b = append(b, string(name.(slip.Symbol))...)
		b = append(b, '\n')
	}
	return b
}

func (p *Prop) init(t *Type) {
	p.pkg = t.pkg
	if 0 < len(p.group) {
		for _, gp := range p.group {
			gp.init(t)
		}
		sortProps(p.group)
		return
	}
	pt := Pkg.FindClass(p.typeName)
	if pt == nil {
		panic(fmt.Sprintf("FHIR type %s property %s specifies an undefined type of %s",
			t.name, p.name, p.typeName))
	}
	p.ftype = pt.(Validator)
}

func (p *Prop) validateValue(value any, onErr OnErrorFunc) bool {
	if 0 < len(p.group) {
		panic("Can only validate a value with a non-group property.")
	}
	data := map[string]any{p.name: value}

	return p.validate(jp.A(), data, onErr)
}

// data is the map the property is in or may be contained in.
func (p *Prop) validate(path jp.Expr, data map[string]any, onErr OnErrorFunc) bool {
	if 0 < len(p.group) {
		return p.validateGroup(path, data, onErr)
	}
	value := data[p.name]
	ppath := append(path, jp.Child(p.name))
	if value == nil {
		if p.required {
			if onErr(ppath, nil, fmt.Sprintf("%s is required yet missing", ppath)) {
				return true
			}
		}
		return false
	}
	if p.array {
		ft := p.ftype.(*Type)
		if array, ok := value.([]any); ok {
			for i, av := range array {
				if ft.validate(append(ppath, jp.Nth(i)), av, onErr) {
					return true
				}
			}
		} else {
			return onErr(ppath, nil, fmt.Sprintf("%s must be an array", ppath))
		}
	} else {
		if ft, ok := p.ftype.(*Type); ok && ft.validate(ppath, value, onErr) {
			return true
		}
		if 0 < len(p.enum) {
			var found bool
			for _, ev := range p.enum {
				if ev == value {
					found = true
					break
				}
			}
			if !found && onErr(ppath, value,
				fmt.Sprintf("%s is not a valid enum value for %s", pretty.SEN(value), ppath)) {
				return true
			}
		}
	}
	return false
}

func (p *Prop) validateGroup(path jp.Expr, data map[string]any, onErr OnErrorFunc) bool {
	var (
		foundData any
		foundProp *Prop
	)
	ppath := append(path, jp.Child(p.name))
	gpath := append(ppath, jp.Child(""))
	for _, gp := range p.group {
		if gp.name[0] == '_' {
			continue
		}
		gpath[len(gpath)-1] = jp.Child(gp.name)
		if dv, has := data[gp.name]; has {
			if foundProp != nil && onErr(gpath, foundProp,
				fmt.Sprintf("Only one %s property allowed. Both %s and %s present", p.name, foundProp.name, gp.name)) {
				return true
			}
			foundProp = gp
			foundData = dv
		}
	}
	if foundProp != nil {
		gpath[len(gpath)-1] = jp.Child(foundProp.name)
		if ft, ok := foundProp.ftype.(*Type); ok && ft.validate(gpath, foundData, onErr) {
			return true
		}
		xname := "_" + foundProp.name
		if dv, has := data[xname]; has {
			// If no group property found the error will be caught in the
			// check for invalid properties.
			if xprop := p.groupFind(xname); xprop != nil {
				gpath[len(gpath)-1] = jp.Child(xname)
				if ft, ok := xprop.ftype.(*Type); ok && ft.validate(gpath, dv, onErr) {
					return true
				}
			}
		}
	}
	return false
}

func (p *Prop) groupFind(name string) (gprop *Prop) {
	for _, gp := range p.group {
		if gp.name == name {
			gprop = gp
			break
		}
	}
	return
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

func propMethodNames() slip.List {
	names := make([]string, 0, len(propMethods))
	for k := range propMethods {
		names = append(names, k)
	}
	sort.Strings(names)
	methods := make(slip.List, len(names))
	for i, name := range names {
		methods[i] = slip.Symbol(name)
	}
	return methods
}
