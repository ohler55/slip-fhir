;;;;
;;;; Capture the summary, modifier, and required flags from profile file from
;;;; the full FHIR spec (https://hl7.org/fhir/fhir-spec.zip). Unzip the spec
;;;; file. If the file is unzipped to spec/r5/full then the profiles will be
;;;; at spec/r5/full/fhir-spec/site/*.profile.json.
;;;;
;;;; Generate the flags file with this command:
;;;;
;;;; slip -e '(capture-the-flag "spec/r5/full/fhir-spec/site/*.profile.json" "spec/r5/flags.sen")' scripts/capture-the-flag.lisp
;;;;

(defun bag-get-list (bag path)
  "Get all the matching values and return as a list."
  (let (lst)
    (bag-walk bag (lambda (v) (addf lst v)) path)
    lst))

(defun expand-x (prop prefix)
  "Expands value[x] patterns into the designated expansions like valueBoolean."
  (setq prefix (subseq prefix 0 (- (length prefix) 3)))
  (let (xa)
    (dolist (suffix (bag-get-list prop "type[*].code"))
      (let ((cap (string-capitalize suffix)))
        (cond ((string= "Datetime" cap) (setq cap "DateTime"))
              ((string= "Codeableconcept" cap) (setq cap "CodeableConcept"))
              ((string= "Sampleddate" cap) (setq cap "SampledDate")))
        (addf xa (join "" prefix cap))))
    xa))

(defun capture-the-flag (src out)
  "Collects summary, modifier, and required flags from resource profile JSON
   files and generates a summaried SEN file from a combination of all the
   resources."
  (let ((flag-map (make-bag "{}")))
    (dolist (path (directory src))
      (let* ((profile (with-open-file (f path :direction :input) (bag-read (make-bag "{}") f)))
             (res-type (bag-get profile "id"))
             (type (bag-get profile "type"))
             (kind (bag-get profile "kind"))
             (res-flags (make-bag (list (cons "resourceType" res-type))))
             modifiers
             (summary '("resourceType"))
             (required '("resourceType")))
        (when (and (string= res-type type) (string= "resource" kind))
          (bag-walk profile (lambda (prop)
                              (let* ((id (bag-get prop "id"))
                                     (sum (bag-get prop "isSummary"))
                                     (mod (bag-get prop "isModifier"))
                                     (min (bag-get prop "min"))
                                     (ida (split id ".")))
                                (when (= 2 (length ida))
                                  (cond ((search "[x]" id)
                                         (when (or sum mod (= 1 min))
                                           (let ((x (expand-x prop (cadr ida))))
                                             (when sum (setq summary (append summary x)))
                                             (when mod (setq modifiers (append modifiers x)))
                                             (when (= 1 min)
                                               (addf required (cadr ida))
                                               (setq required (append required x))))))
                                        (t
                                         (when sum (addf summary (cadr ida)))
                                         (when mod (addf modifiers (cadr ida)))
                                         (when (= 1 min) (addf required (cadr ida))))))))
                    "snapshot.element[*]" t)
          (bag-set res-flags summary "summary")
          (bag-set res-flags modifiers "modifiers")
          (bag-set res-flags required "required")
          (bag-set flag-map res-flags (format nil "['~A']" res-type)))))
    (with-open-file (f out :direction :output :if-exists :supersede :if-does-not-exist :create)
      (bag-write flag-map f)
      (terpri f))))
