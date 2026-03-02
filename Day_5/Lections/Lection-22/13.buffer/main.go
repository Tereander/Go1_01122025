package main

import (
	"fmt"
	"time"
)

// 1. Как блокируется буферизированный канал?
func write(ch chan int) {
	fmt.Println("Write goroutine, Cap:", cap(ch))
	for i := range 5 {
		ch <- i
		fmt.Println("Successfully wrote", i, "to channel.")
	}
	close(ch)
}

// 2. Как только буфер канала заполняется, он блокируется до тех пор,
// пока не будет освобождено место (буфер может быть освобожден не до конца!)

func main() {
	ch := make(chan int, 2)
	fmt.Println("Cap:", cap(ch))
	go write(ch)
	time.Sleep(2 * time.Second) 
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

// 3. Длина и вместимость канала.
// Длина канала len(ch) - сколько сейчас элементов в канале
// Вместимость канала cap(ch) - какой размер буфера у канала