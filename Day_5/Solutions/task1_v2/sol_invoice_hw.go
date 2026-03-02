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
	"fmt"
	"math"
	"os"
	"strings"
)

// Ваш код

type Order struct {
	MerchName    string
	Quantity     float32
	CustomerName string
	PhoneNnumber uint64
	Adress       Adress
}

func (o Order) Print() {
	fmt.Println(" ==========")
	fmt.Println(" - Order  -")
	fmt.Println(" ==========")
	fmt.Println(" - Merch        :", o.MerchName)
	fmt.Println(" - Quantity     :", o.Quantity)
	fmt.Println(" - Customer Name:", o.CustomerName)
	fmt.Println(" - PhoneNnumber :", o.PhoneNnumber)

	o.Adress.Print()
}

type Adress struct {
	Index                     int
	City, Street, House, Room string
}

func (a Adress) Print() {
	fmt.Println(" ==========")
	fmt.Println(" - Adress -")
	fmt.Println(" ==========")
	fmt.Println(" - Index :", a.Index)
	fmt.Println(" - City  :", a.City)
	fmt.Println(" - Street:", a.Street)
	fmt.Println(" - House :", a.House)
	fmt.Println(" - Room  :", a.Room)
	fmt.Println(" ==========")
}

func NewOrder() (order *Order) {
	order = &Order{}

	fmt.Println("Enter Merchandice Name:")
	fmt.Scan(&order.MerchName)

	mercNameLen := len(order.MerchName)
	if mercNameLen < 1 || mercNameLen > 100 {
		fmt.Println("Invalid Merchandice Name length:", mercNameLen)
		os.Exit(2)
	}

	fmt.Println("Enter Merchandice Quantity:")
	fmt.Scan(&order.Quantity)

	if order.Quantity < math.SmallestNonzeroFloat32 {
		fmt.Println("Quantity must be not zero")
		os.Exit(2)
	}

	fmt.Println("Enter Customer Name:")
	fmt.Scan(&order.CustomerName)

	if strings.Contains(order.CustomerName, "0987654321") {
		fmt.Println("Customer name must not contain digits")
		os.Exit(2)
	}

	fmt.Println("Enter Customer Phone:")
	fmt.Scan(&order.PhoneNnumber)

	if order.PhoneNnumber < 100_000_00_00 || order.PhoneNnumber > 999_999_99_99 {
		fmt.Println("Phone number mus be only 10 digits")
		os.Exit(2)
	}

	fmt.Println("Enter Customer Index:")
	fmt.Scan(&order.Adress.Index)

	if order.Adress.Index < 100_000 || order.Adress.Index > 999_999 {
		fmt.Println("Adress Index mus be only 6 digits")
		os.Exit(2)
	}

	fmt.Println("Enter Customer City:")
	fmt.Scan(&order.Adress.City)

	fmt.Println("Enter Customer Street:")
	fmt.Scan(&order.Adress.Street)

	fmt.Println("Enter Customer House:")
	fmt.Scan(&order.Adress.House)

	fmt.Println("Enter Customer Room:")
	fmt.Scan(&order.Adress.Room)

	return order
}

func main() {
	order := NewOrder()

	order.Print()
}
