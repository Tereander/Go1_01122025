/*
3.6 Получить от пользователя натуральное число, посчитать сумму цифр в нём.

	Решить с помощью индексов, т.е. работаем с числом, как со строкой.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string

	fmt.Print("Введите натуральное число: ")
	fmt.Scan(&number)

	// Variant 1
	var sum int
	for idx := range number {
		sum += int(number[idx] - '0') // ASCII code 48, value '0'
	}
	fmt.Println("Сумма цифр:", sum)

	// Variant 2
	var digit, summ int

	for i := 0; i < len(number); i++ {
		digit, _ = strconv.Atoi(string(number[i]))
		summ += digit
	}
	fmt.Println(summ)
}
