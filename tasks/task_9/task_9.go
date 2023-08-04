package task_9

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

// Согласно формулировке задачи, порядок вывода в stdout не важен
func Execute() {
	// Массив (слайс) из чисел, данный по условию
	nums := []int{-1, 0, 1, 100}

	// Канал для передачи чисел x
	chX := make(chan int)

	// Канал для передачи чисел x*2
	chX2 := make(chan int)

	// Переменная wg для ожидания завершения работы горутины, которая вызывает printSlice()
	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	// Вызов горутины, которая передает данные массива chX в канал chX
	go sendSlice(chX, nums)

	// Вызов горутины, которая получает данные из канала chX, умножает их на 2 и передает в канал chX2
	go multiplyX2(chX, chX2)

	// Вызов горутины, которая получает и печатает данные из канала chX2
	go printSlice(chX2, &wg)

	// Ожидание завершения горутины, которая вызвала printSlice()
	wg.Wait()
}

// Функция итерируется по слайсу nums и передает его значения в канал ch
func sendSlice(ch chan<- int, nums []int) {
	for _, v := range nums {
		ch <- v
	}
}

// Функция принимает на вход два канала:
// in_ch - канал, из которого читаются числа
// out_ch - канал, в который пишутся числа из in_ch умноженные на 2
func multiplyX2(in_ch <-chan int, out_ch chan<- int) {
	for v := range in_ch {
		out_ch <- 2 * v
	}
}

// Функция читает из канала ch числа, выводит их в stdout и уменьшает счетчик wg
func printSlice(ch <-chan int, wg *sync.WaitGroup) {
	for v := range ch {
		fmt.Println(v)
		wg.Done()
	}
}
