;;; hello-world.el --- Hello World Exercise (exercism)

;;; Commentary:

;;; Code:

(defun hello (&optional a-string)
  (if a-string
      (concat "Hello, " a-string "!")
    "Hello, World!"))

(provide 'hello-world)
;;; hello-world.el ends here
