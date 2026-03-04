/*
Задача №2

Вход:
Пользователь должен ввести "правильный пароль", состоящий из:
цифр, букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что пароль правильный и он принят.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8

Реализацию оформить через функцию:
1. checkPassword(pass string) (bool, errors <- на усмотрение)
2. main() // для интерактива
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	digits    = "0123456789"
	lowercase = "abcdefghiklmnopqrstvxyz"
	uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
	special   = "_!@#$%^&"
)

func checkPassword(pass string) (bool, string) {
	if len(pass) < 8 {
		return false, fmt.Sprintf(
			"Пароль слишком короткий: %d символов. Минимум 8",
			len(pass),
		)
	}
	if len(pass) > 15 {
		return false, fmt.Sprintf(
			"Пароль слишком длинный: %d символов. Максимум 15",
			len(pass),
		)
	}
	// Флаги для проверки наличия символов из каждого набора
	hasDigit := false
	hasLower := false
	hasUpper := false
	hasSpecial := false

	// итерация по каждому символу

	for _, ch := range pass {
		charStr := string(ch)

		switch {
		case strings.Contains(digits, charStr):
			hasDigit = true
		case strings.Contains(lowercase, charStr):
			hasLower = true
		case strings.Contains(uppercase, charStr):
			hasUpper = true
		case strings.Contains(special, charStr):
			hasSpecial = true
		default:
			return false, fmt.Sprintf(
				"Недопустимый символ: '%s'. Разрешены только цифры, латинские буквы и %s",
				charStr, special,
			)
		}
	}

	// Проверка на все наборы
	if !hasDigit {
		return false, "В пароле должна быть хотя бы одна цифра"
	}
	if !hasLower {
		return false, "В пароле должна быть хотя бы одна строчная буква"
	}
	if !hasUpper {
		return false, "В пароле должна быть хотя бы одна заглавная буква"
	}
	if !hasSpecial {
		return false, fmt.Sprintf("В пароле должен быть хотя бы один специальный символ из: %s", special)
	}

	// Все проверки пройдены
	return true, ""
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	maxAttempts := 5
	var password string

	fmt.Println("Проверка пароля")
	fmt.Printf("Требования к паролю:\n")
	fmt.Printf("1. Длина: от 8 до 15 символов\n")
	fmt.Printf("2. Должен содержать: цифры, строчные буквы, заглавные буквы и спецсимволы %s\n", special)
	fmt.Printf("3. Максимальное количество попыток: %d\n\n", maxAttempts)

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("Попытка %d из %d\n", attempt, maxAttempts)

		if input.Scan() {
			password = input.Text()
		}
		fmt.Println(password)

		isValid, errorMsg := checkPassword(password)

		if isValid {
			fmt.Println("\nПароль правильный! Он принят.")
			fmt.Printf("Ваш пароль: %s\n", password)
			return
		} else {
			fmt.Printf("Ошибка: %s\n", errorMsg)

			if attempt == maxAttempts {
				fmt.Println("\nПревышено максимальное количество попыток. Доступ заблокирован.")
				return
			} else {
				fmt.Printf("Осталось попыток: %d\n\n", maxAttempts-attempt)
			}

		}

	}

}
