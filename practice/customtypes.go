package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64 // 64-битный формат плавающей точки, используемый для хранения вещественных чисел с высокой точностью
	Y float64
}

func (p Point) Abs() float64 { // Создаем пользовательский тип данных
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Point) SetX(newX float64) {
	p.X = newX
}

func main() { // Объявляет переменную и присваивает базовые значения к структуре данных
	p := &Point{1,2}

	p.SetX(100) // Изменяем поинту X значение на 100

	fmt.Println(p.x)
}