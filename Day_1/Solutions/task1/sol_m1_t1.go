// Задача №1
// Программа запрашивает имя пользователя и возраст
// Нужно вывести имя и возраст за вычетом одного года
package main

import "fmt"

const minus = 1


func main() {
	var (
		name string // ""
		age int // 0
	)

	fmt.Println("Enter your name: ")
	counter, _ := fmt.Scan(&name) // create new variable: counter
	fmt.Println(counter)
	fmt.Println("Enter your age: ")
	counter, _ = fmt.Scan(&age)
	fmt.Println(counter)
	fmt.Printf("Your name is %s.\nYour age is %d.\n", name, age - minus)
}