// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import "github.com/ohler55/slip"

// Type is an interface for fhir classes or in FHIR terminology, type.
type Type interface {
	slip.Class
	Validator
}
