# slip-fhir notes

# slip-fhir notes

- plan
 - validation
  - update instance-validate
   - rename to ?? validp
   - on-error docs
  - valid-p (value &key type on-error)

  - complex validation
   - resource
   - datatype
  -
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
