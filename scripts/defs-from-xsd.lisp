;;;;

(defconstant primitive-parents '(("string" . "cl:string")
                                 ("code" . "string")
                                 ("id" . "string")
                                 ("markdown" . "string")
                                 ("integer" . "fixnum")
                                 ("unsignedInt" . "integer")
                                 ("positiveInt" . "integer")
                                 ("integer64" . "fixnum")
                                 ("decimal" . "double-float")
                                 ("boolean" . "symbol")
                                 ("time" . "cl:string")
                                 ("date" . "cl:string")
                                 ("dateTime" . "cl:string")
                                 ("instant" . "cl:string")
                                 ("base64Binary" . "cl:string")
                                 ("uri" . "cl:string")
                                 ("url" . "uri")
                                 ("canonical" . "uri")
                                 ("oid" . "uri")
                                 ("uuid" . "uri")))

(defun fix-type (pname type)
  "TBD"
  (cond ((string= pname "id") "id")
        ((suffixp type "-primitive") (trim-suffix type "-primitive"))
        ((string= type "ResourceContainer") "Resource")
        (t type)))

(defun find-named (schema name key)
  "Search for an element of the schema with the specified name and key where the
   key should match the first car of the element."
  (dolist (element schema)
    (when (and (equal name (cdr (assoc "name" (cadr element))))
               (equal key (car element)))
      (format t "~A~%" element)
      (return element))))

(defun property-from-element (elem)
  "Form a properties list from all the seq of elements or attributes."
  (let* ((px (cadr elem))
         (name (cdr (assoc "name" px)))
         (type (fix-type name (cdr (assoc "type" px))))
         (mn (cdr (assoc "minOccurs" px)))
         (mx (cdr (assoc "maxOccurs" px)))
         (docs (cddr (caddr elem)))
         (prop (make-bag "{}")))

    (bag-set prop name "name")
    (bag-set prop type "type")
    (when docs (bag-set prop (join "\n\n" (mapcar (lambda (doc) (caddr doc)) docs)) "description"))
    (when (and (integerp mn) (/= 0 mn)) (bag-set prop t "required"))
    (when (equal (cdr (assoc "maxOccurs" px)) "unbounded") (bag-set prop t "array"))
    (when (equal "required" (cdr (assoc "use" px)))
      (bag-set prop t "required"))

    ;; TBD need choices (enum), group

    ;; (format t "*** property: ~A~%" (send prop :write nil))
    prop))

(defun primitive-docs (name schema)
  (let ((ct (find-named schema name "complexType")))
    (when ct
      (join "\n\n" (mapcar (lambda (doc) (caddr doc)) (cddr (caddr ct)))))))

(defun form-primitive-node (name element schema)
  "TBD"
  (format t "*** primitive: ~A~%" element)
  (setq name (trim-suffix name "-primitive"))
  (let* ((restriction (cddr (caddr element)))
         (pattern (when restriction (cdaadr (assoc "pattern" restriction))))
         (mn (when restriction (cdaadr (assoc "minLength" restriction))))
         (mx (when restriction (cdaadr (assoc "maxLength" restriction))))
         (pb (make-bag "{}")))
    ;; Restrictions for primitives include "pattern", "minLength",
    ;; "maxLength", and "union". The union sub-element is ignored as
    ;; identifies member types as XSD types.

    (bag-set pb name "name")
    (when pattern (bag-set pb pattern "pattern"))
    (when mn (bag-set pb mn "minLen"))
    (when mx (bag-set pb mx "maxLen"))
    ;; The FHIR XSD jumps back and forth between XSD types and FHIR types. For
    ;; primitives, the FHIR hierarchy is ignored in favor of XSD
    ;; types. Alternatively if using the matching complexType element the base
    ;; is always "Element" which also contradicts the web pages. The FHIR
    ;; types are needed for the defs JSON file so the parents are manually
    ;; set.
    (bag-set pb (cdr (assoc name primitive-parents)) "parent")

    ;; Get the description from the associated complexType. All except for
    ;; SampledDataDataType which is not on the web site but is in the XSD ad
    ;; has no documentation.
    (unless (string= "SampledDataDataType" name)
      (bag-set pb (primitive-docs name schema) "description"))

    pb))

