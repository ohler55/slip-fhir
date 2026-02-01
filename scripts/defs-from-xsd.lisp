;;;;


(defun find-named (schema name)
  (dolist (element schema)
    (when (and (equal name (cdr (assoc "name" (cadr element))))
               (equal "complexType" (car element)))
      (format t "~A~%" element)
      )))


(defun defs-from-xsd (input-filename output-filename)
  "TBD."
  (let ((schema (cddr (assoc "schema" (with-open-file (f input-filename :direction :input) (xml-read f)))))
        (defs (make-bag "{}")))

    (do ((name (caar schema) (caar schema)))
        ((not (or (equal :comment name) (equal "import" name))))
      (setq schema (cdr schema)))


    ;; (format t "~A~%" schema)

    ;;(find-named schema "Patient")
    (dolist (element schema)
      (format t "~A ~A~%" (car element) (cdr (assoc "name" (cadr element)))))
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
