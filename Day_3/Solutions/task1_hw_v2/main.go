package main


// Ваш код
import "fmt"

func main(){
	var input [5]int
	var step int
	var index int

	for index = 0; index < 5; index ++ {
		fmt.Printf("Enter %d-s value: ", index)
		fmt.Scan(&input[index])
	}

	var rising, wrising, descent, wdescent, constant bool = true, true, true, true, true
	
	step = input[0]

	for index = 1; index < 5; index++ {
		if step != input[index] {
			constant = false
		}

		if step >= input[index] {
			rising = false
		}

		if step > input[index] {
			wrising = false
		}
		
		if step < input[index] {
			descent = false
		}

		if step <= input[index] {
			wdescent = false
		}
		
		step = input[index]
	}

	if rising {
		fmt.Print("ASCENDING ")
	}

	if wrising {
		fmt.Print("WEAK ASCENDING ")
	}

	if descent {
		fmt.Print("DESCENDING ")
	}

	if wdescent {
		fmt.Print("WEAK DESCENDING ")
	}

	if constant {
		fmt.Print("CONSTANT ")
	}

	if (rising || wrising || descent || wdescent || constant) == false {
		fmt.Print("RANDOM ")
	}
	fmt.Println()
}
