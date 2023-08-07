package task_24

import (
	"fmt"
	"wb_test_2/tasks/task_24/point"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

func Execute() {
	// Исходные координаты двух точек
	x1 := 5
	y1 := 13
	x2 := 17
	y2 := 10

	// Вычисление расстояния между координатами
	distance := calculateDistance(x1, y1, x2, y2)

	// Вывод результата
	fmt.Println(distance)
}

// Функция для вычисления расстояния между координатами
func calculateDistance(x1, y1, x2, y2 int) float64 {
	// Структура точки 1
	p1 := point.NewPoint(x1, y1)

	// Структура точки 2
	p2 := point.NewPoint(x2, y2)

	// Вычисление расстояния между точками p1 и p2
	return p1.GetDistance(p2)
}
