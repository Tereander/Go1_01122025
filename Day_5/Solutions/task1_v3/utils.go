package main

import (
	"bufio"
	"os"
)

func getStr() string {
    var str string
    input := bufio.NewScanner(os.Stdin)
    if input.Scan() {
        str = input.Text()
    }
    return str
}
