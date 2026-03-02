package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	formated := "Глокая куздра штеко бокра будланула"
	fmt.Printf("3.7 Введен абзац из 5 слов: %s\n", formated)
	formated += " " // Для отработки алгоритма с последним словом
	var word string = ""
	var maxWord string
	maxlen := -1

	for _, char := range formated {
		if char != ' ' {
			word += string(char)
		} else {
			if utf8.RuneCountInString(word) > maxlen {
				maxWord = word
				maxlen = utf8.RuneCountInString(word)
			}
			word = ""
		}
	}

	fmt.Printf("Самое длинное слово '%s' длинной %d\n", maxWord, maxlen)
}