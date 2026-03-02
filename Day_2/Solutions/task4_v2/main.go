package main

import "fmt"

func main() {
        var number, dig1, dig2, dig3, dig4 int

        fmt.Println("Введите 4-х значное число:")
        fmt.Scan(&number)

        dig1 = number / 1000
        dig2 = number / 100 % 10
        dig3 = number / 10 % 10
        dig4 = number % 10

        if dig1 == dig4 && dig2 == dig3 {
                fmt.Println("Палиндром")
                return
        }
        fmt.Println("Не палиндром")
}