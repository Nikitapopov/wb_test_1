package task_7

import (
	"fmt"
	"math/rand"
	"sync"

	"wb_test_2/tasks/task_7/concurrentMap"
)

// Реализовать конкурентную запись данных в map.

// В качестве ключа и значения мапы использован тип int

// Используется для тестирования записи значений
const mapSize = 5

func Execute() {
	// Количество писателей-горутин, которые будут писать в мапу
	writersNum := 100

	// Реализация с использованием блокировки мьютекса на время записи
	fillMap1(writersNum)

	// Реализация со использованием конкурентной обертки над мапой
	fillConcurrentMap(writersNum)
}

// Функция для записи в мапу mp значений, которые записывают конкурентно горутины в количестве writersNum.
// Все горутины используют блокировку мьютекстом перед записью в мапу
func fillMap1(writersNum int) {
	// Мапа, в которую будет вестись конкурентная запись
	mp := make(map[int]int, mapSize)

	// Мьютекс для ограничения гонки данных
	mtx := sync.Mutex{}

	// Переменная wg для ожидания завершения работы всех горутин
	wg := sync.WaitGroup{}
	wg.Add(writersNum)

	// Запуск писателей в мапу в горутинах в количестве writersNum
	for i := 0; i < writersNum; i++ {
		go writeToMap(mp, &mtx, &wg)
	}
	wg.Wait()

	// Проверка на потерю данных
	sum := 0
	for _, value := range mp {
		sum += value
	}

	fmt.Printf("Sum expected: %d, got: %d\n", writersNum, sum)
}

// Функция для записи в мапу mp. Запись сопровождается блокировкой мьютекса mtx. Уменьшается счетчик wg.
// Запись ведется по ключу, представляющем из себя целое числов интервале [0, 5).
func writeToMap(mp map[int]int, mtx *sync.Mutex, wg *sync.WaitGroup) {
	mtx.Lock()
	defer wg.Done()
	defer mtx.Unlock()
	ranNum := rand.Intn(mapSize)
	mp[ranNum]++
}

func fillConcurrentMap(writersNum int) {
	// Структура обертка с мьютексом над мапой
	mp := concurrentMap.NewConcurrentMap()

	// Переменная wg для ожидания завершения работы всех горутин
	wg := sync.WaitGroup{}
	wg.Add(writersNum)

	// Запуск писателей в мапу в горутинах в количестве writersNum
	for i := 0; i < writersNum; i++ {
		go writeToConcurrentMap(mp, &wg)
	}
	wg.Wait()

	// Проверка на потерю данных
	sum := 0
	for item := range mp.Iter() {
		sum += item.Value
	}

	fmt.Printf("Sum expected: %d, got: %d\n", writersNum, sum)
}

// Функция для записи в конкурентно безопасную мапу mp. Уменьшается счетчик wg.
// Запись ведется по ключу, представляющем из себя целое числов интервале [0, 5).
func writeToConcurrentMap(mp concurrentMap.IConcurrentMap, wg *sync.WaitGroup) {
	ranNum := rand.Intn(mapSize)
	mp.Incr(ranNum)
	wg.Done()
}
