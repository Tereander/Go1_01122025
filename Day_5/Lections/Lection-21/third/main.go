package main

import (
	"fmt"
)

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (st Student) Describe() {
	fmt.Printf("%s and %d years old.\n", st.name, st.age)
}

func (st Student) GetName() {
	fmt.Println("Name:", st.name)
}

func TypeFinder(i any) {
	// type assertion вместе с switch
	// Присваиваем переменной v значение, лежащее под преполагаемым интерфейсом
	switch v := i.(type) {
	case string:
		fmt.Println("This is string.")
	case int:
		fmt. Println("This is int.")
	case Describer: // Вывод - с типом интерфейса можно
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}


func FindStudent(i any) {
	// type assertion вместе с switch
	// Присваиваем переменной v значение, лежащее под преполагаемым интерфейсом
	switch v := i.(type) {
	case Student:
		fmt.Println("I'm Student.")
		v.GetName() // Используем метод, который есть только у типа Student
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	std := Student{"Fedya", 20}
	TypeFinder(10)
	TypeFinder("hello")
	TypeFinder(std)
	TypeFinder(nil)
	var desc Describer
	desc = Student{"Mark", 25}
	FindStudent(desc)
	FindStudent("hello students")
}
