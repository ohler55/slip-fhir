# Slip-FHIR

A FHIR client plugin for SLIP.

The purpose of a Slip FHIR plugin is to provide a FHIR client that can
be used with Lisp. This package includes said client as well as
classes and functions to supporting constructing FHIR resources to be
sent as part of a request sent to a FHIR server. Functions are also
provided to access elements of a resource received from a FHIR server.

An extensive set of documentation for resources, datatypes, and
primitives is included.

Details of the design and decisions made in the process of designing
this plugin are in the [design document](design.md).

## Getting Started

 1. Install
 2. Run
 3. Explore
 4. Make HTTP requests to a FHIR server

### Install

Slip-FHIR is a plugin for the [Slip](https://github.com/ohler55/slip)
Lisp environment. It can be imported with the Lisp `require` function
or as an alternative the slap-fhir application can be built. The
slap-fhir application is a standalone version of Slip with the
slip-fhir plugins already imported making for a simplier way to get
the environment up and running.

```
> go install slip-fhir/cmd/slap-fhir@latest
```

A third option is to checkout the
[slap](https://github.com/ohler55/slap) repository and build from the
master branch by typing:

```
> make
```

The slap applicaiton in the top level directory ready to be used or
copied your choice of a `bin` directory.

### Run

Just run the slap application.

```
> slap-fhir
```

The Slip REPL will start and be ready for commands.

### Explore

Help functions with examples can be viewed by calling the `http-help` function.

```lisp
▶ (http-help)
```

Or call with a topic.

```lisp
▶ (http-help summary)
```

The ``http-help` topics can be used to view lists of the primitives, datatypes, and resource type.

```lisp
▶ (http-help datatypes)
```

The help displays display the type names. To access the types using
Lisp code the list function `fhir-resources`, `fhir-datatypes`,
`fhir-primitives`, and `fhir-types` can be use.

```lisp
▶ (format t "~{~A~%~}~%" (mapcar (lambda (c) (class-name c)) (fhir-datatypes 'fhir5)))
address
age
annotation
...
triggerdefinition
usagecontext
virtualservicedetail
```

The details of each type are availble in a format similar to the FHIR
web pages for each type. The `describe-type` function will display
then details of a given type. There are a few keyword options for the
function. Taking advantage of the tab completion in the REPL start by
typing `(des` then press tab. That will expand to `(describe`. Next
type `-t` which will expand to `(describe-type`. If the package is the
`fhir5` package type ` 'fhir5:` followed by the start of the type to
describe such as `Acc` for account and press tab again. That will
expand to `(describe-type 'fhir5:Account_`. Delete the `_` and close
with a `)` then press return.

### HTTP Requests

First select a FHIR server to access and get the URL and
authentication information ready. The help function provides a guide
for making HTTP requests. Try `(http-help read-example)` to see a
detailed walk through.
