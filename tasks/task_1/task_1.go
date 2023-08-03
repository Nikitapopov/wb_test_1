package task_1

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)

import "fmt"

type human struct {
	name string
	age  uint16
}

func (h *human) changeAge(age uint16) {
	fmt.Printf("Возможен вызов метода changeAge структуры human, как и в перемнной типа human, так и типа action\n")
	h.age = age
}

func (h *human) printName() {
	fmt.Printf("Имя human - %s\n", h.name)
}

// Реализуем встраивание структуры human в структуру action.
// В следствии этого можно использовать поля и методы структуры human напрямую у переменной структуры action.
type action struct {
	human
	actionDesc string
}

func (a *action) doAction() {}

func (a *action) printName() {
	fmt.Printf("Имя action - %s\n", a.name)
}

func Execute() {
	a := action{
		human:      human{"Mike", 5},
		actionDesc: "This is a description",
	}

	a.changeAge(10)
	// Приоритет вызова функции printName, которая реализована и в human, и в action структурах будет у action (дочерней)
	a.printName()
}
