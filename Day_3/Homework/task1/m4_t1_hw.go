/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с "00" -> 'a' и до "25" -> 'z', "26" -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через другую строку.
Рекомендую использовать мапы, будет лучше, если вы их создадите с помощью цикла

Задача №1.1
Реализовать и функцию зашифровки

codeToString(code) -> "???????'

stringToCode("hello") -> "??????"
*/
package main

import (
	"fmt"
	"strconv"
)

func codeToString(code string) string {
	codeMap := make(map[int]rune)

	for i := 0; i < 26; i++ {
		codeMap[i] = rune(i + 97)
	}
	codeMap[26] = ' '

	var result []rune

	for i := 0; i < len(code); i += 2 {
		twoDigits := code[i : i+2]

		num, err := strconv.Atoi(twoDigits)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			return ""
		}

		if char, exists := codeMap[num]; exists {
			result = append(result, char)
		}
	}

	return string(result)
}

func stringToCode(text string) string {
	charMap := make(map[rune]int)

	for i := 0; i < 26; i++ {
		charMap[rune(i+97)] = i
	}
	charMap[' '] = 26

	var result string

	for _, char := range text {
		if code, exists := charMap[char]; exists {
			codeStr := fmt.Sprintf("%02d", code)
			result += codeStr
		}
	}

	return result
}

func main() {
	// Тестируем расшифровку
	code := "220411112603141304"
	decoded := codeToString(code)
	fmt.Printf("Зашифровано: %s\n", code)
	fmt.Printf("Расшифровано: %s\n", decoded)

	// Тестируем шифровку
	text := "hello"
	encoded := stringToCode(text)
	fmt.Printf("\nТекст: %s\n", text)
	fmt.Printf("Зашифровано: %s\n", encoded)

	fmt.Printf("\nПроверка: %s -> %s -> %s\n",
		text,
		encoded,
		codeToString(encoded))
}
