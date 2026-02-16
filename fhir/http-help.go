// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"bytes"
	"io"
	"sort"
	"strings"

	"github.com/ohler55/slip"
)

var helpTop = []string{
	`HTTP access to a FHIR service has a multitude of facets. This help function
summarizes some of the features by topic. Start by calling this function; __http-help__ with a
topic as an argument such as _(http-help summary)_.

`,
	`__functions__ the __fhir__ package includes functions to build FHIR HTTP request and assemble responses.`,
	`__resources__ are one of the fundamentals of FHIR. For a list of the Resource types select this topic.`,
	`__datatypes__ are the elements used to compose Resources.`,
	`__backbones__ are also used to compose Resources but represent elements unique to a specific Resource type.`,
	`__primitives__ represent simple values such as integers, strings, and other single value types.`,
	`__explore__ the types, functions, and instances with the _describe_ and _describe-type_ functions.`,
	`__summary__ for a summary of the HTTP requests supported matching the table at
https://www.hl7.org/fhir//http.html#summary.`,
	`__headers__ describes the various headers that may be used by FHIR servers for both requests and responses.`,
	`__parameters__ describes the parameters accepted by confomant FHIR servers.`,
	`__search__ summarizes the basics of searching and provides a summary of search comparators.`,
	`__history__ provides a description of retrieving history.`,
	`__compartment__ outlines the use of compartments as described in
https://www.hl7.org/fhir//compartmentdefinition.html.`,
	`__read-example__ provides a walk through of using the __http-read__ function.`,
	`__create-example__ provides a walk through of using the __http-create__ function.`,
	`__update-example__ provides a walk through of using the __http-update__ function.`,
	`__delete-example__ provides a walk through of using the __http-delete__ function.`,
	`__patch-example__ provides a walk through of using the __http-patch__ function.`,
	`__batch-example__ provides a walk through of using the __http-batch__ function.`,
	// `__graphql__ describes the FHIR graphql operations and schemas.`,
	// `__jet-help__ similar to this help function but for the FHIR messaging using NATS JetStream.`,
	// `__mllp-help__ similar to this help function but for the FHIR messaging using MLLP protocol.`,
}

var topicHelp = map[string][]string{
	"functions": []string{
		`While the HTTP functions don't match the FHIR interactions exactly as described on
https://www.hl7.org/fhir//http.html they do cover all the interactions. The arguments to each function include
a _base_ which is either a URL or a property list that includes a _:url_ as well as a default lambda list
that are the defaults for the _&key_ arguments. The return values for the functions that do not just return
__nil__ are a list of three elements: then HTTP status in the response, a Resource, and the returned headers
as an association list. A more detailed description is available by calling then __describe__ function like:`,
		` ▶ (describe 'http-read)`,
		`The functions in the __fhir__ package are:

`,
		`__http-read__ covers the FHIR _read_ and _vread_ interactions but can also be used for any of the other
interactions as the result is a list of status, Resource, and headers.`,
		`__http-each__ is intended for use with requests that return a Bundle and potentially include a link for
additional results. It can also be used for single non-Bundle responses. A function is called for each Resource
in the Bundle and subsequent Bundles until a limit is reached.`,
		`__http-capabilities__ sends a request for the capabilities of a FHIR server. The response is expected
to be a CapabilityStatement although any returned Resource is unmarshalled and returned in the standard three
element result list this package uses.`,
		`__http-create__ sends an HTTP POST as described by the _create_ interaction described at
https://www.hl7.org/fhir//http.html#create. Headers and parameters described can be included as arguments to the
function.`,
		`__http-update__ sends an HTTP PUT as described by the _update_ interaction described at
https://www.hl7.org/fhir//http.html#update. Headers and parameters described can be included as arguments to the
function.`,
		`__http-delete__ sends an HTTP DELETE as described by the _delete_ interaction described at
https://www.hl7.org/fhir//http.html#delete. Headers and parameters described can be included as arguments to the
function.`,
		`__http-patch__ sends an HTTP PATCH as described by the _patch_ interaction described at
https://www.hl7.org/fhir//http.html#patch. Headers and parameters described can be included as arguments to the
function. Unlike other functions the body of the request will not be a Resource but instead a
_application/json-patch+json.`,
		`__http-search__ sends an HTTP POST as described by the _search_ interaction described at
https://www.hl7.org/fhir//http.html#search. Headers and parameters described can be included as arguments to the
function. The GET search interaction can use the __http-read__ function. The __http-search__, like the
__http-each__ function, expects a callback function that is called for each matching Resource in the returned
Bundle and linked page Bundles.`,
		`__http-batch__ TBD`,
		`__http-operation__ TBD`,
		`__http-compartment__ TBD`,
	},
	"resources": []string{
		`TBD`,
	},
	"datatypes": []string{`TBD`},
	"backbones": []string{`TBD`},
	"primitives": []string{
		`The details of each primitive type can be viewed using the __describe__ function. The primitive
types that are the classes of the simple values in the FHIR datatypes are:
`,
	},
	"explore":        []string{`TBD`},
	"summary":        []string{`TBD`},
	"headers":        []string{`TBD`},
	"parameters":     []string{`TBD`},
	"search":         []string{`TBD`},
	"history":        []string{`TBD`},
	"compartment":    []string{`TBD`},
	"read-example":   []string{`TBD`},
	"create-example": []string{`TBD`},
	"update-example": []string{`TBD`},
	"delete-example": []string{`TBD`},
	"patch-example":  []string{`TBD`},
	"batch-example":  []string{`TBD`},
	// "graphql": []string{`TBD`},
	// "jet-help": []string{`TBD`},
	// "mllp-help": []string{`TBD`},
}

