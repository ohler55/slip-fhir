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
  "Build a base oe hierarchy node from the provided schema definition."
  (let ((hb (make-bag "{}"))
        val)
    (bag-set hb name "name")
    (when (setq val (bag-get def "description"))
      (send hb :set (replace-all val "\r" "\n") "description"))
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
  "The FHIR schema file defines pimitives but then ignores those definitions in
   property definitions and resorts to regexp patterns instead. This function
   corrects that and sets the property type according to what should have been
   in the schema file."
  (cond ((equal type "string")
         (cond ((suffixp name "Instant") "instant")
               ((suffixp name "DateTime") "dateTime")
               ((suffixp name "Time") "time")
               ((suffixp name "Integer") "integer")
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
               (t "string")))
        ((equal type "number")
         (cond ((suffixp name "PositiveInt") "positiveInt")
               ((suffixp name "UnsignedInt") "unsignedInt")
               ((suffixp name "Integer") "integer")
               ((suffixp name "Integer64") "integer64")
               ((suffixp name "Decimal") "decimal")
               (t type)))
        ((prefixp name "_") "Extension") ;; Element in the file when it should be Extension
        (t type)))

(defun determine-type (name pdef prop)
  "Determine the type of a property from a FHIR schema definition. The type can
   appear is 'type', '$ref', 'items.$ref', 'items.enum', or 'const'. It each
   case some additional guesses need to be made depending on the property name
   since the FHIR schema file does not match the published schema on the web
   site."
  (let ((type (or (bag-get pdef "type") (bag-get pdef "const")))
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
                    (bag-set prop "string" "type")))))
          (ref
           (setq type (car (last (split ref "/"))))
           (bag-set prop (correct-type name type) "type"))
          (enum
           (bag-set prop "code" "type")
           (bag-set prop enum "enum"))
          (t
           (bag-set prop (correct-type name type) "type")))))

