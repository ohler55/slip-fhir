// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
)

// Instance is an instance of a FHIR type.
type Instance struct {
	class  *Type
	data   map[string]any
	locker slip.Locker
}

// String representation of the Object.
func (inst *Instance) String() string {
	return string(inst.Append([]byte{}))
}

// Append a buffer with a representation of the Object.
func (inst *Instance) Append(b []byte) []byte {
	b = append(b, "#<"...)
	b = append(b, inst.class.pkg.Name...)
	b = append(b, ':')
	b = append(b, inst.class.name...)
	b = append(b, ' ')
	b = strconv.AppendUint(b, inst.ID(), 16)
	return append(b, '>')
}

// ID returns unique ID for the instance.
func (inst *Instance) ID() uint64 {
	return uint64(uintptr(unsafe.Pointer(inst)))
}

// Simplify by returning the string representation of the flavor.
func (inst *Instance) Simplify() interface{} {
	return inst.data
}

// Hierarchy returns the class hierarchy as symbols for the instance.
func (inst *Instance) Hierarchy() []slip.Symbol {
	return inst.class.Hierarchy()
}

// IsA return true if the instance is of a type that inherits from the
// provided flavor.
func (inst *Instance) IsA(class string) bool {
	fc := fmt.Sprintf("%s:%s", inst.class.pkg.Name, class)
	for _, sym := range inst.class.Hierarchy() {
		if strings.EqualFold(fc, string(sym)) || strings.EqualFold(class, string(sym)) {
			return true
		}
	}
	return false
}

// SlotNames returns a list of the slots names for the instance.
func (inst *Instance) SlotNames() (names []string) {
	return inst.class.VarNames()
}

// SlotValue return the value of an instance variable.
func (inst *Instance) SlotValue(sym slip.Symbol) (value slip.Object, has bool) {
	var v any
	inst.locker.Lock()
	if v, has = jp.C(string(sym)).FirstFound(inst.data); has {
		value = slip.SimpleObject(v)
	}
	inst.locker.Unlock()
	return
}

// SetSlotValue sets the value of an instance variable.
func (inst *Instance) SetSlotValue(sym slip.Symbol, value slip.Object) (has bool) {
	var data any
	switch ta := value.(type) {
	case *flavors.Instance:
		if bag.Flavor() != ta.Type {
			slip.TypePanic(slip.NewScope(), 0, "value", ta,
				"fhir:instance", "bag", "string", "fixnum", "float", "boolean", "nil")
		}
		data = ta.Any
	case *Instance:
		data = ta.data
	default:
		data = slip.Simplify(ta)
	}
	prop := inst.class.propMap[strings.ToLower(string(sym))]
	if prop == nil {
		panic(fmt.Sprintf("%s does nat have a %s property.", inst, sym))
	}
	prop.ftype.Validate(data, func(p jp.Expr, v any, message string) bool {
		panic(fmt.Sprintf("Value at %s, %s: %s.", p, pretty.SEN(v), message))
	})

	inst.locker.Lock()
	_ = jp.C(string(sym)).Set(inst.data, data)
	inst.locker.Unlock()

	return true
}

// Init the instance slots from the provided args list.
func (inst *Instance) Init(scope *slip.Scope, args slip.List, depth int) {
	var (
		data  map[string]any
		onErr slip.Caller
		skip  bool
	)
	if value, has := slip.GetArgsKeyValue(args, slip.Symbol(":data")); has {
		if bg, ok := value.(*flavors.Instance); ok && bag.Flavor() == bg.Type {
			data, _ = bg.Any.(map[string]any)
		}
		if data == nil {
			slip.TypePanic(scope, depth, ":data", args[1], "bag")
		}
	}
	if value, has := slip.GetArgsKeyValue(args, slip.Symbol(":on-error")); has {
		onErr = resolveToOnError(scope, value, depth)
	}
	if value, has := slip.GetArgsKeyValue(args, slip.Symbol(":no-validation")); has && value != nil {
		skip = true
	}
	if data != nil && !skip {
		var onErrFn OnErrorFunc
		if onErr == nil {
			onErrFn = func(p jp.Expr, v any, message string) bool {
				panic(fmt.Sprintf("Value at %s, %s: %s.", p, pretty.SEN(v), message))
			}
		} else {
			onErrFn = func(p jp.Expr, v any, message string) bool {
				cargs := slip.List{bag.Path(p), objectify(v), slip.String(message)}
				return onErr.Call(scope, cargs, depth) == slip.True
			}
		}
		if inst.class.Validate(data, onErrFn) {
			data = nil
		}
	}
	if data == nil {
		data = map[string]any{}
		if inst.class.parent == "DomainResource" {
			data["resourceType"] = inst.class.name
		}
	}
	inst.data = data
}

