module github.com/ohler55/slip-fhir

go 1.25

require github.com/ohler55/slip v1.4.0

require github.com/ohler55/ojg v1.27.0

require (
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/term v0.34.0 // indirect
	golang.org/x/text v0.28.0 // indirect
)

replace github.com/ohler55/slip => ../slip
