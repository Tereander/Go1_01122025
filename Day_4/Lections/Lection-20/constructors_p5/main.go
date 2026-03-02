package main

/*
	Конструктор
*/

import (
	"fmt"
)

// 1. Создадим тип Rectangle
type Rectangle struct {
	length, width int
}

// 2. Создадим конструктор для типа Rectangle
// Для имен конструкторов в Go договорились использовать функции со следующим названием:
// * New() - если данных конструктор на файл один (в файле присутствует только одна структура)
// * NewStructName - если в файле присутствуют разные структуры

// 3. В Go принято возвращать из конструктора не сам экземпляр, а ссылку на него, т.е. указатель
func NewRectangle(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

// 4. Добавим пару методов
func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := NewRectangle(10, 20)
	fmt.Printf("Type of r: %T and value: %v and as is: %#v\n", r, r, r)
	fmt.Println("Perimeter:", r.Perimeter())
	fmt.Println("Area:", r.Area())
}