;;; The FHIR heirarchy is defined as:
;;; Base
;;;   Element
;;;     BackboneElement - not is schema file, same as BackboneType though
;;;     DataType
;;;       <-- most datatypes go here (e.g., Coding)
;;;       PrimitiveType <-- not really here in non-XML implementations
;;;       BackboneType
;;;   Resource
;;;     DomainResource
;;;       <-- most resources go here
;;;
(defun form-hierarchy-node (name element)
  "Build a base hierarchy node from the provided schema definition."
  ;; (format t "----------------~%~A~%" element)
  (let* ((hb (make-bag "{}"))
         (cc (caddr (assoc "complexContent" (cddr element))))
         (anno (cddr (assoc "annotation" (cddr element))))
         (super (cdaadr cc))
         (seq (when cc (cddr (assoc "sequence" (cddr cc)))))
         (attr (when cc (assoc "attribute" (cddr cc))))
         properties)
    (bag-set hb name "name")
    (bag-set hb super "parent")

    (when anno (bag-set hb (join "\n\n" (mapcar (lambda (doc) (caddr doc)) anno)) "description"))

    (dolist (elem seq)
      ;; (format t "*** ~A~%" elem)
      (case (car elem)
        ("element"
         (setq properties (add properties (property-from-element elem))))
        ("choice"
         ;; TBD
         (format t "*** choice: ~A~%" elem))
        ))
    (when attr
      (setq properties (add properties (property-from-element attr))))

    (when properties (bag-set hb properties "properties"))
    hb))

(defun defs-from-xsd (input-filename output-filename)
  "TBD."
  (let ((schema (cddr (assoc "schema" (with-open-file (f input-filename :direction :input) (xml-read f)))))
        (defs (make-bag "{}"))
        primitives datatypes hierarchy backbones resources)

    (do ((name (caar schema) (caar schema)))
        ((not (or (equal :comment name) (equal "import" name))))
      (setq schema (cdr schema)))

    (dolist (element schema)
      (let ((name (cdr (assoc "name" (cadr element)))))
        (cond ((member name '("Base"
                              "Element"
                              "DataType"
                              "BackboneType"
                              "Resource"
                              "DomainResource"))
               (setq hierarchy (add hierarchy (form-hierarchy-node name element))))
              ((suffixp name "-primitive")
               (setq primitives (add primitives (form-primitive-node name element schema))))
              (t
               ;; (format t "~A ~A~%" (car element) name)
               nil
              ))))


    ;; (format t "~A~%" schema)

    ;; (dolist (element schema)
    ;;   (format t "~A ~A~%" (car element) (cdr (assoc "name" (cadr element)))))

      ;; Add each type list to the new schema being constructed.
    (send defs :set primitives "primitives")
    (send defs :set hierarchy "hierarchy")
    (send defs :set datatypes "datatypes")
    (send defs :set backbones "backbones")
    (send defs :set resources "resources")

    (with-open-file (f output-filename :direction :output :if-exists :supersede :if-does-not-exist :create)
      (send defs :write f :pretty t :json t :depth 1))
  ))

