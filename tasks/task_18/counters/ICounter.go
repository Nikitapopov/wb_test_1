package counters

// Интерфейс для счетчика
type ICounter interface {
	Incr()
	Get() int32
}
