# slip-fhir notes

# slip-fhir notes

- get specs from https://www.hl7.org/fhir//downloads.html

- plan
 - property class
 - unit tests
 - use xml schema instead
 - client
  - read, etc
 - design.md update

- enum
 - manually link when needed


- type property access in lisp
 - fhir::property - similar to instance but no way to create
  - methods and functions to get data
  + type-properties => list of properties
  + type-property (name) => property
  + property-name
  + property-type
  + property-group
  + property-enum
  + property-cardinality
  - property-validate (path value &optional on-error)

- navigating type definitions
 - options
  - fully expand all types
   - all there for eye navigation
   - likely too large for most screen bot verical and horizontal
    - at least 6 levels deep
  - interactive
   - move cursor up and down, hit return to switch to child, escape to go back up
    - similar to tab completion
    - need slip support (dialog or navigator)
     - provide lines as (text dig-function)
     - an back or up function on the line set

- allow call to create a fhir package from a file, don't add it to user explicitly
