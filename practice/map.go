package main

import "fmt"

func main() {
	strToNum := map[string]int{ // Словарь
		"zero": 0,
		"one": 1,
		"two": 2,
		"three": 3,
	}

	if value, ok := strToNum["zero"]; ok { // В услвоиях мы также можем присавивать значения аргументам
		fmt.Println("Zero is inside map and it's value is", value)
	} else {
		fmt.Println("Zero is missing")
	}
}