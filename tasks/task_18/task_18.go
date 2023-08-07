package task_18

import (
	"fmt"
	"sync"
	"wb_test_2/tasks/task_18/counters"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

func Execute() {
	// Количество горутин
	workersNum := 30

	// Количество инкрементирований, которое выполняет каждая горутина
	increasesNum := 10

	// Счетчик, использующий мютекс
	mutexCounter := counters.NewMutexCounter()
	fmt.Println("Mutex counter test:")
	testCounter(mutexCounter, workersNum, increasesNum)

	// Счетчик, использующий атомик
	fmt.Println("Atomic counter test:")
	atomicCounter := counters.NewMutexCounter()
	testCounter(atomicCounter, workersNum, increasesNum)
}

// Функция для проверки корректности работы счетчика counter
// на горутинах количеством workersNum с количеством инекрементаций increasesNum
func testCounter(counter counters.ICounter, workersNum, increasesNum int) {
	// Переменная wg для ожидания завершения работы всех горутин
	wg := sync.WaitGroup{}
	wg.Add(workersNum)

	// Запуск горутин
	for i := 0; i < workersNum; i++ {
		// горутина для инкрементации счетчика
		go increaseCounter(counter, increasesNum, &wg)
	}

	// Ожидание заврешения работы всех горутин
	wg.Wait()

	// Вывод результатов
	fmt.Printf("Expected: %d, got: %d\n", workersNum*increasesNum, counter.Get())
}

// Функция для инкрементации счетчика counter increasesNum раз. Уменьшение счетчика wg на 1.
func increaseCounter(counter counters.ICounter, increasesNum int, wg *sync.WaitGroup) {
	for i := 0; i < increasesNum; i++ {
		counter.Incr()
	}

	wg.Done()
}
