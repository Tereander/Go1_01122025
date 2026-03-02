package main

import (
	"fmt"
	"time"
)

// 1. Запустим несколько горутин и посмотрим, как они бьются за ресурсы

// 2. Определим первую горутину
func printEvenNumbers() {
	for i := 1000; i < 1020; i += 2 {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

// 3. Определим вторую горутину
func printOddNumbers() {
	for i := 1; i < 20; i += 2 {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

// 4. Определим main горутину
func main() {
	go printEvenNumbers() // сразу идём дальше, запуск функции будет происходить в отдельной горутине
	go printOddNumbers() //  также идем дальше
	time.Sleep(7000 * time.Millisecond) // тормозим основную горутину для отработки вышезапущенных
	fmt.Println("Main goroutine died.")
}

// 5. Таким образом, горутины работают следующим образом: через context switch