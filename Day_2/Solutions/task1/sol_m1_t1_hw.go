/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*Проверка условий по желанию
*/
package main

import "fmt"

func main() {
        const price float32 = 48
        var distance, rate, cost float32

        fmt.Println("Введите расстояние от 50 до 10000 км:")
        fmt.Scan(&distance)
        fmt.Println("Введите расход бензина в литрах (5-25 литров) на 100км:")
        fmt.Scan(&rate)
        cost = price * distance * rate / 100
        fmt.Printf("Стоимость поездки: %.2f рублей.\n", cost)
}