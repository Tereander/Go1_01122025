/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/
package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите четырехзначное число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}

	// Проверяем, что число четырехзначное
	if number < 1000 || number > 9999 {
		fmt.Println("Ошибка: число должно быть четырехзначным")
		return
	}

	// Преобразуем число в строку
	str := fmt.Sprintf("%d", number)

	// Проверяем, является ли число палиндромом
	if str[0] == str[3] && str[1] == str[2] {
		fmt.Printf("%d - палиндром\n", number)
	} else {
		fmt.Printf("%d - не палиндром\n", number)
	}
}
