package main

import "fmt"

type Point struct { // Пользовательский тип данных (структура)
	X int
	Y int
}

func main() {
	p := Point{X: 1}

	fmt.Println("%#+v\n",p)
}