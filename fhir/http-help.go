// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir

import (
	"bytes"
	"io"
	"sort"
	"strings"

	"github.com/ohler55/slip"
)

var helpTop = []string{
	`HTTP access to a FHIR service has a multitude of facets. This help function
summarizes some of the features by topic. Start by calling this function; __http-help__ with a
topic as an argument such as _(http-help summary)_.

`,
	`__functions__ the __fhir__ package includes functions to build FHIR HTTP request and assemble responses.`,
	`__resources__ are one of the fundamentals of FHIR. For a list of the Resource types select this topic.`,
	`__datatypes__ are the elements used to compose Resources.`,
	`__backbones__ are also used to compose Resources but represent elements unique to a specific Resource type.`,
	`__primitives__ represent simple values such as integers, strings, and other single value types.`,
	`__explore__ the types, functions, and instances with the _describe_ and _describe-type_ functions.`,
	`__summary__ for a summary of the HTTP requests supported matching the table at
https://www.hl7.org/fhir//http.html#summary.`,
	`__headers__ describes the various headers that may be used by FHIR servers for both requests and responses.`,
	`__parameters__ describes the parameters accepted by confomant FHIR servers.`,
	`__search__ summarizes the basics of searching and provides a summary of search comparators.`,
	`__history__ provides a description of retrieving history.`,
	`__compartment__ outlines the use of compartments as described in
https://www.hl7.org/fhir//compartmentdefinition.html.`,
	`__read-example__ provides a walk through of using the __http-read__ function.`,
	`__create-example__ provides a walk through of using the __http-create__ function.`,
	`__update-example__ provides a walk through of using the __http-update__ function.`,
	`__delete-example__ provides a walk through of using the __http-delete__ function.`,
	`__search-example__ provides a walk through of using the __http-search__ function.`,
	`__patch-example__ provides a walk through of using the __http-patch__ function.`,
	`__batch-example__ provides a walk through of using the __http-batch__ function.`,
	// `__graphql__ describes the FHIR graphql operations and schemas.`,
	// `__jet-help__ similar to this help function but for the FHIR messaging using NATS JetStream.`,
	// `__mllp-help__ similar to this help function but for the FHIR messaging using MLLP protocol.`,
}

