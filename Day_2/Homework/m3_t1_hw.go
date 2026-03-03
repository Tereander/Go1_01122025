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
In: 11 9 4 2 -1  Out: DESCENDING
In: 3 8 8 11 12  Out: WEAKLY ASCENDING
In: 2 -1 7 21 1  Out: RANDOM
In: 5 5 5 5 5     Out: CONSTANT

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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите последовательность из 5 чисел: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")

	numbersArr := make([]int, len(parts))

	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		numbersArr[i] = num
	}

	result := determineSequenceType(numbersArr)
	fmt.Println(result)

}

func determineSequenceType(nums []int) string {
	hasIncrease := false
	hasDecrease := false
	hasEqual := false

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			hasIncrease = true
		} else if nums[i] > nums[i+1] {
			hasDecrease = true
		} else if nums[i] == nums[i+1] {
			hasEqual = true
		}
	}
	if hasIncrease && !hasDecrease && !hasEqual {
		return "ASCENDING"
	} else if hasIncrease && !hasDecrease && hasEqual {
		return "WEAKLY ASCENDING"
	} else if !hasIncrease && hasDecrease && !hasEqual {
		return "DESCENDING"
	} else if !hasIncrease && hasDecrease && hasEqual {
		return "WEAKLY DESCENDING"
	} else if !hasIncrease && !hasDecrease && hasEqual {
		return "CONSTANT"
	} else {
		return "RANDOM"
	}
}
