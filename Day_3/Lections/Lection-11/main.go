/*
Слайсы
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// обычно так
	var array [3]int

	array[0] = 100
	array[1] = 200
	array[2] = 300

	fmt.Printf("Array's values: %v, array as is: %#v\n", array, array)

	const a = 6
	b := 6
	fmt.Println("a and b:", a, b)

	// При создании массива необходимо определить размер на этапе компиляции
	var array1 [a]int // Так можно
	// var array2 [b]int // Так нельзя

	fmt.Printf("Array's values: %v, array as is: %#v\n", array1, array1)

	// Слайсы (они же срезы)
	// 1. Слайс - это динамическая обвязка над массивом
	startArr := [4]int{10, 20, 30, 40}
	fmt.Printf("Array's values: %v, array as is: %#v\n", startArr, startArr)
	var startSlice []int = startArr[0:2] // Слайс инициализируется пустыми квадратными скобками
	fmt.Println("Slice[0:2]:", startSlice)
	// Создали слайс на основе уже существующего массива

	// 2. Создание слайса без явной инициализации массива
	secondSlice := []int{15, 20, 25, 40, 50}
	fmt.Println("secondSlice:", secondSlice)

	// 3. Изменение элементов среза
	originArr := [...]int{30, 40, 50, 60, 70, 80}
	thirdSlice := originArr[1:4] // Это набор ссылок на элементы нижележащего массива
	for i := range thirdSlice {
		thirdSlice[i]++
	}
	fmt.Println("originArr:", originArr)
	fmt.Println("thirdSlice:", thirdSlice)

	// 4. Один массив и два производных слайса
	fSlice := originArr[:]
	sSlice := originArr[2:5]
	fmt.Println("Before changing. originArr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After changing. originArr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)

	// 5. Срез как встроенный тип
	/*
		type slice struct {
			Length int
			Capacity int
			ZeroElement *byte
		}
	*/

	// 6. Длина и ёмкость массива
	wordSlice := []string{"one", "two", "three"}
	fmt.Println("wordSlice:", wordSlice, "Length:", len(wordSlice), "Capacity:", cap(wordSlice))
	// Добавляем новый элемент
	wordSlice = append(wordSlice, "four")
	fmt.Println("wordSlice:", wordSlice, "Length:", len(wordSlice), "Capacity:", cap(wordSlice))

	// 7. Capacity(cap) или ёмкость слайса - это значение, которое показывает сколько итого элементов 
	// можно хранить в слайсе, не выделяя дополнительной памяти
	// Если нет места для нового элемента, то выделяется новая память, равная n,
	// где n - это размер ёмкости до изменения, по общей формуле: n * 2
	// cap = 3: 3 -> 6 -> 12 -> ...
	numerics := []int{1, 2}
	for i := 0; i < 200; i++ {
		if i % 5 == 0 {
			fmt.Println("Current len:", len(numerics), "Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
	}

	// Важно. После выделения памяти под новый массив, ссылки со старого будут перенесены в новый.
	// Пример
	numArr := []int{1, 2}
	numSlice := numArr[:]

	// Пример работы с внутренним устройством слайса
	arrayAddress0 := uintptr(unsafe.Pointer(&numArr[0]))
	fmt.Printf("numArr address: 0x%x\n", arrayAddress0)
	arrayAddress1 := uintptr(unsafe.Pointer(&numSlice[0]))
	fmt.Printf("numSlice address: 0x%x\n", arrayAddress1)

	numSlice = append(numSlice, 3) // В этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 0

	// Пример работы с внутренним устройством слайса
	arrayAddress2 := uintptr(unsafe.Pointer(&numSlice[0]))
	fmt.Printf("new numSlice address: 0x%x\n", arrayAddress2)
	numSlicePtr := unsafe.SliceData(numSlice)
	fmt.Printf("pointer: %p\n", numSlicePtr)

	fmt.Println("numArr:", numArr)
	fmt.Println("numSlice:", numSlice)
	
	// Добавление первого элемента через append
	var numSliceChanged []int // nil
	numSliceChanged = append(numSliceChanged, 3) 
	fmt.Println("numSliceChanged:", numSliceChanged)
	ptr := unsafe.SliceData(numSliceChanged)
	fmt.Printf("pointer: %p, %v\n", ptr, ptr == nil)

	// 7. Как создавать слайсы наиболее эффективно
	// make() - это функция, которая позволяет более детально создавать срезы
	sl := make([]int, 10, 15)

	// []int - тип коллекции
	// 10 - длина
	// 15 - ёмкость
	// Сначала инициализируется arr [15]int
	// Затем по нему делается срез arr[0:10]
	// После этого, он возвращается
	fmt.Println("slice:", sl, "Length:", len(sl), "Capacity:", cap(sl))

	// 8. Добавление элементов в срез
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	anotherWords := []string{"four", "five", "six"}
	myWords = append(myWords, anotherWords...)
	myWords = append(myWords, "seven", "eight")
	fmt.Println("myWords after changing:", myWords)

	// 9. Многомерный срез (срез срезов)
	twoDimSlice := [][]int{
		{1, 2},
		{1, 2, 3, 4},
		{10, 20, 30},
		{},
	}
	fmt.Printf("Values of slice: %v, slice as is: %#v\n", twoDimSlice, twoDimSlice)

	// 10. Слайсы рун
	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	myStrFromRunes := string(runeHexSlice)
	fmt.Println("Runes:", runeHexSlice, "myStr:", myStrFromRunes)

	// 10.1 Руны как литералы
	runeLiteralSlice := []rune{'V', 'a', 's', 't'}
	myStrFromLiteralSlice := string(runeLiteralSlice)
	fmt.Println("Runes:", runeLiteralSlice, "myStr:", myStrFromLiteralSlice)

	// 10.2 Вычисление ёмкости строки бессмысленно, т.к. строка - это базовый тип
	fmt.Println(cap([]rune("Вася")))

	// 10.3 Строки - неизменяемые(unmutable), а слайсы рун - изменяемые(mutable)
	firstName := "Anna"
	fmt.Println("firstName:", firstName)
	// firstName[0] = "B" // error

	firstSliceName := []rune("Петя")
	firstSliceName[0] = 'Ф'
	fmt.Println("slice:", firstSliceName, "slice as string after change:", string(firstSliceName))

	// 10.4 Syntax sugar (можно использовать десятичное представление байтов)
	myDecimalByteSlice := []byte{100, 101, 102, 103}
	myDecimalStr := string(myDecimalByteSlice)
	fmt.Println("myStr:", myDecimalStr)
}