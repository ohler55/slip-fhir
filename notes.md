# slip-fhir notes

# slip-fhir notes

- plan
 - some enums are missing
  - read valuesets.json file
   - for lists of enum values with first being name or id or maybe a list of resource and property name
   - try matching based on name, may not work
  - 2 steps, build enum maps then load that in convert
  https://www.hl7.org/fhir//downloads.html

 - script
  - check cardinality for matching
   - Appointment_RecurrenceTemplate occurrence[x]

 - validation
  - complex validation
   - resource
   - datatype
  - type
   - use propMap to determine if all tree elements are ok
   - walk props
    - lookup propert
    - ask property to validate
     - consider cardinality
     - validate with type
     - consider group
     - consider enum

 - property class
 - unit tests
 - client
  - read, etc
 - design.md update

- validate function
 - return indicates continue or stop, let fun determine later if it was a pass or fail
  - if panic then let onErr panic

- type property access in lisp
 - fhir::property - similar to instance but no way to create
  - methods and functions to get data
  -
 - type-properties => returns list of propery objects

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
