package main


import (
	"fmt"
)

func main() {
	// Switch
	var price int
	fmt.Scan(&price)

	// В switch-case запрещены дублирующие условия в case'ах
	switch price {
	case 100: 
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("Third case")
	default:
		// если ни один из вышеперечисленных case'ов не отработал, отработает default
		fmt.Println("Default case")
	}

	// Case с множеством варинтов
	var ageGroup string = "u" // Возрастные группы: "a", "b", "c", "d", ...
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("First group")
	case "d", "e":
		fmt.Println("Second group")
	default:
		fmt.Println("Third group")
	}

	// Case с выражениями
	var age int
	fmt.Scan(&age)
	switch {
	case age <= 18:
		fmt.Println("Too young")
	case age > 18 && age < 40:
		fmt.Println("Middle group")
	case age >= 40 && age <= 100:
		fmt.Println("Too old")
	default:
		fmt.Println("Who are you?")
	}


	// Case с проваливаниями. Проваливания выполняют даже ЛОЖНЫЕ КЕЙСЫ.
	// В момент проваливания fallthrough у следующего кейса не проверяется условие,
	// a сразу выполняется тело.
	var number int
	fmt.Scan(&number)

outer:
	switch {
	case number < 100:
		fmt.Printf("%d less than 100\n", number)
		if number % 2 == 0 {
			break outer
		}
		fallthrough
	case number > 200:
		fmt.Printf("%d more than 200\n", number)
		fallthrough
	case number > 1000:
		fmt.Printf("%d more than 1000\n", number)
		fallthrough
	default:
		fmt.Println("Default:", number)
	}
	
	// Нехорошая конструкция
	var i int
outloop:
	for {
		fmt.Scan(&i)
		switch {
		case i % 2 == 0:
			fmt.Printf("Value i is %d and it's even\n", i)
			break outloop
		}
	}
	fmt.Println("THE END.")
}