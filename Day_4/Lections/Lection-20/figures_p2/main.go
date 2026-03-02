package main

import (
	"fmt"
	"math"
)

// 4. В чем преимущество методов над функциями
// Во-первых: наличие методов повышает читаемость кода ("консистентность")
// Во-вторых: в Go запрещается создавать функции с одинаковыми названиями, в то время
// как методы для различных структур могут иметь одинаковые имена.

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

// 5. Создадим функцию и метод Perimeter для структуры Circle
func PerimeterCircle(c Circle) float64 {
	return 2 * c.radius * math.Pi
}

// Это метод, (c Circle - это получатель метода) 
func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

// 6. Создадим функцию и метод Perimeter для структуры Rectangle
func PerimeterRectangle(r Rectangle) int {
	return 2 * (r.length + r.width)
}

// Это метод, (r Rectangle - это получатель метода) 
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	c := Circle{10.5}
	fmt.Println("Call function for Circle:", PerimeterCircle(c))
	fmt.Println("Call method for Circle:", c.Perimeter())

	r := Rectangle{10, 20}
	fmt.Println("Call function for Rectangle:", PerimeterRectangle(r))
	fmt.Println("Call method for Rectangle:", r.Perimeter())
}