package task_8

import (
	"errors"
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func Execute() {
	// Переменная данная по условию
	var num int64 = -1

	// i-й бит данный по условию
	numberOfBit := 0

	// Установка i-го бита в 1 или 0
	err := setIBit(&num, numberOfBit, false)
	if err != nil {
		fmt.Printf("error in setIBit: %v\n", err)
		return
	}

	// Вывод обновленного значения
	fmt.Printf("Res: %d\n", num)
}

// Функция для установки i-го бита у заданной переменной
// num - указатель на заданную переменную
// bit - номер i-го бита
// setToOne - Если true, то замена бита на 1, иначе на 0
func setIBit(num *int64, bit int, setToOne bool) error {
	// Валидация номера бита, т.к. у типа int64 64 бита
	if bit > 64 || bit < 1 {
		return errors.New("pos should be in interval [1, 64]")
	}

	// Маска для преобразования i-го бита
	var mask int64 = 1 << (64 - bit)
	if setToOne {
		// Если замена на 1, то применяем поразрядную сложение
		*num |= mask
	} else {
		// Иначе, применяем сброса бита
		*num &^= mask
	}

	return nil
}
