(ns anagram
  (:require [clojure.string :as str]))

(defn anagrams-for [word candidates]
  (let [same-letters (comp sort seq str/lower-case)
        same-word #(= (str/lower-case word) (str/lower-case  %))]
    (filter #(and  (not (same-word %))
                   (= (same-letters word) (same-letters %)))
            candidates)))