(defun add-property-enum (prop container name)
  (unless (< 0 (length (bag-get prop "enum")))
    (let ((crx (join "" "(?i)^" container "-"))
          (nrx (join "" "-" name "$"))
          (frx (join "" "(?i)^" container "-" name "$"))
          found)
      ;; explain
      (dolist (enum *enum-map*)
        (when (regex-match nrx (car enum))
          (setq found (add found enum))))
      (cond ((null found) nil)
            ((= 1 (length found)) (bag-set prop (cdar found) "enum"))
            ;; more than one candidate enum found
            ((dolist (fe found)
               (when (regex-match frx (car fe))
                 (bag-set prop (cdr fe) "enum")
                 (return t))) nil)
            (t ;; TBD manually set up links between enums and elements
             ;; (format t "*** too many for enum matches for ~A.~A: ~A~%" container name (mapcar #'car found))
             nil
             )))))

(defun form-property (container name def req)
  "Forms a property node from the provided definition. Since the indicator that
   the property is a required property is outside the property definition it
   is determined outside this function and the indicator pass in as an
   argument."
  (let ((prop (make-bag "{}"))
        type
        val)
    (bag-set prop name "name")
    (determine-type name def prop)
    (setq type (bag-get prop "type"))
    (when (setq val (bag-get def "description"))
      ;; To stay consistent returns are replaced with newlines.
      (bag-set prop (replace-all val "\r" "\n") "description"))
    (cond ((and (setq val (bag-get def "pattern")) (equal type "string"))
           (bag-set prop val "pattern"))
          ((equal name "resourceType")
           (bag-set prop (list type) "enum")
           (bag-set prop "code" "type"))
          ((equal type "code")
           (add-property-enum prop container name)))
    (when req (bag-set prop t "required"))
    prop))

(defun add-properties (rb def ignore)
  "Scans a type definition for properties and adds the nodes from those
   properties to a list which is added to the resource bag at the end of the
   function."
  (let ((reqs (bag-get def "required"))
        props)
    (bag-scan (or (bag-get def "properties" t) (make-bag '()))
              (lambda (p def)
                (when (< 2 (length p))
                    (setq p (subseq p 2))
                    (unless (or (containsp p ".") (containsp p "[")) ;; only top level nodes
                      (unless (member p ignore)
                        (setq props (add props (form-property (bag-get rb "name") p def (member p reqs)))))))))

    (when props (bag-set rb props "properties"))))

(defun form-datatype-node (name def)
  "Forms a DataType node from the provided definition."
  (let ((dt (make-bag "{}"))
        val)
    (bag-set dt name "name")
    (when (setq val (bag-get def "description"))
      ;; To stay consistent returns are replaced with newlines.
      (bag-set dt (replace-all val "\r" "\n") "description"))
    (add-properties dt def '("id" "extension"))
    (bag-set dt "Element" "parent") ;; The web pages show Element for individual view yet DataType in the model.
    dt))

(defun form-backbone-node (name def)
  "Forms a BackboneType node from a schema definition."
  (let ((dt (make-bag "{}"))
        val)
    (bag-set dt name "name")
    (when (setq val (bag-get def "description"))
      ;; To stay consistent returns are replaced with newlines.
      (bag-set dt (replace-all val "\r" "\n") "description"))
    (add-properties dt def '("id" "extension" "modifierExtension"))
    (bag-set dt "BackboneType" "parent")
    dt))

(defun form-resource-node (name def)
  "Forms a resource node from a schema definition."
  (let ((dt (make-bag "{}"))
        val)
    (bag-set dt name "name")
    (when (setq val (bag-get def "description"))
      ;; To stay consistent returns are replaced with newlines.
      (bag-set dt (replace-all val "\r" "\n") "description"))
    (add-properties dt def '("id" "meta" "implicitRules" "_implicitRules"
                             "language" "_language" "text" "_text" "contained" "modifierExtension"))
    (bag-set dt "DomainResource" "parent")
    dt))

(defun prop-in-groups-p (prop groups)
  (let ((pname (bag-get prop "name")))
    (dolist (group groups)
      (when (dolist (gp (cdr group))
              (when (equal pname (bag-get gp "name"))
                (return t)))
        (return t)))))

(defun get-group-prop (group)
  (dolist (p (cdr group))
    (unless (prefixp (bag-get p "name") "_")
      (return p))))

(defun discover-groups-in-type (type-node patterns)
  "In the model definitions on the FHIR web site some properties are listed as
   groups with a notation such as 'value[x]' followed by the acceptable
   sub-type such as 'valueString' or 'valueInteger'. The schema file does not
   include that information. This function looks for property names that match
   the pattern of a prefix followed by a type suffix. All must have the same
   cardinality. When found those are grouped together."
  (let (matches ;; property list
        groups) ;; assoc list
    (bag-walk type-node (lambda (p)
                          ;;(format t "*** prop: ~A ~A ~A~%" (bag-get p "name") (bag-get p "required") (bag-get p "array")))
                          (let* ((name (bag-get p "name"))
                                 (req (bag-get p "required"))
                                 (ary (bag-get p "array")))

                            (dolist (rx patterns)
                              (when (regex-match rx name)
                                (let* ((prefix (subseq name 0 (+ (- (length name) (length rx)) 3)))
                                       (key (join "" prefix (if req "1" "0") (if ary "x" "1")))
                                       (lst (getf matches key)))
                                  ;; add but keep going to match paterns like DateTime and Time
                                  (setf (getf matches key) (add lst p)))))))
              "properties[*]" t)
    (dotimes (i (/ (length matches) 2))
      (let* ((key (nth (* 2 i) matches))
             (prefix (subseq key 0  (- (length key) 2)))
             (group (nth (1+ (* 2 i)) matches)))
        (when (and (< 1 (length group)) (not (prefixp prefix "_")))
          (setq groups (add groups (cons prefix (append (getf matches (join "" "_" prefix)) group)))))))
    (when (< 0 (length groups))
      (let (props)
        (bag-walk type-node (lambda (p)
                              (unless (prop-in-groups-p p groups)
                                (setq props (add props p))))
                  "properties[*]" t)
        (dolist (group groups)
          (let ((gp (make-bag "{}"))
                (np (get-group-prop group)))
            (bag-set gp (join "" (car group) "[x]") "name")
            (bag-set gp (cdr group) "group")
            (bag-set gp (bag-get np "description") "description")
            (when (bag-get np "array")  (bag-set gp t "array"))
            (when (bag-get np "required")  (bag-set gp t "required"))
            (dolist (p (cdr group))
              (unless (prefixp (bag-get p "name") "_")
                (bag-remove p "description")))
            (setq props (add props gp))))
        (bag-set type-node props "properties")))))

(defun discover-groups (schema)
  "Groups such as value[x] can be found in resources, backbones, and
   datatype. Search each type, discover groups, and then replace sets of
   properties with a group."

  ;; Groups only have primitve or datatype suffixes.
  (let (patterns)
    ;; Build a set of regex patterns to use when checking candidates for
    ;; grouping.
    (bag-walk schema (lambda (name)
                       (setq patterns
                             (add patterns
                                  (format nil ".+~A$" (string-upcase name :start 0 :end 1)))))
              "primitives[*].name")
    (bag-walk schema (lambda (name)
                       (setq patterns (add patterns (format nil ".+~A$" name))))
              "datatypes[*].name")

    ;; Armed with a set of patterns check each non-primitive type for groups
    ;; and modify.
    (bag-walk schema
              (lambda (type-node) (discover-groups-in-type type-node patterns))
              "resources[*]" t)
    (bag-walk schema
              (lambda (type-node) (discover-groups-in-type type-node patterns))
              "backbones[*]" t)
    (bag-walk schema
              (lambda (type-node) (discover-groups-in-type type-node patterns))
              "datatypes[*]" t)
    ))

(defun convert-fhir-schema (input-filename output-filename)
  "Convert a fhir.schema.json file to a schema file suitable for loading by this
   package."
  (let ((fhir-schema (load-bag input-filename))
        (resource-schema (load-bag "spec/resource-schema.sen"))
        (domain-resource-schema (load-bag "spec/domainresource-schema.sen"))
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
                                   (send pb :set (replace-all val "\r" "\n") "description"))
                                 (when (setq val (bag-get def "pattern"))
                                   (send pb :set val "pattern"))
                                 (send pb :set p "name")
                                 ;; Since most primitives need some
                                 ;; customization a case statement is used.
                                 (case p
                                   ("string"
                                    ;;(send pb :set "string" "name")
                                    (send pb :set "cl:string" "parent"))
                                   ("integer"
                                    (send pb :set "fixnum" "parent"))
                                   ("unsignedInt" (send pb :set "integer" "parent"))
                                   ("positiveInt" (send pb :set "integer" "parent"))
                                   ("integer64" (send pb :set "fixnum" "parent"))
                                   ("decimal" (send pb :set "double-float" "parent"))
                                   ("boolean" (send pb :set "symbol" "parent"))
                                   ("time" (send pb :set "cl:string" "parent"))
                                   ("date" (send pb :set "cl:string" "parent"))
                                   ((or "instant" "dateTime") (send pb :set "cl:time" "parent"))
                                   ((or "id" "code" "markdown") (send pb :set "string" "parent"))
                                   (t (send pb :set "string" "parent")))

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

      ;; The language enum are not listed in the schema file so they are added here.
      (bag-set resource-schema language-codes "properties[?@.name == 'language'].enum")
      ;; (bag-set resource-schema (cdr (assoc "languages" *enum-map*)) "properties[?@.name == 'language'].enum")
      (setq hierarchy (add hierarchy resource-schema domain-resource-schema))

      ;; Add each type list to the new schema being constructed.
      (send schema :set hierarchy "hierarchy")
      (send schema :set datatypes "datatypes")
      (send schema :set backbones "backbones")
      (send schema :set resources "resources")

      ;; There are no indicators for groups in the schema file so discover
      ;; them and update the new schema.
      (discover-groups schema)

      (with-open-file (f output-filename :direction :output :if-exists :supersede :if-does-not-exist :create)
        (send schema :write f :pretty t :json t :depth 1)))))
