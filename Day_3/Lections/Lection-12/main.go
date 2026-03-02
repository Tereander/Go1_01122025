package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var name string

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		// Команда захвата потока ввода и сохранение в буфер
		// (захват идёт до символа окончания строки)
		name = input.Text()
		// Метод Text() возвращает элементы, помещенные в буфер в виде строки(string)
	}
	fmt.Println(name)

	// Вариант с циклом
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is", result)
		}
	}

	// Преобразование строкового литерала к какому-нибудь числовому типу
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr) // Atoi -> Anything to Int (именно int)
	fmt.Printf("Value is %v and type of %T\n", numInt, numInt)

	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	numFloat32, _ := strconv.ParseFloat(numStr, 32)
	// Но это 64-разрядное число будет ГАРАНТИРОВАННО без проблем приводится к float32
	fmt.Println(numInt64, numFloat32)
	fmt.Printf("Value: %.3f and %T\n", numFloat32, float32(numFloat32))

	// Пример с bufio.NewReader
	fmt.Println("bufio.NewReader example.")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("line as is %#v\n", line)
	line = strings.Trim(line, "\n")
	fmt.Printf("After trim line as is %#v\n", line)
	data := strings.Fields(line) // разделим исходную строку на слайс строк по пробелу
	fmt.Printf("data value: %v and data as is: %#v\n", data, data)

}
