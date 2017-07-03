(ns word-count
  (:require [clojure.string :as string]
            [clojure.string :as str]))

(defn word-count [a-string]
  (let [cleaned-up-string (-> a-string
                              (str/lower-case)
                              (string/split #"[^\w]+"))]
    (frequencies cleaned-up-string)))

