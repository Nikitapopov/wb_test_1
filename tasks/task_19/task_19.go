package task_19

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func Execute() {
	// Исходная строка
	str := "главрыба"

	// Переворот строки
	revert(&str)

	// Вывод результата
	fmt.Println(str)
}

// Функция для переворота строки
func revert(str *string) {
	// Конвертация строки в слайс рун
	strAsRunes := []rune(*str)

	// Длина слайса
	size := len(strAsRunes)

	// Итерация по половине слайса
	for i := 0; i < size/2; i++ {
		// Обмен значений первого и последнего элемента, второго и предпоследнего и т.д.
		strAsRunes[i], strAsRunes[size-1-i] = strAsRunes[size-1-i], strAsRunes[i]
	}

	// Назначение переменной str новой перевернутой строки
	*str = string(strAsRunes)
}
