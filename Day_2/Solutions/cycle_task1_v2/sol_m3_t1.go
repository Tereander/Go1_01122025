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

// Ваш код

import "fmt"

func main(){
	var input, summ int
	var expression string

	fmt.Print("Enter number: ")
	fmt.Scan(&input)

	for digit := 0; input > 0; input /= 10 {
		digit = input % 10

		summ += digit

		if len(expression) == 0 {
			expression = fmt.Sprintf("%d = ", digit)
		} else {
			expression = fmt.Sprintf("%d + ", digit) + expression
		}
	}
	fmt.Printf("Sum: %v\n", summ)
	fmt.Printf("Expression: %s%v\n", expression, summ)
}

