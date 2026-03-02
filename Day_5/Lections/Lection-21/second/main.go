package main

import "fmt"

// 1. А что если создать интерфейс, в котором вообще нет требований к поведению?
type Empty interface{}  // сейчас есть алиас any

// 2. Вопрос в том, кто удовлетворяет этому интерфейсу? Ответ - кто угодно

// 3. Реализуем функцию, которая может принимать что угодно
func Describer(pretendent any) {
	fmt.Printf("Type: %T and value %v\n", pretendent, pretendent)
}

// 4. Type assertion
func SemiGeneric(pretendent any) {
	value, ok := pretendent.(int) // Проверка того, что pretendent типа int
	fmt.Println("Value:", value, "OK?", ok)
}

type Student struct {
	name string
}


func main() {
	// 5. Передача параметров без предварительного присваивания и с.
	Describer("Hello world")

	intSample := 200
	Describer(intSample)

	Describer(Student{"Vasya"})

	// 6. Пример работы с функцией SemiGeneric
	SemiGeneric(10)
	SemiGeneric(Student{"Petr"})
	SemiGeneric("Hello world")
}