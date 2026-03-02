/* 
	Задача №1
	Написать функцию, которая расшифрует строку. 
	code = "220411112603141304"
	Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
	Отчет с "00" -> 'a' и до "25" -> 'z', "26" -> ' '(пробел).
	Вход: строка из цифр. Выход: Текст.
	Проверка работы функции выполняется через другую строку.
	Рекомендую использовать мапы, будет лучше, если вы их создадите с помощью цикла

	Задача №1.1
	Реализовать и функцию зашифровки

	codeToString(code) -> "???????' 

	stringToCode("hello") -> "??????"
*/
package main


import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var menu int
    runeByCodeMap := make(map[string]rune)
    codeByRuneMap := make(map[rune]string)

	// Наполняем мапы нужными парами
    
    for char, i := 'a', 0; char <= 'z'; char, i = char+1, i+1 {
        runeByCodeMap[fmt.Sprintf("%02d", i)] = char
        codeByRuneMap[char] = fmt.Sprintf("%02d", i)
        
    }
    runeByCodeMap["26"] = ' '
    codeByRuneMap[' '] = "26"

    for {
        fmt.Printf("\n\nВведите:\n 1 для расшифровки кода\n 2 для зашифровки\n 0 для выхода\n")
        fmt.Scan(&menu)
        if menu == 1 {
            fmt.Println("Расшифровка:", codeToString(runeByCodeMap, getStr()))
        } else if menu == 2 {
            fmt.Println("Код:", stringToCode(codeByRuneMap, getStr()))
        } else {
            break
        }
    }
}

func getStr() string {
    var str string
    fmt.Println("Введите строку:")
    input := bufio.NewScanner(os.Stdin)
    if input.Scan() {
        str = input.Text()
    }
    return str
}


func codeToString(runeByCodeMap map[string]rune, str string) string {
    var code, result string
    runeArray := []rune(str)
    for i, value := range runeArray {
        if i % 2 == 0 {
            code = string(value)
        } else {
            code += string(value)
            if char, ok := runeByCodeMap[code]; ok {
                result += string(char)
            }
        }
    }
    return result
}

func stringToCode(codeByRuneMap map[rune]string, str string) string {
    var result string
    runeArray := []rune(str)
    for _, value := range runeArray {
        if code, ok := codeByRuneMap[value]; ok {
            result += code
        }
    }
    return result
}
