package main

import "fmt"

func main() {
	var input int

	fmt.Printf("Enter 3 digit code: ")
	fmt.Scan(&input)

	d1, d2, d3 := input/100%10, input/10%10, input%10
	result := d3 * 100 + d2 * 10 + d1

	fmt.Printf("Result: %03d\n", result)
}