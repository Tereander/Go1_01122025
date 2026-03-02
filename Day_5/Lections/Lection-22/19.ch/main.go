package main

import (
	"fmt"
	"sync"
)

// также мы может разрешить состояник гонки через каналы,
// т.к. каналы более детальный иструмент коммуницирования

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<- ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1) // buffer channel, нет блокировки
	for range 1000 {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("Final value of x:", x)
}