(defvar pat '("complexType" (("name" . "Patient"))
              ("annotation" ()
               ("documentation" (("lang" . "en"))
                "Demographics and other administrative information about an individual or animal receiving care or other health-related services.")
               ("documentation" (("lang" . "en")) "If the element is present, it must have either a @value, an @id, or extensions"))
              ("complexContent" ()
               ("extension" (("base" . "DomainResource"))
                ("sequence" ()
                            ("element"
                             (("name" . "identifier") ("minOccurs" . "0") ("maxOccurs" . "unbounded")
                              ("type" . "Identifier"))
                             ("annotation" ()
                                           ("documentation" (("lang" . "en")) "An identifier for this patient.")))
                            ("element" (("name" . "active") ("minOccurs" . "0") ("maxOccurs" . "1") ("type" . "boolean"))
                                       ("annotation" ()
                                                     ("documentation" (("lang" . "en"))
                                                                      "Whether this patient record is in active use.
Many systems use this property to mark as non-current patients, such as those that have not been seen for a period of time based on an organization's business rules.

It is often used to filter patient lists to exclude inactive patients

Deceased patients may also be marked as inactive for the same reasons, but may be active for some time after death.")))
                            ("element"
                             (("name" . "name") ("minOccurs" . "0") ("maxOccurs" . "unbounded") ("type" . "HumanName"))
                             ("annotation" ()
                                           ("documentation" (("lang" . "en"))
                                                            "A name associated with the individual.")))
                            ("element"
                             (("name" . "telecom") ("minOccurs" . "0") ("maxOccurs" . "unbounded")
                              ("type" . "ContactPoint"))
                             ("annotation" ()
                                           ("documentation" (("lang" . "en"))
                                                            "A contact detail (e.g. a telephone number or an email address) by which the individual may be contacted.")))
                            ("element"
                             (("name" . "gender") ("minOccurs" . "0") ("maxOccurs" . "1")
                              ("type" . "AdministrativeGender"))
                             ("annotation" ()
                                           ("documentation" (("lang" . "en"))
                                                            "Administrative Gender - the gender that the patient is considered to have for administration and record keeping purposes.")))
                            ("element" (("name" . "birthDate") ("minOccurs" . "0") ("maxOccurs" . "1") ("type" . "date"))
                                       ("annotation" ()
                                                     ("documentation" (("lang" . "en"))
                                                                      "The date of birth for the individual.")))
                            ("choice" (("minOccurs" . "0") ("maxOccurs" . "1"))
                                      ("annotation" ()
                                                    ("documentation" (("lang" . "en"))
                                                                     "Indicates if the individual is deceased or not."))
                                      ("element" (("name" . "deceasedBoolean") ("type" . "boolean")))
                                      ("element" (("name" . "deceasedDateTime") ("type" . "dateTime"))))
                            ("element"
                             (("name" . "address") ("minOccurs" . "0") ("maxOccurs" . "unbounded") ("type" . "Address"))
                             ("annotation" () ("documentation" (("lang" . "en")) "An address for the individual.")))
                            ("element"
                             (("name" . "maritalStatus") ("minOccurs" . "0") ("maxOccurs" . "1")
                              ("type" . "CodeableConcept"))
                             ("annotation" ()
                                           ("documentation" (("lang" . "en"))
                                                            "This field contains a patient's most recent marital (civil) status.")))
                            ("choice" (("minOccurs" . "0") ("maxOccurs" . "1"))
                                      ("annotation" ()
                                                    ("documentation" (("lang" . "en"))
                                                                     "Indicates whether the patient is part of a multiple (boolean) or indicates the actual birth order (integer)."))
                                      ("element" (("name" . "multipleBirthBoolean") ("type" . "boolean")))
                                      ("element" (("name" . "multipleBirthInteger") ("type" . "integer"))))
                            ("element"
                             (("name" . "photo") ("minOccurs" . "0") ("maxOccurs" . "unbounded") ("type" . "Attachment"))
                             ("annotation" () ("documentation" (("lang" . "en")) "Image of the patient.")))
                            ("element"
                             (("name" . "contact") ("type" . "Patient.Contact") ("minOccurs" . "0")
                              ("maxOccurs" . "unbounded"))
                             ("annotation" ()
                                           ("documentation" (("lang" . "en"))
                                                            "A contact party (e.g. guardian, partner, friend) for the patient.")))
                            ("element"
                             (("name" . "communication") ("type" . "Patient.Communication") ("minOccurs" . "0")
                              ("maxOccurs" . "unbounded"))
                             ("annotation" ()
                                           ("documentation" (("lang" . "en"))
                                                            "A language which may be used to communicate with the patient about his or her health.")))
                            ("element"
                             (("name" . "generalPractitioner") ("minOccurs" . "0") ("maxOccurs" . "unbounded")
                              ("type" . "Reference"))
                             ("annotation" ()
                                           ("documentation" (("lang" . "en"))
                                                            "Patient's nominated care provider.")))
                            ("element"
                             (("name" . "managingOrganization") ("minOccurs" . "0") ("maxOccurs" . "1")
                              ("type" . "Reference"))
                             ("annotation" ()
                                           ("documentation" (("lang" . "en"))
                                                            "Organization that is the custodian of the patient record.")))
                            ("element"
                             (("name" . "link") ("type" . "Patient.Link") ("minOccurs" . "0") ("maxOccurs" . "unbounded"))
                             ("annotation" ()
                                           ("documentation" (("lang" . "en"))
                                                            "Link to a Patient or RelatedPerson resource that concerns the same actual individual."))))))))
