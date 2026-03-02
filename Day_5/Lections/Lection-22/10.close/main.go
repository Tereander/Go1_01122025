package main

import "fmt"


// 1. Закрытие каналов и итерирование
// Со стороны получателя можно использовать следующий синтаксис
// value, ok := <- ch
// где, value - это значение, помещенное в канал отправителем
// ok - true/false в зависимости от того, открыл канал или закрыт отправителем
// Если канал закрыт, то в value будет записано zeroValue для типа канала

func generator(ch chan int) {
	for i := range 20 {
		ch <- i
	}
	close(ch)
	// ch <- 100 // panic
}

func main() {
	ch := make(chan int)
	go generator(ch)

	// for {
	// 	value, ok := <- ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("Received from channel:", value)
	// }

	// Конструкцию можно упростить с помощью итерирования по каналу:
	for value := range ch {
		fmt.Println("Received from channel:", value)
	}
}
