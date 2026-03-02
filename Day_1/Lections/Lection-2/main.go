package main

import "fmt"

func main() {
	// Пример однострочного комментария. Ctrl + /
	// var <variable_name> <variable_type> + auto initialization
	var age int
	fmt.Println("My age is", age)

	// var b int = nil // compile-error
	// fmt.Print(b)

	// Декларирование и инициализация пользовательским значением
	var height int = 183 // + manual initialization
	fmt.Println("My height is:", height)

	// В чем "полустрогость" типизации?
	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid)

	// Множественное присваивание через :=
	aArg, bArg := 10, 30
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg)
	// aArg, bArg := 10, 30 // error
	// fmt.Println(aArg, bArg)

	// Добавляем новую переменую
	bArg, cArg := 300, 400
	fmt.Println(bArg, cArg)

	// Константы
	const price, tax float32 = 275.00, 27.50 // Типизированные константы
	const quantity, inStock = 2, true // Нетипизированные константы
	fmt.Println("Total:", 2 * quantity * (price + tax))

	// Variables
	var cost float32 = 275.00
	var total float32 = 27.50
	cost = 300
	fmt.Println(cost + total)

	// Error block
	// var value = 275.00
	// var taxes float32 = 27.50
	// fmt.Println(value + taxes)

	// Work block
	var value = 275.00
	var taxes float32 = 27.50
	fmt.Println(value + float64(taxes))

	// Короткое присваивание
	result := false
	value, new_value := 3.12, 121 // value exists but new_value in new
	fmt.Println(value, new_value, result)

	// Ввод данных
	var (
		number int
		s string
	)

	fmt.Scan(&number, &s)
	fmt.Println(number, s)
}
