def is_isogram(word):

    word = word.replace(" ", "")
    return sorted(list(set(word.lower()))) == sorted(list(word.lower()))
