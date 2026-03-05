/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес: индекс(ровно 6 цифр), город, улица, дом, квартира

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

Реализовать конструктор и несколько методов у типа "Накладная"

Пример:
invoice = NewInvoice()

или

order = NewOrder()

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Address Структура для адреса
type Address struct {
	Index     string
	City      string
	Street    string
	House     string
	Apartment string
}

// Invoice Основная структура - Накладная
type Invoice struct {
	ProductName  string
	Quantity     int
	CustomerName string
	Phone        string
	Address      Address
}

// NewInvoice возвращаем указатель
func NewInvoice() *Invoice {
	return &Invoice{}
}

// Метод для проверки, является ли строка только буквами
func (i *Invoice) isOnlyLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && r != ' ' {
			return false
		}
	}
	return true
}

// Метод для проверки, является ли строка только цифрами
func (i *Invoice) isOnlyDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// InputData Метод для ввода данных (с валидацией)
func (i *Invoice) InputData() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("СОЗДАНИЕ НАКЛАДНОЙ")
	fmt.Println(strings.Repeat("-", 40))

	// Ввод названия товара
	for {
		fmt.Print("Наименование товара (1-100 символов): ")
		scanner.Scan()
		product := strings.TrimSpace(scanner.Text())

		if product == "" {
			fmt.Println("Название товара не может быть пустым!")
			continue
		}

		if len([]rune(product)) > 100 {
			fmt.Println("Название товара слишком длинное!")
			continue
		}

		i.ProductName = product
		break
	}

	// Ввод количества
	for {
		fmt.Print("Количество: ")
		scanner.Scan()
		quantityStr := strings.TrimSpace(scanner.Text())

		if quantityStr == "" {
			fmt.Println("Количество не может быть пустым!")
			continue
		}

		var quantity int
		_, err := fmt.Sscanf(quantityStr, "%d", &quantity)
		if err != nil || quantity <= 0 {
			fmt.Println("Введите корректное положительное число!")
			continue
		}

		i.Quantity = quantity
		break
	}

	// Ввод ФИО
	for {
		fmt.Print("ФИО получателя (только буквы): ")
		scanner.Scan()
		name := strings.TrimSpace(scanner.Text())

		if name == "" {
			fmt.Println("ФИО не может быть пустым!")
			continue
		}

		if !i.isOnlyLetters(name) {
			fmt.Println("ФИО должно содержать только буквы!")
			continue
		}

		i.CustomerName = name
		break
	}

	// Ввод телефона
	for {
		fmt.Print("Телефон (10 цифр): ")
		scanner.Scan()
		phone := strings.TrimSpace(scanner.Text())

		// Очищаем от возможных символов форматирования
		phone = strings.ReplaceAll(phone, "+", "")
		phone = strings.ReplaceAll(phone, "-", "")
		phone = strings.ReplaceAll(phone, "(", "")
		phone = strings.ReplaceAll(phone, ")", "")
		phone = strings.ReplaceAll(phone, " ", "")

		if phone == "" {
			fmt.Println("Телефон не может быть пустым!")
			continue
		}

		if !i.isOnlyDigits(phone) {
			fmt.Println("Телефон должен содержать только цифры!")
			continue
		}

		if len(phone) != 10 {
			fmt.Println("Телефон должен содержать ровно 10 цифр!")
			continue
		}

		i.Phone = phone
		break
	}

	// Ввод адреса
	fmt.Println("\nАДРЕС ДОСТАВКИ:")

	// Индекс
	for {
		fmt.Print("  Индекс (6 цифр): ")
		scanner.Scan()
		index := strings.TrimSpace(scanner.Text())

		if index == "" {
			fmt.Println("Индекс не может быть пустым!")
			continue
		}

		if !i.isOnlyDigits(index) {
			fmt.Println("Индекс должен содержать только цифры!")
			continue
		}

		if len(index) != 6 {
			fmt.Println("Индекс должен содержать ровно 6 цифр!")
			continue
		}

		i.Address.Index = index
		break
	}

	// Город
	for {
		fmt.Print("  Город: ")
		scanner.Scan()
		city := strings.TrimSpace(scanner.Text())

		if city == "" {
			fmt.Println("Город не может быть пустым!")
			continue
		}

		if !i.isOnlyLetters(city) {
			fmt.Println("Город должен содержать только буквы!")
			continue
		}

		i.Address.City = city
		break
	}

	// Улица
	for {
		fmt.Print("  Улица: ")
		scanner.Scan()
		street := strings.TrimSpace(scanner.Text())

		if street == "" {
			fmt.Println("Улица не может быть пустой!")
			continue
		}

		i.Address.Street = street
		break
	}

	// Дом
	for {
		fmt.Print("  Дом: ")
		scanner.Scan()
		house := strings.TrimSpace(scanner.Text())

		if house == "" {
			fmt.Println("Дом не может быть пустым!")
			continue
		}

		i.Address.House = house
		break
	}

	// Квартира
	for {
		fmt.Print("  Квартира: ")
		scanner.Scan()
		apt := strings.TrimSpace(scanner.Text())

		if apt == "" {
			fmt.Println("Квартира не может быть пустой!")
			continue
		}

		if !i.isOnlyDigits(apt) {
			fmt.Println("Квартира должна содержать только цифры!")
			continue
		}

		i.Address.Apartment = apt
		break
	}
}

