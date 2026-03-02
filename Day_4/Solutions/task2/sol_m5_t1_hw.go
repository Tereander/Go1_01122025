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

// Ваш код
import (
	"fmt"
	"strings"
)

const strdigits string = "0123456789"
const strlowercase string = "abcdefghijklmnopqrstuvwxyz"
const struppercase string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const strspecial string = "_!@#$%^&"

const passwordlenmin int = 8
const passwordlenmax int = 15

/* Было так:
const flagdigits int = 1
const flaglowercase int = 2
const flaguppercase int = 4
const flagspecial int = 8
const flagshort int = 16
const flaglong int = 32
*/

// Стало так:
const (
	flagDigits int = 1 << iota
	flagLowercase
	flagUppercase
	flagSpecial
	flagShort
	flagLong
)

func main() {
	var password string
	var retryleft int = 5

	fmt.Println(`Введите новый пароль состоящий из:
цифр, строчных и прописных букв латинского алфавита и 
специальных символов - _!@#$%^&`)

	for ; retryleft > 0; retryleft-- {
		fmt.Printf("Введите новый пароль. У вас осталось %d попыток:\n", retryleft)
		fmt.Scan(&password)

		errors := checkPassword(password)

		if errors == 0 {
			fmt.Println("Хороший пароль")
			break
		}

		fmt.Println("Слабый пароль:")

		if errors&flagShort != 0 {
			fmt.Println(" - пароль должен содержать минимум 8 символов,")
		}
		if errors&flagLong != 0 {
			fmt.Println(" - пароль должен содержать не более 15 символов,")
		}
		if errors&flagDigits != 0 {
			fmt.Println(" - пароль должен содержать хотябы одну цифру 0-9,")
		}
		if errors&flagUppercase != 0 {
			fmt.Println(" - пароль должен содержать хотябы одну заглавную букву A-Z,")
		}
		if errors&flagLowercase != 0 {
			fmt.Println(" - пароль должен содержать хотябы одну строчную букву a-z,")
		}
		if errors&flagSpecial != 0 {
			fmt.Println(" - пароль должен содержать хотябы один специальный символ,")
		}
	}

	if retryleft == 0 {
		fmt.Println("Вы исчерпали все попытки.")
	}
}

func checkPassword(password string) (errorflags int) {

	var countdigits, countlowercase, countuppercase, countspecial int

	// Length check
	if len(password) < passwordlenmin {
		errorflags += flagShort
	}

	if len(password) > passwordlenmax {
		errorflags += flagLong
	}

	// Alphabet check
	for _, symbole := range password {
		if strings.ContainsRune(strdigits, symbole) {
			countdigits++
		}
		if strings.ContainsRune(strlowercase, symbole) {
			countlowercase++
		}
		if strings.ContainsRune(struppercase, symbole) {
			countuppercase++
		}
		if strings.ContainsRune(strspecial, symbole) {
			countspecial++
		}
	}

	if countdigits == 0 {
		errorflags += flagDigits
	}
	if countlowercase == 0 {
		errorflags += flagLowercase
	}
	if countuppercase == 0 {
		errorflags += flagUppercase
	}
	if countspecial == 0 {
		errorflags += flagSpecial
	}

	return
}
