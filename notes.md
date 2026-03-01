# slip-fhir notes

# slip-fhir notes

- get specs from https://www.hl7.org/fhir//downloads.html

- plan

 - https://fire.ly/wp-content/uploads/2023/11/FHIR-R5_Nov2023.pdf

 - http-client-functions (https://www.hl7.org/fhir//http.html)
  + http-read
  + http-each
  + http-capabilities
  + http-create
  + http-update
  + http-delete
  + http-search
  + :compartment
  + http-operation
  + http-patch (https://www.hl7.org/fhir//fhirpatch.html and https://www.hl7.org/fhir//http.html#patch)

  - http-batch
   - POST
    - list of
     - method (or read, update, etc)
     - keys and params to form url
     - headers If-None-Match, If-Match, If-Noe-Exist, or If-Modified-Since
     - resource
    - need operation also
    - need search also
    - patch?? not listed but maybe



 - future
  - graphql https://fhir.hl7.org/fhir/graphql.html
   - maybe help building the query response template
    - (id (code (coding system code)) (subject (reference type)))
     - although spec calls for resource(type: Patient) { birthDate }
      - not supported directly, custom
   - no directives like @flatten or @first are supported in building
   - mutations are supported with the assumption that the server will accept a resource as if it were an inpt type
    - create, update, delete
   - graphql-query (base &key type id headers timeout post url-query-field)
    - [post is a boolean for a POST vs GET]
    - url-query-field or implied-field?
     - t - Patient/id/$graphql
     - nil - Patient(id:"xxx")
    - non-standard approach of partly url and partly query
    - could also assume patient(id: String!)
   - graphql-mutation (base &key type id headers timeout post query-encoding)

  - message
   - mllp mllp-read, etc
   - jetstream jet-read, etc
  - subscriptions
   - just jetstream for now
   - related resources
    - Subscription
    - SubscriptionTopic
    - SubscriptionStatus

- sample fhir servers
 - https://server.fire.ly (best) https://server.fire.ly/r5
 - http://hapi.fhir.org/baseR4 (down often and only r4)
 - http://test.fhir.org/r4
 - info https://confluence.hl7.org/spaces/FHIR/pages/35718859/Public+Test+Servers

- xml schema (fhir-single.xsd)
 - no better, missing enums as well

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
