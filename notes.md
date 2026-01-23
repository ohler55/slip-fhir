# slip-fhir notes

# slip-fhir notes

- add bag-load (filename) => bag

- update script to pull in hierarchy

- Element
 - add id and extension properties
 - others can skip inherited properties

- allow call to create a fhir package from a file, don't add it to user explicitly

- slap - repo for slip + plugins and embeded lisp code

- convert script
 - fstring parent should be string
 - time
 - xhtml

- schema
 - get list of resources from discriminator.mapping
 - walk definitions
  - ignore ResourceList
  - types with lower case, form primitives map
   - name
   - description
   - type (number, string, ??)
   - pattern (regex)
  - Capitalized
   - in resource list
    - class
   - with _ then a backbone element
    - class with different precedence
   - no _
    - datatype class
  - special cases for


 - simplify spec version with script
 - Coding
  - description
  - elements or properties or props []
   - system
    - description
    - optional bool
    - array bool
    - type (primitive or class name)
    - choices

Base
  Element
    BackboneElement
    DataType
      PrimitiveType
      BackboneType
  Resource
    DomainResource
      CanonicalResource <- interfaces, don't count
        MetadataResource <- interfaces, don't count
      Account
      Patient
      etc

----------
- define classes
 - primitives
  - can not create instances of or maybe need a coerce or maybe just for validation
 - datatypes
  - class with precedence of (Datatype Element Base)
 - hard code Datatype, Element, and Base

- deviate from spec with primitive precedence
 - describe inheritance tree somewhere

- primitives inherit from slip primitives conceptually


-------
- primitives are just go structs that support ValidateFhir(v any) error
 - or should ValidateFhir just panic
 - or just named Validate(v any)
- define classes
 - metaclass fhir-datatype[-class] and fhir-resource[-class]
 - fhir:Base class empty but with some common methods
  - slip.class with instance slip.Object
  - fhir classes use the names, instances are fhir.Instance (a slip.Instance)
 - fhir:Element
 - fhir:DataType
 - fhir:Coding
 - make-datatype
 - make-resource

 - define Base as base for all precedence (t)
 - define DataType (print capitalized but search lowercase)
