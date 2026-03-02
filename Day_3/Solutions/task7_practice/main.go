/*
	3.7 Вводим строку без знаков препинания(5 слов).
		Найти самое длинное слово в строке и вывести сколько в нем букв.

	Пример:
	In: Скажи как дела в учебе
	Out: учебе, 5

	In: Закрепляем циклы в языке Golang
	Out: Закрепляем, 10
*/

package main

import "fmt"

func main() {
	    var words [5]string
        fmt.Println("Задача 3.7")
        fmt.Println("Введите фразу из пяти слов:")
        fmt.Scan(&words[0], &words[1], &words[2], &words[3], &words[4])
        word_max_length := 0
        word_with_max_length := ""
        for i := range 5 {
			if word_max_length < len(words[i]) {
				word_max_length = len(words[i])
				word_with_max_length = words[i]
			}
        }
        fmt.Printf("Самое длинное слово в строке: %s, число букв: %d\n", word_with_max_length, word_max_length)
}
