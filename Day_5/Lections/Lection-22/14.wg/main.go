package main

/*
	WaitGroup
*/

// 1. Еще один инструмент для управления горутинами - это WaitGroup
// По сути WaitGroup - это счетчик горутин
// Когда горутина запускается, выполняется WaitGroup++
// Когда горутина завершается, выполняется WaitGroup--
// Таким образом, когда значение WaitGroup == 0, все горутины отработали!

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // WaitGroup--
}

func main() {
	no := 5
	var wg sync.WaitGroup
	for i := range no {
		wg.Add(1) // WaitGroup++
		go process(i, &wg)
	}
	wg.Wait() // if WaitGroup == 0 ? До тех пор пока условие не выполнено
	// мы будем заблокированы в данной строке для горутины main()
	fmt.Println("All goroutines have finished.")
}