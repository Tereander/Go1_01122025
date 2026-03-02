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
    "regexp"
    "strconv"
)

type Address struct {
    index int
    city string
    street string
    house int
    apartment int
}

type Customer struct {
    name string
    phone int
    Address
}

type Invoice struct {
    product string
    productQty int
    Customer
}

func NewInvoice(
    newProduct string,
    newProductQty int,
    newCustomerName string,
    newCustomerPhone int,
    newCustomerAddressIndex int,
    newCustomerAddressCity string,
    newCustomerAddressStreet string,
    newCustomerAddressHouse int,
    newCustomerAddressApartment int,
) *Invoice {
    return &Invoice{
        product: newProduct,
        productQty: newProductQty,
        Customer: Customer{
            name: newCustomerName,
            phone: newCustomerPhone,
            Address: Address{
                index: newCustomerAddressIndex,
                city: newCustomerAddressCity,
                street: newCustomerAddressStreet,
                house: newCustomerAddressHouse,
                apartment: newCustomerAddressApartment,
            },
        },
    }
}

func (i Invoice) DisplayInfo() {
    fmt.Println("Product:", i.product)
    fmt.Println("ProductQty:", i.productQty)
    fmt.Println("Customer name:", i.Customer.name)
    fmt.Println("Customer phone:", i.Customer.phone)
    fmt.Println("Customer index:", i.Customer.Address.index)
    fmt.Println("Customer city:", i.Customer.Address.city)
    fmt.Println("Customer street:", i.Customer.Address.street)
    fmt.Println("Customer house:", i.Customer.Address.house)
    fmt.Println("Customer apartment:", i.Customer.Address.apartment)
}

func getStr() string {
    var str string
    input := bufio.NewScanner(os.Stdin)
    if input.Scan() {
        str = input.Text()
    }
    return str
}

func main() {
    var str string
    var match bool
    var err error
    var newProduct, newCustomerName, newCustomerAddressCity, newCustomerAddressStreet string
    var newProductQty, newCustomerPhone, newCustomerAddressIndex, newCustomerAddressHouse, newCustomerAddressApartment int

    fmt.Println("Добавление накладной. Для выхода нажмите CTRL+C.")
    for {
        if newProduct == "" {
            fmt.Println("Введите наименование продукта:")
            str = getStr()
            if len(str) < 1 || len(str) > 100 {
                fmt.Println("Недопустимая длина наименования товара! Должна быть от 1 до 100.")
                fmt.Println("Попробуйте еще раз.")
                continue
            } else {
                newProduct = str
            }
        }
        if newProductQty == 0 {
            fmt.Println("Введите количество продукта целым числом:")
            str = getStr()
            newProductQty, err = strconv.Atoi(str)
            if err != nil {
                fmt.Println("Ошибка в количестве товара:", err)
                fmt.Println("Попробуйте еще раз.")
                continue
            }
        }
        if newCustomerName == "" {
            fmt.Println("Введите ФИО покупателя:")
            str = getStr()
            match, err = regexp.MatchString(`^[a-zA-Z]+$`, str)
            if !match {
                fmt.Println("ФИО должно содержать только буквы.")
                fmt.Println("Попробуйте еще раз.")
                continue
            } else {
                newCustomerName = str
            }
        }
        if newCustomerPhone == 0 {
            fmt.Println("Введите номер телефона целым числом:")
            str = getStr()
            match, err = regexp.MatchString(`^\d{10}$`, str)
            if !match {
                fmt.Println("Телефонный номер должен содержать только 10 цифр!")
                fmt.Println("Попробуйте еще раз.")
                newCustomerPhone = 0
                continue
            } else {
              newCustomerPhone, err = strconv.Atoi(str)
            }
        }
        if newCustomerAddressIndex == 0 {
            fmt.Println("Введите почтовый индекс целым числом:")
            str = getStr()
            match, err = regexp.MatchString(`^\d{6}$`, str)
            if !match {
                fmt.Println("В почтовом индексе должно быть 6 цифр")
                fmt.Println("Попробуйте еще раз.")
                newCustomerAddressIndex = 0
                continue
            } else {
                newCustomerAddressIndex, err = strconv.Atoi(str)
            }
        }
        if newCustomerAddressCity == "" {
            fmt.Println("Введите город:")
            str = getStr()
            if len(str) == 0 {
                fmt.Println("Название города не может быть пустым!")
                fmt.Println("Попробуйте еще раз.")
                continue
            } else {
                newCustomerAddressCity = str
            }
        }
        if newCustomerAddressStreet == "" {
            fmt.Println("Введите улицу:")
            str = getStr()
            if len(str) == 0 {
                fmt.Println("Название улицы не может быть пустым!")
                fmt.Println("Попробуйте еще раз.")
                continue
            } else {
                newCustomerAddressStreet = str
            }
        }
        if newCustomerAddressHouse == 0 {
            fmt.Println("Введите номер дома целым числом:")
            str = getStr()
            newCustomerAddressHouse, err = strconv.Atoi(str)
            if err != nil {
                fmt.Println("Ошибка при вводе номера дома:", err)
                fmt.Println("Попробуйте еще раз.")
                newCustomerAddressHouse = 0
                continue
            }
        }
        if newCustomerAddressApartment == 0 {
            fmt.Println("Введите номер квартиры целым числом:")
            str = getStr()
            newCustomerAddressApartment, err = strconv.Atoi(str)
            if err != nil {
                fmt.Println("Ошибка при вводе номера квартиры:", err)
                fmt.Println("Попробуйте еще раз.")
                newCustomerAddressApartment = 0
                continue
            }
        }
        break
    }

    fmt.Println("Ввод накладной успешно завершен:")
    fmt.Println("===========================")
    invoice := NewInvoice(
        newProduct,
        newProductQty,
        newCustomerName,
        newCustomerPhone,
        newCustomerAddressIndex,
        newCustomerAddressCity,
        newCustomerAddressStreet,
        newCustomerAddressHouse,
        newCustomerAddressApartment,
    )
    invoice.DisplayInfo()
}