var topicHelp = map[string][]string{
	"functions": []string{
		`While the HTTP functions don't match the FHIR interactions exactly as described on
https://www.hl7.org/fhir//http.html they do cover all the interactions. The arguments to each function include
a _base_ which is either a URL or a property list that includes a _:url_ as well as a default lambda list
that are the defaults for the _&key_ arguments. The return values for the functions that do not just return
__nil__ are a list of three elements: then HTTP status in the response, a Resource, and the returned headers
as an association list. A more detailed description is available by calling then __describe__ function like:`,
		` ▶ (describe 'http-read)`,
		`The client functions in the __fhir__ package are:

`,
		`__http-read__ covers the FHIR _read_ and _vread_ interactions but can also be used for any of the other
interactions as the result is a list of status, Resource, and headers.`,
		`__http-each__ is intended for use with requests that return a Bundle and potentially include a link for
additional results. It can also be used for single non-Bundle responses. A function is called for each Resource
in the Bundle and subsequent Bundles until a limit is reached.`,
		`__http-capabilities__ sends a request for the capabilities of a FHIR server. The response is expected
to be a CapabilityStatement although any returned Resource is unmarshalled and returned in the standard three
element result list this package uses.`,
		`__http-create__ sends an HTTP POST as described by the _create_ interaction described at
https://www.hl7.org/fhir//http.html#create. Headers and parameters described can be included as arguments to the
function.`,
		`__http-update__ sends an HTTP PUT as described by the _update_ interaction described at
https://www.hl7.org/fhir//http.html#update. Headers and parameters described can be included as arguments to the
function.`,
		`__http-delete__ sends an HTTP DELETE as described by the _delete_ interaction described at
https://www.hl7.org/fhir//http.html#delete. Headers and parameters described can be included as arguments to the
function.`,
		`__http-patch__ sends an HTTP PATCH as described by the _patch_ interaction described at
https://www.hl7.org/fhir//http.html#patch. Headers and parameters described can be included as arguments to the
function. Unlike other functions the body of the request will not be a Resource but instead a
_application/json-patch+json.`,
		`__http-search__ sends an HTTP POST as described by the _search_ interaction described at
https://www.hl7.org/fhir//http.html#search. Headers and parameters described can be included as arguments to the
function. The GET search interaction can use the __http-read__ function. The __http-search__, like the
__http-each__ function, expects a callback function that is called for each matching Resource in the returned
Bundle and linked page Bundles.`,
		`__http-batch__ helps to build and submit a Bundle for a batch or transaction.`,
		`__http-operation__ send a request to invoke an operation either using a GET or POST HTTP request.`,
	},
	"resources": []string{
		`Resource are the leaves or concrete types of the FHIR inheritance tree. They all inherit from the
DoaminResource which inherits from Resource. Resources have named properties with each property being of a
specific type. Each property also has a cardinality which defines a minimum of either 0 or 1 and a maximum
of 1 or unlimited which is denoted by a * in the descriptions.
`,
		`The Resource types are:
`,
	},
	"datatypes": []string{
		`The FHIR specification defines a type hierarchy on https://fhir.hl7.org/fhir/datatypes.html. It also
describes each types but indicates all DataTypes inherit from just Element. The types in the specification
also deviate from what is in the schema files used to dynamically build the types for the imported packages.
As an example, the specifcation identifies a MoneyQuantity but the schema calls that same type, Money. Other
than minor inconsistencies such as those the schema does match the FHIR specification web pages.
`,
		`The DataTypes are:
`,
	},
	"backbones": []string{
		`Backbone types are embedded in other types as reflected in the name of the type. As an example,
then Patient_Communication is a backbone type in the Patient resource. The FHIR specification in the framework
diagram (https://fhir.hl7.org/fhir/types.html#2.1.27.0) shows a BackboneElement and BackboneType. They are
effectively identical. In the description of a resource the embedded elements are described as BackboneElement
in most cases but in a few BackboneType is specified. For this, __fhir__ package, BackboneType is used in
all cases.
`,
		`The Backbone types are:
`,
	},
	"primitives": []string{
		`The details of each primitive type can be viewed using the __describe__ function. The primitive type
framework is somewhat convoluted possible due to the XML heritage of then FHIR specification. While a primitive
type represents a single value that single value is also described as having an __id__ and __extension__ field.
The FHIR specification partially works around this disconnect by providing a mechanism in any container that has
a property that is a primitive type. Property names that start with an underscore character are considered
extensions of a property with the same name if the leading underscore is ignored. The specification also assumes
all primitive types are based on XSD types. In this, __fhir__ package, primitive types are built on base Lisp
typesof classes such as __fixnum__, __string__, etc.
`,
		`The primitive types that are the classes of the simple values in the FHIR datatypes are:
`,
	},
	"explore": []string{
		`This, __fhir__ package, can be used as an alternative to or an offline version of the FHIR web pages.
In addition to the __http-help__ topics, the __describe__ and __describe-type__ functions are available for
types, functions, and instances.
`,
		`The type description format is similar to the FHIR web pages and includes property names, cardinality,
type, and a description. The __describe-type__ has options for a full, expanded display and for alternating
backgrounds to make property separation more clear. When displaying the full description all inherited properties
are shown in addition to extentions and a search parameter table.
`,
		`An example of the __describe-type__ output but cut off in after a few properties is:
`,
		`▶ (describe-type 'fhir5:basic)
__fhir5:Basic__ is a FHIR Resource:
  Documentation:
    Basic is used for handling concepts not yet defined in FHIR, narrative-only
    resources that don't map to an existing resource, and custom resources not
    appropriate for inclusion in the FHIR specification.
  Direct Ancestor: DomainResource
  Class precedence list: fhir:Basic fhir5:DomainResource fhir5:Resource fhir5:Base t
  Properties:
    __Name__          __Card__. __Type__             __Description__
    resourceType  1..1  code             This is a Basic resource
    author        0..1  Reference        Indicates who was responsible for
                                         creating the resource instance.
    code          1..1  CodeableConcept  Identifies the 'type' of resource -
                                         equivalent to the resource name for
                                         other resources.
    ...
`,
		`With the __:full__ option the extensions and search parameters are listed as well.
`,
		`    ...
    _ _text              0..*  Extension        Extensions for text.
  Search Parameters:
    Name        Type       Description                         Expression
    author      reference  Who created                         Basic.author
    code        token      Kind of Resource                    Basic.code
    ...
`,
		`Inspecting an instance shows the properties set in a Simple Encoding Notation (SEN) format
as defined at https://github.com/ohler55/ojg/blob/develop/sen.md.
`,
		`▶ (describe (make-instance 'fhir5:Patient :data "{resourceType:Patient id:p001 name:[{given:[Quinn] family:Quux}]}"))

__#<fhir5:Patient 488285c08900>__, an instance of __fhir5:Patient__,
  {
    id: p001
    name: [
      {family: Quux given: [Quinn]}
    ]
    resourceType: Patient
  }
`,
	},
	"summary": []string{
		`The table that follows is based on https://www.hl7.org/fhir//http.html#summary and is a summary of the
requests and responses with a FHIR server.
`,
		` __Interaction     Function           Path                                 Verb    Response Body__`,
		`^ read            http-read          /[type]/[id]                         GET     Resource
 vread           http-read          /[type]/[id]/ history/[vid]          GET     Resource
 update          http-update        /[type]/[id]                         PUT     Resource
 patch           http-patch         /[type]/[id]                         PATCH   Resource
 delete          http-delete        /[type]/[id]                         DELETE
 create          http-create        /[type]                              POST    Resource
 search-type     http-read          /[type]?                             GET     Bundle
                 http-search        /[type]_search                       POST    Bundle
 search-system   http-read          /?                                   GET     Bundle
                 http-search        /_search                             POST    Bundle
 search-         http-read          /[compartment]/[id]/*?               GET     Bundle
 compartment     http-read          /[compartment]/[id]/[type]?          GET     Bundle
                 http-search        /[compartment]/[id]/_search?         POST    Bundle
                 http-search        /[compartment]/[id]/[type]/_search?  POST    Bundle
 capabilities    http-capabilities  /metadata                            GET     CapabilityStatement
 transaction     http-batch         /                                    POST    Bundle
 batch           http-batch         /                                    POST    Bundle
 history-inst    http-history       /[type]/[id]_history                 GET     Bundle
 history-type    http-history       /[type]/_history                     GET     Bundle
 history-system  http-history       /[type]/_history                     GET     Bundle
 (operation)     http-operation     /$[name]                          GET/POST   Parameters/Resource
                 http-operation     /[type]/$[name]                   GET/POST   Parameters/Resource
                 http-operation     /[type]/[id]/$[name]              GET/POST   Parameters/Resource
`,
	},
	"headers": []string{
		`FHIR server interactions make use of HTTP headers to convey some information such as conditions to check before
taking action. Not all servers support all headers. A summary of https://www.hl7.org/fhir//http.html#Http-Headers is:

`,
		`_Request Headers:_`,
		`__Accept__ used to request that the server response be of one of the specified mime types.`,
		`__If-Match__ specifies a condition where the provided ETag formatted value must match the resource being acted on.`,
		`__If-Modified-Since__ specifies a time based condition for a read. The format of the time must be
<day-name>, <day> <month> <year> <hour>:<minute>:<second> GMT.`,
		`__If-None-Exist__ specifies a condition for a create using parameters encoded search url encoding.`,
		`__If-None-Match__ specifies a condition that compares resources to multiple ETag formatted values`,
		`__Prefer__ specifies a preference for the return value such as return=minimal.`,
		`__X-Forwarded-For__ an optional header that identifies the origin of a forwarded request.`,
		`__X-Forwarded-Host__ an optional header that identifies the host of forwarded request.

`,

		`_Response Headers:_`,
		`__ETag__ used to relay the resource version as a weak (W/) ETag such as W/"123".`,
		`__Last-Modified__ used to relay the resource meta.lastUpdates time.`,
		`__Location__ used to relay the resource location as a URL.`,
		`__Content-Location__ for async use to indicate the location of the response.

`,
		`_Optional Bidirectional Headers:_`,
		`__X-Request-Id__ an id assigned by the client if on a request or by the server if in a response header.`,
		`__X-Correlation-Id__ is a client assigned id that is passed back in a response.`,
		`__X-Intermediary__ identifies an intermediary agent that modifies either a request or a response.`,
		`__X-Forwarded-Proto__ identifies the original protocol used by a client that has connected to an intermediary.`,
		`__X-Forwarded-Port__ identifies the intermediary port that it accepts connections on.`,
		`__X-Forwarded-Prefix__ a non-standard header used with a proxy.`,
	},
	"parameters": []string{
		`General paramters such as _ format_, _ summary_, _ pretty_, and _ elements_ apply to any request. Others
can be more restrictive. Parameters are:

`,
		`_ _at [date(Time)] only include resource versions that were current at some point during the time period
specified in the date time value.`,
		`_ _contained [true|fales|both] Whether to return resources contained in other resources in the
search matches.`,
		`_ _containedType [container|contained] If returning contained resources, whether to return the
contained or container resources`,
		`_ _content [string] Text search against the entire resource.`,
		`_ _count [integer] the maximum number of search results on a page, excluding related resources included
by _include or _revinclude or OperationOutcomes. The server is not bound to return the number requested, but
cannot return more `,
		`_ _elements [string,...] Request that only a specific set of elements be returned for resources.`,
		`_ _filter [special] Filter search parameter which supports a more sophisticated grammar for
searching. See https://www.hl7.org/fhir//search_filter.html.`,
		`_ _format Supersedes the Accept header. For this package a format of other that application/fhir+json
is unsupported.`,
		`_ _graph [reference] Include additional resources according to a GraphDefinition.`,
		`_ _has [special] Provides limited support for reverse chaining.
See https://www.hl7.org/fhir//search.html#has`,
		`_ _id [token] Resource id (not a full URL)`,
		`_ _in [reference] Group, List, or CareTeam membership.`,
		`_ _include [string] Other resources to include in the search results that search matches point to.`,
		`_ _language [token] Language of the resource content. (Resource.language)`,
		`_ _lastUpdated [date] Date last updated. Server has discretion on the boundary precision.
(Resource.meta.lastUpdated)`,
		`_ _list [reference] All resources in nominated list (by id, not a full URL).`,
		`_ _maxresults [number] Hint to a server that only the first 'n' results will ever be processed.`,
		`_ _pretty [true|false] Indicates whether the return should be pretty or not. Since the JSON content is
parsed into a instance on receiving a response the parameter has no visible effect on the response other than
to change the content length.`,
		`_ _profile [reference] Search for all resources tagged with a profile. (Resource.meta.profile)`,
		`_ _query [string] Custom named query.`,
		`_ _revinclude [string] Other resources to include in the search results when they refer to search matches.`,
		`_ _score [true|false] Request match relevance in results.`,
		`_ _security [token] Search by a security label. (Resource.meta.security)`,
		`_ _since [instant] Only include resource versions that were created at or after the given instant in time.`,
		`_ _sort_ [choice] allowed sort values are limited to:`,
		`^    -_lastUpdate (default) - sort in descending lastUpdated order.`,
		`^    _lastUpdate - sort in ascending lastUpdated order.`,
		`^    none - data will have no defined sort order.`,
		`_ _source [uri] Search by where the resource comes from. (Resource.meta.source)`,
		`_ _summary [true|text|data|count|false] A request to the server to limit the fields in the response.
It can be one of:`,
		`^    true -  Return a limited subset of elements from the resource. This subset SHOULD
            consist solely of all supported elements that are marked as "summary" in
            the base definition of the resource(s) `,
		`^    text -  Return only the text, id, meta, and top-level mandatory elements (these
            mandatory elements are included to ensure that the payload is valid FHIR;
            servers MAY omit elements within these sub-trees as long as they ensure
            that the payload is valid). Servers MAY return extensions, but clients
            SHOULD NOT rely on extensions being present and SHOULD use another search
            mode if data contained in extensions is required.`,
		`^    data -  Remove the text element.`,
		`^    count - Search only: just return a count of the matching resources, without
            returning the actual matches.`,
		`^    false - Return all parts of the resource(s).`,
		`_ _tag [token] Search by a resource tag (Resource.meta.tag)`,
		`_ _text [string] Text search against the narrative.`,
		`_ _total [none|estimate|accurate] Request a precision of the total number of results for a request.`,
		`_ _type [token] Is used to allow filtering of types in searches that are performed across multiple
resource types (e.g., searches across the server root).`,

		`Search parameters follow the general param=value format but with a variety of refinements for both the
param and value to build more complex queries. A rather extensive description of the search language is
https://www.hl7.org/fhir//search.html#3.2.1`,
	},
	"search": []string{
		`Queries for searching are defined by the query parameters in a URL or encoded in a POST request. They
follow the general param=value format but with a variety of refinements for both the parameter key and value
to build more complex queries. A rather extensive description of the search language is
https://www.hl7.org/fhir//search.html`,
		`While including a complete description of the search facility isn't feasible, summaries of the
prefixes and modifies are included here.`,
		` __Prefix
  Code  Description__`,
		`^ eq    The resource value is equal to or fully contained by the parameter value.
 ne    The resource value is not equal to the parameter value.
 gt    The resource value is greater than the parameter value.
 lt    The resource value is less than the parameter value.
 ge    The resource value is greater or equal to the parameter value.
 le    The resource value is less or equal to the parameter value.
 sa    The resource value starts after the parameter value.
 eb    The resource value ends before the parameter value.
 ap    The resource value is approximately the same to the parameter value.
       Note that the recommended value for the approximation is 10% of the stated
       value (or for a date, 10% of the gap between now and the date), but systems
       may choose other values where appropriate the range of the parameter value
       overlaps with the range of the resource value`,
		` __Modifier       Types      Description__`,
		`^ above          reference  Tests whether the value in a resource is or subsumes
                token      the supplied parameter value (is-a, or hierarchical
                uri        relationships).`,
		`^ below          reference  Tests whether the value in a resource is or is
                token      subsumed by the supplied parameter value (is-a, or hierarchical
                uri        relationships).`,
		`^ code-text      reference  Tests whether the textual display value in a resource
                token      (e.g., CodeableConcept.text, Coding.display, or
                           Reference.display) matches the supplied parameter value.`,
		`^ contains       string     Tests whether the value in a resource includes the
                uri        supplied parameter value anywhere within the field being searched.`,
		`^ exact          string     Tests whether the value in a resource exactly matches
                           the supplied parameter value (the whole string, including casing
                           and accents).`,
		`^ identifier     reference  Tests whether the Reference.identifier in a resource
                           (rather than the Reference.reference) matches the supplied
                           parameter value.`,
		`^ in             token      Tests whether the value in a resource is a member of
                           the supplied parameter ValueSet.`,
		`^ iterate        n/a (not   The search parameter indicates an inclusion directive
                allowed    (_include, _revinclude) that is applied to an included
                anywhere   resource instead of the matching resource.
                by default)`,
		`^ missing        date       Tests whether the value in a resource is present
                number     (when the supplied parameter value is true) or absent (when
                quantity   the supplied parameter value is false).
                reference
                string
                token
                uri`,
		`^ not            token      Tests whether the value in a resource does not match
                           the specified parameter value. Note that this includes resources
                           that have no value for the parameter.`,
		`^ not-in         reference  Tests whether the value in a resource is not a member
                token      of the supplied parameter ValueSet.`,
		`^ of-type        token      Tests whether the Identifier value in a resource
                (only      matches the supplied parameter value.
                Identifier)`,
		`^ text           reference  Tests whether the textual value in a resource
                token      (e.g., CodeableConcept.text, Coding.display, Identifier.type.text,
                           or Reference.display) matches the supplied parameter value using
                           basic string matching (begins with or is, case-insensitive).`,
		`^ text           string     The search parameter value should be processed as input
                           to a search with advanced text handling.`,
		`^ text-advanced  reference  Tests whether the value in a resource matches the supplied
                token      parameter value using advanced text handling that searches text
                           associated with the code/value - e.g., CodeableConcept.text,
                           Coding.display, or Identifier.type.text.`,

		`^ [type]         reference  Tests whether the value in a resource points to a resource
                           of the supplied parameter type. Note: a concrete ResourceType is
                           specified as the modifier (e.g., not the literal :[type], but a
                           value such as :Patient).`,
	},
	"history": []string{
		`Retrieving the history of a resource, if the resource type supports history, is best done using
the __http-each__ function. The __http-each__ function iterates over all returned resources and follows
pages transparently. By including the _:history_ argument to __t__, _ _history is appended to the URL
used to start fetching previous versions of a resource.`,
		`Deleted resources may be in the history responses but will show up as __bag__ instances since
no resourceType is included in the entry for a deleted resource.`,
		`Parameters that can be used when retrieving history include _ _count, _ _since, _ _at, _ _list,
and _ _sort. More details regarding history can be found on at https://www.hl7.org/fhir//http.html#history.`,
	},
	"compartment": []string{
		`The __http-each__ or the __http-search__ functions can be used for compartment searches. Providing
a value for the _:compartment_ keyword option will form one of the compartment URLs. If a type of nil
is specified or no type is specified then either the pattern ending in "/*?" or "/_search?" is used.
Compartments are defined at https://www.hl7.org/fhir//compartmentdefinition.html. The patterns for
compartment searching are:`,
		`^ GET  /[compartment]/[id]/*?`,
		`^ GET  /[compartment]/[id]/[type]?`,
		`^ POST /[compartment]/[id]/_search?`,
		`^ POST /[compartment]/[id]/[type]/_search?`,
		`Defined compartments are Patient, Encounter, RelatedPerson, Practitions, and Device.`,
	},
	"read-example": []string{
		`Reading from a FHIR server is one of the most common uses of the server. This example covers making a read
request with the __http-read__ function to access a Patient resource with an id of "P001". __http-read__ requires at
least one argument, the _base_ which can be either a URL as a string or a property list that includes a URL targeting
a FHIR server plus default values for the other optional key arguments the function accepts.
`,
		`For this example the fictitious FHIR server has at http://fire.fake:8080. For purposes of this example, the
server expects authorization with a bearer token of "access-token". Instead of having to add that information on every
call it can be placed in a property list _base_. Other default such as a timeout and the default FHIR package can
also be included.`,
		`^
▶ (defvar fire-base '(:url "http://fire.fake:8080"
                    :headers ("Authentication" "Bearer access-token")
                    :timeout 5
                    :fhir-package fhir5))
fire-base`,
		`The read request is then send and the response bound to a variable.`,
		`^
▶ (defvar resp (http-read fire-base :type "Patient" :id "P001"))
resp`,
		`A quick check to verify the request returned a 200 HTTP success status code.`,
		`^
▶ (car resp)
200`,
		`There are a few useful pieces of information in the returned headers: Location reiterates then URL to the
returned resource, ETag identifies the version which is also in the resource meta.version field, and Last-Modified
which is also in the resource meta field as meta.lastUpdated.`,
		`^
▶ (cadr (assoc "Location" (nth 2 resp)))
http://fire.fake:8080/Patient/P001`,
		`(nth 2 resp) returns the headers. Using the assoc function a list of word "Location" and the value are
returned. Taking the cadr or that is the value of the location. The same approach can be used to find the ETag and
Last-Modified if they are present.`,
		`^
▶ (cadr (assoc "ETag" (nth 2 resp)))
W/"v3"
▶ (cadr (assoc "Last-Modified" (nth 2 resp)))
"Mon, 05 Jan 2026 22:33:44 GMT"`,
		`The __describe__ function can be used to see all the properties in the returned resource.`,
		`^
▶ (describe (cadr resp))
#<fhir5:Patient 27fcb054220>, an instance of fhir5:Patient,
  {
    birthDate: "1969-01-02"
    id: P001
    meta: {
      lastUpdated: "2026-01-05T22:33:44.123Z"
      versionId: "v3"
    }
    name: [
      {family: Racoon given: [Rocky]}
    ]
    resourceType: Patient
  }
`,
		`Individual element of the returned resource can be accessed with the __instance-get__ function which
utilizes JSONPath to navigate the resource.`,
		`^
▶ (instance-get (cadr resp) "name[*].given[0]")
"Rocky"`,
	},
	"create-example": []string{
		`Creating a new resource starts with building a resource of the correct type. That resource must include the
resourceType property.`,
		`^
▶ (defvar pat (make-instance 'fhir5:patient :data "{resourceType: Patient, name:[{given:[Rocky] family:Racoon}]}"))
pat`,
		`It's best to validate the resource before sending a create request to the FHIR server. Instances in the
__fhir__ package are a Flavor/CLOS belnd referred to as FLOS types. That allows an Instance to be validated
either using the __valid-p__ function or to send the instance a request to validate itself.`,
		`^
▶ (send pat :valid-p)
t
▶ (valid-p pat)
t`,
		`With confidence that the Instance is a valid Patient a request is sent to the FHIR server to create the
Instance. Note that any id property will be ignored as will the meta.versionId and meta.lastUpdated properties.
Like the read-example a fire-base is defined first.`,
		`^
▶ (defvar fire-base '(:url "http://fire.fake:8080"
                    :headers ("Authentication" "Bearer access-token")
                    :timeout 5
                    :fhir-package fhir5))
fire-base
▶ (defvar create-resp (http-create pat fire-base))
create-resp`,
		`Following the read-example inspection of a response, the ETag, Last-Modified, and Location headers can be
viewed and should match the meta.versionId, meta.lastUpdated, and the newly assigned id property.`,
		`^
▶ (car create-resp)
201
▶ (cadr (assoc "Location" (nth 2 create-resp)))
http://fire.fake:8080/Patient/P002/_history/v01
▶ (cadr (assoc "ETag" (nth 2 create-resp)))
W/"v01"
▶ (cadr (assoc "Last-Modified" (nth 2 create-resp)))
"Mon, 05 Jan 2026 22:33:44 GMT"
▶ (describe (cadr create-resp))
#<fhir5:Patient 27fcb054220>, an instance of fhir5:Patient,
  {
    id: P002
    meta: {
      lastUpdated: "2026-01-05T22:33:44.123Z"
      versionId: "v01"
    }
    name: [
      {family: Racoon given: [Rocky]}
    ]
    resourceType: Patient
  }`,
	},
	"update-example": []string{
		`Updating a resource starts with building or downloading a resource of the correct type. If built from scratch
the resource must include the resourceType and id properties.`,
		`^
▶ (defvar fire-base '(:url "http://fire.fake:8080"
                    :headers ("Authentication" "Bearer access-token")
                    :timeout 5
                    :fhir-package fhir5))
fire-base
▶ (defvar pat001 (cadr (http-read fire-base :type "Patient" :id "P001")))
pat001`,
		`We took a shortcut by not checking the response status and assuming the returned resource was a Patient.
Next a change is made to the returned resource by setting the active property using the __instance-set__ function.
An alterenative would be to send the resource a __:set__ method with the name and value similar to the using the
__instance-set__ function.`,
		`^
▶ (instance-set pat001 "active" t)
nil`,
		`The __http-update__ function is used to update the P001 Patient.`,
		`^
▶ (defvar update-resp (http-update pat fire-base))
update-resp`,
		`Following the read-example inspection of a response, the ETag, Last-Modified, and Location headers can be
viewed and should match the meta.versionId, meta.lastUpdated, and the id property.`,
		`^
▶ (car update-resp)
200
▶ (cadr (assoc "Location" (nth 2 update-resp)))
http://fire.fake:8080/Patient/P001/_history/v02
▶ (cadr (assoc "ETag" (nth 2 update-resp)))
W/"v02"
▶ (cadr (assoc "Last-Modified" (nth 2 update-resp)))
"Mon, 05 Jan 2026 22:33:44 GMT"
▶ (describe (cadr update-resp))
#<fhir5:Patient 27fcb054220>, an instance of fhir5:Patient,
  {
    active: true
    id: P001
    meta: {
      lastUpdated: "2026-01-05T22:33:44.123Z"
      versionId: "v02"
    }
    name: [
      {family: Racoon given: [Rocky]}
    ]
    resourceType: Patient
  }`,
	},
	"delete-example": []string{
		`Resource can be removed from a FHIR server using the __http-delete__ function. A standard delete requires a
resource type and an id. More advanced delete operation can make use of a resource type and a condition. The
simple delete starts off like all other http operations with the creation of a base.`,
		`^
▶ (defvar fire-base '(:url "http://fire.fake:8080"
                    :headers ("Authentication" "Bearer access-token")
                    :timeout 5
                    :fhir-package fhir5))
fire-base
`,
		`The __http-delete__ function is then called and the response checked. From the response the status code
is of interest and is expected to be 204 if there is no content. If content such as an OperationOutcome is
included then the status on success will be 200.`,
		`^
▶ (defvar delete-resp (http-delete fire-base :type "Patient" :id "P001"))
delete-resp
▶ (car update-resp)
204
`,
	},
	"search-example": []string{
		`Searches are done using either a GET or a POST. The difference as far as the __http-search__ function
is concerned is simply changing from using the _:params_ keyword to _:query_. __http-search__ makes use of
callbacks. The same approach is used for __http-each__ and __http-history__. The first argument to __http-search__
is a function that is called for each resource found by the search. Since a FHIR search returns a Bundle
resource, the entries (entry property) in the Bundle is iterated over with each resource created and passed
as the sole argument to the callback function. A _:limit_ option will abort the search when the limit is
reached. If the Bundle has a _next_ property the URL in the _next_ property is used to request the next
batch. This process continues until either no more results are available or the limit is reached.`,
		`To get started a _base_ is set up:`,
		`^
▶ (defvar fire-base '(:url "http://fire.fake:8080"
                    :headers ("Authentication" "Bearer access-token")
                    :timeout 5
                    :fhir-package fhir5))
fire-base
`,
		`In this example a limit of 2 is set and the callback function prints a description of each resource.`,
		`^
▶ (http-search (lambda (r) (send r :describe)) fire-base :type "Patient" :params '("given" "Pete"))
#<fhir5:Patient 1ceb160a0200>, an instance of fhir5:Patient,
  {
    id: "P001"
    meta: {
      lastUpdated: "2026-02-23T23:54:59.123Z"
      versionId: "v02"
    }
    name: [
      {family: Parrot given: [Pete]}
    ]
    resourceType: Patient
  }
#<fhir5:Patient 1ceb160a0240>, an instance of fhir5:Patient,
  {
    id: "P003"
    meta: {
      lastUpdated: "2026-02-23T23:54:43.337Z"
      versionId: "v003"
    }
    name: [
      {family: Porcupine given: [Pete]}
    ]
    resourceType: Patient
  }
`,
	},
	"patch-example": []string{
		`TBD`,
	},
	"batch-example": []string{
		`TBD`,
	},
	// "graphql": []string{`TBD`},
	// "jet-help": []string{`TBD`},
	// "mllp-help": []string{`TBD`},
}

