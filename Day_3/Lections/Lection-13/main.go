package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Map - это набор ключ:значение
	// Инициализация пустой мапы
	mapper := make(map[string]int)
	fmt.Println("Show empty map:", mapper)

	// 2. Добавление элемента в мапу
	mapper["Petr"] = 32
	mapper["Anna"] = 41
	fmt.Println("After adding new elemens(pairs):", mapper)

	// 3. Инициализация мапы с указанием пар
	newMapper := map[string]int{
		"Alice": 32,
		"Bob": 123,
	}
	newMapper["Nikolay"] = 99
	fmt.Println("After adding new elemens(pairs):", newMapper)

	// 4. Что может быть ключом в мапе?
	// 4.1 ВАЖНО: Мапа в Golang не упорядочена
	// 4.2 Ключом в мапе может быть только сравнимый тип (==, !=)

	// 5. Нулевое значени для map
	// var mapZeroValue map[string]int // mapZeroValue == nil
	// mapZeroValue["Alice"] = 12 // error

	// 6. Получение значения из мапы
	testPerson := "Alice"
	// 6.1 Получение значения, которое есть в мапе
	fmt.Println("Salary of Alice:", newMapper[testPerson])
	// 6.2 Получение значения, которого нет в мапе
	testPerson = "Петя"
	fmt.Println("Salary of Петя:", newMapper[testPerson]) // 0, ошибки нет
	fmt.Println("newMapper:", newMapper)

	// 7. Проверка нахождения ключа
	employee := map[string]int{
		"Denis": 0,
		"Vasya": 0,
		"Petya": 0,
	}

	// 7.1 При обращении по ключу, возвращается два значения
	if value, ok := employee["Denis"]; ok { // ;ok -> ok == true
		fmt.Println("Denis exists as key and it's value:", value)
	} else {
		fmt.Println("Denis does not exist in map")
		
	}

	if value, ok := employee["Roman"]; ok { // ;ok -> ok == false
		fmt.Println("Roman exists as key and it's value:", value)
	} else {
		fmt.Println("Roman does not exist in map")
		
	}

	// 8. Перебор значений мапы
	fmt.Println(strings.Repeat("=", 30))

	for key, value := range employee {
		fmt.Printf("Key: %v and value: %v\n", key, value)
	}

	// 9. Как удалять пары из мапы
	// 9.1 Удаление существующей пары
	fmt.Println("Before deleting item:", employee)
	delete(employee, "Vasya")
	fmt.Println("After deleting item:", employee)

	// 9.2 Удаление несуществующей пары
	if _, ok := employee["Alice"]; ok {
		delete(employee, "Alice") // очень дорогая опрация
	}

	// 10. Количество пар == длина мапы
	fmt.Println("Pairs amount in map:", len(employee))

	// 11. Map(как и слайс) - это ссылочный тип
	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}
	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "One")

	fmt.Printf("words: %v and it's pointer: %p\n", words, &words)
	fmt.Printf("newWords: %v and it's pointer: %p\n", newWords, &newWords)


	// 12. Сравнение мап
	// 12.1 Сравнение массивов (массив можно использовать как ключ в мапе)
	arr := [3]int{1, 2, 3}
	if [3]int{1, 2, 3} == arr {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	// 12.2 Сравнение слайсов
	// Из-за того, что слайсы - это ссылочный тип, его можно сравнить только с nil
	// if []int{1, 2, 3} == []int{1, 2, 3} {
	// 	fmt.Println()
	// }

	// 12.3 Сравнение мап
	// Из-за того, что мапы - это ссылочный тип, его можно сравнить только с nil
	aMap := map[string]int{
		"a": 1,
	}
	var bMap map[string]int

	if aMap == nil {
		fmt.Println("Zero value map")
	}

	if bMap == nil {
		fmt.Println("Zero value map: bMap")
	}

	// Следствие
	// Если слайс или мапа являются составляющими какой-либо структуры,
	// структура автоматически не сравнима!
}	