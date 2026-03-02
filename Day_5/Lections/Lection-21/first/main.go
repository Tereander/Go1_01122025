package main

import (
	"fmt"
	"unicode/utf8"
)

// 1. Описание интерфейса
type BigWord interface {
	IsBig() bool
}

// 2. Наш тип, который должен иметь метод IsBig()
type MySuperString string

// 3. Реализация IsBig()
func (mss MySuperString) IsBig() bool {
	return utf8.RuneCountInString(string(mss)) > 10 
}


func main() {
	sample := MySuperString("hello")
	var interfaceSample BigWord // объявление переменной с типом BigWord
	fmt.Printf("Value: %v and type: %T\n", sample, sample)
	fmt.Printf("Value: %v and type: %T\n", interfaceSample, interfaceSample)

	interfaceSample = sample // так можно сказать, т.к. sample удовлетваряет интерфейсу BigWord
	fmt.Printf("After value: %v and type: %T\n", interfaceSample, interfaceSample)
	fmt.Println("IsBig?", interfaceSample.IsBig())

	// 4. Попытка присвоить значение с типом, который не отвечает интерфейсу, приводит к ошибке
	// var interfaceBadSample BigWord
	// interfaceBadSample = "long string" // string не имеет реализации IsBig()

}