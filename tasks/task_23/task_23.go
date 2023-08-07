package task_23

import "errors"

// Удалить i-ый элемент из слайса.

func Execute() {
	// Исходный слайс
	slc := []int{3, 4, 5, 6, 7, 8}

	// Номер удаляемого элемента
	deletedNum := 3

	// Удаление элемента из слайса
	deleteElem(&slc, deletedNum)
}

// Удаление элемента под номером num из слайса slc
func deleteElem(slc *[]int, num int) error {
	// Валидация удаляемого номера
	if num < 1 || num > len(*slc) {
		return errors.New("number exceeds slice range")
	}

	// Удаление из слайса элемента под номером num путем помещения элементов слайса слева от номера и справа в слайс
	*slc = append((*slc)[:num-1], (*slc)[num-1+1:]...)
	return nil
}
