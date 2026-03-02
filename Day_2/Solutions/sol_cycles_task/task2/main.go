/*
3.2 Доработать предыдущую задачу так, чтобы выводились только числа,

	делящиеся на 5 без остатка.
*/
package main

import "fmt"


func main(){
	var start, end, index int

	fmt.Printf("Задача3.2\nВведи берега: ")
	fmt.Scan(&start, &end)

	if start > end {
		fmt.Println("Берега попутал, братик, но ничего")
		start, end = end, start
	}
	fmt.Println("1-st variant")
	for index = start + 1; index < end; index++ {
		if index % 5 == 0 {
			fmt.Printf("%d ", index)
		}
	}
	fmt.Println()

	// Variant 2
	fmt.Println("2-nd variant")
	for index = (start / 5 + 1 )  * 5; index < end; index += 5 {
		fmt.Printf("%d ", index)
	}
	fmt.Println()
}