// HasMethod returns true if the instance handles the named method.
func (inst *Instance) HasMethod(method string) (has bool) {
	return typeMethods[strings.ToLower(method)] != nil
}

// GetMethod returns the method if it exists.
func (inst *Instance) GetMethod(name string) *slip.Method {
	return typeMethods[strings.ToLower(name)]
}

// MethodNames returns a sorted list of the methods of the instance.
func (inst *Instance) MethodNames() slip.List {
	return typeMethodNames()
}

// Receive a method invocation from the send function. Not intended to be
// called by any code other than the send function but is public to allow it
// to be over-ridden.
func (inst *Instance) Receive(s *slip.Scope, message string, args slip.List, depth int) slip.Object {
	method := typeMethods[strings.ToLower(message)]
	if method == nil {
		slip.InvalidMethodPanic(s, depth,
			inst, nil, slip.Symbol(message), "%s does not include the %s method.", inst.class.Name(), message)
	}
	if method.Combinations[0].Primary == nil {
		slip.InvalidMethodPanic(s, depth,
			inst, nil, slip.Symbol(message), "Can not evaluate the %s %s method.", inst.class.Name(), message)
	}
	return method.Combinations[0].Primary.Call(s, append(slip.List{inst}, args...), depth)
}

// Equal returns true if this Object and the other are equal in value.
func (inst *Instance) Equal(other slip.Object) bool {
	if inst == other {
		return true
	}
	if oi, ok := other.(*Instance); ok && inst.class == oi.class {
		return alt.Compare(inst.data, oi.data) == nil
	}
	return false
}

// Eval returns self.
func (inst *Instance) Eval(s *slip.Scope, depth int) slip.Object {
	return inst
}

// Describe the instance in detail.
func (inst *Instance) Describe(b []byte, indent, right int, ansi bool) []byte {
	b = append(b, indentSpaces[:indent]...)
	if ansi {
		b = append(b, bold...)
		b = inst.Append(b)
		b = append(b, colorOff...)
	} else {
		b = inst.Append(b)
	}
	b = append(b, ", an instance of "...)
	if ansi {
		b = append(b, bold...)
		b = append(b, inst.class.pkg.Name...)
		b = append(b, ':')
		b = append(b, inst.class.name...)
		b = append(b, colorOff...)
	} else {
		b = append(b, inst.class.pkg.Name...)
		b = append(b, ':')
		b = append(b, inst.class.name...)
	}
	b = append(b, ",\n  "...)
	data := strings.ReplaceAll(pretty.SEN(inst.data), "\n", "  \n")
	b = append(b, data...)

	return append(b, '\n')
}

// Class returns the flavor of the instance.
func (inst *Instance) Class() slip.Class {
	return inst.class
}

// Dup returns a duplicate of the instance.
func (inst *Instance) Dup() slip.Instance {
	inst.locker.Lock()
	dup := Instance{
		class: inst.class,
		data:  alt.Dup(inst.data).(map[string]any),
	}
	inst.locker.Unlock()
	if _, ok := inst.locker.(*sync.Mutex); ok {
		dup.locker = &sync.Mutex{}
	} else {
		dup.locker = slip.NoOpLocker{}
	}
	return &dup
}

// LoadForm returns a form that can be evaluated to create the object.
func (inst *Instance) LoadForm() slip.Object {
	return slip.InstanceLoadForm(inst)
}

// SetSynchronized set the synchronized mode of the instance.
func (inst *Instance) SetSynchronized(on bool) {
	if on {
		inst.locker = &sync.Mutex{}
	} else {
		inst.locker = slip.NoOpLocker{}
	}
}

// Synchronized returns true if the instance is in synchronized mode.
func (inst *Instance) Synchronized() bool {
	_, ok := inst.locker.(*sync.Mutex)
	return ok
}
