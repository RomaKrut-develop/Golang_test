package main

import ( // Импорт модулей
	"bufio"
	"fmt"
	"os"
)

func main() { // Точка входа в процедуру
	var input = readLine("Enter yes or no: ") // Объявляем переменную, которая будет получать данные от пользователя

	switch input { // Конструкция switch вместо if/else 
	case "yes":
		fmt.Println("You're agreed")
	case "no":
		fmt.Println("You're disagreed")	
	default: // На случай, если ни одно из условий не было выполнено
			fmt.Println("Unknown")
	}
}

func readLine(greeting string) string {
	fmt.Print(greeting)
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	return string(text)
}