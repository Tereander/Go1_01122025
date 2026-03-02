package main

import (
	"fmt"
	"time"
)

// 1. select - это инструмент, позволящий выбирать из множества канальных операций(чтение/запись) для множества каналов
// Если из 10 каналов что-то пришло в один из них - select выбирает его
// Если из 10 каналов что-то пришло сразу в два и более - select выбирает случайный


func server1(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func server3(ch chan string) {
	time.Sleep(3500 * time.Millisecond)
	ch <- "from server3"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2: // select выбирает этот case, потому что он отработает быстрее
		fmt.Println(s2)
	}

	// Выполняем какие-то действия до получения данных из канала
	ch := make(chan string)
	go server3(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <- ch:
			fmt.Println("received value:", v)
			return
		default:
			fmt.Println("no value received")	
		}
	}
}