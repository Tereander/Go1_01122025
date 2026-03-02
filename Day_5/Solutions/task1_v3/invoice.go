package main

import "fmt"




type Invoice struct {
    product string
    productQty int
    Customer
}

// NewInvoice - create new Invoice
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


// DisplayInfo - show information about Invoice
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