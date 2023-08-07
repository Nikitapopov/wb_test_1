package task_26

import "unicode"

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func Execute() {
	// Исходная строка
	str := "aabcd"

	// Проверка на регистронезависимую уникальность символов
	checkUniqueSymbols(str)
}

// Функция для проверки регистронезависимой уникальности символов
func checkUniqueSymbols(str string) bool {
	// Сет символов, входящих в строку
	set := map[rune]struct{}{}

	// Итерация по строке
	for _, symbol := range str {
		// Конвертация символа в строчный
		lcSymbol := unicode.ToLower(symbol)

		// Проверка, есть ли этот символ в сете
		_, ok := set[lcSymbol]

		// Если есть, то возврат отрицательного значения
		if ok {
			return false
		}

		// Помещение символа в сет
		set[lcSymbol] = struct{}{}
	}

	// Если все значения в строке регистронезависимо уникальные, то возврат положительного значения
	return true
}
