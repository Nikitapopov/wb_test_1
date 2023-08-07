package task_20

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func Execute() {
	// Исходная строка
	str := "snow dog sun"

	// Переворот слов в строке
	revertWords(&str)
}

//
// В случае нескольких идущих друг за другом пробелов, функция считает их за один.
// По условию любой символ кроме пробела это часть слова.
func revertWords(str *string) string {
	// массив из слов в str по порядку
	words := [][]rune{}

	// Текущее слово при итерации по str
	curWord := []rune{}

	// Количество рун в словах
	wordsRunesSize := 0

	// Итерация по str
	for _, symbol := range *str {
		// Если текущая руна равна пробелу
		if symbol == ' ' {
			// И если текущее слово не пустое, то добавление в массив words и обнуление текущего
			if len(curWord) > 0 {
				words = append(words, curWord)
				curWord = []rune{}
			}
			// И продолжение итерации
			continue
		}

		// Добавление руны в текущее слово и увеличение количества рун в словах на 1
		curWord = append(curWord, symbol)
		wordsRunesSize++
	}

	// В конце строки str обычно не стоит пробел, поэтому добавляем последнее текущее слово
	if len(curWord) > 0 {
		words = append(words, curWord)
	}

	if wordsRunesSize == 0 {
		return ""
	}

	// Выделение памяти под слайс рун с перевернутыми словами
	resStrAsRunes := make([]rune, 0, 2*wordsRunesSize-1)

	// Заполнение слайса рун с перевернутыми словами
	for i := len(words) - 1; i >= 0; i-- {
		resStrAsRunes = append(resStrAsRunes, words[i]...)

		// Добавление пробелов между словами
		if i > 0 {
			resStrAsRunes = append(resStrAsRunes, rune(' '))
		}
	}

	// Возврат строки с перевернутыми словами
	return string(resStrAsRunes)
}
