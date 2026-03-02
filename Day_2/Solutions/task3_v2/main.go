package main

// Ваш код
import "fmt"

func main() {
	var input1, input2, input3, temp int

	fmt.Printf("Enter 3 numbers: ")
	fmt.Scan(&input1, &input2, &input3)

	if input1 > input2 {
		temp = input1
		input1 = input2
		input2 = temp
	}

	if input2 > input3 {
		input2, input3 = input3, input2
	}

	if input1 > input2 {
		temp = input1
		input1 = input2
		input2 = temp
	}

	fmt.Printf("Numbers: %d %d %d\n", input1, input2, input3)

}
