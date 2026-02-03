# slip-fhir notes

# slip-fhir notes

- get specs from https://www.hl7.org/fhir//downloads.html

- plan
 - unit tests
  - property validate
 - PropertyClass
 - does there need to be a property-class?
  - how to get describe-method to work without a class?
 - add vanilla methods to prop
  - propIDMethod
  - propClassMethod
  - propDescribeMethod
  - propPrintSelfMethod
  - propWhichOperationsMethod
  - propOperationHandledPMethod
  - propEqualMethod

 - use xml schema instead
 - client
  - read, etc
 - design.md update

- load-fhir or build-fhir or ???
 - load a defs file
  - use for testing
  - option for new package



- enum
 - manually link when needed

- inspect - interactive
 - list top level slots
 - move cursor (hi-lighted) up and down (arrow and ^p ^n)
  - on slot, x for expand, -> or ^f to open that slot and replace display
  - <- or ^b to go back up
  - esc to exit inspector
 - need some kind of general dialog handler
  - given list of text, x command, right, left commands as well
  - or maybe require a tree and hardcode the navigation except for leaves
- edit
 - same as inspect but all modifications
  - primitive, replace or edit with normal repl commands
  - complex, add member from pick-list
  - for a list, add {}


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
