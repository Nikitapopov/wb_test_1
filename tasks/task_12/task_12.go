package task_12

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func Execute() {
	// Исходная последовательность строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Сет строк
	set := map[string]struct{}{}

	// Итерация по sequence и заполнение сета set
	for _, str := range sequence {
		set[str] = struct{}{}
	}

	// Вывод сета
	fmt.Println(set)
}
