package main

/*
	Структуры (struct)
*/

import (
	"fmt"
)

// 1. Структура - заименованный набор полей(состояний), определяющий новый тип данных

// 2. Определение структуры(явное определение)
type Student struct {
	firstName string
	lastName  string
	age       int
}

// 3. Если имеется ряд состояний одинакового типа, то можно перечислить их в одной строке
type AnotherStudent struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}

func PrintStudent(stud Student) {
	fmt.Println("=======================")
	fmt.Println("FirstName:", stud.firstName)
	fmt.Println("LastName:", stud.lastName)
	fmt.Println("Age:", stud.age)
	fmt.Println("=======================")
}

// 11. Структура с анонимными полями
type Human struct {
	firstName string
	lastName  string
	string
	int
	bool
}

func main() {
	// 4. Создание представителя (экземпляра) структуры
	// Порядок полей не важен, есть имена
	student := Student{
		firstName: "Vasya",
		age:       21,
		lastName:  "Petrov",
	}
	PrintStudent(student)

	student2 := Student{"Petr", "Vasechkin", 19} // Порядок указания свойств важен, такой же как в структуре
	PrintStudent(student2)

	// 5. Что если не все поля определены при создании экземпляра?
	student3 := Student{
		firstName: "Fedya",
	}
	PrintStudent(student3)

	// 6. Анонимные структуры
	anonStudent := struct {
		age           int
		groupID       int
		professorName string
	}{
		age:           23,
		groupID:       2,
		professorName: "Ivan",
	}
	fmt.Printf("AnonStudent: %v and as is: %#v\n", anonStudent, anonStudent)

	// 7. Доступ к полям и их модификация
	studMark := Student{"Mark", "Twen", 19}
	fmt.Println("FirstName:", studMark.firstName)
	fmt.Println("LastName:", studMark.lastName)
	fmt.Println("Age:", studMark.age)
	studMark.age += 2
	fmt.Println("New age:", studMark.age)

	// 8. Инициализация пустой структуры
	emptyStudent1 := Student{}
	var emptyStudent2 Student
	PrintStudent(emptyStudent1)
	PrintStudent(emptyStudent2)

	// 9. Указатели на экземпляры структуры
	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Sidorov",
		age:       23,
	}
	fmt.Printf("Value of studPointer: %v and address: %p\n", studPointer, studPointer)

	secondPointer := studPointer
	(*secondPointer).age += 20
	fmt.Println("Value of studPointer after changing:", studPointer)
	studPointerNew := new(Student)
	fmt.Println(studPointerNew)

	// 10. Работа с доступом к полям через указатель
	fmt.Println("Age via (*...).age:", (*studPointer).age)
	// Неявно происходит разыменование  и обращение к соответствующему полю
	fmt.Println("Age via .age:", studPointer.age)

	// 12. Создание экземпляра структуры с анонимными полями
	human := Human{
		firstName: "Nikolay",
		lastName: "Svetlov",
		string: "Additional information",
		int: -1,
		bool: true,
	}
	fmt.Println(human)
	fmt.Println("Anon field string:", human.string)
	fmt.Println("Anon field int:", human.int)
	fmt.Println("Anon field bool:", human.bool)
}
