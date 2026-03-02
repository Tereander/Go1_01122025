package main

import (
	"fmt"
	"runtime"
	"time"
)

func newGoroutine() {
	fmt.Println("Hello, I'm new Goroutine")
}


func main() {
	go newGoroutine() 
	fmt.Printf("Number of active goroutines: %d\n", runtime.NumGoroutine())
	// немного тормозим выполнение main(), для того, чтобы дать время отработать newGoroutine()
	time.Sleep(1 * time.Second) 
	fmt.Printf("After Sleep, number of active goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println("Main goroutine work...")
}