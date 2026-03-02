package main

import "fmt"

/*
	Enums - перечисления (iota)
*/

func main() {
	const (
		var0 = iota // старт с 0
		var1 // 1
		var2 // 2
		var3 // 3
	)
	fmt.Println(var0, var1, var2, var3)

	const (
		_ = iota // старт с 0
		param1 = 1 + iota // 1 + 1 = 2
		param2 = 10 // 10
		param3 // не указан ни тип, ни значение, поэтому смотрим на param2 и получаем 10
		param4 = iota // 4
		param5 // 5
	)
	fmt.Println(param1, param2, param3, param4, param5)
	
	const (
		value1 = 1 << iota  // 1 << 0 = 1
		value2 // 1 << 1 = 2
		value3 // 4
		value4 // 8
	)

	fmt.Println(value1, value2, value3, value4)
}