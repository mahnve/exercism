(ns rna-transcription)

(defn to-single-rna [code]
  (case code
    \A    \U
    \G    \C
    \T    \A
    \C    \G
    (throw (AssertionError.))))

(defn to-rna [code]
  (apply str (map to-single-rna code)))
