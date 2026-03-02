/* 
	3.4 В цикле получать от пользователя оценки по четырём экзаменам.
		Вывести сумму набранных им баллов.
		Функцию fmt.Scan() в коде используем только один раз.
*/
package main

import "fmt"

func main(){
	var value, summ int

	fmt.Print("3.4 Экзамены:\n")

	for index := 1; index <= 4; index++ {
		fmt.Printf("Введи оценку за %d экзамен: ", index)
		fmt.Scan(&value)
        summ += value
	}

	fmt.Printf("Итоговый балл за экзамены: %d\n", summ)
}