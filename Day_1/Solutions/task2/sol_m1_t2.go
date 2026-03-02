// Задача №2
// Вход: трехзначное число. 
// Найти первую, вторую и последнюю цифры заданного трехзначного числа.
package main

import "fmt"

func main() {
		// Variant 1
        var num, dig1, dig2, dig3 int

        fmt.Println("Enter 3-digits number: ")
        fmt.Scan(&num)

        dig1 = num / 100
        dig2 = num / 10 % 10
        dig3 = num % 10

        fmt.Println("Digits:", dig1, dig2, dig3)

		// Variant 2
		var input uint16

		fmt.Printf("Enter 3 digit code: ")
		fmt.Scan(&input)

		d1, d2, d3 := input / 100, input / 10 % 10, input % 10
		fmt.Printf("First digit %d, Second digit %d, last digit %d\n", d1, d2, d3)
}