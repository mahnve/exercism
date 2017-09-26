def is_isogram(word):

    word = word.replace("-", "")
    word = word.replace(" ", "")
    return sorted(set(word.lower())) == sorted(list(word.lower()))
