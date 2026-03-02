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

import "fmt"

func main() {
        var numbers [5]int
        var asc, desc, cnst int

        fmt.Println("Введите последовательность из 5 целых чисел через пробел.")
        fmt.Scanln(&numbers[0], &numbers[1], &numbers[2], &numbers[3], &numbers[4])

        num_prev := numbers[0]
        for i := 1; i < 5; i++ {
			if numbers[i] == num_prev {
				cnst++
			} else if numbers[i] > num_prev {
				asc++
			} else if numbers[i] < num_prev {
				desc++
			}
			num_prev = numbers[i]
        }

        fmt.Printf("ASC: %d, DESC: %d, CONST: %d\n", asc, desc, cnst)

        switch {
        case asc > 0 && desc == 0 && cnst == 0:
            fmt.Println("ASCENDING")
        case asc > 0 && desc == 0 && cnst > 0:
            fmt.Println("WEAKLY ASCENDING")
        case asc == 0 && desc > 0 && cnst == 0:
            fmt.Println("DESCENDING")
        case asc == 0 && desc > 0 && cnst > 0:
            fmt.Println("WEAKLY DESCENDING")
        case asc > 0 && desc > 0 && cnst >= 0:
            fmt.Println("RANDOM")
        case asc == 0 && desc == 0 && cnst > 0:
            fmt.Println("CONSTANT")
        default:
            fmt.Println("UNKNOWN")
        }
}
