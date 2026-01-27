// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/slip"
)

// Validator is an interface for validating values for an specific type.
type Validator interface {
	slip.Class

	// Validate should return without panicing if the value is acceptable for
	// the instance and panics otherwise.
	Validate(value any) bool

	// Validate the provided data and call the onErr function on a validation
	// error. If all validation rules succeed then true is returned else false is
	// returned. The result of the onErr call should be one of:
	// :continue - indicates validation should continue and the error ignored.
	// :reject - indicates validation should continue but validation should fail after completion.
	// :raise - indicates validation should stop and an error raised.
	// Validate(value any, onErr func(path jp.Expr, value any) slip.Symbol) bool
}
