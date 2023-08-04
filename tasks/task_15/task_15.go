package task_15

import (
	"fmt"
	"math/rand"
	"time"
	// "golang.org/x/exp/utf8string"
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
// 1. justString находится в глобальной памяти, поэтому память будет выделена в куче,
// 		что имеет небольшой влияние на производительность
// 2. Срез строки выполняется как слайс байтов, в том время, как символы строки могут занимать больше одного байта
// 3. Плохой нейминг

func Execute() {
	var justString string
	setRandom100Symbols(&justString)
	fmt.Println(justString)
}

// Функция для создания случайной строки и взятия и вывода от нее первых 100 символов
func setRandom100Symbols(justString *string) {
	// Создание строки
	hugeStr := createHugeString(1 << 20)

	// TODO разобраться, имеет ли это решение преимущество
	// slc := utf8string.NewString(hugeStr).Slice(0, 100)

	// Преобразование строки к слайсу рун
	hugeStrAsRunes := []rune(hugeStr)

	// Взятие первых ста символов строки
	*justString = string(hugeStrAsRunes[:100])
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
		if r.Intn(10) == 0 {
			symbol = ' '
		} else {
			symbol = rune('А' + r.Intn(65))
		}
		bs = append(bs, symbol)
	}

	// Возвращение слайса рун строкой
	return string(bs)
}
