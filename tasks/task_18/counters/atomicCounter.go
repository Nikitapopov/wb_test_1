package counters

import "sync/atomic"

// Счетчик
type atomicCounter struct {
	counter int32
}

// Конструктор
func NewAtomicCounter() ICounter {
	return &atomicCounter{
		counter: 0,
	}
}

// Метод инкрементирующий значение с помощью атомика
func (c *atomicCounter) Incr() {
	atomic.AddInt32(&c.counter, 1)
}

// Геттер значения счетчика
func (c *atomicCounter) Get() int32 {
	return c.counter
}
