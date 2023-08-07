package task_17

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func Execute() {
	// Исходный отсортированный массив
	arr := []int{} //1, 2, 3, 4, 5, 6
	target := 6

	// бинарный поиск
	index, ok := search(arr, target)
	if ok {
		fmt.Printf("Index: %d\n", index)
	} else {
		fmt.Println("Elem not found")
	}
}

// Функция реализующая бинарный поиск значения target в массиве arr.
// Возвращает индекс найденного элемента, если поиск успешен, и флаг, успешен ли поиск.
func search(arr []int, target int) (int, bool) {
	// Если массив пустой, то возврат отрицательного ответа
	if len(arr) == 0 {
		return 0, false
	}

	return binarySearch(arr, target, 0, len(arr)-1)
}

// Функция реализующая бинарный поиск значения target в массиве arr в промежутке от low до high индексов.
// Возвращает индекс найденного элемента, если поиск успешен, и флаг, успешен ли поиск.
func binarySearch(arr []int, target int, low, high int) (int, bool) {
	// Среднее, округленное в меньшую сторону, между значениями элементов с индексами low и high
	mean := (high + low) / 2

	// Если среднее равно target, то возврат этого значения
	if arr[mean] == target {
		return mean, true
	}

	// Если в массиве не осталось элементов, то возврат отрицательного ответа
	if low == high {
		return 0, false
	}

	// Если target меньше среднего, то поиск в массиве arr слева от элемента с индексом mean, иначе справа
	if target < mean {
		return binarySearch(arr, target, low, mean-1)
	} else {
		return binarySearch(arr, target, mean+1, high)
	}
}
