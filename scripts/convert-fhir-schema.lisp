;;;;
;;;; Converts a FHIR JSON schema file into a JSON file suitable for loading
;;;; into the slip-fhir package. Some modifications to type names are made
;;;; during the conversion to avoid name clashes with built in class names.
;;;;
;;;; To generate a fhir5.json file in the fhir directory run the following.
;;;;
;;;; slip -e '(convert-fhir-schema "spec/fhir.schema.json" "fhir/fhir5.json")' scripts/convert-fhir-schema.lisp
;;;;

(defun convert-fhir-schema (input-filename output-filename)
  "Convert a fhir.schema.json file to a schema file suitable for loading by this
   package."
  (let ((fhir-schema (make-bag nil))
        (schema (make-bag "{}")))
    ;; Load the FHIR schema file into a bag (json in memory).
    (with-open-file (f input-filename :direction :input)
      (send fhir-schema :read f))

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
                               ;; (format t "*** resource: ~A~%" p)
                               )

                              ;; Backbone definitions all include an _ character.
                              ((containsp p "_")
                               ;; (format t "*** backbone: ~A~%" p)
                               )

                              ;; Hierarchy types are handled individually.
                              ((member p '("Base"
                                           "Element"
                                           "BackboneElement"
                                           "DataType"
                                           "PrimitiveType"
                                           "BackboneType"
                                           "Resource"
                                           "DomainResource"))
                               (setq hierarchy (add hierarchy (form-hierarchy-node p def))))

                              ;; Everything else is a FHIR DataType.
                              (t
                               ;; (format t "*** datatype: ~A~%" p)
                               ;; special case
                               )))))))
      (send schema :set primitives "primitives")
      ;; TBD add missing hierarchy types
      ;; TBD add datatypes, hierarchy, backbones, and resources
      (with-open-file (f output-filename :direction :output :if-exists :supersede :if-does-not-exist :create)
        (send schema :write f :pretty t :json t :depth 1)))))

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
  ;; (format t "*** hierarchy: ~A~%" p)
  (let ((hb (make-bag "{}")))
    (bag-set hb name "name")
    hb))
