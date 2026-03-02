package main

import (
	"fmt"
	"strings"
	// "strings"
)

func main() {
	// Шаблон цикла for как for
	// for init; condition; post {
	// init - блок инициализации переменных цикла
	// condition - условие (если условие верно, цикл продолжается, если - нет, цикл завершается)
	// post - изменение переменной цикла (инкрементарное или декрементарное действие)
	// }
	fmt.Println("First cycle.")
	for i := 0; i <= 5; i++ {
		fmt.Println("Current value of i:", i)
	}

	// Важный момент. В качестве init можно использовать ТОЛЬКО краткую запись :=

	// break - команда, которая прерывает текущее выполнение цикла
	// и передает управление следующим за циклом инструкциям
	fmt.Println("Second cycle.")
	for i := 11; i < 22; i++ {
		if i > 15 {
			break
		}
		fmt.Println("Current value of i:", i)
	}
	fmt.Println("After for loop with BREAK.")

	// continue - команда, которая прерывает текущее выполнение цикла 
	// и передает управление следующей итерации цикла
	fmt.Println("Third cycle.")
	for i := 105; i <= 120; i++ {
		if i >= 110 && i <= 115 {
			continue
		}
		fmt.Println("Current value of i:", i)
	}
	fmt.Println("After for loop with CONTINUE.")

	// Вложенные циклы и лейблы
	for i := 0; i < 10; i++ { // for i := range 10 {
		for j := 0; j <= i; j ++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("Выше должен быть треугольник.")

	// Бывают ситуации, когда нужно прервать сразу оба цикла
	// Здесь помогут лейблы. Лейблы - это синтаксический сахар(syntax sugar) для оператора goto
	fmt.Println("Label example.")

outer:
	for i := 0; i < 2; i++ {
		for j := 0; j <= 3; j ++ {
			fmt.Printf("i: %d and j: %d and sum of i+j: %d\n", i, j, i+j)
			if i + j > 3 {
				break outer // Хочу, чтобы все прекратили работу, даже внешний
			}
		}
	}
	fmt.Println("After outer stop.")

	// Шаблон цикла for как while (модификация for)
	// 1. Классический while do
	var loopVar int = 0
	// while (loopVar < 10) {
	// ...
	// loopVar ++
	// }
	fmt.Println("for like while.")
	for loopVar < 10 { // по факту это -> for ;loopVar < 10; {}
		fmt.Println("In while like loop:", loopVar)
		loopVar++
	}

	// loopVar - переменная цикла, а не функции main
	// for loopVar := 0; loopVar < 10; {
	// 	fmt.Println("In while like loop:", loopVar)
	// 	loopVar++
	// }
	// fmt.Println(loopVar) // 0

	// 2. Классический бесконечный цикл
	var password string

outer2:
	for { // the same as for ;; {}
		fmt.Println("Enter a password:")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password. Try again")
		} else {
			fmt.Println("Password accepted")
			break outer2
		}
	}

	// 3. Цикл с множественными перемеными цикла
	for x, y := 0, 1; x < 5 && y < 8; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}

	// 4. x1 - это переменная main, а не цикла
	x1 := 5
	for x1 = 10; x1 < 15; x1 += 1 {
		fmt.Println(x1)
	}
	fmt.Println("Result:", x1)
}
