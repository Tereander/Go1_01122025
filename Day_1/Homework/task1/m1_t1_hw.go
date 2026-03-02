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

import (
	"fmt"
)

// стоимость бенза
const fuelPrice float64 = 48.0

var distance float64
var consumption float64

func main() {
	// стоимость бенза

	fmt.Println("Расчет стоимости поездки")
	fmt.Println("______________________________")

	//вводим растояние и проверяем
	for {
		fmt.Println("Введите растояние от 50 до 10 000 км")
		fmt.Scan(&distance)

		if 50 < distance && distance < 10000 {
			break
		}
		fmt.Println("Ошибка! Растояние должно быть от 50 до 10 000")
	}

	//вводим расход топлива и проверяем
	for {
		fmt.Println("Введите расход топлива от 5л./100 км до 25л./100 км")
		fmt.Scan(&consumption)

		if 5 < consumption && consumption < 100 {
			break
		}
		fmt.Println("Ошибка! Расход должен быт ьв диапазоне от 5 до 25 литров на сотку")
	}
	// Расчет стоимости поездки
	totalFuel := (distance / 100) * consumption
	totalCost := totalFuel * fuelPrice

	// печатаем в консоль
	fmt.Printf("\nРезультаты расчета:")
	fmt.Printf("Расстояние: %.2f км\n", distance)
	fmt.Printf("Расход топлива: %.2f л/100км\n", consumption)
	fmt.Printf("Цена бензина: %.2f руб/л\n", fuelPrice)
	fmt.Printf("Всего топлива: %.2f л\n", totalFuel)
	fmt.Printf("Стоимость поездки: %.2f руб\n", totalCost)
}
