// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/slip"
)

// OnErrorFunc is the function type used for validation notification. A return
// of true indicates validation should halt. A return of false indicates
// validation should continue. The path is the normalized path to the property
// while the value is the value that failed validation.
type OnErrorFunc func(path jp.Expr, value any, message string) bool

// Validator is an interface for validating values for an specific type.
type Validator interface {
	slip.Class

	// Validate should return without panicing if the value is acceptable for
	// the instance and panics otherwise.
	Validate(value any, onErr OnErrorFunc) bool
}
