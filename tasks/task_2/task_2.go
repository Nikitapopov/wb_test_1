package task_2

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

func Execute() {
	// Массив (слайс) чисел, для которых вычисляются квадраты
	nums := []int{2, 4, 6, 8, 10}

	calculateSquares(nums)

	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	nums2 := []int{2, 4, 6, 8, 10}
	squares2 := calculateSquaresImmunably(nums2)

	for _, v := range squares2 {
		fmt.Printf("%d ", v)
	}
}

// Функция конкурентно вычисляет квадраты чисел из поданного на вход массива nums.
// Результаты вычислений записываются в тот же массив в соответствующие элементы.
func calculateSquares(nums []int) {
	// Объявляется переменная wg для того, чтобы последующие горутины сработали, как одна группа
	wg := sync.WaitGroup{}

	// Счетчику переменной wg назначается значение длины массива чисел
	wg.Add(len(nums))

	// В цикле для каждого числа из массива запускается по горутине. Через горутину вызывается функция вычисления квадрата числа
	for i, _ := range nums {
		go calculateSquare(&nums[i], &wg)
	}

	// Ожидание, пока все горутины не закончат работу
	wg.Wait()
}

// Функция вычисляет квадрат числа указателя num, кладет результат в ту же область память num и вызывает уменьшение счетчика переменной wg
func calculateSquare(num *int, wg *sync.WaitGroup) {
	*num = *num * *num
	wg.Done()
}

// Анлог функции calculateSquares без сайд эффектов
func calculateSquaresImmunably(nums []int) []int {
	// Объявляется переменная wg для того, чтобы последующие горутины сработали, как одна группа
	wg := sync.WaitGroup{}

	// Счетчику переменной wg назначается значение длины массива чисел
	wg.Add(len(nums))

	res := make([]int, len(nums))

	// В цикле для каждого числа из массива запускается по горутине. Через горутину вызывается функция вычисления квадрата числа
	for i, v := range nums {
		go calculateSquareImmunably(v, &res[i], &wg)
	}

	// Ожидание, пока все горутины не закончат работу
	wg.Wait()

	return res
}

// Функция вычисляет квадрат числа указателя num, кладет результат в ту же область память num и вызывает уменьшение счетчика переменной wg
func calculateSquareImmunably(num int, res *int, wg *sync.WaitGroup) {
	*res = num * num
	wg.Done()
}
