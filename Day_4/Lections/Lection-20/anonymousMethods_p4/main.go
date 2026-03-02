package main

import (
	"fmt"
)

type University struct {
	city string
	name string
}

// 1. Данный метод будет связан только со структурой University
func (u *University) FullInfoUniversity() {
	fmt.Printf("Uni name: %v and City: %v\n", u.name, u.city)
}

// 2. В структуру Professor встроены поля University (все методы тоже переходят)
type Professor struct {
	fullName string
	age int
	University
}

func (p Professor) Info() {
	fmt.Println("Fullname:", p.fullName)
	fmt.Println("Age:", p.age)
}

func main() {
	professor := Professor {
		fullName: "Mikhail Nikolaev",
		age: 25,
		University: University{
			city: "Moscow",
			name: "BMSTU",
		},
	}

	// 3. Вызываем метод структуры University через экземпляр профессора
	professor.FullInfoUniversity()

	// 4. Вызываем "свой" метод
	professor.Info()
}