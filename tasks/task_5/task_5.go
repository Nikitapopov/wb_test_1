package task_5

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func Execute() {
	// Количество секунд, через которое программа завершается
	timeoutInSecs := 2

	// Канал для передачи значений
	ch := make(chan any)

	// Создание писателя значений для канала ch с таймаутом
	go produce(ch, timeoutInSecs)

	// Создание читателя значения из канала ch
	consume(ch)
}

// Функция писатель в канал ch
func produce(ch chan<- any, timeoutInSecs int) {
	// Создание канала timeout, который определяет, когда писатель закончит писать в канал и закроет канал
	timeout := time.After(time.Duration(timeoutInSecs) * time.Second)

	// Бесконечный цикл
	for {
		select {
		// При наступлении таймаута канал ch закрывается и функция возвращает значение
		case <-timeout:
			close(ch)
			return
		// Если таймаут еще не наступил, то продолжаем передавать данные в канал ch
		default:
			ch <- generateValue()
		}
	}
}

// Функция для создания значения для канала
func generateValue() interface{} {
	return struct{ x int }{x: 1}
}

func consume(ch <-chan any) {
	for event := range ch {
		fmt.Println(event)
	}
}
