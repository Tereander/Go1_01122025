/*
Задача №1
Вводим любое натуральное число.
Нужно посчитать сумму цифр числа, с помощью цикла for
Пример:
In: 4521
Out: 12
*Задание 1.1: 4 + 5 + 2 + 1 = 12 - добавить к выводу сумму как выражение
*/
package main

import (
	"fmt"
	"strings"
	// "strings"
)

func main() {
	var user_value, sum int64
	var digits string
	fmt.Println("Enter your digital value: ")
	fmt.Scanln(&user_value)

	for {
		if user_value == 0 {
			break
		}
		var digit int64 = 0
		digit = user_value % 10
		sum += digit
		digits = fmt.Sprintf("%d + ", digit) + digits
		user_value /= 10
	}
	// Когда переменная должна быть создана, но она не нужна,
	// можно использовать заглушку в виде _
	digits, _ = strings.CutSuffix(digits, " + ")

	fmt.Println(digits, "=", sum)
}
