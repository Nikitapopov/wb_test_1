package task_3

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2 + 3^2 + 4^2...) с использованием конкурентных вычислений

import (
	"fmt"
	"runtime"
	"sync"
)

// В качестве решения приведено две функции squaresSum и squaresSumWithSmartGorutines.
// Первый вариант использует конкурентное вычисление квадратов элементов массива и в отдельной горутине выполняет их суммирование.
// Второй вариант, на мой взгяд практичнее, так как в нем суммирование элементов частично проходит в самих горутинах,
// количество которых ограничено имеющимися в системе ресурсами.
// Таким образом количество горутин минимально, а суммирование распараллеливается.

func Execute() {
	// Массив (слайс) чисел, для которых вычисляются квадраты
	nums := []int{2, 4, 6, 8, 10}

	// Вычисление суммы квадратов элементов массива nums
	sum := squaresSum(nums)

	// Вывод результата вычислений суммы квадратов элементов массива nums первым способом
	fmt.Printf("Sum = %d\n", sum)

	// Вычисление суммы квадратов элементов массива nums вторым способом
	sum = squaresSumWithSmartGorutines(nums)

	// Вывод результата вычислений суммы квадратов элементов массива nums вторым способом
	fmt.Printf("Sum = %d\n", sum)
}

// Вычисление суммы квадратов элементов массива с вызовом горутины для каждого элемента
func squaresSum(nums []int) int {
	// Канал, через который передаются квадраты чисел
	elemSquareCh := make(chan int)

	// Канал, через который возвращается сумма квадратов чисел
	sumCh := make(chan int)

	// В горутине запускается функция суммирования, которая получает квадраты чисел через канал elemSquareCh
	// и возвращает их сумму через канал sumCh
	go sum(elemSquareCh, sumCh)
	defer close(sumCh)

	// Использование типа WaitGroup для того, чтобы дождаться окончания вычисления всех квадратов чисел
	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	// Итерация по массиву (слайсу) чисел и вычисление квдарата для каждого из элементов в отдельной горутине
	for _, num := range nums {
		go square(num, elemSquareCh, &wg)
	}
	wg.Wait()
	close(elemSquareCh)

	// Возврат суммы квадратов из канала sumCh
	return <-sumCh
}

// Функция получает число num, вычисляем от него квадрат и прокидывает в канал ch.
// Уменьшает на 1 счетчик wg
func square(num int, ch chan<- int, wg *sync.WaitGroup) {
	ch <- num * num
	wg.Done()
}

// Функция получает канал ch, из которого получает числа.
// После закрытия канала ch прокидывает сумму этих чисел в канал sumCh
func sum(ch <-chan int, sumCh chan<- int) {
	sum := 0
	for num := range ch {
		sum += num
	}
	sumCh <- sum
}

// Вычислени суммы квадратов элементов массива с использованием горутин, в количестве имеющихся на ВМ ресурсов
func squaresSumWithSmartGorutines(nums []int) int {
	// Получение количества возможных процессов
	procNum := runtime.NumCPU()

	// Количество частей, на которые будет распараллелено вычисление
	chunkNum := procNum

	// Если количество элементов в массиве меньше количества процессов, то уменьшаем количество используемых чанков
	if procNum > len(nums) {
		chunkNum = len(nums)
	}

	// Канал для передачи из горутин частичных сумм квадратов
	partSumCh := make(chan int)

	// Вызов горутин в количестве chunkNum для вычисления частичных сумм квадратов
	for i := 0; i < chunkNum; i++ {
		go sumConcreteElemsOfSlice(i, chunkNum, nums, partSumCh)
	}

	sum := 0
	for i := 0; i < chunkNum; i++ {
		sum += <-partSumCh
	}
	close(partSumCh)

	return sum
}

// Функция для вычисления частичных сумм слайса, те суммы квадратов только тех элементов слайса,
// которые входят в итерацию по начальному элементу i и шагу step
// i - Индекс слайса, с которого начинается суммирование
// step - Шаг, с которым проходит итерация при суммировании
// slc - Слайс
// ch - Канал, в который передается частичная сумма
func sumConcreteElemsOfSlice(i int, step int, slc []int, ch chan<- int) {
	sum := 0
	for i := i; i < len(slc); i += step {
		sum += slc[i] * slc[i]
	}
	ch <- sum
}
