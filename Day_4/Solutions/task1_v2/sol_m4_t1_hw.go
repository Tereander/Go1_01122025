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

// Ваш код
/*
создаем мапу, ключ - число от 00 до 25, значение - буква a-z.
То есть находим код "а", исходим из того, что это 0б далее в цикле
Далее идем в цикле по строке с шагом 2, берем по 2 элемента, берем из мапы значение.
Пишем в переменную - печатаем расшифрованное.
Для шифрования нужна обратная мапка.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	const code = "220411112603141304"
	fmt.Println("Result:", decode(code))

}
func decode(code string) string{
	dictionary := make(map[int]string)

	for i := range 26 {
		dictionary[i] = string(rune(97 + i))
	}
	dictionary[26] = " "

	var digit_char string
	var phrase string
	
	for i := 0; i < len(code); i += 2 {
		digit_char = string(code[i]) + string(code[i+1])
		int_char, _ := strconv.Atoi(digit_char)
		phrase += dictionary[int_char]
	}
	return phrase
}

