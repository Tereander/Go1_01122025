
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
    "fmt"
    "strings"
)

func main() {
    var passwd, err string
    var result bool

    for i := 5; i > 0; i-- {
        fmt.Println("Введите пароль или CTRL+C для выхода.")
        fmt.Println("Осталось попыток:", i)
        fmt.Scan(&passwd)
        if i == 0 {
			fmt.Println("Попытки закончились, приходите завтра.")
            return
        }
        result, err = checkPassword(passwd)
        if result {
            fmt.Println("ОК. Пароль соответствует всем требованиям безопасности.")
            return
        } else {
            fmt.Println("Пароль не соответствует требованиям безопасности!")
            fmt.Println(err)
        }
    }
}

func checkPassword(passwd string) (bool, string) {

    var dFlag, lFlag, uFlag, sFlag, result bool
    var err string

    if len(passwd) < 8 || len(passwd) > 15 {
        return result, "Длина пароля должна быть от 8 до 15 символов!"
    }

    digits := "0123456789"
    lowercase := "abcdefghijklmnopqrstuvwxyz"
    uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    special := "_!@#$%^&"

    for _, value := range passwd {

        if strings.ContainsRune(digits, value) {
            dFlag = true
        }
        if strings.ContainsRune(lowercase, value) {
            lFlag = true
        }
        if strings.ContainsRune(uppercase, value) {
            uFlag = true
        }
        if strings.ContainsRune(special, value) {
            sFlag = true
        }

        if dFlag && lFlag && uFlag && sFlag {
            break
        }
    }

    switch {
    case !dFlag: // dFlag == false
        err = "Должен содержать хотя бы одну цифру!"
    case !lFlag:
        err = "Должен содержать хотя бы одну букву в нижнем регистре!"
    case !uFlag:
        err = "Должен содержать хотя бы одну букву в верхнем регистре!"
    case !sFlag:
        err = "Должен содержать хотя бы один специальный символ!"
    default:
        result = true
    }

    return result, err
}