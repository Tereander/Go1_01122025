package main

import (
	"fmt"
)

// 1. Созададим чуть более полезную программу
// Есть число, например 456
// (4^2 + 5^2 + 6^2) + (4^3 + 5^3 + 6^3) 
// сложим суммы квадратов и кубов каждого разряда числа
// Алгоритм
// * main goroutine получает число и вызывает две других горутины, получает результаты, складывает их и выводит на консоль
// * squaresGoroutine считает квадраты и записывает в канал
// * cubesGoroutine считает кубы и записывает в канал

func squaresGoroutine(num int, squareChan chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	squareChan <- sum
}

func cubesGoroutine(num int, squareChan chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	squareChan <- sum
}

func main() {
	number := 839472994
	squareChan := make(chan int)
	cubesChan := make(chan int)
	go squaresGoroutine(number, squareChan)
	go cubesGoroutine(number,cubesChan)
	squareSum, cubesSum := <- squareChan, <- cubesChan
	fmt.Println("Total result is:", squareSum + cubesSum)
}
