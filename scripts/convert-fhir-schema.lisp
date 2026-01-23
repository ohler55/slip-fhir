;;;;
;;;; Converts a FHIR JSON schema file into a JSON file suitable for loading
;;;; into the slip-fhir package. Some modifications to type names are made
;;;; during the conversion to avoid name clashes with built in class names.
;;;;
;;;; To generate a fhir5.json file in the fhir directory run the following.
;;;;
;;;; slip -e '(convert-fhir-schema "spec/fhir.schema.json" "fhir/fhir5.json")' scripts/*.lisp
;;;;

;;; The FHIR heirarchy is defined as:
;;; Base
;;;   Element
;;;     BackboneElement - not is schema file, same as BackboneType though
;;;     DataType
;;;       <-- most datatypes go here (e.g., Coding)
;;;       PrimitiveType <-- not really here in non-XML implementations
;;;       BackboneType
;;;   Resource - not in the schema file
;;;     DomainResource - not in the schema file
;;;       <-- most resources go here
;;;
(defun form-hierarchy-node (name def)
  (let ((hb (make-bag "{}"))
        val)
    (bag-set hb name "name")
    (when (setq val (bag-get def "description"))
      (send hb :set val "description"))
    (if (equal "Element" name)
        (add-properties hb def nil)
        (add-properties hb def '("id" "extension")))
    (send hb :set (case name
                    ("Base" nil)
                    ("Element" "Base")
                    ("DataType" "Element")
                    ("BackboneType" "DataType")) "parent")
    hb))

(defun correct-type (name type)
  (cond ((equal type "string")
         (cond ((suffixp name "Instant") "instant")
               ((suffixp name "DateTime") "dateTime")
               ((suffixp name "Time") "ftime")
               ((suffixp name "Integer") "integer32")
               ((suffixp name "Integer64") "integer64")
               ((suffixp name "Decimal") "decimal")
               ((suffixp name "Base64Binary") "base64Binary")
               ((suffixp name "Canonical") "canonical")
               ((suffixp name "Code") "code")
               ((suffixp name "Date") "date")
               ((suffixp name "Id") "id")
               ((suffixp name "Markdown") "markdown")
               ((suffixp name "Oid") "oid")
               ((suffixp name "PositiveInt") "positiveInt")
               ((suffixp name "UnsignedInt") "unsignedInt")
               ((suffixp name "Uri") "uri")
               ((suffixp name "Url") "url")
               ((suffixp name "Uuid") "uuid")
               (t "fstring")))
        ((equal type "time") "ftime")
        ((equal type "integer") "integer32")
        ((equal type "number")
         (cond ((suffixp name "PositiveInt") "positiveInt")
               ((suffixp name "UnsignedInt") "unsignedInt")
               ((suffixp name "Integer") "integer32")
               ((suffixp name "Integer64") "integer64")
               ((suffixp name "Decimal") "decimal")
               (t type)))
        ((prefixp name "_") "Extension")
        (t type)))

(defun determine-type (name pdef prop)
  "Determine the type of a property from a FHIR schema definition. The type can
   appear is 'type', '$ref', 'items.$ref', or 'items.enum'. It each case some
   additional guesses need to be made depending on the property name since the
   FHIR schema file does not match the published schema on the web site."
  (let ((type (bag-get pdef "type"))
        (ref (bag-get pdef "['$ref']"))
        (enum (bag-get pdef "enum")))
    (cond ((equal type "string")
           (bag-set prop (correct-type name type) "type"))
          ((equal type "array")
           (bag-set prop t "array")
           (let ((ienum (bag-get pdef "items.enum"))
                 (iref (bag-get pdef "['items']['$ref']")))
             (cond (ienum
                    (bag-set prop "code" "type")
                    (bag-set prop ienum "enum"))
                   (iref
                    (setq type (car (last (split iref "/"))))
                    (bag-set prop (correct-type name type) "type"))
                   (t
                    (bag-set prop "fstring" "type")))))
          (ref
           (setq type (car (last (split ref "/"))))
           (bag-set prop (correct-type name type) "type"))
          (enum
           (bag-set prop "code" "type")
           (bag-set prop enum "enum"))
          (t
           (bag-set prop (correct-type name type) "type")))))

(defun form-property (name def req)
  (let ((prop (make-bag "{}"))
        val)
    (bag-set prop name "name")
    (determine-type name def prop)
    (when (setq val (bag-get def "description"))
      (bag-set prop val "description"))
    (when (and (setq val (bag-get def "pattern")) (equal (bag-get prop "type") "fstring"))
      (bag-set prop val "pattern"))
    (when req (bag-set prop t "required"))
    prop))

(defun add-properties (rb def ignore)
  (let ((reqs (bag-get def "required"))
        props)
    (bag-scan (or (bag-get def "properties" t) (make-bag '()))
              (lambda (p def)
                (when (< 2 (length p))
                    (setq p (subseq p 2))
                    (unless (or (containsp p ".") (containsp p "[")) ;; only top level nodes
                      (unless (member p ignore)
                        (setq props (add props (form-property p def (member p reqs)))))))))

    (when props (bag-set rb props "properties"))))

(defun form-datatype-node (name def)
  (let ((dt (make-bag "{}"))
        val)
    (bag-set dt name "name")
    (when (setq val (bag-get def "description"))
      (bag-set dt val "description"))
    (add-properties dt def '("id" "extension"))
    (bag-set dt "DataType" "parent")
    dt))

(defun form-backbone-node (name def)
  (let ((dt (make-bag "{}"))
        val)
    (bag-set dt name "name")
    (when (setq val (bag-get def "description"))
      (bag-set dt val "description"))
    (add-properties dt def '("id" "extension" "modifierExtension"))
    (bag-set dt "BackboneType" "parent")
    dt))

(defun form-resource-node (name def)
  (let ((dt (make-bag "{}"))
        val)
    (bag-set dt name "name")
    (when (setq val (bag-get def "description"))
      (bag-set dt val "description"))
    (add-properties dt def '("id" "meta" "implicitRules" "_implicitRules"
                             "language" "_language" "text" "_text" "contained" "modifierExtension"))
    (bag-set dt "DomainResource" "parent")
    dt))

(defun load-sen (filename)
  (let ((bag (make-bag nil)))
    (with-open-file (f filename :direction :input)
      (send bag :read f))
    bag))

(defun convert-fhir-schema (input-filename output-filename)
  "Convert a fhir.schema.json file to a schema file suitable for loading by this
   package."
  (let ((fhir-schema (load-sen input-filename))
        (resource-schema (load-sen "spec/resource-schema.sen"))
        (domain-resource-schema (load-sen "spec/domainresource-schema.sen"))
        (schema (make-bag "{}")))

    (send schema :set (car (last (split (send fhir-schema :get "id") "/"))) "version")

    ;; Grab the definitions.
    (let* ((defs (send fhir-schema :get "definitions" t))
           (res-map (send fhir-schema :get "discriminator.mapping" t))
           primitives datatypes hierarchy backbones resources)
      ;; Scan all nodes in the defs and process all the top level nodes,
      (bag-scan defs
                (lambda (p def)
                  (when (< 2 (length p))
                    (setq p (subseq p 2))
                    (unless (containsp p ".") ;; only top level nodes
                      ;; Categorize definitions into primary, resource,
                      ;; datatypes, and backbones.
                      (let ((c0 (char p 0)))
                        ;; Primitives definitions all start with a lowercase character.
                        (cond ((and (char<= c0 #\z) (char>= c0 #\a))
                               (let ((pb (make-bag "{}"))
                                     val)
                                 (when (setq val (bag-get def "description"))
                                   (send pb :set val "description"))
                                 (when (setq val (bag-get def "pattern"))
                                   (send pb :set val "pattern"))
                                 (send pb :set p "name")
                                 ;; Since most primitives need some
                                 ;; customization a case statement is used.
                                 (case p
                                   ("string"
                                    (send pb :set "fstring" "name")
                                    (send pb :set "string" "parent"))
                                   ("integer"
                                    (send pb :set "integer32" "name")
                                    (send pb :set "fixnum" "parent"))
                                   ("unsignedInt" (send pb :set "integer32" "parent"))
                                   ("positiveInt" (send pb :set "integer32" "parent"))
                                   ("integer64" (send pb :set "fixnum" "parent"))
                                   ("decimal" (send pb :set "double-float" "parent"))
                                   ("boolean" (send pb :set "symbol" "parent"))
                                   ("time"
                                    (send pb :set "ftime" "name")
                                    (send pb :set "string" "parent"))
                                   ("date" (send pb :set "string" "parent"))
                                   ((or "instant" "dateTime")
                                    (send pb :set "time" "parent"))
                                   (t (send pb :set "fstring" "parent")))

                                 (setq primitives (add primitives pb))))

                              ;; Resource definition are in the discriminator.mapping element.
                              ((bag-has res-map p)
                               (setq resources (add resources (form-resource-node p def))))

                              ;; Backbone definitions all include an _ character.
                              ((containsp p "_")
                               (setq backbones (add backbones (form-backbone-node p def))))

                              ;; BackboneElement is not used and is the same as BackboneType.
                              ((string= "BackboneElement" p) nil)
                              ;; PrimitiveType is handled independent of the definitions.
                              ((string= "PrimitiveType" p) nil)

                              ;; Hierarchy types are handled individually.
                              ((member p '("Base"
                                           "Element"
                                           "DataType"
                                           "BackboneType"))
                               (setq hierarchy (add hierarchy (form-hierarchy-node p def))))

                              ;; Everything else is a FHIR DataType.
                              (t
                               (setq datatypes (add datatypes (form-datatype-node p def))))
                              ))))))
      (send schema :set primitives "primitives")

      (bag-set resource-schema language-codes "properties[?@.name == 'language'].enum")
      (setq hierarchy (add hierarchy resource-schema domain-resource-schema))
      (send schema :set hierarchy "hierarchy")

      (send schema :set datatypes "datatypes")

      (send schema :set backbones "backbones")

      (send schema :set resources "resources")

      (with-open-file (f output-filename :direction :output :if-exists :supersede :if-does-not-exist :create)
        (send schema :write f :pretty t :json t :depth 1)))))
