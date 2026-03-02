/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример: 
вход: 346, выход: 643
вход: 120, выход: 021
вход: 100, выход: 001
*/
package main

import "fmt"

func main() {
        var number, dig1, dig2, dig3 int

        fmt.Println("Введите трехзначное число:")
        fmt.Scan(&number)

        dig1 = number / 100
        dig2 = number / 10 % 10
        dig3 = number % 10

        fmt.Printf("Реверсная запись: %d%d%d\n", dig3, dig2, dig1)
}