/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/
package main

// Ваш код
import "fmt"

func main() {
	var input int

	fmt.Printf("Enter 4 digit code: ")
	fmt.Scan(&input)

	d1, d2, d3, d4 := input/1000%10, input/100%10, input/10%10, input%10

	if d1 == d4 && d2 == d3 {
		fmt.Printf("%4d палиндром\n", input)
	} else {
		fmt.Printf("%4d не палиндром\n", input)
	}
}
