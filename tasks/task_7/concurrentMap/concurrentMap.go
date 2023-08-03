// Пакет частично реализующий потоко безопасную мапу ПБМ
package concurrentMap

import "sync"

// Структура ПБМ
// Представляет собой обертку на мапой mp, вся работа с которой происходит с использованием мьютекса mtx
type concurrentMap struct {
	mtx sync.Mutex
	mp  map[int]int
}

// Интерфейс ПБМ
type IConcurrentMap interface {
	Incr(int)
	Add(int, int)
	Get(int) int
	Iter() <-chan concurrentMapItem
}

// Структура, возвращаемая во время итерации по ПБМ с помощью метода Iter
type concurrentMapItem struct {
	Key   int
	Value int
}

// Конструктор
func NewConcurrentMap() IConcurrentMap {
	return &concurrentMap{
		mtx: sync.Mutex{},
		mp:  map[int]int{},
	}
}

// Метод инкрементирующий значение по ключу key
func (m *concurrentMap) Incr(key int) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.mp[key]++
}

// Метод для добавления ключа key со значением value
func (m *concurrentMap) Add(key, value int) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.mp[key] = value
}

// Метод для получения значения по ключу key
func (m *concurrentMap) Get(key int) int {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	return m.mp[key]
}

// Метод для итерации по ПБМ
func (m *concurrentMap) Iter() <-chan concurrentMapItem {
	ch := make(chan concurrentMapItem)

	go func() {
		m.mtx.Lock()
		defer m.mtx.Unlock()
		for index, value := range m.mp {
			ch <- concurrentMapItem{index, value}
		}
		close(ch)
	}()

	return ch
}
