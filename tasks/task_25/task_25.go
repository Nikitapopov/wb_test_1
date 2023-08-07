package task_25

import "time"

// Реализовать собственную функцию sleep

func Execute() {
	// Количество наносекунд, на которое произойден засыпание
	var ns int64 = 1e9

	// Засыпание с помощью time.NewTimer
	sleepWithNewTimer(ns)

	// Засыпание с помощью time.After
	sleepWithTimeAfter(ns)
}

// Функция ожидающая ns наносекунд с помощью time.NewTimer
func sleepWithNewTimer(ns int64) {
	// Если ns меньше нуля, то сразу возврат из функции
	if ns <= 0 {
		return
	}

	// Запуск таймера
	timer := time.NewTimer(time.Nanosecond * time.Duration(ns))

	// Ожидание передачи данных в канал по истечению ns наносекунд
	<-timer.C
}

// Функция ожидающая ns наносекунд с помощью time.After, те тоже с помощью time.NewTimer
func sleepWithTimeAfter(ns int64) {
	// Если ns меньше нуля, то сразу возврат из функции
	if ns <= 0 {
		return
	}

	// Ожидание срабатывания таймаута
	<-time.After(time.Nanosecond * time.Duration(ns))
}
