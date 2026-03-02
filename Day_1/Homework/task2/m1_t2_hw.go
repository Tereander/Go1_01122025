/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 120, выход: 021
вход: 100, выход: 001
*/
package main

import (
	"fmt"
)

func reverseNumber(n int) string {
	edenichki := n % 10
	tens := (n / 10) % 10
	hundreds := n / 100

	return fmt.Sprintf("%d%d%d", edenichki, tens, hundreds)
}

func main() {
	var number int
	fmt.Println("Введите трехзначное число: ")

	_, err := fmt.Scan(&number)
	//проверку на пустоту предложила ide , в целом я согласен :)
	if err != nil {
		return
	}

	if number < 100 || number > 999 {
		fmt.Println("Ошибка: Число должно быть трехзначным")
		return
	}

	result := reverseNumber(number)
	fmt.Printf("Результат: %s\n", result)
}
