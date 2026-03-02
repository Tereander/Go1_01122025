/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1 9 2
Выход: 1 2 9

Про sort мы пока ничего не знаем!
Это задача на условный оператор
*/
package main

import "fmt"

func main() {
        var dig1, dig2, dig3 int
        var dig1_str, dig2_str, dig3_str, res string

        fmt.Println("Введите первое число:")
        fmt.Scan(&dig1)
        fmt.Println("Введите второе число:")
        fmt.Scan(&dig2)
        fmt.Println("Введите третье число:")
        fmt.Scan(&dig3)

        dig1_str = fmt.Sprint(dig1)
        dig2_str = fmt.Sprint(dig2)
        dig3_str = fmt.Sprint(dig3)

        if dig1 < dig2 && dig1 < dig3 {
                res = dig1_str + " "
                if dig2 < dig3 {
                        res += dig2_str + " " + dig3_str
                } else {
                        res += dig3_str + " " + dig2_str
                }
        } else if dig2 < dig1 && dig2 < dig3 {
                res = dig2_str + " "
                if dig1 < dig3 {
                        res += dig1_str + " " + dig3_str
                } else {
                        res += dig3_str + " " + dig1_str
                }
        } else {
                res = dig3_str + " "
                if dig1 < dig2 {
                        res += dig1_str + " " + dig2_str
                } else {
                        res += dig2_str + " " + dig1_str
                }
        }

        fmt.Println("Числа в порядке возрастания: ", res)
}
