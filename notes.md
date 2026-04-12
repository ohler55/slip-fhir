# slip-fhir notes

- specs from https://www.hl7.org/fhir//downloads.html


- article - Playing with FHIR

- plan

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
  - https://fire.ly/wp-content/uploads/2023/11/FHIR-R5_Nov2023.pdf
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
