# SLIP FHIR Design Decisions

The design of the SLIP FHIR plugin is based on then [FHIR release 5
specification](https://www.hl7.org/fhir). Some liberties were taken in
order to find compromised where needed as the FHIR spec has some
creative notions about type and scalars that are not alway consistent
with object systems such as CLOS or Flavors in LISP nor are they
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

- what's missing
 - no indication of type other than "object"
 - some types such Resource and DomainResource that are used for every resource
- primitive type mapping to lisp/slip types
- had to rename string and time to avoid conflict

### Primitive Types

- largely hand coded to fit into a lisp or slip class hierarchy
- string, time, integer, Condition, and Rotio are shadowed so use fhir:string for example

### Models and Inheritance

- mix of types ad classes/types and interfaces that are not really used
- scalar can have id and extensions
 - most likely due to the XML background in an attempt to define an object model with XML at it's core
 - use of _foo as recognized work around
- start with some key types (some not defined in the schema file from fhir) then add to based on schema

## Data

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