var topicHelpExtras = map[string]func(b []byte, right int, ansi bool) []byte{
	"resources":  helpResourcesExtra,
	"datatypes":  helpDatatypesExtra,
	"backbones":  helpBackbonesExtra,
	"primitives": helpPrimitivesExtra,
}

func initHTTPHelp() {
	slip.Define(
		func(args slip.List) slip.Object {
			f := HTTPHelp{Function: slip.Function{Name: "http-help", Args: args, SkipEval: []bool{true}}}
			f.Self = &f
			return &f
		},
		&slip.FuncDoc{
			Name: "http-help",
			Args: []*slip.DocArg{
				{Name: "&optional"},
				{
					Name: "topic",
					Type: "string|symbol",
					Text: `Names a topic to display a description for.`,
				},
			},
			Return: "nil",
			Text:   `__http-help__ displays descriptions for various topics related to FHIR HTTP use.`,
		}, &Pkg)
}

// HTTPHelp represents the http-help function.
type HTTPHelp struct {
	slip.Function
}

// Call the the function with the arguments provided.
func (f *HTTPHelp) Call(s *slip.Scope, args slip.List, depth int) slip.Object {
	slip.CheckArgCount(s, depth, f, args, 0, 1)
	ansi := s.Get("*print-ansi*") != nil
	right := int(s.Get("*print-right-margin*").(slip.Fixnum))
	var (
		b     []byte
		extra func(b []byte, right int, ansi bool) []byte
	)
	help := helpTop
	if 0 < len(args) {
		topic := slip.MustBeString(args[0], "topic")
		if h := topicHelp[strings.ToLower(topic)]; 0 < len(help) {
			help = h
		}
		extra = topicHelpExtras[topic]
	}
	b = appendHelpDoc(b, help, right, ansi)
	if extra != nil {
		b = extra(b, right, ansi)
	}
	b = append(b, '\n')

	so := s.Get("*standard-output*")
	w := so.(io.Writer)
	if _, err := w.Write(b); err != nil {
		ss, _ := so.(slip.Stream)
		slip.StreamPanic(s, depth, ss, "write failed: %s", err)
	}
	return nil
}