var topicHelpExtras = map[string]func(b []byte, right int, ansi bool) []byte{
	"resources":  helpResourcesExtra,
	"datatypes":  helpDatatypesExtra,
	"backbones":  helpBackbonesExtra,
	"primitives": helpPrimitivesExtra,
}

func initHTTPHelp() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := HTTPHelp{Function: slip.Function{Name: "http-help", Args: args, SkipEval: []bool{true}}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "http-help",
			Args: []*slip.DocArg{
				{Name: "&optional"},
				{
					Name: "topic",
					Type: "string|symbol",
					Text: `Names a topic to display a description for.`,
				},
			},
			Return: "nil",
			Text:   `__http-help__ displays descriptions for various topics related to FHIR HTTP use.`,
		}, &Pkg)
}

// HTTPHelp represents the http-help function.
type HTTPHelp struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPHelp) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 0, 1)
	ansi := s.Get("*print-ansi*") != nil
	right := int(s.Get("*print-right-margin*").(slip.Fixnum))
	var (
		b     []byte
		extra func(b []byte, right int, ansi bool) []byte
	)
	help := helpTop
	if 0 < len(args) {
		topic := slip.MustBeString(args[0], "topic")
		if h := topicHelp[strings.ToLower(topic)]; 0 < len(help) {
			help = h
		}
		extra, _ = topicHelpExtras[topic]
	}
	b = appendHelpDoc(b, help, right, ansi)
	if extra != nil {
		b = extra(b, right, ansi)
	}
	b = append(b, '\n')

	so := s.Get("*standard-output*")
	w := so.(io.Writer)
	if _, err := w.Write(b); err != nil {
		ss, _ := so.(slip.Stream)
		slip.StreamPanic(s, depth, ss, "write failed: %s", err)
	}
	return nil
}

func appendHelpDoc(b []byte, help []string, right int, ansi bool) []byte {
	for i, h := range help {
		if 0 < i {
			b = append(b, '\n')
		}
		if h[0] == '_' {
			b = slip.AppendDoc(b, h, 4, right, ansi, 2)
		} else {
			b = append(b, '\n')
			b = slip.AppendDoc(b, h, 0, right, ansi)
		}
	}
	return append(b, '\n')
}

func helpResourcesExtra(b []byte, right int, ansi bool) []byte {

	return b
}

func helpDatatypesExtra(b []byte, right int, ansi bool) []byte {

	return b
}

func helpBackbonesExtra(b []byte, right int, ansi bool) []byte {

	return b
}

func helpPrimitivesExtra(b []byte, right int, ansi bool) []byte {
	var words []string

	for _, p := range slip.AllPackages() {
		for _, class := range p.AllClasses() {
			if t, _ := class.(*Type); t != nil {
				name := class.Name()
				if 'a' <= name[0] && name[0] <= 'z' {
					words = append(words, t.pkg.Name+":"+name)
				}
			}
		}
	}
	// b = append(b, bold...)
	// b = appendWords(b, words, right)
	// return append(b, colorOff...)

	return appendWords(b, words, right)
}

func appendWords(b []byte, words []string, right int) []byte {
	sort.Strings(words)
	width := right - 2
	var ww int
	for _, word := range words {
		if ww < len(word) {
			ww = len(word)
		}
	}
	ww += 2
	colCnt := width / ww
	var cnt int
	for _, word := range words {
		if cnt == 0 {
			b = append(b, '\n', ' ', ' ')
		}
		cnt++
		b = append(b, word...)
		b = append(b, bytes.Repeat([]byte{' '}, ww-len(word))...)
		if colCnt <= cnt {
			cnt = 0
		}
	}
	return append(b, '\n')
}
