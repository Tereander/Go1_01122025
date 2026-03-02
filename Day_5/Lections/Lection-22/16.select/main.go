package main

import "fmt"

// 1. select - это инструмент защиты от deadlock
// Добавление default в select'e страхует нас от появления deadlock'a и берет на себя всю работу
// смотрим пример с default'ом и без


func main() {
	var ch chan string

	select {
	case v := <- ch:
		fmt.Println("value:", v)
	default:
		fmt.Println("default case executed.")
	}
}