// DisplayInfo Метод для отображения информации
func (i Invoice) DisplayInfo() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("НАКЛАДНАЯ НА ОТПРАВКУ ЗАКАЗА")
	fmt.Println(strings.Repeat("=", 50))

	fmt.Printf("Товар: %s\n", i.ProductName)
	fmt.Printf("Количество: %d шт.\n", i.Quantity)
	fmt.Printf("Получатель: %s\n", i.CustomerName)
	fmt.Printf("Телефон: +7%s\n", i.Phone)

	fmt.Println("\nАДРЕС ДОСТАВКИ:")
	fmt.Printf("   %s, %s\n", i.Address.Index, i.Address.City)
	fmt.Printf("   ул. %s, д. %s, кв. %s\n",
		i.Address.Street, i.Address.House, i.Address.Apartment)
	fmt.Println(strings.Repeat("=", 50))
}

// UpdateQuantity Метод с PointerReceiver для изменения данных
func (i *Invoice) UpdateQuantity(newQuantity int) {
	if newQuantity > 0 {
		i.Quantity = newQuantity
		fmt.Printf("Количество изменено на %d\n", newQuantity)
	} else {
		fmt.Println("Количество должно быть положительным!")
	}
}

// FormattedPhone Метод для форматирования телефона
func (i Invoice) FormattedPhone() string {
	if len(i.Phone) == 10 {
		return fmt.Sprintf("+7 (%s) %s-%s-%s",
			i.Phone[:3], i.Phone[3:6], i.Phone[6:8], i.Phone[8:])
	}
	return i.Phone
}

// Validate Метод для проверки корректности всех данных
func (i *Invoice) Validate() (bool, []string) {
	var errors []string

	// Проверка товара
	if i.ProductName == "" {
		errors = append(errors, "не указано наименование товара")
	} else if len([]rune(i.ProductName)) > 100 {
		errors = append(errors, "название товара слишком длинное")
	}

	// Проверка количества
	if i.Quantity <= 0 {
		errors = append(errors, "количество должно быть положительным")
	}

	// Проверка ФИО
	if i.CustomerName == "" {
		errors = append(errors, "не указано ФИО")
	} else if !i.isOnlyLetters(i.CustomerName) {
		errors = append(errors, "ФИО содержит недопустимые символы")
	}

	// Проверка телефона
	if i.Phone == "" {
		errors = append(errors, "не указан телефон")
	} else if !i.isOnlyDigits(i.Phone) || len(i.Phone) != 10 {
		errors = append(errors, "телефон должен содержать 10 цифр")
	}

	// Проверка адреса
	if i.Address.Index == "" {
		errors = append(errors, "не указан индекс")
	} else if !i.isOnlyDigits(i.Address.Index) || len(i.Address.Index) != 6 {
		errors = append(errors, "индекс должен содержать 6 цифр")
	}

	if i.Address.City == "" {
		errors = append(errors, "не указан город")
	} else if !i.isOnlyLetters(i.Address.City) {
		errors = append(errors, "город содержит недопустимые символы")
	}

	if i.Address.Street == "" {
		errors = append(errors, "не указана улица")
	}

	if i.Address.House == "" {
		errors = append(errors, "не указан дом")
	}

	if i.Address.Apartment == "" {
		errors = append(errors, "не указана квартира")
	} else if !i.isOnlyDigits(i.Address.Apartment) {
		errors = append(errors, "квартира должна содержать только цифры")
	}

	return len(errors) == 0, errors
}

func main() {
	// Создаем накладную через конструктор
	invoice := NewInvoice()

	// Заполняем данные
	invoice.InputData()

	// Проверяем валидность
	if valid, errors := invoice.Validate(); !valid {
		fmt.Println("\nОБНАРУЖЕНЫ ОШИБКИ:")
		for _, err := range errors {
			fmt.Printf("  - %s\n", err)
		}
		fmt.Println("\nПожалуйста, исправьте ошибки и создайте накладную заново.")
		return
	}

	invoice.DisplayInfo()

	fmt.Printf("\nТелефон в красивом формате: %s\n", invoice.FormattedPhone())

	fmt.Println("\nОбновление количества:")
	invoice.UpdateQuantity(5)
	fmt.Printf("Новое количество: %d шт.\n", invoice.Quantity)

}
