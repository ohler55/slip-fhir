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
fhir.schema.json specification files. Those files are shy of some key
details but for the most part cover enough so that with some help from
a conversion script to convert and enhance the schema the classes and
code for this plugin can be driven by a schema file to create
resources and datatypes used in the plugin.

### Schema Conversion and Enhancement

The FHIR JSON schema could be read in an customized with Go code but
Lisp is better suited, in my opinion, for processing data such as the
JSON schema file. The JSON support in SLIP are particularly useful in
the processing.

The approach taken is to read in the JSON schema file and then produce
a description file more suitable for loading into the slip-fhir
package to dynamically create all the FHIR types according to a
specific FHIR version.

There quite a bit missing in the schema file. Some of the missing data
has been pulled from the https://hl7.org/fhir web site. In some cases
assumptions are made based on descriptions from the web site.

An example of an assumption made and verified with a sampling from the
web site is the inheritance of Resources. All Resources inherit from
the DomainResource type yet the schema lists the type of Resources
only as "object". The generated model description file specifies the
inheritance of Resources as DomainResource.

Primitive types in the schema limit the type of primitives as
"string", "number", and "boolean" while the web site defines a more
nested model. That model is used to clean up the primitive type
specification in the generated model file. After cleaning up the model
the mapping from FHIR PrimitiveTypes to SLIP primitives is specified.

These are just a few of the schema cleanup performed in the
[convert-fhir-schema.lisp](scripts/convert-fhir-schema.lisp) file.

### Model

- Inheritance weirdness
- mix of types ad classes/types and interfaces that are not really used
- scalar can have id and extensions
 - most likely due to the XML background in an attempt to define an object model with XML at it's core
 - use of _foo as recognized work around
- start with some key types (some not defined in the schema file from fhir) then add to based on schema

- primitive, complete, meta
- purpose validation, building and explicit validation

#### Meta (`fhir::Type`)

#### Primitive Types

- largely hand coded to fit into a lisp or slip class hierarchy
- string, time, integer, Condition, and Ratio are shadowed so use fhir:string for example

#### Complex Types

-

## Data

- instance - variation of class, flavor, and struct

- scalars
- complex types
 - too many and use is sparse so creating a class to store data for a resource was rejected
 - instead store as a tree similar the slip bag type
 - classes would provide validation and functions for getting and setting elements
 - instances
  - have a reference to a class
  - have a member for the element or resource data similar to the flavors implementation
  - are a receiver
  - no daemons
- make use of validation when setting or when validate is called
 - panic on error

## Client

## Versions

- load different versions as whole package such as fhir4:integer32
