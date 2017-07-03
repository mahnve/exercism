;;; bob.el --- Bob exercise (exercism)

;;; Commentary:

;;; Code:
(defun lastchar (str char)
  (string= (substring str -1) char))

(defun caps(str)
  (string= (upcase str) str))

(defun response-for (str)
  (if (caps str)
      "Whoa, chill out!"
    (if (lastchar str "?")
        "Sure.")
    "Whatever."))

(provide 'bob)
;;; bob.el ends here
