package main

import "fmt"

// 1. Решим задачу про сумму квадратов и кубов с помощью закрытия каналов

func digits(num int, dataChan chan int) {
	for num != 0 {
		digit := num % 10
		dataChan <- digit
		num /= 10
	}
	close(dataChan)
}

func calcSquares(num int, squareCh chan int){
	sum := 0
	dch := make(chan int)
	go digits(num, dch) // отдельная горутина для генерации цифр
	for digit := range dch {
		sum += digit * digit
	}
	squareCh <- sum
	close(squareCh)

}

func calcCubes(num int, cubesCh chan int){
	sum := 0
	dch := make(chan int)
	go digits(num, dch) // отдельная горутина для генерации цифр
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubesCh <- sum
	close(cubesCh)
}

func main() {
	number := 839472994
	squareChan := make(chan int)
	cubesChan := make(chan int)
	go calcSquares(number, squareChan)
	go calcCubes(number,cubesChan)
	squareSum, cubesSum := <- squareChan, <- cubesChan
	fmt.Println("Total result is:", squareSum + cubesSum)
}
