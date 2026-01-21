;;;;
;;;;

(defun convert-fhir-schema (input-filename output-filename)
  (let ((fhir-schema (make-bag nil))
        (schema (make-bag "{}")))
    (with-open-file (f input-filename :direction :input)
      (send fhir-schema :read f))

    (send schema :set (car (last (split (send fhir-schema :get "id") "/"))) "version")
    ;;"id": "http://hl7.org/fhir/json-schema/5.0",


    (let* ((defs (send fhir-schema :get "definitions" t))
           (res-map (send fhir-schema :get "discriminator.mapping" t))
           primitives)
      ;; Categorize definitions into primary, resource, datatypes, and
      ;; backbones.
      (dolist (adef (send defs :native))
        (let* ((name (car adef))
               (c0 (char name 0)))
          (cond ((and (char<= c0 #\z) (char>= c0 #\a))
                 (let ((pb (make-bag "{}"))
                       (fields (cdr adef))
                       val)
                   (when (setq val (cdr (assoc "description" fields)))
                     (send pb :set val "description"))
                   (when (setq val (cdr (assoc "pattern" fields)))
                     (send pb :set val "pattern"))
                   (send pb :set name "name")
                   ;; Since most primitives need some customization a case
                   ;; statement is used.
                   (case name
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

                ;; TBD other types
                (t
                 ;; (format t "*** datatype: ~A~%" name)
                 ))))
      (send schema :set primitives "primitives")
      (with-open-file (f output-filename :direction :output :if-exists :supersede :if-does-not-exist :create)
        (send schema :write f :pretty t :json t :depth 1)))))
