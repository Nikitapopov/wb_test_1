package task_13

import "fmt"

// Поменять местами два числа без создания временной переменной.

func Execute() {
	// Исходные меременные
	var1 := 1
	var2 := 2

	// Назначение новых значений
	var1, var2 = var2, var1

	// Вывод значений
	fmt.Printf("var 1: %d, var 2: %d", var1, var2)
}
