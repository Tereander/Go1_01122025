package main

import (
	"fmt"
)

// 1. Каналы - средство для коммуникаций между горутинами.
// Каналы можно рассматривать как соединительные трубы,
// через которые горутины общаются между собой.

// 2. Объявление канала
// Каналы по умолчанию имеют zeroValue - nil. Поэтому их создают с помощью make

func main() {
	var a chan int // объявляем канал, через который будут передаваться данные типа int
	if a == nil {
		fmt.Println("channel is nil. Let's define it.")
		a = make(chan int)
		fmt.Printf("Type of a: %T and as is: %#v\n", a, a)
	}
}
