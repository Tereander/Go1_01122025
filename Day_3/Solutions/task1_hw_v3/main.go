/*
Задача №1.
Программа получает на вход последовательность из 5 целых чисел.
Вам нужно определить вид последовательности:
  - возрастающая
  - убывающая
  - случайная
  - постоянная

В качестве ответа следуют выдать прописными латинскими буквами тип последовательности:
1. ASCENDING (строго возрастающая)
2. WEAKLY ASCENDING (нестрого возрастающая, то есть неубывающая)
3. DESCENDING (строго убывающая)
4. WEAKLY DESCENDING (нестрого убывающая, то есть невозрастающая)
5. CONSTANT (постоянная)
7. RANDOM (случайная)
Примеры входных и выходных данных:
In: 11 9 4 2 -1 Out: DESCENDING
In: 3 8 8 11 12 Out: WEAKLY ASCENDING
In: 2 -1 7 21 1 Out: RANDOM
In: 5 5 5 5 5 Out: CONSTANT
Подсказка: используем метод строки strings.Split()
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your sequence like \\\\\\x y z q w e r t y\\\\\\")
    reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	// input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
	sequence := strings.Split(input, " ")
	var seq_desc = [...]string{
		"", "DESCENDING", "CONSTANT", "WEAKLY DESCENDING", "ASCENDING", 
		"RANDOM", "WEAKLY ASCENDING", "RANDOM",
	}
								
	fmt.Printf("%s\n", sequence)
	var a, b, c int
	for idx := 0; idx < len(sequence)-1; idx++ {
		a, _ = strconv.Atoi(string(sequence[idx]))
		b, _ = strconv.Atoi(string(sequence[idx+1]))
		if a > b {
			c |= 1 << 0 // 1
		}
		if a == b {
			c |= 1 << 1 // 2
		}
		if a < b {
			c |= 1 << 2 // 4
		}
	} 

	fmt.Printf("Your sequence is %s\n", seq_desc[c])
}
