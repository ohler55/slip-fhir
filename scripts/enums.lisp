
(defconstant *enum-map*
  ;; id of enum followed by the values
  '(
    ("example" "chol-mmol" "chol-mass" "chol")
    ("nhin-purposeofuse" "TREATMENT" "PAYMENT" "OPERATIONS" "SYSADMIN" "FRAUD" "PSYCHOTHERAPY" "TRAINING" "LEGAL"
                     "MARKETING" "DIRECTORY" "FAMILY" "PRESENT" "EMERGENCY" "DISASTER" "PUBLICHEALTH" "ABUSE"
                     "OVERSIGHT" "JUDICIAL" "LAW" "DECEASED" "DONATION" "RESEARCH" "THREAT" "GOVERNMENT" "WORKERSCOMP"
                     "COVERAGE" "REQUEST")
    ("example-metadata" "A" "B" "C")
    ("example-metadata-2" "A" "B" "C" "D")
    ("fhir-types" "Base")
    ("fhirpath-types" "http://hl7.org/fhirpath/System.String" "http://hl7.org/fhirpath/System.Boolean"
                  "http://hl7.org/fhirpath/System.Date" "http://hl7.org/fhirpath/System.DateTime"
                  "http://hl7.org/fhirpath/System.Decimal" "http://hl7.org/fhirpath/System.Integer"
                  "http://hl7.org/fhirpath/System.Time")
    ("administrative-gender" "male" "female" "other" "unknown")
    ("binding-strength" "required" "extensible" "preferred" "example")
    ("cdshooks-indicator" "info" "warning" "critical")
    ("concept-map-relationship" "related-to" "not-related-to")
    ("document-reference-status" "current" "superseded" "entered-in-error")
    ("FHIR-version" "0.01" "0.05" "0.06" "0.11" "0.0" "0.4" "0.5" "1.0" "1.1" "1.4" "1.6" "1.8" "3.0" "3.3" "3.5" "4.0"
                "4.1" "4.2" "4.3" "4.4" "4.5" "4.6" "5.0")
    ("note-type" "display" "print" "printoper")
    ("operation-outcome" "DELETE_MULTIPLE_MATCHES" "MSG_AUTH_REQUIRED" "MSG_BAD_FORMAT" "MSG_BAD_SYNTAX"
                     "MSG_CANT_PARSE_CONTENT" "MSG_CANT_PARSE_ROOT" "MSG_CREATED" "MSG_DATE_FORMAT" "MSG_DELETED"
                     "MSG_DELETED_DONE" "MSG_DELETED_ID" "MSG_DUPLICATE_ID" "MSG_ERROR_PARSING" "MSG_ID_INVALID"
                     "MSG_ID_TOO_LONG" "MSG_INVALID_ID" "MSG_JSON_OBJECT" "MSG_LOCAL_FAIL" "MSG_NO_EXIST"
                     "MSG_NO_MATCH" "MSG_NO_MODULE" "MSG_NO_SUMMARY" "MSG_OP_NOT_ALLOWED" "MSG_PARAM_CHAINED"
                     "MSG_PARAM_INVALID" "MSG_PARAM_MODIFIER_INVALID" "MSG_PARAM_NO_REPEAT" "MSG_PARAM_UNKNOWN"
                     "MSG_REMOTE_FAIL" "MSG_RESOURCE_EXAMPLE_PROTECTED" "MSG_RESOURCE_ID_FAIL"
                     "MSG_RESOURCE_ID_MISMATCH" "MSG_RESOURCE_ID_MISSING" "MSG_RESOURCE_NOT_ALLOWED"
                     "MSG_RESOURCE_REQUIRED" "MSG_RESOURCE_TYPE_MISMATCH" "MSG_SORT_UNKNOWN"
                     "MSG_TRANSACTION_DUPLICATE_ID" "MSG_TRANSACTION_MISSING_ID" "MSG_UNHANDLED_NODE_TYPE"
                     "MSG_UNKNOWN_CONTENT" "MSG_UNKNOWN_OPERATION" "MSG_UNKNOWN_TYPE" "MSG_UPDATED" "MSG_VERSION_AWARE"
                     "MSG_VERSION_AWARE_CONFLICT" "MSG_VERSION_AWARE_URL" "MSG_WRONG_NS" "SEARCH_MULTIPLE"
                     "SEARCH_NONE" "UPDATE_MULTIPLE_MATCHES")
    ("publication-status" "draft" "active" "retired" "unknown")
    ("relationship" "1" "2" "3" "4" "5")
    ("remittance-outcome" "complete" "error" "partial")
    ("search-param-type" "number" "date" "string" "token" "reference" "composite" "quantity" "uri" "special")
    ("usage-context-agreement-scope" "realm-base" "knowledge" "domain" "community" "system-design" "system-implementation")
    ("restful-interaction" "read" "vread" "update" "patch" "delete" "history" "create" "search" "capabilities" "transaction"
                       "batch" "operation")
    ("safety-entries" "life-cycle" "modifiers" "modifier-extensions" "must-support" "identity" "current" "error-checks"
                  "link-merge" "cs-declare" "valid-checked" "obs-focus" "time-zone" "date-rendering" "cross-resource"
                  "display-warnings" "search-parameters" "missing-values" "default-filters" "deletion-check"
                  "deletion-replication" "deletion-support" "check-consent" "distribute-aod" "check-clocks"
                  "check-dns-responses" "use-encryption" "use-tls" "use-smime" "use-tls-per-bcp195" "use-ouath"
                  "use-openidconnect" "use-rbac" "use-labels" "render-narratives" "check=validation" "use-provenance"
                  "enable-cors" "use-json" "json-for-errors" "use-format-header" "use-operation-outcome")
    ("concept-properties" "status" "inactive" "effectiveDate" "deprecated" "deprecationDate" "retirementDate"
                      "notSelectable" "parent" "child" "partOf" "synonym" "comment" "itemWeight")
    ("w3c-provenance-activity-type" "Generation" "Usage" "Communication" "Start" "End" "Invalidation" "Derivation"
                                "Revision" "Quotation" "Primary-Source" "Attribution" "Collection")
    ("extra-activity-type" "aggregate" "compose" "label")
    ("resource-status" "error" "proposed" "planned" "draft" "requested" "received" "declined" "accepted" "arrived" "active"
                   "suspended" "failed" "replaced" "complete" "inactive" "abandoned" "unknown" "unconfirmed"
                   "confirmed" "resolved" "refuted" "differential" "partial" "busy-unavailable" "free" "on-target"
                   "ahead-of-target" "behind-target" "not-ready" "transduc-discon" "hw-discon")
    ("tldc" "draft" "pending" "active" "review" "cancelled" "rejected" "retired" "terminated")
    ("etsi-signature-type" "ProofOfOrigin" "ProofOfReceipt" "ProofOfDelivery" "ProofOfSender" "ProofOfapproval"
                       "ProofOfCreation")
    ("astm-signature-type" "1.2.840.10065.1.12.1.1" "1.2.840.10065.1.12.1.2" "1.2.840.10065.1.12.1.3"
                       "1.2.840.10065.1.12.1.4" "1.2.840.10065.1.12.1.5" "1.2.840.10065.1.12.1.6"
                       "1.2.840.10065.1.12.1.7" "1.2.840.10065.1.12.1.8" "1.2.840.10065.1.12.1.9"
                       "1.2.840.10065.1.12.1.10" "1.2.840.10065.1.12.1.11" "1.2.840.10065.1.12.1.12"
                       "1.2.840.10065.1.12.1.13" "1.2.840.10065.1.12.1.14" "1.2.840.10065.1.12.1.15"
                       "1.2.840.10065.1.12.1.16" "1.2.840.10065.1.12.1.17" "1.2.840.10065.1.12.1.18")
    ("sample-security-structural-roles" "regulated-health-professionals" "non-regulated-health-professionals")
    ("fhir-format-type" "xml" "json" "ttl")
    ("conceptmap-properties" "relationshipRefinement")
    ("fhir-old-types" "BodySite" "CatalogEntry" "Conformance" "DataElement" "DeviceComponent" "DeviceUseRequest"
                  "DeviceUseStatement" "DiagnosticOrder" "DocumentManifest" "EffectEvidenceSynthesis"
                  "EligibilityRequest" "EligibilityResponse" "ExpansionProfile" "ImagingManifest"
                  "ImagingObjectSelection" "Media" "MedicationOrder" "MedicationUsage" "MedicinalProduct"
                  "MedicinalProductAuthorization" "MedicinalProductContraindication" "MedicinalProductIndication"
                  "MedicinalProductIngredient" "MedicinalProductInteraction" "MedicinalProductManufactured"
                  "MedicinalProductPackaged" "MedicinalProductPharmaceutical" "MedicinalProductUndesirableEffect"
                  "Order" "OrderResponse" "ProcedureRequest" "ProcessRequest" "ProcessResponse" "ReferralRequest"
                  "RequestGroup" "ResearchDefinition" "ResearchElementDefinition" "RiskEvidenceSynthesis" "Sequence"
                  "ServiceDefinition" "SubstanceSpecification")
    ("color-names" "aliceblue" "antiquewhite" "aqua" "aquamarine" "azure" "beige" "bisque" "black" "blanchedalmond" "blue"
               "blueviolet" "brown" "burlywood" "cadetblue" "chartreuse" "chocolate" "coral" "cornflowerblue"
               "cornsilk" "crimson" "cyan" "darkblue" "darkcyan" "darkgoldenrod" "darkgray" "darkgreen" "darkgrey"
               "darkkhaki" "darkmagenta" "darkolivegreen" "darkorange" "darkorchid" "darkred" "darksalmon"
               "darkseagreen" "darkslateblue" "darkslategray" "darkslategrey" "darkturquoise" "darkviolet" "deeppink"
               "deepskyblue" "dimgray" "dimgrey" "dodgerblue" "firebrick" "floralwhite" "forestgreen" "fuchsia"
               "gainsboro" "ghostwhite" "gold" "goldenrod" "gray" "green" "greenyellow" "grey" "honeydew" "hotpink"
               "indianred" "indigo" "ivory" "khaki" "lavender" "lavenderblush" "lawngreen" "lemonchiffon" "lightblue"
               "lightcoral" "lightcyan" "lightgoldenrodyellow" "lightgray" "lightgreen" "lightgrey" "lightpink"
               "lightsalmon" "lightseagreen" "lightskyblue" "lightslategray" "lightslategrey" "lightsteelblue"
               "lightyellow" "lime" "limegreen" "linen" "magenta" "maroon" "mediumaquamarine" "mediumblue"
               "mediumorchid" "mediumpurple" "mediumseagreen" "mediumslateblue" "mediumspringgreen" "mediumturquoise"
               "mediumvioletred" "midnightblue" "mintcream" "mistyrose" "moccasin" "navajowhite" "navy" "oldlace"
               "olive" "olivedrab" "orange" "orangered" "orchid" "palegoldenrod" "palegreen" "paleturquoise"
               "palevioletred" "papayawhip" "peachpuff" "peru" "pink" "plum" "powderblue" "purple" "rebeccapurple"
               "red" "rosybrown" "royalblue" "saddlebrown" "salmon" "sandybrown" "seagreen" "seashell" "sienna"
               "silver" "skyblue" "slateblue" "slategray" "slategrey" "snow" "springgreen" "steelblue" "tan" "teal"
               "thistle" "tomato" "turquoise" "violet" "wheat" "white" "whitesmoke" "yellow" "yellowgreen")
    ("conformance-expectation" "SHALL" "SHOULD" "MAY" "SHOULD-NOT")
    ("related-artifact-type-expanded" "reprint" "reprint-of")
    ("narrative-status" "generated" "extensions" "additional" "empty")
    ("identifier-use" "usual" "official" "temp" "secondary" "old")
    ("quantity-comparator" "<" "<=" ">=" ">" "ad")
    ("name-use" "usual" "official" "temp" "nickname" "anonymous" "old")
    ("address-use" "home" "work" "temp" "old" "billing")
    ("address-type" "postal" "physical" "both")
    ("contact-point-system" "phone" "fax" "email" "pager" "url" "sms" "other")
    ("contact-point-use" "home" "work" "temp" "old" "mobile")
    ("event-timing" "MORN" "MORN.early" "MORN.late" "NOON" "AFT" "AFT.early" "AFT.late" "EVE" "EVE.early" "EVE.late" "NIGHT"
                "PHS" "IMD")
    ("timing-abbreviation" "C")
    ("days-of-week" "mon" "tue" "wed" "thu" "fri" "sat" "sun")
    ("virtual-service-type" "zoom" "ms-teams" "whatsapp")
    ("price-component-type" "base" "surcharge" "deduction" "discount" "tax" "informational")
    ("contributor-type" "author" "editor" "reviewer" "endorser")
    ("sort-direction" "ascending" "descending")
    ("operation-parameter-use" "in" "out")
    ("related-artifact-type" "documentation" "justification" "citation" "predecessor" "successor" "derived-from"
                         "depends-on" "composed-of" "part-of" "amends" "amended-with" "appends" "appended-with" "cites"
                         "cited-by" "comments-on" "comment-in" "contains" "contained-in" "corrects" "correction-in"
                         "replaces" "replaced-with" "retracts" "retracted-by" "signs" "similar-to" "supports"
                         "supported-with" "transforms" "transformed-into" "transformed-with" "documents"
                         "specification-of" "created-with" "cite-as")
    ("citation-artifact-classifier" "audio" "D001877" "cds-artifact" "D016420" "common-share" "D019991" "D064886"
                                "dataset-unpublished" "Electronic" "Electronic-eCollection" "Electronic-Print"
                                "executable-app" "fhir-resource" "image" "interactive-form" "D016428" "D016422"
                                "machine-code" "medline-base" "prediction-model" "D000076942" "Print"
                                "Print-Electronic" "project-specific" "protocol" "pseudocode" "D016425"
                                "standard-specification" "terminology" "D059040" "webpage")
    ("trigger-type" "named-event" "periodic" "data-changed" "data-accessed" "data-access-ended")
    ("constraint-severity" "error" "warning")
    ("resource-slicing-rules" "closed" "open" "openAtEnd")
    ("resource-aggregation-mode" "contained" "referenced")
    ("property-representation" "xmlAttr" "xmlText" "typeAttr" "cdaText" "xhtml")
    ("reference-version-rules" "either" "independent" "specific")
    ("discriminator-type" "value" "exists" "pattern" "type" "profile" "position")
    ("additional-binding-purpose" "maximum" "minimum" "required" "extensible" "candidate" "current" "preferred" "ui"
                              "starter" "component")
    ("event-status" "preparation" "in-progress" "not-done" "on-hold" "stopped" "completed" "entered-in-error" "unknown")
    ("request-status" "draft" "active" "on-hold" "revoked" "completed" "entered-in-error" "unknown")
    ("request-intent" "proposal" "plan" "directive" "order" "option")
    ("request-priority" "routine" "urgent" "asap" "stat")
    ("product-status" "active" "entered-in-error")
    ("resource-validation-mode" "create" "update" "delete" "profile")
    ("version-algorithm" "semver" "integer" "alpha" "date" "natural")
    ("flag-status" "active" "inactive" "entered-in-error")
    ("allergy-intolerance-type" "allergy" "intolerance")
    ("allergy-intolerance-category" "food" "medication" "environment" "biologic")
    ("allergy-intolerance-criticality" "low" "high" "unable-to-assess")
    ("reaction-event-severity" "mild" "moderate" "severe")
    ("care-team-status" "proposed" "active" "suspended" "inactive" "entered-in-error")
    ("capability-statement-kind" "instance" "capability" "requirements")
    ("restful-capability-mode" "client" "server")
    ("restful-security-service" "OAuth" "SMART-on-FHIR" "NTLM" "Basic" "Kerberos" "Certificates")
    ("versioning-policy" "no-version" "versioned" "versioned-update")
    ("conditional-read-status" "not-supported" "modified-since" "not-match" "full-support")
    ("conditional-delete-status" "not-supported" "single" "multiple")
    ("reference-handling-policy" "literal" "logical" "resolves" "enforced" "local")
    ("message-transport" "http" "ftp" "mllp")
    ("event-capability-mode" "sender" "receiver")
    ("document-mode" "producer" "consumer")
    ("detectedissue-status" "mitigated")
    ("detectedissue-severity" "high" "moderate" "low")
    ("udi-entry-type" "barcode" "rfid" "manual" "card" "self-reported" "electronic-transmission" "unknown")
    ("device-status" "active" "inactive" "entered-in-error")
    ("device-availability-status" "lost" "damaged" "destroyed" "available")
    ("device-nametype" "registered-name" "user-friendly-name" "patient-reported-name")
    ("device-category" "active" "communicating" "dme" "home-use" "implantable" "in-vitro" "point-of-care" "single-use"
                   "reusable" "software")
    ("device-specification-category" "communication" "performance" "measurement" "risk-class" "electrical" "material"
                                 "exchange")
    ("device-operation-mode" "normal" "demo" "service" "maintenance" "test")
    ("deviceusage-status" "active" "completed" "not-done" "entered-in-error" "intended" "stopped" "on-hold")
    ("deviceusage-adherence-code" "always" "never" "sometimes")
    ("deviceusage-adherence-reason" "lost" "stolen" "prescribed" "broken" "burned" "forgot")
    ("sequence-type" "aa" "dna" "rna")
    ("orientation-type" "sense" "antisense")
    ("strand-type" "watson" "crick")
    ("diagnostic-report-status" "registered" "partial" "final" "amended" "cancelled" "entered-in-error" "unknown")
    ("citation-summary-style" "vancouver" "ama11" "apa7" "apa6" "asa6" "mla8" "cochrane" "elsevier-harvard" "nature" "acs"
                          "chicago-a-17" "chicago-b-17" "ieee" "comppub")
    ("citation-classification-type" "citation-source" "medline-owner" "fevir-platform-use")
    ("citation-status-type" "pubmed-pubstatus-received" "pubmed-pubstatus-accepted" "pubmed-pubstatus-epublish"
                        "pubmed-pubstatus-ppublish" "pubmed-pubstatus-revised" "pubmed-pubstatus-aheadofprint"
                        "pubmed-pubstatus-retracted" "pubmed-pubstatus-ecollection" "pubmed-pubstatus-pmc"
                        "pubmed-pubstatus-pmcr" "pubmed-pubstatus-pubmed" "pubmed-pubstatus-pubmedr"
                        "pubmed-pubstatus-premedline" "pubmed-pubstatus-medline" "pubmed-pubstatus-medliner"
                        "pubmed-pubstatus-entrez" "pubmed-pubstatus-pmc-release" "medline-completed"
                        "medline-in-process" "medline-pubmed-not-medline" "medline-in-data-review" "medline-publisher"
                        "medline-medline" "medline-oldmedline" "pubmed-publication-status-ppublish"
                        "pubmed-publication-status-epublish" "pubmed-publication-status-aheadofprint")
    ("cited-artifact-status-type" "created" "submitted" "withdrawn" "pre-review" "under-review" "post-review-pre-published"
                              "rejected" "published-early-form" "published-final-form" "accepted" "archived"
                              "retracted" "draft" "active" "approved" "unknown")
    ("title-type" "primary" "official" "scientific" "plain-language" "subtitle" "short-title" "acronym" "earlier-title"
              "language" "autotranslated" "human-use" "machine-use" "duplicate-uid")
    ("cited-artifact-abstract-type" "primary-human-use" "primary-machine-use" "truncated" "short-abstract" "long-abstract"
                                "plain-language" "different-publisher" "language" "autotranslated" "duplicate-pmid"
                                "earlier-abstract")
    ("cited-artifact-part-type" "pages" "sections" "paragraphs" "lines" "tables" "figures" "supplement" "supplement-subpart"
                            "article-set")
    ("published-in-type" "D020492" "D019991" "D001877" "D064886")
    ("cited-medium" "internet" "print" "offline-digital-storage" "internet-without-issue" "print-without-issue"
                "offline-digital-storage-without-issue")
    ("artifact-url-classifier" "abstract" "full-text" "supplement" "webpage" "file-directory" "code-repository" "restricted"
                           "compressed-file" "doi-based" "pdf" "json" "xml" "version-specific" "computable-resource"
                           "not-specified")
    ("cited-artifact-classification-type" "publication-type" "mesh-heading" "supplemental-mesh-protocol"
                                      "supplemental-mesh-disease" "supplemental-mesh-organism" "keyword"
                                      "citation-subset" "chemical" "publishing-model" "knowledge-artifact-type"
                                      "coverage")
    ("artifact-contribution-type" "conceptualization" "data-curation" "formal-analysis" "funding-acquisition"
                              "investigation" "methodology" "project-administration" "resources" "software"
                              "supervision" "validation" "visualization" "writing-original-draft"
                              "writing-review-editing")
    ("contributor-role" "publisher" "author" "reviewer" "endorser" "editor" "informant" "funder")
    ("artifact-contribution-instance-type" "reviewed" "approved" "edited")
    ("contributor-summary-type" "author-string" "contributorship-list" "contributorship-statement" "acknowledgement-list"
                            "acknowledgment-statement" "funding-statement" "competing-interests-statement")
    ("contributor-summary-style" "a1full" "a1init" "a3full" "a3init" "a6full" "a6init" "aallfull" "aallfullwithand"
                             "aallfullwithampersand" "aallinit" "aallinitwithand" "aallinitwithampersand"
                             "contr-full-by-person" "contr-init-by-person" "contr-full-by-contr" "contr-init-by-contr")
    ("contributor-summary-source" "publisher-data" "article-copy" "citation-manager" "custom")
    ("evidence-report-type" "classification" "search-results" "resources-compiled" "text-structured")
    ("focus-characteristic-code" "citation" "clinical-outcomes-observed" "population" "exposure" "comparator" "outcome"
                             "medication-exposures" "study-type")
    ("report-relation-type" "replaces" "amends" "appends" "transforms" "replacedWith" "amendedWith" "appendedWith"
                        "transformedWith")
    ("evidence-report-section" "Evidence" "Intervention-group-alone-Evidence" "Intervention-vs-Control-Evidence"
                           "Control-group-alone-Evidence" "EvidenceVariable" "EvidenceVariable-observed"
                           "EvidenceVariable-intended" "EvidenceVariable-population" "EvidenceVariable-exposure"
                           "EvidenceVariable-outcome" "Efficacy-outcomes" "Harms-outcomes" "SampleSize" "References"
                           "Assertion" "Reasons" "Certainty-of-Evidence" "Evidence-Classifier" "Warnings"
                           "Text-Summary" "SummaryOfBodyOfEvidenceFindings" "SummaryOfIndividualStudyFindings" "Header"
                           "Tables" "Table" "Row-Headers" "Column-Header" "Column-Headers")
    ("evidence-classifier-code" "COVID19Specific" "COVID19Relevant" "COVID19HumanResearch" "OriginalResearch"
                            "ResearchSynthesis" "Guideline" "ResearchProtocol" "NotResearchNotGuideline" "Treatment"
                            "PreventionAndControl" "Diagnosis" "PrognosisPrediction" "RatedAsYes" "RatedAsNo"
                            "NotAssessed" "RatedAsRCT" "RatedAsControlledTrial" "RatedAsComparativeCohort"
                            "RatedAsCaseControl" "RatedAsUncontrolledSeries" "RatedAsMixedMethods" "RatedAsOther"
                            "RiskOfBias" "NoBlinding" "AllocConcealNotStated" "EarlyTrialTermination" "NoITT"
                            "Preprint" "PreliminaryAnalysis" "BaselineImbalance" "SubgroupAnalysis")
    ("composition-status" "registered" "partial" "final" "amended" "cancelled" "entered-in-error" "deprecated" "unknown")
    ("composition-attestation-mode" "personal" "professional" "legal" "official")
    ("catalogType" "medication" "device" "protocol")
    ("document-relationship-type" "replaces" "transforms" "signs" "appends" "incorporates" "summarizes")
    ("encounter-status" "planned" "in-progress" "on-hold" "discharged" "completed" "cancelled" "discontinued"
                    "entered-in-error" "unknown")
    ("encounter-reason-use" "CC" "HC" "AD" "RV" "HM")
    ("encounter-diagnosis-use" "working" "final")
    ("encounter-location-status" "planned" "active" "reserved" "completed")
    ("history-status" "partial" "completed" "entered-in-error" "health-unknown")
    ("goal-status" "proposed" "planned" "accepted" "cancelled" "entered-in-error" "rejected")
    ("graph-compartment-use" "where" "requires")
    ("graph-compartment-rule" "identical" "matching" "different" "custom")
    ("group-type" "person" "animal" "practitioner" "device" "careteam" "healthcareservice" "location" "organization"
              "relatedperson" "specimen")
    ("group-membership-basis" "definitional" "enumerated")
    ("imagingselection-status" "available" "entered-in-error" "unknown")
    ("imagingselection-2dgraphictype" "point" "polyline" "interpolated" "circle" "ellipse")
    ("imagingselection-3dgraphictype" "point" "multipoint" "polyline" "polygon" "ellipse" "ellipsoid")
    ("imagingstudy-status" "registered" "available" "cancelled" "entered-in-error" "unknown")
    ("spdx-license" "not-open-source" "0BSD" "AAL" "Abstyles" "Adobe-2006" "Adobe-Glyph" "ADSL" "AFL-1.1" "AFL-1.2"
                "AFL-2.0" "AFL-2.1" "AFL-3.0" "Afmparse" "AGPL-1.0-only" "AGPL-1.0-or-later" "AGPL-3.0-only"
                "AGPL-3.0-or-later" "Aladdin" "AMDPLPA" "AML" "AMPAS" "ANTLR-PD" "Apache-1.0" "Apache-1.1" "Apache-2.0"
                "APAFML" "APL-1.0" "APSL-1.0" "APSL-1.1" "APSL-1.2" "APSL-2.0" "Artistic-1.0-cl8" "Artistic-1.0-Perl"
                "Artistic-1.0" "Artistic-2.0" "Bahyph" "Barr" "Beerware" "BitTorrent-1.0" "BitTorrent-1.1" "Borceux"
                "BSD-1-Clause" "BSD-2-Clause-FreeBSD" "BSD-2-Clause-NetBSD" "BSD-2-Clause-Patent" "BSD-2-Clause"
                "BSD-3-Clause-Attribution" "BSD-3-Clause-Clear" "BSD-3-Clause-LBNL"
                "BSD-3-Clause-No-Nuclear-License-2014" "BSD-3-Clause-No-Nuclear-License"
                "BSD-3-Clause-No-Nuclear-Warranty" "BSD-3-Clause" "BSD-4-Clause-UC" "BSD-4-Clause" "BSD-Protection"
                "BSD-Source-Code" "BSL-1.0" "bzip2-1.0.5" "bzip2-1.0.6" "Caldera" "CATOSL-1.1" "CC-BY-1.0" "CC-BY-2.0"
                "CC-BY-2.5" "CC-BY-3.0" "CC-BY-4.0" "CC-BY-NC-1.0" "CC-BY-NC-2.0" "CC-BY-NC-2.5" "CC-BY-NC-3.0"
                "CC-BY-NC-4.0" "CC-BY-NC-ND-1.0" "CC-BY-NC-ND-2.0" "CC-BY-NC-ND-2.5" "CC-BY-NC-ND-3.0"
                "CC-BY-NC-ND-4.0" "CC-BY-NC-SA-1.0" "CC-BY-NC-SA-2.0" "CC-BY-NC-SA-2.5" "CC-BY-NC-SA-3.0"
                "CC-BY-NC-SA-4.0" "CC-BY-ND-1.0" "CC-BY-ND-2.0" "CC-BY-ND-2.5" "CC-BY-ND-3.0" "CC-BY-ND-4.0"
                "CC-BY-SA-1.0" "CC-BY-SA-2.0" "CC-BY-SA-2.5" "CC-BY-SA-3.0" "CC-BY-SA-4.0" "CC0-1.0" "CDDL-1.0"
                "CDDL-1.1" "CDLA-Permissive-1.0" "CDLA-Sharing-1.0" "CECILL-1.0" "CECILL-1.1" "CECILL-2.0" "CECILL-2.1"
                "CECILL-B" "CECILL-C" "ClArtistic" "CNRI-Jython" "CNRI-Python-GPL-Compatible" "CNRI-Python"
                "Condor-1.1" "CPAL-1.0" "CPL-1.0" "CPOL-1.02" "Crossword" "CrystalStacker" "CUA-OPL-1.0" "Cube" "curl"
                "D-FSL-1.0" "diffmark" "DOC" "Dotseqn" "DSDP" "dvipdfm" "ECL-1.0" "ECL-2.0" "EFL-1.0" "EFL-2.0"
                "eGenix" "Entessa" "EPL-1.0" "EPL-2.0" "ErlPL-1.1" "EUDatagrid" "EUPL-1.0" "EUPL-1.1" "EUPL-1.2"
                "Eurosym" "Fair" "Frameworx-1.0" "FreeImage" "FSFAP" "FSFUL" "FSFULLR" "FTL" "GFDL-1.1-only"
                "GFDL-1.1-or-later" "GFDL-1.2-only" "GFDL-1.2-or-later" "GFDL-1.3-only" "GFDL-1.3-or-later" "Giftware"
                "GL2PS" "Glide" "Glulxe" "gnuplot" "GPL-1.0-only" "GPL-1.0-or-later" "GPL-2.0-only" "GPL-2.0-or-later"
                "GPL-3.0-only" "GPL-3.0-or-later" "gSOAP-1.3b" "HaskellReport" "HPND" "IBM-pibs" "ICU" "IJG"
                "ImageMagick" "iMatix" "Imlib2" "Info-ZIP" "Intel-ACPI" "Intel" "Interbase-1.0" "IPA" "IPL-1.0" "ISC"
                "JasPer-2.0" "JSON" "LAL-1.2" "LAL-1.3" "Latex2e" "Leptonica" "LGPL-2.0-only" "LGPL-2.0-or-later"
                "LGPL-2.1-only" "LGPL-2.1-or-later" "LGPL-3.0-only" "LGPL-3.0-or-later" "LGPLLR" "Libpng" "libtiff"
                "LiLiQ-P-1.1" "LiLiQ-R-1.1" "LiLiQ-Rplus-1.1" "Linux-OpenIB" "LPL-1.0" "LPL-1.02" "LPPL-1.0" "LPPL-1.1"
                "LPPL-1.2" "LPPL-1.3a" "LPPL-1.3c" "MakeIndex" "MirOS" "MIT-0" "MIT-advertising" "MIT-CMU" "MIT-enna"
                "MIT-feh" "MIT" "MITNFA" "Motosoto" "mpich2" "MPL-1.0" "MPL-1.1" "MPL-2.0-no-copyleft-exception"
                "MPL-2.0" "MS-PL" "MS-RL" "MTLL" "Multics" "Mup" "NASA-1.3" "Naumen" "NBPL-1.0" "NCSA" "Net-SNMP"
                "NetCDF" "Newsletr" "NGPL" "NLOD-1.0" "NLPL" "Nokia" "NOSL" "Noweb" "NPL-1.0" "NPL-1.1" "NPOSL-3.0"
                "NRL" "NTP" "OCCT-PL" "OCLC-2.0" "ODbL-1.0" "OFL-1.0" "OFL-1.1" "OGTSL" "OLDAP-1.1" "OLDAP-1.2"
                "OLDAP-1.3" "OLDAP-1.4" "OLDAP-2.0.1" "OLDAP-2.0" "OLDAP-2.1" "OLDAP-2.2.1" "OLDAP-2.2.2" "OLDAP-2.2"
                "OLDAP-2.3" "OLDAP-2.4" "OLDAP-2.5" "OLDAP-2.6" "OLDAP-2.7" "OLDAP-2.8" "OML" "OpenSSL" "OPL-1.0"
                "OSET-PL-2.1" "OSL-1.0" "OSL-1.1" "OSL-2.0" "OSL-2.1" "OSL-3.0" "PDDL-1.0" "PHP-3.0" "PHP-3.01"
                "Plexus" "PostgreSQL" "psfrag" "psutils" "Python-2.0" "Qhull" "QPL-1.0" "Rdisc" "RHeCos-1.1" "RPL-1.1"
                "RPL-1.5" "RPSL-1.0" "RSA-MD" "RSCPL" "Ruby" "SAX-PD" "Saxpath" "SCEA" "Sendmail" "SGI-B-1.0"
                "SGI-B-1.1" "SGI-B-2.0" "SimPL-2.0" "SISSL-1.2" "SISSL" "Sleepycat" "SMLNJ" "SMPPL" "SNIA" "Spencer-86"
                "Spencer-94" "Spencer-99" "SPL-1.0" "SugarCRM-1.1.3" "SWL" "TCL" "TCP-wrappers" "TMate" "TORQUE-1.1"
                "TOSL" "Unicode-DFS-2015" "Unicode-DFS-2016" "Unicode-TOU" "Unlicense" "UPL-1.0" "Vim" "VOSTROM"
                "VSL-1.0" "W3C-19980720" "W3C-20150513" "W3C" "Watcom-1.0" "Wsuipa" "WTFPL" "X11" "Xerox" "XFree86-1.1"
                "xinetd" "Xnet" "xpp" "XSkat" "YPL-1.0" "YPL-1.1" "Zed" "Zend-2.0" "Zimbra-1.3" "Zimbra-1.4"
                "zlib-acknowledgement" "Zlib" "ZPL-1.1" "ZPL-2.0" "ZPL-2.1")
    ("guide-page-generation" "html" "markdown" "xml" "generated")
    ("guide-parameter-code" "apply" "path-resource" "path-pages" "path-tx-cache" "expansion-parameter" "rule-broken-links"
                        "generate-xml" "generate-json" "generate-turtle" "html-template")
    ("linkage-type" "source" "alternate" "historical")
    ("list-status" "current" "retired" "entered-in-error")
    ("list-mode" "working" "snapshot" "changes")
    ("list-item-flag" "01" "02" "03" "04" "05" "06")
    ("location-status" "active" "suspended" "inactive")
    ("location-mode" "instance" "kind")
    ("location-characteristic" "wheelchair" "has-translation" "has-oxy-nitro" "has-neg-press" "has-iso-ward" "has-icu")
    ("medication-admin-status" "in-progress" "not-done" "on-hold" "completed" "entered-in-error" "stopped" "unknown")
    ("reason-medication-not-given-codes" "a" "b" "c" "d")
    ("administration-subpotent-reason" "partialdose" "coldchainbreak" "recall" "adversestorage" "expired")
    ("medicationdispense-status" "preparation" "in-progress" "cancelled" "on-hold" "completed" "entered-in-error" "stopped"
                             "declined" "unknown")
    ("medicationdispense-status-reason" "frr01" "frr02" "frr03" "frr04" "frr05" "frr06" "altchoice" "clarif" "drughigh"
                                    "hospadm" "labint" "non-avail" "preg" "saig" "sddi" "sdupther" "sintol" "surg"
                                    "washout" "outofstock" "offmarket")
    ("medicationdispense-admin-location" "inpatient" "outpatient" "community")
    ("medicationrequest-status" "active" "on-hold" "ended" "entered-in-error" "draft" "unknown")
    ("medicationrequest-intent" "proposal" "plan" "order" "option")
    ("medication-intended-performer-role" "registerednurse" "oncologynurse" "paincontrolnurse" "physician" "pharmacist")
    ("medication-dose-aid" "blisterpack" "dosette" "sachets")
    ("medication-statement-status" "recorded" "entered-in-error" "draft")
    ("medication-statement-adherence" "taking" "not-taking" "unknown")
    ("medication-status" "active" "inactive" "entered-in-error")
    ("medication-ingredientstrength" "qs" "trace")
    ("response-code" "ok" "transient-error" "fatal-error")
    ("observation-triggeredbytype" "reflex" "repeat" "re-run")
    ("observation-status" "registered" "preliminary" "final" "amended" "cancelled" "entered-in-error" "unknown")
    ("observation-referencerange-normalvalue" "negative" "absent")
    ("observation-statistics" "average" "maximum" "minimum" "count" "total-count" "median" "std-dev" "sum" "variance"
                          "20-percent" "80-percent" "4-lower" "4-upper" "4-dev" "5-1" "5-2" "5-3" "5-4" "skew"
                          "kurtosis" "regression")
    ("issue-severity" "fatal" "error" "warning" "information" "success")
    ("issue-type" "invalid" "security" "processing" "transient" "informational" "success")
    ("link-type" "replaced-by" "replaces" "refer" "seealso")
    ("device-action" "implanted" "explanted" "manipulated")
    ("servicerequest-orderdetail-parameter-code" "catheter-insertion" "body-elevation" "device-configuration"
                                             "device-settings")
    ("provenance-entity-role" "revision" "quotation" "source" "instantiates" "removal")
    ("item-type" "group" "display" "question")
    ("questionnaire-enable-operator" "exists" "=" "!=" ">" "<" ">=" "<=")
    ("questionnaire-enable-behavior" "all" "any")
    ("questionnaire-disabled-display" "hidden" "protected")
    ("questionnaire-answer-constraint" "optionsOnly" "optionsOrType" "optionsOrString")
    ("questionnaire-answers-status" "in-progress" "completed" "amended" "entered-in-error" "stopped")
    ("audit-event-action" "C" "R" "U" "D" "E")
    ("audit-event-severity" "emergency" "alert" "critical" "error" "warning" "notice" "informational" "debug")
    ("specimen-status" "available" "unavailable" "unsatisfactory" "entered-in-error")
    ("specimen-combined" "grouped" "pooled")
    ("specimen-role" "b" "c" "e" "f" "o" "p" "q" "r" "v")
    ("substance-status" "active" "inactive" "entered-in-error")
    ("filter-operator" "=" "is-a" "descendent-of" "is-not-a" "regex" "in" "not-in" "generalizes" "child-of"
                   "descendent-leaf" "exists")
    ("conceptmap-property-type" "Coding" "string" "integer" "boolean" "dateTime" "decimal" "code")
    ("conceptmap-attribute-type" "code" "Coding" "string" "boolean" "Quantity")
    ("conceptmap-unmapped-mode" "use-source-code" "fixed" "other-map")
    ("slotstatus" "busy" "free" "busy-unavailable" "busy-tentative" "entered-in-error")
    ("appointmentstatus" "proposed" "pending" "booked" "arrived" "fulfilled" "cancelled" "noshow" "entered-in-error"
                     "checked-in" "waitlist")
    ("participationstatus" "accepted" "declined" "tentative" "needs-action")
    ("week-of-month" "first" "second" "third" "fourth" "last")
    ("namingsystem-type" "codesystem" "identifier" "root")
    ("namingsystem-identifier-type" "oid" "uuid" "uri" "iri-stem" "v2csmnemonic" "other")
    ("endpoint-status" "active" "suspended" "error" "off" "entered-in-error")
    ("endpoint-environment" "prod" "staging" "dev" "test" "train")
    ("subscription-status" "requested" "active" "error" "off" "entered-in-error")
    ("subscription-payload-content" "empty" "id-only" "full-resource")
    ("subscription-notification-type" "handshake" "heartbeat" "event-notification" "query-status" "query-event")
    ("subscriptiontopic-cr-behavior" "test-passes" "test-fails")
    ("operation-kind" "operation" "query")
    ("operation-parameter-scope" "instance" "type" "system")
    ("service-mode" "in-person" "telephone" "videoconference" "chat")
    ("coverage-kind" "insurance" "self-pay" "other")
    ("fm-status" "active" "cancelled" "draft" "entered-in-error")
    ("claim-use" "claim" "preauthorization" "predetermination")
    ("datestype" "card-issued" "claim-received" "service-expected")
    ("icd-10-procedures" "123001" "123002" "123003")
    ("claim-outcome" "queued" "complete" "error" "partial")
    ("claim-decision" "denied" "approved" "partial" "pending")
    ("claim-decision-reason" "0001" "0002" "0003" "0004" "0005")
    ("explanationofbenefit-status" "active" "cancelled" "draft" "entered-in-error")
    ("eligibilityrequest-purpose" "auth-requirements" "benefits" "discovery" "validation")
    ("bundle-type" "document" "message" "transaction" "transaction-response" "batch" "batch-response" "history" "searchset"
               "collection" "subscription-notification")
    ("iana-link-relations" "about" "acl" "alternate" "amphtml" "appendix" "apple-touch-icon" "apple-touch-startup-image"
                       "archives" "author" "blocked-by" "bookmark" "canonical" "chapter" "cite-as" "collection"
                       "contents" "convertedFrom" "copyright" "create-form" "current" "describedby" "describes"
                       "disclosure" "dns-prefetch" "duplicate" "edit" "edit-form" "edit-media" "enclosure" "external"
                       "first" "glossary" "help" "hosts" "hub" "icon" "index" "intervalAfter" "intervalBefore"
                       "intervalContains" "intervalDisjoint" "intervalDuring" "intervalEquals" "intervalFinishedBy"
                       "intervalFinishes" "intervalIn" "intervalMeets" "intervalMetBy" "intervalOverlappedBy"
                       "intervalOverlaps" "intervalStartedBy" "intervalStarts" "item" "last" "latest-version" "license"
                       "linkset" "lrdd" "manifest" "mask-icon" "media-feed" "memento" "micropub" "modulepreload"
                       "monitor" "monitor-group" "next" "next-archive" "nofollow" "noopener" "noreferrer" "opener"
                       "openid2.local_id" "openid2.provider" "original" "P3Pv1" "payment" "pingback" "preconnect"
                       "predecessor-version" "prefetch" "preload" "prerender" "prev" "preview" "previous"
                       "prev-archive" "privacy-policy" "profile" "publication" "related" "restconf" "replies"
                       "ruleinput" "search" "section" "self" "service" "service-desc" "service-doc" "service-meta"
                       "sponsored" "start" "status" "stylesheet" "subsection" "successor-version" "sunset" "tag"
                       "terms-of-service" "timegate" "timemap" "type" "ugc" "up" "version-history" "via" "webmention"
                       "working-copy" "working-copy-of")
    ("search-entry-mode" "match" "include" "outcome")
    ("http-verb" "GET" "HEAD" "POST" "PUT" "DELETE" "PATCH")
    ("search-processingmode" "normal" "phonetic" "other")
    ("search-comparator" "eq" "ne" "gt" "lt" "ge" "le" "sa" "eb" "ap")
    ("search-modifier-code" "missing" "exact" "contains" "not" "text" "in" "not-in" "below" "above" "type" "identifier"
                        "of-type" "code-text" "text-advanced" "iterate")
    ("eligibilityresponse-purpose" "auth-requirements" "benefits" "discovery" "validation")
    ("eligibility-outcome" "queued" "complete" "error" "partial")
    ("enrollment-outcome" "queued" "complete" "error" "partial")
    ("payment-kind" "deposit" "periodic-payment" "online" "kiosk")
    ("payment-issuertype" "patient" "insurance")
    ("payment-outcome" "queued" "complete" "error" "partial")
    ("metric-operational-status" "on" "off" "standby" "entered-in-error")
    ("metric-category" "measurement" "setting" "calculation" "unspecified")
    ("metric-calibration-type" "unspecified" "offset" "gain" "two-point")
    ("metric-calibration-state" "not-calibrated" "calibration-required" "calibrated" "unspecified")
    ("identity-assuranceLevel" "level1" "level2" "level3" "level4")
    ("vision-eye-codes" "right" "left")
    ("vision-base-codes" "up" "down" "in" "out")
    ("episode-of-care-status" "planned" "waitlist" "active" "onhold" "finished" "cancelled" "entered-in-error")
    ("structure-definition-kind" "primitive-type" "complex-type" "resource" "logical")
    ("extension-context-type" "fhirpath" "element" "extension")
    ("type-derivation-rule" "specialization" "constraint")
    ("map-model-mode" "source" "queried" "target" "produced")
    ("map-group-type-mode" "types" "type-and-types")
    ("map-input-mode" "source" "target")
    ("map-source-list-mode" "first" "not_first" "last" "not_last" "only_one")
    ("map-target-list-mode" "first" "share" "last" "single")
    ("map-transform" "create" "copy" "truncate" "escape" "cast" "append" "translate" "reference" "dateOp" "uuid" "pointer"
                 "evaluate" "cc" "c" "qty" "id" "cp")
    ("supplyrequest-status" "draft" "active" "suspended" "cancelled" "completed" "entered-in-error" "unknown")
    ("supplydelivery-status" "in-progress" "completed" "abandoned" "entered-in-error")
    ("supplydelivery-supplyitemtype" "medication" "device" "biologicallyderivedproduct")
    ("testscript-scope-conformance-codes" "required" "optional" "strict")
    ("testscript-scope-phase-codes" "unit" "integration" "production")
    ("http-operations" "delete" "get" "options" "patch" "post" "put" "head")
    ("assert-direction-codes" "response" "request")
    ("assert-manual-completion-codes" "fail" "pass" "skip" "stop")
    ("assert-operator-codes" "equals" "notEquals" "in" "notIn" "greaterThan" "lessThan" "empty" "notEmpty" "contains"
                         "notContains" "eval" "manualEval")
    ("assert-response-code-types" "continue" "switchingProtocols" "okay" "created" "accepted" "nonAuthoritativeInformation"
                              "noContent" "resetContent" "partialContent" "multipleChoices" "movedPermanently" "found"
                              "seeOther" "notModified" "useProxy" "temporaryRedirect" "permanentRedirect" "badRequest"
                              "unauthorized" "paymentRequired" "forbidden" "notFound" "methodNotAllowed"
                              "notAcceptable" "proxyAuthenticationRequired" "requestTimeout" "conflict" "gone"
                              "lengthRequired" "preconditionFailed" "contentTooLarge" "uriTooLong"
                              "unsupportedMediaType" "rangeNotSatisfiable" "expectationFailed" "misdirectedRequest"
                              "unprocessableContent" "upgradeRequired" "internalServerError" "notImplemented"
                              "badGateway" "serviceUnavailable" "gatewayTimeout" "httpVersionNotSupported")
    ("report-status-codes" "completed" "in-progress" "waiting" "stopped" "entered-in-error")
    ("report-result-codes" "pass" "fail" "pending")
    ("report-participant-type" "test-engine" "client" "server")
    ("report-action-result-codes" "pass" "skip" "fail" "warning" "error")
    ("account-status" "active" "inactive" "entered-in-error" "on-hold" "unknown")
    ("account-billing-status" "open" "carecomplete-notbilled" "billing" "closed-baddebt" "closed-voided" "closed-completed"
                          "closed-combined")
    ("account-relationship" "parent" "guarantor")
    ("account-aggregate" "patient" "insurance" "total")
    ("account-balance-term" "current" "30" "60" "90" "120")
    ("condition-precondition-type" "sensitive" "specific")
    ("condition-questionnaire-purpose" "preadmit" "diff-diagnosis" "outcome")
    ("contract-status" "amended" "appended" "cancelled" "disputed" "entered-in-error" "executable" "executed" "negotiable"
                   "offered" "policy" "rejected" "renewed" "revoked" "resolved" "terminated")
    ("contract-legalstate" "amended" "appended" "cancelled" "disputed" "entered-in-error" "executable" "executed"
                       "negotiable" "offered" "policy" "rejected" "renewed" "revoked" "resolved" "terminated")
    ("contract-expiration-type" "breach")
    ("contract-scope" "policy")
    ("contract-definition-type" "temp")
    ("contract-definition-subtype" "temp")
    ("contract-publicationstatus" "amended" "appended" "cancelled" "disputed" "entered-in-error" "executable" "executed"
                              "negotiable" "offered" "policy" "rejected" "renewed" "revoked" "resolved" "terminated")
    ("contract-security-classification" "policy")
    ("contract-security-category" "policy")
    ("contract-security-control" "policy")
    ("contract-party-role" "flunky")
    ("contract-decision-mode" "policy")
    ("contract-assetscope" "thing")
    ("contract-assettype" "participation")
    ("contract-assetsubtype" "participation")
    ("contract-assetcontext" "custodian")
    ("asset-availability" "lease")
    ("contract-actionstatus" "complete")
    ("consent-state-codes" "draft" "active" "inactive" "not-done" "entered-in-error" "unknown")
    ("consent-provision-type" "deny" "permit")
    ("consent-data-meaning" "instance" "related" "dependents" "authoredby")
    ("measure-definition-example" "screening" "standardized-depression-screening-tool")
    ("measure-group-example" "primary-rate" "secondary-rate")
    ("measure-aggregate-method" "sum" "average" "median" "minimum" "maximum" "count")
    ("measure-stratifier-example" "age" "gender" "region")
    ("measure-supplemental-data-example" "age" "gender" "ethnicity" "payer")
    ("measure-report-status" "complete" "pending" "error")
    ("measure-report-type" "individual" "subject-list" "summary" "data-exchange")
    ("submit-data-update-type" "incremental" "snapshot")
    ("measurereport-stratifier-value-example" "northwest" "northeast" "southwest" "southeast")
    ("codesystem-hierarchy-meaning" "grouped-by" "is-a" "part-of" "classified-with")
    ("codesystem-content-mode" "not-present" "example" "fragment" "complete" "supplement")
    ("concept-property-type" "code" "Coding" "string" "integer" "boolean" "dateTime" "decimal")
    ("concept-subsumption-outcome" "equivalent" "subsumes" "subsumed-by" "not-subsumed")
    ("compartment-type" "Patient" "Encounter" "RelatedPerson" "Practitioner" "Device" "EpisodeOfCare")
    ("task-status" "draft" "requested" "received" "accepted" "rejected" "ready" "cancelled" "in-progress" "on-hold" "failed"
               "completed" "entered-in-error")
    ("task-status-reason" "missing" "misidentified" "equipment-issue" "environmental-issue" "personnel-issue")
    ("task-intent" "unknown")
    ("task-code" "approve" "fulfill" "instantiate" "abort" "replace" "change" "suspend" "resume")
    ("action-participant-type" "careteam" "device" "group" "healthcareservice" "location" "organization" "patient"
                           "practitioner" "practitionerrole" "relatedperson")
    ("action-code" "send-message" "collect-information" "prescribe-medication" "recommend-immunization" "order-service"
               "propose-diagnosis" "record-detected-issue" "record-inference" "report-flag")
    ("action-reason-code" "off-pathway" "risk-assessment" "care-gap" "drug-drug-interaction" "quality-measure")
    ("action-condition-kind" "applicability" "start" "stop")
    ("action-relationship-type" "before" "concurrent" "after")
    ("action-participant-function" "performer" "author" "reviewer" "witness")
    ("action-grouping-behavior" "visual-group" "logical-group" "sentence-group")
    ("action-selection-behavior" "any" "all" "all-or-none" "exactly-one" "at-most-one" "one-or-more")
    ("action-required-behavior" "must" "could" "must-unless-documented")
    ("action-precheck-behavior" "yes" "no")
    ("action-cardinality-behavior" "single" "multiple")
    ("guidance-module-code" "bmi-calculator" "mme-calculator" "opioid-cds" "anc-cds" "chf-pathway" "covid-19-severity")
    ("guidance-response-status" "success" "data-requested" "data-required" "in-progress" "failure" "entered-in-error")
    ("research-study-prim-purp-type" "treatment" "prevention" "diagnostic" "supportive-care" "screening"
                                 "health-services-research" "basic-science" "device-feasibility")
    ("research-study-phase" "n-a" "early-phase-1" "phase-1" "phase-1-phase-2" "phase-2" "phase-2-phase-3" "phase-3"
                        "phase-4")
    ("research-study-focus-type" "medication" "device" "intervention" "factor")
    ("research-study-classifiers" "fda-regulated-drug" "fda-regulated-device" "mpg-paragraph-23b" "irb-exempt")
    ("research-study-party-role" "sponsor" "lead-sponsor" "sponsor-investigator" "primary-investigator" "collaborator"
                             "funding-source" "general-contact" "recruitment-contact" "sub-investigator"
                             "study-director" "study-chair" "irb")
    ("research-study-party-organization-type" "nih" "fda" "government" "nonprofit" "academic" "industry")
    ("research-study-status" "overall-study" "active" "active-but-not-recruiting" "administratively-completed" "approved"
                         "closed-to-accrual" "closed-to-accrual-and-intervention" "completed" "disapproved"
                         "enrolling-by-invitation" "in-review" "not-yet-recruiting" "recruiting"
                         "temporarily-closed-to-accrual" "temporarily-closed-to-accrual-and-intervention" "terminated"
                         "withdrawn")
    ("research-study-reason-stopped" "accrual-goal-met" "closed-due-to-toxicity" "closed-due-to-lack-of-study-progress"
                                 "temporarily-closed-per-study-design")
    ("research-study-arm-type" "active-comparator" "placebo-comparator" "sham-comparator" "no-intervention" "experimental"
                           "other-arm-type")
    ("research-study-objective-type" "primary" "secondary" "exploratory")
    ("message-significance-category" "consequence" "currency" "notification")
    ("messageheader-response-request" "always" "on-error" "never" "on-success")
    ("adverse-event-actuality" "actual" "potential")
    ("chargeitem-status" "planned" "billable" "not-billable" "aborted" "billed" "entered-in-error" "unknown")
    ("specimen-contained-preference" "preferred" "alternate")
    ("permitted-data-type" "Quantity" "CodeableConcept" "string" "boolean" "integer" "Range" "Ratio" "SampledData" "time"
                       "dateTime" "Period")
    ("observation-range-category" "reference" "critical" "absolute")
    ("examplescenario-actor-type" "person" "system")
    ("code-search-support" "in-compose" "in-expansion" "in-compose-or-expansion")
    ("invoice-status" "draft" "issued" "balanced" "cancelled" "entered-in-error")
    ("organization-role" "provider" "agency" "research" "payer" "diagnostics" "supplier" "HIE/HIO" "member")
    ("verificationresult-status" "attested" "validated" "in-process" "req-revalid" "val-fail" "reval-fail"
                             "entered-in-error")
    ("medicinal-product-type" "MedicinalProduct" "InvestigationalProduct")
    ("medicinal-product-domain" "Human" "Veterinary" "HumanAndVeterinary")
    ("combined-dose-form" "100000073366" "100000073651" "100000073774" "100000073781" "100000073801" "100000073860"
                      "100000073868" "100000073869" "100000073884" "100000073891" "100000073892" "100000073941"
                      "100000073972" "100000073973" "100000073974" "100000073975" "100000073987" "100000073988"
                      "100000073989" "100000073990" "100000073999" "100000074015" "100000074016" "100000074017"
                      "100000074018" "100000074030" "100000074031" "100000074032" "100000074048" "100000074051"
                      "100000074053" "100000074056" "100000074057" "100000074061" "100000074064" "100000075580"
                      "100000075584" "100000075587" "100000116137" "100000116141" "100000116155" "100000116160"
                      "100000116172" "100000116173" "100000116174" "100000116175" "100000116176" "100000116177"
                      "100000116179" "100000125746" "100000125747" "100000125777" "100000136318" "100000136325"
                      "100000136558" "100000136560" "100000136907" "100000143502" "100000143546" "100000143552"
                      "100000156068" "100000157796" "100000164467" "100000169997" "100000170588" "100000171127"
                      "100000171193" "100000171238" "100000171935" "100000174065" "200000002161" "200000002287"
                      "200000004201" "200000004819" "200000004820" "200000005547" "200000010382")
    ("legal-status-of-supply" "100000072076" "100000072077" "100000072078" "100000072079" "100000072084" "100000072085"
                          "100000072086" "100000157313")
    ("medicinal-product-additional-monitoring" "BlackTriangleMonitoring")
    ("medicinal-product-special-measures" "Post-authorizationStudies")
    ("medicinal-product-pediatric-use" "InUtero" "PretermNewborn" "TermNewborn" "Infants" "Children" "Adolescents" "Adults"
                                   "Elderly" "Neonate" "PediatricPopulation" "All" "Prepubertal" "AdultsAndElderly"
                                   "PubertalAndPostpubertal")
    ("medicinal-product-package-type" "100000073490" "100000073491" "100000073492" "100000073493" "100000073494"
                                  "100000073495" "100000073496" "100000073497" "100000073498" "100000073547"
                                  "100000073563" "100000143555")
    ("medicinal-product-contact-type" "ProposedMAH" "ProcedureContactDuring" "ProcedureContactAfter" "QPPV" "PVEnquiries")
    ("medicinal-product-name-type" "BAN" "INN" "INNM" "pINN" "rINN")
    ("medicinal-product-name-part-type" "FullName" "InventedNamePart" "ScientificNamePart" "StrengthPart" "DoseFormPart"
                                    "FormulationPart" "IntendedUsePart" "PopulationPart" "ContainerPart" "DevicePart"
                                    "TrademarkOrCompanyPart" "TimeOrPeriodPart" "FlavorPart" "DelimiterPart"
                                    "LegacyNamePart" "SpeciesNamePart")
    ("medicinal-product-cross-reference-type" "InvestigationalProduct" "VirtualProduct" "ActualProduct" "BrandedProduct"
                                          "GenericProduct" "Parallel")
    ("medicinal-product-confidentiality" "CommerciallySensitive" "NotCommerciallySensitive")
    ("package-type" "MedicinalProductPack" "RawMaterialPackage" "Shipping-TransportContainer")
    ("packaging-type" "100000073490" "100000073491" "100000073492" "100000073493" "100000073494" "100000073495"
                  "100000073496" "100000073497" "100000073498" "100000073499" "100000073500" "100000073501"
                  "100000073502" "100000073503" "100000073504" "100000073505" "100000073506" "100000073507"
                  "100000073508" "100000073509" "100000073510" "100000073511" "100000073512" "100000073513"
                  "100000073514" "100000073515" "100000073516" "100000073517" "100000073518" "100000073519"
                  "100000073520" "100000073521" "100000073522" "100000073523" "100000073524" "100000073525"
                  "100000073526" "100000073527" "100000073528" "100000073529" "100000073530" "100000073531"
                  "100000073532" "100000073533" "100000073534" "100000073535" "100000073536" "100000073537"
                  "100000073538" "100000073539" "100000073540" "100000073541" "100000073542" "100000073543"
                  "100000073544" "100000073545" "100000073546" "100000073547" "100000073548" "100000073549"
                  "100000073550" "100000073551" "100000073552" "100000073553" "100000073554" "100000073555"
                  "100000073556" "100000073557" "100000073558" "100000073559" "100000073560" "100000073561"
                  "100000073562" "100000073563" "100000075664" "100000116195" "100000116196" "100000116197"
                  "100000125779" "100000137702" "100000137703" "100000143554" "100000143555" "100000163233"
                  "100000163234" "100000164143" "100000166980" "100000169899" "100000173982" "100000173983"
                  "100000174066" "100000174067" "100000174068" "100000174069" "100000174070" "200000005068"
                  "200000005585" "200000010647" "200000011726" "200000012539" "200000013191" "200000024874")
    ("package-material" "200000003200" "200000003201" "200000003202" "200000003203" "200000003204" "200000003205"
                    "200000003206" "200000003207" "200000003208" "200000003209" "200000003210" "200000003211"
                    "200000003212" "200000003213" "200000003214" "200000003215" "200000003216" "200000003217"
                    "200000003218" "200000003219" "200000003220" "200000003221" "200000003222" "200000003223"
                    "200000003224" "200000003225" "200000003226" "200000003227" "200000003228" "200000003229"
                    "200000003529" "200000012514" "200000012515" "200000012521" "200000012522" "200000012523"
                    "200000012524" "200000012538" "200000015521" "200000023330" "200000023332" "200000025255"
                    "200000025257")
    ("manufactured-dose-form" "100000073362" "100000073363" "100000073364" "100000073365" "100000073367" "100000073368"
                          "100000073369" "100000073370" "100000073371" "100000073372" "100000073373" "100000073374"
                          "100000073375" "100000073376" "100000073377" "100000073378" "100000073379" "100000073380"
                          "100000073642" "100000073643" "100000073644" "100000073645" "100000073646" "100000073647"
                          "100000073648" "100000073649" "100000073650" "100000073652" "100000073653" "100000073654"
                          "100000073655" "100000073656" "100000073657" "100000073658" "100000073659" "100000073660"
                          "100000073661" "100000073662" "100000073663" "100000073664" "100000073665" "100000073666"
                          "100000073667" "100000073668" "100000073669" "100000073670" "100000073671" "100000073672"
                          "100000073673" "100000073674" "100000073675" "100000073676" "100000073677" "100000073678"
                          "100000073679" "100000073680" "100000073681" "100000073682" "100000073683" "100000073684"
                          "100000073685" "100000073686" "100000073687" "100000073688" "100000073689" "100000073690"
                          "100000073691" "100000073692" "100000073693" "100000073694" "100000073695" "100000073696"
                          "100000073697" "100000073698" "100000073699" "100000073700" "100000073701" "100000073702"
                          "100000073703" "100000073704" "100000073705" "100000073706" "100000073707" "100000073708"
                          "100000073709" "100000073710" "100000073711" "100000073712" "100000073713" "100000073714"
                          "100000073715" "100000073716" "100000073717" "100000073718" "100000073719" "100000073720"
                          "100000073721" "100000073722" "100000073723" "100000073724" "100000073725" "100000073726"
                          "100000073727" "100000073728" "100000073729" "100000073730" "100000073731" "100000073732"
                          "100000073733" "100000073734" "100000073735" "100000073736" "100000073737" "100000073738"
                          "100000073739" "100000073740" "100000073741" "100000073742" "100000073743" "100000073744"
                          "100000073745" "100000073746" "100000073747" "100000073748" "100000073749" "100000073750"
                          "100000073751" "100000073752" "100000073753" "100000073754" "100000073755" "100000073756"
                          "100000073757" "100000073758" "100000073759" "100000073760" "100000073761" "100000073762"
                          "100000073763" "100000073764" "100000073765" "100000073766" "100000073767" "100000073768"
                          "100000073769" "100000073770" "100000073771" "100000073772" "100000073773" "100000073775"
                          "100000073776" "100000073777" "100000073778" "100000073779" "100000073780" "100000073782"
                          "100000073783" "100000073784" "100000073785" "100000073786" "100000073787" "100000073788"
                          "100000073789" "100000073790" "100000073791" "100000073792" "100000073793" "100000073794"
                          "100000073795" "100000073796" "100000073797" "100000073798" "100000073799" "100000073800"
                          "100000073802" "100000073803" "100000073804" "100000073805" "100000073806" "100000073807"
                          "100000073808" "100000073809" "100000073810" "100000073811" "100000073812" "100000073813"
                          "100000073814" "100000073815" "100000073816" "100000073817" "100000073818" "100000073819"
                          "100000073820" "100000073821" "100000073822" "100000073823" "100000073824" "100000073825"
                          "100000073826" "100000073827" "100000073863")
    ("administrable-dose-form" "100000073362" "100000073363" "100000073364" "100000073365" "100000073367" "100000073368"
                           "100000073369" "100000073370" "100000073371" "100000073372" "100000073373" "100000073374"
                           "100000073375" "100000073376" "100000073377" "100000073378" "100000073379" "100000073380"
                           "100000073642" "100000073643" "100000073644" "100000073645" "100000073646" "100000073647"
                           "100000073648" "100000073649" "100000073650" "100000073652" "100000073653" "100000073654"
                           "100000073655" "100000073656" "100000073657" "100000073658" "100000073659" "100000073660"
                           "100000073661" "100000073662" "100000073663" "100000073664" "100000073665" "100000073666"
                           "100000073667" "100000073668" "100000073669" "100000073670" "100000073671" "100000073672"
                           "100000073673" "100000073674" "100000073675" "100000073676" "100000073677" "100000073678"
                           "100000073679" "100000073680" "100000073681" "100000073682" "100000073683" "100000073684"
                           "100000073685" "100000073686" "100000073687" "100000073688" "100000073689" "100000073690"
                           "100000073691" "100000073692" "100000073693" "100000073694" "100000073695" "100000073696"
                           "100000073697" "100000073698" "100000073699" "100000073700" "100000073701" "100000073702"
                           "100000073703" "100000073704" "100000073705" "100000073706" "100000073707" "100000073708"
                           "100000073709" "100000073710" "100000073711" "100000073712" "100000073713" "100000073714"
                           "100000073715" "100000073716" "100000073717" "100000073718" "100000073719" "100000073720"
                           "100000073721" "100000073722" "100000073723" "100000073724" "100000073725" "100000073726"
                           "100000073727" "100000073728" "100000073729" "100000073730" "100000073731" "100000073732"
                           "100000073733" "100000073734" "100000073735" "100000073736" "100000073737" "100000073738"
                           "100000073739" "100000073740" "100000073741" "100000073742" "100000073743" "100000073744"
                           "100000073745" "100000073746" "100000073747" "100000073748" "100000073749" "100000073750"
                           "100000073751" "100000073752" "100000073753" "100000073754" "100000073755" "100000073756"
                           "100000073757" "100000073758" "100000073759" "100000073760" "100000073761" "100000073762"
                           "100000073763" "100000073764" "100000073765" "100000073766" "100000073767" "100000073768"
                           "100000073769" "100000073770" "100000073771" "100000073772" "100000073773" "100000073775"
                           "100000073776" "100000073777" "100000073778" "100000073779" "100000073780" "100000073782"
                           "100000073783" "100000073784" "100000073785" "100000073786" "100000073787" "100000073788"
                           "100000073789" "100000073790" "100000073791" "100000073792" "100000073793" "100000073794"
                           "100000073795" "100000073796" "100000073797" "100000073798" "100000073799" "100000073800"
                           "100000073802" "100000073803" "100000073804" "100000073805" "100000073806" "100000073807"
                           "100000073808" "100000073809" "100000073810" "100000073811" "100000073812" "100000073813"
                           "100000073814" "100000073815" "100000073816" "100000073817" "100000073818" "100000073819"
                           "100000073820" "100000073821" "100000073822" "100000073823" "100000073824" "100000073825"
                           "100000073826" "100000073827" "100000073863")
    ("unit-of-presentation" "200000002108" "200000002109" "200000002110" "200000002111" "200000002112" "200000002113"
                        "200000002114" "200000002115" "200000002116" "200000002117" "200000002118" "200000002119"
                        "200000002120" "200000002121" "200000002122" "200000002123" "200000002124" "200000002125"
                        "200000002126" "200000002127" "200000002128" "200000002129" "200000002130" "200000002131"
                        "200000002132" "200000002133" "200000002134" "200000002135" "200000002136" "200000002137"
                        "200000002138" "200000002139" "200000002140" "200000002141" "200000002142" "200000002143"
                        "200000002144" "200000002145" "200000002146" "200000002147" "200000002148" "200000002149"
                        "200000002150" "200000002151" "200000002152" "200000002153" "200000002154" "200000002155"
                        "200000002156" "200000002157" "200000002158" "200000002159" "200000002163" "200000002164"
                        "200000002165" "200000002166")
    ("target-species" "100000108874" "100000108875" "100000108876" "100000108877" "100000108878" "100000108879"
                  "100000108880" "100000108881" "100000108882" "100000108883" "100000108884" "100000108885"
                  "100000108886" "100000108887" "100000108888" "100000108889" "100000108890" "100000108891"
                  "100000108892" "100000108893" "100000108894" "100000108895" "100000108896" "100000108897"
                  "100000108898" "100000108899" "100000108900" "100000108901" "100000108902" "100000108903"
                  "100000108904" "100000108905" "100000108906" "100000108907" "100000108908" "100000108909"
                  "100000108910" "100000108911" "100000108912" "100000108913" "100000108914" "100000108915"
                  "100000108916" "100000108917" "100000108918" "100000108919" "100000108920" "100000108921"
                  "100000108922" "100000108923" "100000108924" "100000108925" "100000108926" "100000108927"
                  "100000108928" "100000108929" "100000108930" "100000108931" "100000108932" "100000108933"
                  "100000108934" "100000108935" "100000108936" "100000108937" "100000108938" "100000108939"
                  "100000108940" "100000108941" "100000108942" "100000108943" "100000108944" "100000108945"
                  "100000108946" "100000108947" "100000108948" "100000108949" "100000108950" "100000108951"
                  "100000108952" "100000108953" "100000108954" "100000108955" "100000108956" "100000108957"
                  "100000108958" "100000108959" "100000108960" "100000108961" "100000108962" "100000108963"
                  "100000108964" "100000108965" "100000108966" "100000108967" "100000108968" "100000108969"
                  "100000108970" "100000108971" "100000108972" "100000108973" "100000108974" "100000108975"
                  "100000108976" "100000108977" "100000108978" "100000108979" "100000108980" "100000108981"
                  "100000108982" "100000108983" "100000108984" "100000108985" "100000108986" "100000108987"
                  "100000108988" "100000108989" "100000108990" "100000108991" "100000108992" "100000108993"
                  "100000108994" "100000108995" "100000108996" "100000108997" "100000108998" "100000108999"
                  "100000109000" "100000109001" "100000109002" "100000109003" "100000109004" "100000109005"
                  "100000109006" "100000109007" "100000109008" "100000109009" "100000109010" "100000109011"
                  "100000109012" "100000109013" "100000109014" "100000109015" "100000109016" "100000109017"
                  "100000109018" "100000109019" "100000109020" "100000109021" "100000109022" "100000109023"
                  "100000109024" "100000109025" "100000109026" "100000109027" "100000109028" "100000109029"
                  "100000109030" "100000109031" "100000109032" "100000109033" "100000109034" "100000109035"
                  "100000109036" "100000109037" "100000109038" "100000109039" "100000109040" "100000109041"
                  "100000109042" "100000109043" "100000109044" "100000109045" "100000109046" "100000109047"
                  "100000109048" "100000109049" "100000109050" "100000109051" "100000109052" "100000109053"
                  "100000109054" "100000109055" "100000109056" "100000109057" "100000109058" "100000109059"
                  "100000109060" "100000109061" "100000109062" "100000109063" "100000109064" "100000109065"
                  "100000109066" "100000109067" "100000109068" "100000109069" "100000109070" "100000109071"
                  "100000109072" "100000109073")
    ("animal-tissue-type" "100000072091" "100000072092" "100000072093" "100000072094" "100000072095" "100000072096"
                      "100000072104" "100000072105" "100000072106" "100000072107" "100000072108" "100000072109"
                      "100000111053" "100000111054" "100000111055" "100000111056" "100000111057" "100000111058"
                      "100000111059" "100000111060" "100000111061" "100000111062" "100000111063" "100000111064"
                      "100000111065" "100000111066" "100000111067" "100000111068" "100000111069" "100000111070"
                      "100000111071" "100000111072" "100000111073" "100000111074" "100000111075" "100000111076"
                      "100000111077" "100000111078" "100000111079" "100000111080" "100000111081" "100000111082"
                      "100000111083" "100000111084" "100000111085" "100000111086" "100000111087" "100000111088"
                      "100000111089" "100000111090" "100000111091" "100000111092" "100000111093" "100000111094"
                      "100000111095" "100000111096" "100000111097" "100000111098" "100000111099" "100000111100"
                      "100000111101" "100000111102" "100000111103" "100000111104" "100000111105" "100000111106"
                      "100000111107" "100000111108" "100000111109" "100000111110" "100000111111" "100000111112"
                      "100000111113" "100000111114" "100000111115" "100000111116" "100000111117" "100000111118"
                      "100000111119" "100000111120" "100000111121" "100000111122" "100000111123" "100000111124"
                      "100000111125" "100000111126" "100000111127" "100000111128" "100000111129" "100000111130"
                      "100000111131" "100000111132" "100000111133" "100000111134" "100000111135" "100000111136"
                      "100000111137" "100000111138" "100000111139" "100000111140" "100000111141" "100000111142"
                      "100000111143" "100000111144" "100000111145" "100000111146" "100000111147" "100000111148"
                      "100000111149" "100000111150" "100000111151" "100000111152" "100000111153" "100000111154"
                      "100000111155" "100000111156" "100000111157" "100000111158" "100000111159" "100000111160"
                      "100000111161" "100000111162" "100000111163" "100000111164" "100000125717" "100000136180"
                      "100000136181" "100000136182" "100000136183" "100000136184" "100000136185" "100000136186"
                      "100000136187" "100000136188" "100000136189" "100000136190" "100000136191" "100000136192"
                      "100000136193" "100000136194" "100000136195" "100000136196" "100000136197" "100000136198"
                      "100000136199" "100000136200" "100000136201" "100000136202" "100000136203" "100000136204"
                      "100000136205" "100000136206" "100000136207" "100000136208" "100000136209" "100000136210"
                      "100000136211" "100000136212" "100000136213" "100000136214" "100000136215" "100000136216"
                      "100000136217" "100000136218" "100000136219" "100000136220" "100000136221" "100000136222"
                      "100000136223" "100000136224" "100000136225" "100000136226" "100000136227" "100000136228"
                      "100000136229" "100000136230" "100000136231" "100000136232" "100000136233" "100000136234"
                      "100000136235" "100000136236" "100000136237" "100000136247" "100000136248" "100000136554"
                      "100000136555" "100000136556" "100000142485")
    ("regulated-authorization-type" "MarketingAuth" "Orphan" "Pediatric")
    ("product-intended-use" "Prevention" "Treatment" "Alleviation" "Diagnosis" "Monitoring")
    ("regulated-authorization-basis" "Full" "NewSubstance" "KnownSubstance" "SimilarBiological" "Well-establishedUse"
                                 "TraditionalUse" "Bibliographical" "KnownHumanBlood" "TemporaryUse" "ParallelTrade")
    ("regulated-authorization-case-type" "InitialMAA" "Variation" "LineExtension" "PSUR" "Renewal" "Follow-up"
                                     "100000155699" "AnnualReassessment" "UrgentSafetyRestriction"
                                     "PaediatricSubmission" "TransferMA" "LiftingSuspension" "Withdrawal"
                                     "Reformatting" "RMP" "ReviewSuspension" "SupplementalInformation" "RepeatUse"
                                     "SignalDetection" "FLU" "PANDEMIC" "Orphan")
    ("ingredient-role" "100000072072" "100000072073" "100000072082" "100000136065" "100000136066" "100000136178"
                   "100000136179" "100000136561" "200000003427")
    ("ingredient-function" "Antioxidant" "AlkalizingAgent")
    ("ingredient-manufacturer-role" "allowed" "possible" "actual")
    ("substance-grade" "USP-NF" "Ph.Eur" "JP" "BP" "CompanyStandard")
    ("substance-stereochemistry" "ConstitutionalIsomer" "Stereoisomer" "Enantiomer")
    ("substance-optical-activity" "+" "-")
    ("substance-amount-type" "Average" "Approximately" "LessThan" "MoreThan")
    ("substance-structure-technique" "X-Ray" "HPLC" "NMR" "PeptideMapping" "LigandBindingAssay")
    ("substance-form" "salt" "base")
    ("substance-weight-method" "SDS-PAGE" "Calculated" "LighScattering" "Viscosity" "GelPermeationCentrifugation"
                           "End-groupAnalysis" "End-groupTitration" "Size-ExclusionChromatography")
    ("substance-weight-type" "Exact" "Average" "WeightAverage")
    ("substance-representation-type" "Systematic" "Scientific" "Brand")
    ("substance-representation-format" "InChI" "SMILES" "MOLFILE" "CDX" "SDF" "PDB" "mmCIF")
    ("substance-name-type" "Systematic" "Scientific" "Brand")
    ("substance-name-domain" "ActiveIngredient" "FoodColorAdditive")
    ("substance-name-authority" "BAN" "COSING" "Ph.Eur." "FCC" "INCI" "INN" "JAN" "JECFA" "MARTINDALE" "USAN" "USP" "PHF"
                            "HAB" "PhF" "IUIS")
    ("substance-relationship-type" "Salt" "ActiveMoiety" "StartingMaterial" "Polymorph" "Impurity")
    ("substance-source-material-type" "Animal" "Plant" "Mineral")
    ("substance-source-material-genus" "Mycobacterium" "InfluenzavirusA" "Ginkgo")
    ("substance-source-material-species" "GinkgoBiloba" "OleaEuropaea")
    ("substance-source-material-part" "Animal" "Plant" "Mineral")
    ("product-category" "organ" "tissue" "fluid" "cells" "biologicalAgent")
    ("biologicallyderived-productcodes" "e0398" "s1128" "s1194" "s1195" "s1310" "s1398" "s2598" "e4377" "t1396")
    ("biologicallyderived-product-status" "available" "unavailable")
    ("medicationknowledge-status" "active" "entered-in-error" "inactive")
    ("medication-cost-category" "banda" "bandb")
    ("devicedefinition-regulatory-identifier-type" "basic" "master" "license")
    ("devicedefinition-relationtype" "gateway" "replaces" "previous")
    ("device-productidentifierinudi" "lot-number" "manufactured-date" "serial-number" "expiration-date" "biological-source"
                                 "software-version")
    ("device-correctiveactionscope" "model" "lot-numbers" "serial-numbers")
    ("definition-method" "systematic-assessment" "non-systematic-assessment" "mean" "median" "mean-of-mean" "mean-of-median"
                     "median-of-mean" "median-of-median")
    ("characteristic-offset" "UNL" "LNL")
    ("characteristic-combination" "all-of" "any-of" "at-least" "at-most" "statistical" "net-effect" "dataset")
    ("evidence-variable-event" "study-start" "treatment-start" "condition-detection" "condition-treatment"
                           "hospital-admission" "hospital-discharge" "operative-procedure")
    ("variable-handling" "continuous" "dichotomous" "ordinal" "polychotomous")
    ("clinical-use-definition-type" "indication" "contraindication" "interaction" "undesirable-effect" "warning")
    ("clinical-use-definition-category" "Pregnancy" "Overdose" "DriveAndMachines")
    ("therapy-relationship-type" "contraindicated-only-with" "contraindicated-except-with" "indicated-only-with"
                             "indicated-except-with" "indicated-only-after" "indicated-only-before"
                             "replace-other-therapy" "replace-other-therapy-contraindicated"
                             "replace-other-therapy-not-tolerated" "replace-other-therapy-not-effective")
    ("interaction-type" "drug-drug" "drug-food" "drug-test" "other")
    ("interaction-incidence" "Theoretical" "Observed")
    ("undesirable-effect-frequency" "Common" "Uncommon" "Rare")
    ("warning-type" "P313" "P314" "P315" "P320" "P321" "P322" "P330" "P331" "P361" "P363")
    ("study-design" "SEVCO:01001" "SEVCO:01002" "SEVCO:01010" "SEVCO:01023" "SEVCO:01022" "SEVCO:01027" "SEVCO:01028"
                "SEVCO:01045" "SEVCO:01026" "SEVCO:01049" "SEVCO:01042" "SEVCO:01051" "SEVCO:01086" "SEVCO:01087"
                "SEVCO:01060" "SEVCO:01061" "SEVCO:01062" "SEVCO:01063" "SEVCO:01064" "SEVCO:01043" "SEVCO:01052"
                "SEVCO:01053" "SEVCO:01054" "SEVCO:01085" "SEVCO:01089")
    ("statistic-model-code" "oneTailedTest" "twoTailedTest" "zTest" "oneSampleTTest" "twoSampleTTest" "pairedTTest"
                        "chiSquareTest" "chiSquareTestTrend" "pearsonCorrelation" "anova" "anovaOneWay" "anovaTwoWay"
                        "anovaTwoWayReplication" "manova" "anovaThreeWay" "signTest" "wilcoxonSignedRankTest"
                        "wilcoxonRankSumTest" "mannWhitneyUTest" "fishersExactTest" "mcnemarsTest" "kruskalWallisTest"
                        "spearmanCorrelation" "kendallCorrelation" "friedmanTest" "goodmanKruskasGamma" "glm"
                        "glmProbit" "glmLogit" "glmIdentity" "glmLog" "glmGeneralizedLogit" "glmm" "glmmProbit"
                        "glmmLogit" "glmmIdentity" "glmmLog" "glmmGeneralizedLogit" "linearRegression"
                        "logisticRegression" "polynomialRegression" "coxProportionalHazards"
                        "binomialDistributionRegression" "multinomialDistributionRegression" "poissonRegression"
                        "negativeBinomialRegression" "zeroCellConstant" "zeroCellContinuityCorrection" "adjusted"
                        "interactionTerm" "manteHaenszelMethod" "metaAnalysis" "inverseVariance" "petoMethod"
                        "hartungKnapp" "modifiedHartungKnapp" "effectsFixed" "effectsRandom" "chiSquareTestHomogeneity"
                        "dersimonianLairdMethod" "pauleMandelMethod" "restrictedLikelihood" "maximumLikelihood"
                        "empiricalBayes" "hunterSchmidt" "sidikJonkman" "hedgesMethod" "tauDersimonianLaird"
                        "tauPauleMandel" "tauRestrictedMaximumLikelihood" "tauMaximumLikelihood" "tauEmpiricalBayes"
                        "tauHunterSchmidt" "tauSidikJonkman" "tauHedges" "poolMantelHaenzsel" "poolInverseVariance"
                        "poolPeto" "poolGeneralizedLinearMixedModel")
    ("certainty-type" "Overall" "RiskOfBias" "Inconsistency" "Indirectness" "Imprecision" "PublicationBias"
                  "DoseResponseGradient" "PlausibleConfounding" "LargeEffect")
    ("certainty-rating" "high" "moderate" "low" "very-low" "no-concern" "serious-concern" "very-serious-concern"
                    "extremely-serious-concern" "present" "absent" "no-change" "downcode1" "downcode2" "downcode3"
                    "upcode1" "upcode2")
    ("nutritionproduct-status" "active" "inactive" "entered-in-error")
    ("permission-status" "active" "entered-in-error" "draft" "rejected")
    ("permission-rule-combining" "deny-overrides" "permit-overrides" "ordered-deny-overrides" "ordered-permit-overrides"
                             "deny-unless-permit" "permit-unless-deny")
    ("inventoryreport-status" "draft" "requested" "active" "entered-in-error")
    ("inventoryreport-counttype" "snapshot" "difference")
    ("devicedispense-status" "preparation" "in-progress" "cancelled" "on-hold" "completed" "entered-in-error" "stopped"
                         "declined" "unknown")
    ("devicedispense-status-reason" "out-of-stock" "off-market" "contraindication" "incompatible-device" "order-expired"
                                "verbal-order")
    ("artifactassessment-information-type" "comment" "classifier" "rating" "container" "response" "change-request")
    ("artifactassessment-workflow-status" "submitted" "triaged" "waiting-for-input" "resolved-no-change"
                                      "resolved-change-required" "deferred" "duplicate" "applied" "published"
                                      "entered-in-error")
    ("artifactassessment-disposition" "unresolved" "not-persuasive" "persuasive" "persuasive-with-modification"
                                  "not-persuasive-with-modification")
    ("transport-status" "in-progress" "completed" "abandoned" "cancelled" "planned" "entered-in-error")
    ("transport-status-reason" "declined-by-patient" "refused-by-recipient" "patient-not-available" "specimen-not-available"
                           "wrong-request-data" "in-route-accident" "request-not-acknowledged")
    ("transport-intent" "unknown")
    ("transport-code" "approve" "fulfill" "instantiate" "abort" "replace" "change" "suspend" "resume")
    ("genomicstudy-status" "registered" "available" "cancelled" "entered-in-error" "unknown")
    ("genomicstudy-type" "alt-splc" "chromatin" "cnv" "epi-alt-hist" "epi-alt-dna" "fam-var-segr" "func-var"
                     "gene-expression" "post-trans-mod" "snp" "str" "struc-var")
    ("genomicstudy-methodtype" "biochemical-genetics" "cytogenetics" "molecular-genetics" "analyte"
                           "chromosome-breakage-studies" "deletion-duplication-analysis" "detection-of-homozygosity"
                           "enzyme-assay" "fish-interphase" "fish-metaphase" "flow-cytometry" "fish"
                           "immunohistochemistry" "karyotyping" "linkage-analysis" "methylation-analysis" "msi"
                           "m-fish" "mutation-scanning-of-select-exons" "mutation-scanning-of-the-entire-coding-region"
                           "protein-analysis" "protein-expression" "rna-analysis" "sequence-analysis-of-select-exons"
                           "sequence-analysis-of-the-entire-coding-region" "sister-chromatid-exchange"
                           "targeted-variant-analysis" "udp" "aspe" "alternative-splicing-detection"
                           "bi-directional-sanger-sequence-analysis" "c-banding" "cia"
                           "chromatin-immunoprecipitation-on-chip" "comparative-genomic-hybridization" "damid"
                           "digital-virtual-karyotyping" "digital-microfluidic-microspheres" "enzymatic-levels"
                           "enzyme-activity" "elisa" "fluorometry" "fusion-genes-microarrays" "g-banding" "gc-ms"
                           "gene-expression-profiling" "gene-id" "gold-nanoparticle-probe-technology" "hplc" "lc-ms"
                           "lc-ms-ms" "metabolite-levels" "methylation-specific-pcr" "microarray" "mlpa" "ngs-mps"
                           "ola" "oligonucleotide-hybridization-based-dna-sequencing" "other" "pcr"
                           "pcr-with-allele-specific-hybridization" "pcr-rflp-with-southern-hybridization"
                           "protein-truncation-test" "pyrosequencing" "q-banding" "qpcr" "r-banding" "rflp" "rt-lamp"
                           "rt-pcr" "rt-pcr-with-gel-analysis" "rt-qpcr" "snp-detection" "silver-staining" "sky"
                           "t-banding" "ms-ms" "tetra-nucleotide-repeat-by-pcr-or-southern-blot" "tiling-arrays"
                           "trinucleotide-repeat-by-pcr-or-southern-blot" "uni-directional-sanger-sequencing")
    ("genomicstudy-changetype" "DNA" "RNA" "AA" "CHR" "CNV")
    ("genomicstudy-dataformat" "bam" "bed" "bedpe" "bedgraph" "bigbed" "bigWig" "birdsuite-files" "broadpeak" "cbs"
                           "chemical-reactivity-probing-profiles" "chrom-sizes" "cn" "custom-file-formats" "cytoband"
                           "fasta" "gct" "cram" "genepred" "gff-gtf" "gistic" "goby" "gwas" "igv" "loh"
                           "maf-multiple-alignment-format" "maf-mutation-annotation-format" "merged-bam-file" "mut"
                           "narrowpeak" "psl" "res" "rna-secondary-structure-formats" "sam"
                           "sample-info-attributes-file" "seg" "tdf" "track-line" "type-line" "vcf" "wig")
    ("formularyitem-status" "active" "entered-in-error" "inactive")
    ("biologicallyderivedproductdispense-status" "preparation" "in-progress" "allocated" "issued" "unfulfilled" "returned"
                                             "entered-in-error" "unknown")
    ("biologicallyderivedproductdispense-origin-relationship" "autologous" "related" "directed" "allogeneic" "xenogenic")
    ("biologicallyderivedproductdispense-match-status" "crossmatched" "selected" "unmatched" "least-incompatible")
    ("biologicallyderivedproductdispense-performer-function" "group-and-type" "antibody-screen" "antibody-identification"
                                                         "crossmatch" "release" "transport" "receipt")
    ("deviceassociation-status" "implanted" "explanted" "entered-in-error" "attached" "unknown")
    ("deviceassociation-status-reason" "attached" "disconnected" "failed" "placed" "replaced")
    ("deviceassociation-operationstatus" "on" "off" "standby" "defective" "unknown")
    ("inventoryitem-status" "active" "inactive" "entered-in-error" "unknown")
    ("inventoryitem-nametype" "trade-name" "alias" "original-name" "preferred")
    ("message-events")
    ("knowledge-representation-level" "narrative" "semi-structured" "structured" "executable")
    ))
