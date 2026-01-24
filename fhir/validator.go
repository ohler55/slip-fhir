// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import "github.com/ohler55/slip"

// Validator is an interface for validating values for an specific type.
type Validator interface {
	slip.Class
	// Validate should return without panicing if the value is acceptable for
	// the instance and panics otherwise.
	Validate(value any)
}
