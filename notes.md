# slip-fhir notes

# slip-fhir notes

- get specs from https://www.hl7.org/fhir//downloads.html

- plan
 - unit tests
 - use xml schema instead
 - client
  - read, etc
 - design.md update

- load-fhir or build-fhir or ???
 - load a defs file
  - use for testing
  - option for new package

 - does there need to be a property-class?
  - how to get describe-method to work without a class?
- add vanilla methods to prop
 - propIDMethod
 - propTypeMethod
 - propClassMethod
 - propDescribeMethod
 - propPrintSelfMethod
 - propWhichOperationsMethod
 - propOperationHandledPMethod
 - propEqualMethod


- enum
 - manually link when needed


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
