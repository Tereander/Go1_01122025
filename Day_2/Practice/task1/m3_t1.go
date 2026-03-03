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

import "fmt"

func main() {
	var inputNumber, finalSum int
	var expression string

	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&inputNumber)
	if err != nil {
		return
	}

	for inputNumber > 0 {
		digit := inputNumber % 10
		finalSum += digit

		if len(expression) == 0 {
			expression += fmt.Sprint(digit)
		} else {
			expression += " + " + fmt.Sprint(digit)
		}

		inputNumber /= 10
	}
	fmt.Println("Сумма чисел:", finalSum)

	fmt.Println("Полное выражение:", expression, "=", finalSum)
}
