/*
3.3 Пользователь вводит число. Вывести таблицу умножения на это число (от 1 до 10)
*/
package main

import "fmt"

func main() {
	var value int

	fmt.Printf("Задача 3.3\nВведи число: ")
	fmt.Scan(&value)

	for index := 1; index <= 10; index++ {
		fmt.Printf(" %d * %d = %d\n", value, index, index*value)
	}
}