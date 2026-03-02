package main

/*
	Встраиваемые и вложенные структуры
*/

import (
	"fmt"
)

// 1. Вложенные структуры
// Это использование одной структуры, как тип поля в другой структуре
type University struct {
	age int
	yearBased int
	infoShort string
	infoLong string
}

type Student struct {
	firstName string
	lastName string
	university University 
}

// 4. Пример встроенной структуры
type Professor struct {
	firstName string
	lastName string
	age int
	greatWorks string
	University // встраивание, в этом месте происходит добавление все полей структуры University
}

func main() {
	// 2. Создание экземпляра структур с вложением
	// Имена(тип нужен) либо порядок(тип не нужен)
	student := Student{
		firstName: "Anna",
		lastName: "Ivanova",
		university: University{
			yearBased: 1991,
			infoShort: "good University",
			infoLong: "It's very good University",
		},
	}

	// 3. Получение доступа к полям вложенной структуры
	fmt.Println("FirstName:", student.firstName)
	fmt.Println("lastName:", student.lastName)
	fmt.Println("yearBased of Uni:", student.university.yearBased)
	fmt.Println("infoShort of Uni:", student.university.infoShort)

	// 5. Создание экземпляра встроенной структуры
	professor := Professor {
		firstName: "Anatoly",
		lastName: "Smirnov",
		age: 39,
		greatWorks: "Ultimate Go programming",
		// papers map[string]string - добавление этого поля делает структуру несравнимой
		University: University{
			yearBased: 1974,
			infoShort: "short info",
			age: 2025 - 1974,
			infoLong: "long info",
		},
	}

	// 6. Получение доступа к полям встроенной структуры
	fmt.Println("FirstName:", professor.firstName)
	fmt.Println("lastName:", professor.lastName)
	fmt.Println("yearBased of Uni:", professor.yearBased)  // получили доступ через University
	fmt.Println("infoShort of Uni:", professor.infoShort)
	fmt.Println("Age of professor:", professor.age) // 39, Получим доступ к полю ВЫШЕСТОЯЩЕЙ структуры
	fmt.Println("Age of Uni:", professor.University.age) // 51

	// 7. Сравнение экземпляров
	profLeft := Professor{}
	profRight := Professor{}
	fmt.Println(profLeft == profRight) // true
	fmt.Println(profLeft == professor) // false

	// 8. Если хотя бы одно из полей структуры несравнимо - то и вся структура тоже несравнима!
}