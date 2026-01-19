;;;;
;;;;

(defun convert-fhir-schema (input-filename output-filename)
  (let ((fhir-schema (make-bag nil))
        (schema (make-bag "{}")))
    (with-open-file (f input-filename :direction :input)
      (send fhir-schema :read f))
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
                   (cond ((string= "string" name) (setq name "fstring"))
                         ((string= "integer" name) (setq name "integer32")))
                   (send pb :set name "name")
                   (when (setq val (cdr (assoc "description" fields)))
                     (send pb :set val "description"))
                   (when (setq val (cdr (assoc "pattern" fields)))
                     (send pb :set val "pattern"))
                   (setq val (cdr (assoc "type" fields)))
                   (send pb :set
                         (cond ((equal val "string") "fstring")
                               ((equal val "boolean") "symbol") ;; TBD what should this be?
                               ((equal val "number")
                                (ecase name
                                  ("unsignedInt" "integer32")
                                  ("positiveInt" "integer32")
                                  ("integer32" "integer")
                                  ("decimal" "double-float"))))
                         "parent")
                   (setq primitives (add primitives pb))))

                ;; TBD other types
                (t
                 ;; (format t "*** datatype: ~A~%" name)
                 ))))
      (send schema :set primitives "primitives")
      (with-open-file (f output-filename :direction :output :if-exists :supersede :if-does-not-exist :create)
        (send schema :write f :pretty t :json t)))))
