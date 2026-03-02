package main

import "fmt"

// Ваш код
func main() {
	const petrol_price float32 = 48.0
	var distance, gas_mileage uint16

	fmt.Print("Введите пробег, км: ")
	fmt.Scan(&distance)
	fmt.Printf("Вы ввели пробег %d км\n", distance)

	if distance < 50 || distance > 10000 {
		fmt.Print("Пробег должен быть в диапазоне 50-10000 км\n")
		return
	}

	fmt.Print("Введите расход топлива, литров \\ 100 км: ")
	fmt.Scan(&gas_mileage)
	fmt.Printf("Вы ввели расход топлива, %d литров \\ 100 км\n", gas_mileage)

	if gas_mileage < 5 || gas_mileage > 25 {
		fmt.Print("Расход топлива быть в диапазоне 5-25 литров \\ 100 км\n")
		return
	}

	fmt.Printf("Цена 1 литра бензина сооставляет %.2f рублей \\ литр\n", petrol_price)

	journey_cost := float32(distance) / 100 * float32(gas_mileage) * petrol_price

	fmt.Printf("Цена всей поездки сооставляет %.2f рублей\n", journey_cost)
}