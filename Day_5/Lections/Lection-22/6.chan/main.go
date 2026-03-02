package main

import (
	"fmt"
	"time"
)

// 1. Немного изменим последнюю программу, чтобы лучше увидеть как работает процесс блокировки

func newGoroutine(done chan bool) {
	fmt.Println("Hello, I'm Goroutine and I'm going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("newGoroutine awaked and going to send data to channel")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("I'm main goroutine and going to call newGoroutine")
	go newGoroutine(done)
	<- done // в это точке выполнение main горутины останавливается до получения данных
	fmt.Println("I'm main and received data and going to die.")
}