func appendHelpDoc(b []byte, help []string, right int, ansi bool) []byte {
	for i, h := range help {
		if 0 < i {
			b = append(b, '\n')
		}
		switch h[0] {
		case '_':
			b = slip.AppendDoc(b, h, 4, right, ansi, 2)
		case '^':
			b = append(b, h[1:]...)
		default:
			b = append(b, '\n')
			b = slip.AppendDoc(b, h, 0, right, ansi)
		}
	}
	return append(b, '\n')
}

func helpResourcesExtra(b []byte, right int, ansi bool) []byte {
	var words []string

	for _, p := range slip.AllPackages() {
		for _, class := range p.AllClasses() {
			if t, _ := class.(*Type); t != nil {
				if t.parent == "DomainResource" {
					name := class.Name()
					words = append(words, t.pkg.Name+":"+name)
				}
			}
		}
	}
	return appendWords(b, words, right)
}

func helpDatatypesExtra(b []byte, right int, ansi bool) []byte {
	var words []string

	for _, p := range slip.AllPackages() {
		for _, class := range p.AllClasses() {
			if t, _ := class.(*Type); t != nil {
				if t.parent == "Element" {
					name := class.Name()
					words = append(words, t.pkg.Name+":"+name)
				}
			}
		}
	}
	return appendWords(b, words, right)
}

