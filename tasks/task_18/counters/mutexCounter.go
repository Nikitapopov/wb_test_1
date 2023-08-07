package counters

import "sync"

// Счетчик с мьютексом
type mutexCounter struct {
	mtx     sync.Mutex
	counter int32
}

// Конструктор
func NewMutexCounter() ICounter {
	return &mutexCounter{
		mtx:     sync.Mutex{},
		counter: 0,
	}
}

// Метод инкрементирующий значение с помощью мьютекса
func (c *mutexCounter) Incr() {
	// Блокировка мьютекса
	c.mtx.Lock()

	// Разблокировка мьютекса по окончанию функции
	defer c.mtx.Unlock()

	//Инкрементация счетчика
	c.counter++
}

// Геттер значения счетчика
func (c *mutexCounter) Get() int32 {
	return c.counter
}
