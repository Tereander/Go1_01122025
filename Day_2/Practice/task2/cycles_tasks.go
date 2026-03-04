/*
=======
Задачи:
=======

3.1 Пользователь вводит числа a и b (b больше a).

	Вывести все целые числа, расположенные между ними.

3.2 Доработать предыдущую задачу так, чтобы выводились только числа,

	делящиеся на 5 без остатка.

3.3 Пользователь вводит число. Вывести таблицу умножения на это число (от 1 до 10)

3.4 В цикле получать от пользователя оценки по четырём экзаменам.

	Вывести сумму набранных им баллов.
	Функцию fmt.Scan() в коде используем только один раз.

3.5 В бесконечном цикле приложение запрашивает у пользователя числа.

	Ввод завершается, как только пользователь вводит число "-1".
	После завершения ввода приложение выводит сумму чисел.
	Используем конструкцию:
	for {
	  // body
	}

3.6 Получить от пользователя натуральное число, посчитать сумму цифр в нём.

	Решить с помощью индексов, т.е. работаем с числом, как со строкой.

3.7 Вводим строку без знаков препинания(5 слов).

	Найти самое длинное слово в строке и вывести сколько в нем букв.

Пример:
In: Скажи как дела в учебе
Out: учебе, 5

In: Закрепляем циклы в языке Golang
Out: Закрепляем, 10
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func task31() {
	var a, b int

	fmt.Print("Введите числа a и b (b > a)")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	if a > b {
		fmt.Print("Ну я же просил чтобы b было больше a....")
		return
	}

	fmt.Printf("Числа между %d и %d: ", a, b)
	for i := a + 1; i < b; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func task32() {
	var a, b, j int

	fmt.Print("Введите числа a и b (b > a)")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	if a > b {
		fmt.Print("Ну я же просил чтобы b было больше a....")
		return
	}

	fmt.Printf("Числа между %d и %d, делящиеся на 5: ", a, b)

	j = 0 // счетчик цикла
	for i := a + 1; i < b; i++ {
		if i%5 == 0 {
			fmt.Print(i, " ")
			j += 1
		}
	}
	if j == 0 {
		fmt.Print("Таких чисел нет :(")
	}
	fmt.Println()
}

func task33() {
	var n int

	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	fmt.Printf("Таблица умножения на %d:\n", n)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%2d * %2d = %3d\n", i, n, i*n)
	}
}

func task34() {
	var grade int
	total := 0

	fmt.Print("Введите оценки по 4 экзаменам")

	for i := 1; i <= 4; i++ {
		fmt.Printf("Экзамен %d: ", i)
		_, err := fmt.Scan(&grade)
		if err != nil {
			return
		}
		total += grade
	}

	fmt.Printf("Сумма баллов: %d\n", total)
}

func task35() {
	var num int
	sum := 0

	fmt.Println("Вводите числа (для завершения введите -1):")
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}

		if num == -1 {
			break
		}
		sum += num
	}

	fmt.Printf("Сумма введенных чисел: %d\n", sum)
}

func task36() {
	var numStr string
	sum := 0
	fmt.Print("Введите натуральное число: ")
	_, err := fmt.Scan(&numStr)
	if err != nil {
		return
	}

	fmt.Print("Сумма цифр: ")
	for i := 0; i < len(numStr); i++ {
		digit := int(numStr[i] - '0')
		sum += digit

		fmt.Printf("%c", numStr[i])
		if i < len(numStr)-1 {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

func task37() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите 5 слов: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Fields(input)

	mostLongestWord := ""
	maxLength := 0

	for _, word := range words {
		runeCount := len([]rune(word))
		if runeCount > maxLength {

			maxLength = runeCount
			mostLongestWord = word
		}
	}

	fmt.Printf("%s, %d\n", mostLongestWord, maxLength)
}

func main() {
	// Вызываем все задачи по порядку

	//fmt.Println("Задача 3.1: Числа между a и b")
	//task31()
	//
	//fmt.Println("\nЗадача 3.2: Числа, делящиеся на 5")
	//task32()
	//
	//fmt.Println("\nЗадача 3.3: Таблица умножения")
	//task33()
	//
	//fmt.Println("\nЗадача 3.4: Сумма оценок")
	//task34()

	fmt.Println("\nЗадача 3.5: Бесконечный цикл до -1")
	task35()

	fmt.Println("\nЗадача 3.6: Сумма цифр (строка)")
	task36()
	// fmt.Scanln()
	fmt.Println("\nЗадача 3.7: Самое длинное слово")
	task37()

}