func helpBackbonesExtra(b []byte, right int, ansi bool) []byte {
	var words []string

	for _, p := range slip.AllPackages() {
		for _, class := range p.AllClasses() {
			if t, _ := class.(*Type); t != nil {
				if t.parent == "BackboneType" {
					name := class.Name()
					words = append(words, t.pkg.Name+":"+name)
				}
			}
		}
	}
	return appendWords(b, words, right)
}

func helpPrimitivesExtra(b []byte, right int, ansi bool) []byte {
	var words []string

	for _, p := range slip.AllPackages() {
		for _, class := range p.AllClasses() {
			if t, _ := class.(*Type); t != nil {
				name := class.Name()
				if 'a' <= name[0] && name[0] <= 'z' {
					words = append(words, t.pkg.Name+":"+name)
				}
			}
		}
	}
	return appendWords(b, words, right)
}

func appendWords(b []byte, words []string, right int) []byte {
	sort.Strings(words)
	width := right - 2
	var ww int
	for _, word := range words {
		if ww < len(word) {
			ww = len(word)
		}
	}
	ww += 2
	colCnt := width / ww
	var cnt int
	for _, word := range words {
		if cnt == 0 {
			b = append(b, '\n', ' ', ' ')
		}
		cnt++
		b = append(b, word...)
		b = append(b, bytes.Repeat([]byte{' '}, ww-len(word))...)
		if colCnt <= cnt {
			cnt = 0
		}
	}
	return append(b, '\n')
}
