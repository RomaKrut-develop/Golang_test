package main

import "fmt"

type Point struct { // Создаем пользовательский тип данных
	X, Y int
}

func main() {
	p := Point{} // Объявляем переменную p и присваиваем ей структуру XY

	fill(p) // Вызываем функцию которая заполнит структуру (пользовательский тип данных Point)

	fmt.Println(p)
}

func fill(p Point) { // Функция заполняет структуру XY значениями
	p.X = 1
	p.Y = 2
}