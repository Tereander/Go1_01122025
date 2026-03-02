package main

import "fmt"

// 1. Каналы могут иметь направление
// То, что мы использовали до этого, было ДВУНАПРАВЛЕННЫЙ канал
// (в него можно писать и из него можно читать).
// Можно создать канал только на отправку:
// sendChan := make(chan<- int)

// Можно создать канал только на чтение
// sendChan := make(<-chan int)

func sender(sendChan chan<- int) {
	sendChan <- 10 // Тут всё OK
}

func main() {
	sendChan := make(chan<- int)
	go sender(sendChan)
	// <- sendChan // тут не ок, т.к. канал только на отправку данных, но не на чтение

	intCh := make(chan int, 2)
	go factorial(5, intCh)
	fmt.Println(<-intCh)
	fmt.Println("The End")

}

// 2. Использование однонаправленных каналов никак не сказывается на производительности,
// а служит лишь для логического разделения кода

func factorial (n int, ch chan<- int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result // Только отправка (запись)
}