;;;;

(defun extract-enums (input-filename output-filename)
  "TBD"
  (let ((valuesets (load-bag input-filename))
        (*print-right-margin* 120))
    (with-open-file (f output-filename :direction :output :if-exists :supersede :if-does-not-exist :create)
      (write-string "
(defconstant *enum-map*
  ;; id of enum followed by the values
  '(
" f)
      (bag-walk valuesets (lambda (x)
                            (when (equal "complete" (bag-get x "resource.content"))
                              (let (values)
                                (bag-walk x (lambda (y) (setq values (add values y)))
                                          "resource.concept[*].code")
                                (format f "    ~A~%" (cons (bag-get x "resource.id") values)))))
                "entry[?@.resource.resourceType == 'CodeSystem']" t)
      (write-string "    ))
" f))))
