module github.com/ohler55/slip-fhir

go 1.27

require (
	github.com/ohler55/ojg v1.28.5
	github.com/ohler55/slip v1.5.1
	golang.org/x/term v0.45.0
)

require (
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)

replace github.com/ohler55/slip => ../slip
