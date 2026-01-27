// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/slip"
	"github.com/ohler55/slip/pkg/bag"
	"github.com/ohler55/slip/pkg/flavors"
)

var (
	instanceGetMethod = slip.Method{
		Name: ":get",
		Doc: &slip.FuncDoc{
			Name: ":get",
			Args: []*slip.DocArg{
				{
					Name: "path",
					Type: "string|bag-path",
					Text: "Path to the property value to retrieve.",
				},
			},
			Return: "object",
			Text: `__:get__ returns the first value at the provided _path_. The _path_ must be a
JSONPath as either a __string__ or a __bag-path__. The returned value will be a __bag__ if the
value is a complex value otherwise a build-in primitive such as a __string__, __fixnum__, or
other is returned.`,
		},
		Combinations: []*slip.Combination{{From: &blankType, Primary: &instanceGet{}}},
	}
)

func initInstanceGet() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := instanceGet{Function: slip.Function{Name: "instance-get", Args: args}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "instance-get",
			Args: []*slip.DocArg{
				{
					Name: "instance",
					Type: "instance",
					Text: "An instance.",
				},
				{
					Name: "path",
					Type: "string|bag-path",
					Text: "Path to the property value to retrieve.",
				},
			},
			Return: "object",
			Text: `__instance-get__ returns the first value at the provided _path_. The _path_ must be a
JSONPath as either a __string__ or a __bag-path__. The returned value will be a __bag__ if the
value is a complex value otherwise a build-in primitive such as a __string__, __fixnum__, or
other is returned.`,
		}, &Pkg)
}

type instanceGet struct {
	slip.Function
}

// Call the the function with the arguments provgeted.
func (f *instanceGet) Call(s *slip.Scope, args slip.List, depth int) (result slip.Object) {
	slip.CheckArgCount(s, depth, f, args, 2, 2)
	inst, ok := args[0].(*Instance)
	if !ok {
		slip.TypePanic(s, depth, "instance", args[0], "fhir:instance")
	}
	var path bag.Path
	switch ta := args[1].(type) {
	case slip.String:
		path = bag.Path(jp.MustParseString(string(ta)))
	case slip.List:
		path = make(bag.Path, len(ta))
		for i, e := range ta {
			switch te := e.(type) {
			case nil:
				path[i] = jp.Wildcard('*')
			case slip.String:
				path[i] = jp.Child(te)
			case slip.Symbol:
				path[i] = jp.Child(te)
			case slip.Integer:
				path[i] = jp.Nth(te.Int64())
			default:
				slip.TypePanic(s, depth, "path fragment", e, "string", "symbol", "integer", "nil")
			}
		}
	default:
		slip.TypePanic(s, depth, "path", args[0], "string", "list")
	}

	value := jp.Expr(path).First(inst.data)
	switch tv := value.(type) {
	case nil:
		// result remains nil
	case map[string]any, []any:
		bg := bag.Flavor().MakeInstance().(*flavors.Instance)
		bg.Any = tv
		result = bg
	default:
		result = slip.SimpleObject(value)
	}
	return
}
