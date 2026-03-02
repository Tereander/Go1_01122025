package main

import (
	"fmt"
	"go1_01122025/Day_5/Lections/Lection-23/rectangle"
)


// 1. Два файлы помещены в одну директорию (находятся в одном пакете, main)


func main() {
	// данные видны компилятору
	resAdd := add(4, 5)
	fmt.Println("resAdd:", resAdd)

	resSub := Sub(9, 3)
	fmt.Println("resSub:", resSub)

	// Если имя сущности написано с маленькой буквы, то ее нельзя передать за пределы пакеты.
	// Если имя сущности написано с большой буквы, то ее можно экспортировать за пределы пакета


	r := rectangle.New(10, 20, "green")
	fmt.Println("Perimeter:", r.Perimeter())
}