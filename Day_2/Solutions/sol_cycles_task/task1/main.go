/*
3.1 Пользователь вводит числа a и b (b больше a).

	Вывести все целые числа, расположенные между ними.
*/
package main

import "fmt"

func main() {
	var a, b int
	// Multiline string, используем backticks, символы ``
	fmt.Print(
		`Задача 3.1
Введите первое число: `)
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Println("Все целые числа, расположенные между ними:")

	for i := a + 1; i < b; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
