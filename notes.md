# slip-fhir notes

# slip-fhir notes

- get specs from https://www.hl7.org/fhir//downloads.html

- plan

 - http-client-functions (https://www.hl7.org/fhir//http.html)
  - http-read (url &key type id version headers params timeout mime-type)
   - test http-read

  - http-each (base function &key...)
  - http-update (url resource &key version condition headers params timeout)
  - http-patch (url patch &key type id condition headers params timeout)
  - http-delete (url &key type id condition headers params timeout)
  - http-create (url resource &key condition headers params timeout)
  - http-search (url query &key type id headers params timeout)
   - handle type, system, and compartment (same as id as far as the client is concerned?)
  - http-page (url &key query-id page-id headers params timeout)
  - http-capabilities (url &key headers params timeout)
  - http-batch (url bundle &key headers params timeout)
  - http-history (url &key type id headers params timeout)
   - instance, type, all
   - params must of a property list (or an assoc?)
  - graphql-query (base &key type id headers timeout post url-query-field)
   - [post is a boolean for a POST vs GET]
   - url-query-field or implied-field?
    - t - Patient/id/$graphql
    - nil - Patient(id:"xxx")
   - non-standard approach of partly url and partly query
   - could also assume patient(id: String!)
  - graphql-mutation (base &key type id headers timeout post query-encoding)
 - http- prefix leaves the option open for other ways of making requests
 - maybe allow url to be base with url, initial-headers, initial-params, default-timeout
  - as property list (same as lambda list)
 - future
  - graphql https://fhir.hl7.org/fhir/graphql.html
   - maybe help building the query response template
    - (id (code (coding system code)) (subject (reference type)))
     - although spec calls for resource(type: Patient) { birthDate }
      - not supported directly, custom
   - no directives like @flatten or @first are supported in building
   - mutations are supported with the assumption that the server will accept a resource as if it were an inpt type
    - create, update, delete
  - message
   - mllp
   - jetstream
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
