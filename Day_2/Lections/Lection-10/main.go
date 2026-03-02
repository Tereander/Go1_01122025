package main


import "fmt"


func main() {
	// Побитовые операции
	// & - bit and
	// | - bit or
	// ^ - bit xor: 1 ^ 0 = 1, 0 ^ 1 = 1, 1 ^ 1 = 0
	// >> - сдвиг вправо
	// << - сдвиг влево
	// Syntax sugar: &=, |=, ^=

	// Examples
	a := 5  // bin: 0101
	b := 10 // bin: 1010
	c := a & b // 0000
	d := a | b // 1111
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("5 ^ 3 =", 5 ^ 3) // 101 ^ 011 = 110 (6)
	fmt.Println("4 ^ 7 =", 4 ^ 7) // 100 ^ 111 = 011 (3)
}
