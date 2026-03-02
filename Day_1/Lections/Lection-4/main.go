package main

import "fmt"


func main() {

	fmt.Printf("Binary: %b\\%b\n", 4, 5)

	fmt.Printf("%d %%\n", 50) 

	array := []int64{0, 1}
	fmt.Printf("array value: %v\n",array)
	fmt.Printf("array value as is: %#v\n",array)
	fmt.Printf("type value: %T\n",array)

	s := "café"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)

	fmt.Printf("\\caf\u00e9\n")
}
