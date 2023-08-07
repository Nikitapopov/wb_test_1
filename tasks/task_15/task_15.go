package task_15

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/exp/utf8string"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.
//
// var justString string
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }
// func main() {
// 	someFunc()
// }

// Последствия фрагмента кода:
// 1. justString находится в глобальной памяти, поэтому память будет выделена в куче, что имеет влияние на
// 		производительность
// 2. Срез строки выполняется как слайс байтов, в том время, как символы строки могут занимать больше одного байта
// 3. Плохой нейминг

func Execute() {
	var justString string
	setRandom100SymbolsManually(&justString)
	fmt.Println(justString)
	setRandom100SymbolsWithUtf8stringPkg(&justString)
	fmt.Println(justString)
}

// Функция для создания случайной строки и взятия и вывода от нее первых 100 символов с помощью конвертации
// в слайс рун и взятие среза от него.
// Переименованный из someFunc.
func setRandom100SymbolsManually(justString *string) {
	// Создание строки
	hugeStr := createHugeString(1 << 10)

	// Преобразование строки к слайсу рун
	hugeStrAsRunes := []rune(hugeStr)

	// Взятие первых ста символов строки
	*justString = string(hugeStrAsRunes[:100])
}

// Функция для создания случайной строки и взятия и вывода от нее первых 100 символов с помощью пакета utf8string.
// Переименованный из someFunc.
func setRandom100SymbolsWithUtf8stringPkg(justString *string) {
	// Создание строки
	hugeStr := createHugeString(1 << 10)

	// Использование пакета utf8string, чтобы обработать строку и взять с 0 по 100 символы
	*justString = utf8string.NewString(hugeStr).Slice(0, 100)
}

// Функция для генерации строки длиной size, состоящей из кириллицы
func createHugeString(size int) string {
	// Создание и запуск генератора
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Слай из рун вместимостью size, куда будет генерироваться случайные руны
	bs := make([]rune, 0, size)

	// Заполнение bs символами кириллицы с пробелами
	for i := 0; i < size; i++ {
		var symbol rune
		// С вероятностью 10%, если это не первый символ, он приравнивается к пробелу
		if r.Intn(10) == 0 && i > 0 && bs[i-1] != ' ' {
			symbol = ' '
		} else {
			symbol = rune('А' + r.Intn(64))
		}
		bs = append(bs, symbol)
	}

	// Возвращение слайса рун строкой
	return string(bs)
}
