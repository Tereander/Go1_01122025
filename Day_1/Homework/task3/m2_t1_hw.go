/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1 9 2
Выход: 1 2 9

Про sort мы пока ничего не знаем!
Это задача на условный оператор
*/
package main

import (
	"fmt"
)

func main() {
	var x, y, z int

	fmt.Println("Введите три числа для сортировки: ")
	_, err := fmt.Scan(&x, &y, &z)

	/*
		Хочу вот так:
		var testArr []string
		_, err := fmt.Scan(&testArr)
		fmt.Printf("Тестовый массив: %v\n", testArr)
	*/

	if err != nil {
		return
	}

	numbers := []int{x, y, z}

	var sortArr []int
	sortArr = bubbleSort(numbers)
	fmt.Printf("Отсортированный массив: %v\n", sortArr)
}

func bubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
