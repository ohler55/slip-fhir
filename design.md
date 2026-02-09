# SLIP FHIR Design Decisions

The design of the SLIP FHIR plugin is based on then [FHIR release 5
specification](https://www.hl7.org/fhir). Some liberties were taken in
order to find compromised where needed as the FHIR spec has some
creative notions about type and scalars that are not alway consistent
with object systems such as CLOS or Flavors in Lisp nor are they
consistent with most other object systems found in languages such as
Go, Python, Java, Ruby, etc.

## Schema File Driven

Multiple versions of the FHIR spec are supported by using the
fhir.schema.json and valuesets.json specification files. Those files
are shy of some key details but for the most part cover enough so that
with some help from a conversion script to convert and enhance the
schema the classes and code for this plugin can be driven by the
schema files to create resources and datatypes used in the plugin.

### Schema Conversion and Enhancement

The FHIR JSON schema could be read in and customized with Go code but
Lisp is better suited, in my opinion, for processing data such as the
JSON schema file. The JSON support in SLIP are particularly useful in
the processing.

---
**note**

The fhir-single.xsd file was considered but rejected in favor of the
JSON files. The XSD file had some additional information but it did
not always coincide with the FHIR web pages. Of course the JSON
versions were sometime out of sync as well but then CSD seemed to be
worse with additional types such as SampledDataDataType. Finally the
XSD scattered type information across multiple elements. The integer64
type for example can be found as simpleType integer64-primitive while
the documentation for that type is in a complexType named just
integer64. After attempting to use each, JSON and XSD the JSON version
seemed like the most straight forward to use and easier to modify and
extend for customization.

---

The approach taken is to read in the JSON schema files and from those
produce a definitions file more suitable for loading into the
slip-fhir package. The FHIR types, Lisp psuedo classes, are then
dynamically created according to the specific FHIR version.

There is quite a bit missing in the schema file. Some of the missing
data has been pulled from the https://hl7.org/fhir web site and added
maually. In some cases assumptions are made based on descriptions from
the web site to accomplish this.

An example of an assumption made and verified with a sampling from the
web site is the inheritance of Resources. All Resources inherit from
the DomainResource type yet the schema lists the type of Resources
only as "object". The generated model definitions file specifies the
inheritance of Resources as DomainResource.

Primitive types in the schema limit the type of primitives as
"string", "number", and "boolean" while the web site defines a more
nested model. The type framework extracted from the web site is used
to clean up the primitive type specification in the generated model
file. After cleaning up the model the mapping from FHIR PrimitiveTypes
to SLIP primitives is specified.

These are just a few of the schema cleanup performed in the
[convert-fhir-schema.lisp](scripts/convert-fhir-schema.lisp) file.

### Model

The FHIR type framework is somewhat schizophrenic in that it sometimes
used XML as a pattern, sometimes uses a traditional model, sometimes
interfaces, and sometime another unique derivative that allows
attributes to be associated with primitive type such as an
integer. This package attempts to normalize those views into a simple
single parent inheritance model. Since the FHIR type framework only
describes types and fields and not methods there is no need for
anything more exotic.

The FHIR team has done some of the work in normalizing the various
models used by defining a JSON format for each type. The JSON format
does define the inheritance hierarchy but it does address the XML view
by specifying a field with a underscore prefix are extensions to that
field. This also allows primitive fields to have extensions of that
field in the containing type. Not exactly attributes of the primitive
value itself but a creative way to accomplish the same effect to some
degree.

The issue of having interfaces in the FHIR type framework can be
ignored as each type is fully defined on it's own or through simple
inheritance.

The FHIR schema files are missing some type definitions such as
Resource and DomainResource. In other cases the parent or super type
is specified incorrectly or abstractly with terms like "string"
instead of "code" as seen on the web pages. Assuming that the FHIR web
pages are the most accurate source of truth those inaccuracies and
omissions are corrected manually in the code that creates the JSON
definitions file.

The normalized model definitions are used to dynamically generate the
FHIR types as psuedo Lisp classes or types. These Lisp types are used
to validate data and to provide documentation in the REPL. That type
documentation attempts to provide a view similar to the FHIR we pages
but with some additional filtering functionality on inherited and
extension properties. The types also aid in building an instance for
sending to a FHIR server.

### Meta (`fhir::Type`)

There are two metaclasses in the package; `fhir:Type` and
`fhir:Property`. Both are part of the `fhir` package as indicated by
the `fhir:` prefix. All primitive types, datatypes, backbone types,
and resources are represented by class (type) that is an instance of
the `fhir:Type`. Complex types, any non-primitive type has properties
that are represented by instances of the `fhir:Property` metaclass.

#### Type

Similar to CLOS, Flavors, and structs, `fhir:Type` classes or types can
be the target of class related functions such as `class-name`,
`class-metaclass`, `find-class`, or `make-instance`. Some such as
`change-class` are not supported. The intent being to make `fhir:Type`
classes blend in with the rest of the Slip or Lisp environment. As
such, what FHIR refers to as properties, Lisp calls slots or variables
of an instance.

As with the other classes, `fhir:Type` classes inherit properties from
super classes. A resource type such as Patient inherits properties
from DomainResource which inherits properties from
Resource. Inheritance is limited to a single parent similar to structs
as that is consistent with the FHIR type framework.

As with other classes, the `make-instance` function makes a new
`fhir:Instance` and returns the new instance after initializing with
the `:init` method. Since a `fhir:Instance` is like any other instance
of a CLOS or Flavors class the `:init` method documentation can be
viewed by calling:

```lisp
▶ (describe-method 'patient :init)
```

#### Instance

Consistent with other instance types such as those for CLOS, Flavors,
or structs, the `fhir:Instance` has a reference to it's class which is
a `fhir:Type`. Like Flavors Instances, the `fhir:Instance` include a
field for it's data which is name `data` and is an `any`. A lock field
is also included to be used if needed for concurrent access to the
instance with the Slip `synchronize` function.

The `fhir:Instance` data is access using `:set` and `:get` methods or
the `instance-set` and `instance-get` functions. Sets are by validated
and raise an error if the value is not compatible with the FHIR
specifications. The get accessors are more flexible and take a JSON
Path to access direct properties or nested properties. The path is the
same as used by the Slip bag package or the `bag-get` function. There
is backdoor to the instance data as well. The 'fhir:Instance` `:data`
method can be called to wrap the instance's data in a Slip bag. Then
all the bag package functions can be used to access and modify the
data. As long as the top level element is not replaced the changes are
made to the instance's data. The `instance-data` function can be used
instead of the `:data` method to accomplish the same bag wrap
functionality.

As might be expected with the mention of methods for interacting with
a `fhir:Instance`, Flavors like method invocations are supported. To
see which methods are supported use the `describe` function. It will
list the supported operations for the instance's type just as with
other classes.

```lisp
▶ (describe 'fhir:type)
```

Alternatively this will return a list of the methods:

```lisp
▶ (send (make-instance 'patient) :which-operations)
```

Both functions and methods are provided for most operations following
the Slip FLOS (Flavor and CLOS) approach. The user can choose the
approach they prefer.

#### Property

Non primitive or complex `fhir:Type` include properties which are
instances of `fhir:Property`. Each `fhir:Property` instance includes a
name, a type, cardinality flags, a description, and optionally enum
values and groups of other properties if the property is something
live value[x] in the FHIR specification. Information about properties
can be viewed using the `describe` function. Properties are always
contained in a `fhir:Type` and can be found using the `type-property`
function. Together a property can be describe.

```lisp
▶ (describe (type-property 'patient 'gender))
```

Like `fhir:Type` instances `fhir:Property` instance have methods which
can be viewed with the `describe` function.

```lisp
▶ (describe 'fhir:Property)
```

#### Primitive Types

While primitive types are listed in the FHIR schema files the
information about the inheritance hierarchy is missing but is added
manually by the script that generates the definition file this package
loads when required by the `require` function. All primitives are
based off a Slip or Lisp type. For example a `fhir:string` inherits or
is based on `cl:string`. The `fhir:code` is based on `fhir:string`
which, as noted, is based on `fhir:string`.

### Error Handling

In Lisp errors are raised when they occur. Slip follows that approach
but it's version of raising an error is a Go panic. When an error
occurs in this package a panic is raised and can be caught with a
`recover` function.

One of the primary purposes, other than documentation, of the
`fhir:Type` and `fhir:Property` are to aid in data validation. The
appraoch used to report validation violations is to use an on-error
callback. Validation violations include a JSON path to identify where
the violation was encountered, the value at the location or in the
case of a set, the provided value. An error message is also
included. The on-error function provide to validation functions or
methods must take three arguments; a path, value, and message. The
function can then return __t__ to abort and fail validation or __nil__
to continue. Of course the option to panic is also available.

### Shadowing

Some of the FHIR type names collide with Lisp core types. When loading these type do not shadow the built in types but instead a warning message is displayed when loading that those `fhir` package types are shadowned by the common list package. To specify those types the package name must be included. For example to descrine a `fhir:string` this would be the call.

```lisp
▶ (describe 'fhir:string)
```

The shadowed types are:

  `fhir:string`
  `fhir:integer`
  `fhir:ratio`
  `fhir:time`
  `fhir:condition`

## Client

TBD

## Versions

- load different versions as whole package such as fhir4:integer32
