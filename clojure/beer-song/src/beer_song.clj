(ns beer-song)

(defn bottles [number]
  (if (= number 1)
    "bottle"
    "bottles"))

(defn first-line [number]
  (str number " bottles of beer on the wall, " number " bottles of beer.\n"))

(defn second-line [number]
  (str "Take one down and pass it around, " number " bottles of beer on the wall.\n"))

(defn verse [number]
  (str (first-line number) (second-line (dec number))))

(defn sing [start end])
