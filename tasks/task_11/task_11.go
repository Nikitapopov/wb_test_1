package task_11

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func Execute() {
	// Сет с неупорядоченным множеством 1
	set1 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}

	// Сет с неупорядоченным множеством 2
	set2 := map[int]struct{}{
		2: {},
		3: {},
		4: {},
	}

	// Сет с пересечением двух множеств set1 и set2
	intersection := map[int]struct{}{}

	// Добавление множества set1 в intersection
	for key := range set1 {
		intersection[key] = struct{}{}
	}

	// Добавление множества set2 в intersection
	for key := range set2 {
		intersection[key] = struct{}{}
	}

	// Вывод пересечения
	for key, _ := range intersection {
		fmt.Printf("%d ", key)
	}
}
