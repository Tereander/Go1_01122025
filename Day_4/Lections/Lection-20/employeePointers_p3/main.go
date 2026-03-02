package main

import (
	"fmt"
)

type Employee struct {
	name string
	salary int
}

// CRUD(Create Read(Get) Update(Set) Delete)
// 1. Метод, в котором получатель копируется и в его теле происходит работа с локальной копией
func (e Employee) SetName(newName string) {
	e.name = newName 
	fmt.Println("Show new name:", e) // видим, что изменения есть
}

// 2. Метод в котором получатель передается по ссылке и в теле метода работаем с ссылкой на экземляр
func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

// 4. Используйте методы с PointerReceiver'ом в ситуациях когда:
// 1) Изменения в работе метода над экземляром дожны быть видны на вызывающей стороне
// 2) Когда экземпляр "тяжелый", т.е. копировать его "дорого" с точки зрения расходов ресурсов
// 3) С ними может работать любой вид экземпляра

func main() {
	e := Employee {"Aleksey", 3000}
	fmt.Println("Before setting new name:", e)
	e.SetName("Yan")
	fmt.Println("After setting new name:", e) // изменений нет


	// 3. Вызов метода у ссылки на сотрудника
	fmt.Println("Before setting new salary:", e)
	e.SetSalary(5000)
	// 5. Аналогично явному вызову: (&e).SetSalary(5000)
	fmt.Println("After setting new salary:", e)

}
