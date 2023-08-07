package task_22

import (
	"fmt"
	"math"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

func Execute() {
	// Исходное число a
	a := int64(math.Pow(2, 63)) - 5
	fmt.Println("a", a)

	// Исходное число b
	b := int64(math.Pow(2, 63)) - 7
	fmt.Println("b", b)

	// Выполнение арифметических операций
	doArifmeticOperations(a, b)
}

// Выполнение умножения, деления, суммирование и вычитание чисел a и b
func doArifmeticOperations(a, b int64) {
	// преобразование числа a в тип big.Int
	var aBig big.Int
	aBig.SetInt64(a)

	// преобразование числа b в тип big.Int
	var bBig big.Int
	bBig.SetInt64(b)

	// Переменная для результата арифметической операции
	var res big.Int

	// Перемножение aBig и bBig, помещение результата в res, вывод результата
	res.Mul(&aBig, &bBig)
	fmt.Println("Результат умножения", res.String())

	// Деление aBig на bBig, помещение результата в res, вывод результата
	res.Div(&aBig, &bBig)
	fmt.Println("Результат деления", res.String())

	// Суммирование aBig и bBig, помещение результата в res, вывод результата
	res.Add(&aBig, &bBig)
	fmt.Println("Результат сложения", res.String())

	// Разность aBig и bBig, помещение результата в res, вывод результата
	res.Sub(&aBig, &bBig)
	fmt.Println("Результат вычитания", res.String())
}
