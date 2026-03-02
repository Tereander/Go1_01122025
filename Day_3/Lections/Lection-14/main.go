// Показать документацию пакета

/*
	DRY - don't repeat yourself
	1. Явная функция - это локально-определенный блок кода, который имеет имя (явное определение)
	Функцию нужно определить(создать) + функцию нужно вызвать(calling)
	2. Сигнатура функций и их определение
	Шаблон:
	func functionName(params type) typeReturnValue {
		function body
	}
*/
package main


import (
	"fmt"
)	

func main(){
	fmt.Println("call function")
	// 4. Вызов функции
	res := add(5, 8)
	fmt.Println("Result of add(5, 8) is", res)

	multRes := mult(1, 2, 3, 4)
	fmt.Println("Result of mult(1, 2, 3, 4) is", multRes)

	per, sq := rectangleParameters(10, 20)
	fmt.Println("Perimeter is", per, "and square is", sq)

	onlyPerimeter, _ := rectangleParameters(15, 30)
	fmt.Println("Perimeter is",onlyPerimeter)

	perNamed, squareNamed := namedReturn(10, 20)
	fmt.Println("Perimeter is", perNamed, "and square is", squareNamed)

	fmt.Println(funcWithReturn(14, 9))

	emptyReturn(100)

	helloVariadic()
	helloVariadic(70)
	helloVariadic(20, 50, 67, 89)

	checkStrings(3, 5)
	checkStrings(2, 7, "hello", "age", "golang", "program")

	resSum1 := sumVariadic(10, 20, 30, 40)
	fmt.Println("resSum1:", resSum1)

	sliceNumber := []int{5, 7, 8, 9}
	fmt.Println(sumSlice(sliceNumber))

	resSum2 := sumVariadic(sliceNumber...)
	fmt.Println("resSum2:", resSum2)

}

// 3. Простейшая функция (определить функцию можно как до момента ее вызова в функции main),
// так и в любом месте пакета, главное, чтобы она была определана в принципе где-нибудь

// add function return sum of two numbers
func add(a int, b int) int {
	result := a + b
	return result
}

// 5. Функция с однотипными параметрами
func mult(a, b, c, d int) int {
	return a * b * c * d  // можно указать и выражение
}

// 6. Возврат больше, чем одного значения (ReturnType1, ReturnType2, ...)
func rectangleParameters(a, b int) (int32, int) {
	perimetr := int32(2 * (a + b))
	square := a * b
	return perimetr, square
}

// 7. Именованный возврат значений
func namedReturn(a, b int) (perimetr int32, square int) {
	perimetr = int32(2 * (a + b))
	square = a * b
	return // не нужно указывать возвращаемый значения
}

// 8. При вызове оператора return функция прекращает свое выполнение и что-то возвращает
func funcWithReturn(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}

	if a == b {
		return a, true
	}

	return 0, false
}

// 9. Что вернёт функция в случае, если return не указан(или он пустой)
func emptyReturn(a int) {
	fmt.Println("I'm emptyReturn with parameters...", a)
}

// 10. Variadic parameters (континуальные параметры)
func helloVariadic(a ...int) {
	fmt.Printf("value is %v and type is %#v\n", a, a)
}

// 11. Смешивание параметров с variadic
// 	1. Континуальный параметр всегда самый последний
//	2. Variadic parametr - один на всю функцию
func checkStrings(a, b int, words ...string) {
	fmt.Println("Parametr a:", a)
	fmt.Println("Parametr b:", b)
	fmt.Printf("value of words: %v and type: %#v\n", words, words)

	var counter int
	for _, word := range words {
		if len(word) > a && len(word) < b {
			fmt.Println("Needed word", word)
			counter++
		}
	}
	fmt.Println("Counter:", counter)
}

// 12. Передача слайса
func sumVariadic(nums ...int) int {
	var sum int
	for _, number := range nums {
		sum += number
	}
	return sum
}

func sumSlice(nums []int) int {
	var sum int
	for _, number := range nums {
		sum += number
	}
	return sum
}