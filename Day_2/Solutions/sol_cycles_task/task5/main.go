/*
3.5 В бесконечном цикле приложение запрашивает у пользователя числа.

	Ввод завершается, как только пользователь вводит число "-1".
	После завершения ввода приложение выводит сумму чисел.
	Используем конструкцию:
	for {
	// body
	}
*/
package main

import "fmt"

func main(){
    var sum, value int
    for {
        fmt.Print("Enter a number but -1 is quit. ")
        fmt.Scan(&value)
        if value == -1 {
            break
        }
        sum += value
    }
    fmt.Println("sum:", sum)
}