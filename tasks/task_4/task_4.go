package task_4

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора
// количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы
// всех воркеров.

func Execute() {
	// Получение количества воркеров
	workersNum := scanNum()

	// Канал для передачи данных в воркеры
	ch := make(chan interface{})

	// Запуск воркеров в горутинах в количестве workersNum
	for i := 0; i < workersNum; i++ {
		go handleEvent(ch)
	}

	// Объявление канала, отслеживающего SIGINT сигнал
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	//Флаг, определяющий был ли вызван SIGINT сигнал
	isStopped := false

	// Если флаг положительный, значит необходимо завершить генерацию событий и закрыть канал ch
	// Иначе продолжаем генерировать события
	for !isStopped {
		select {
		case <-signalCh:
			close(ch)
			isStopped = true
		default:
			ch <- generateEvent()
		}
	}

	// Завершение работы всех воркеров произойдет благодаря закрытию канала ch.
	// После закрытия воркеры выйдут из цикла
	// for event := range ch
	// и завершатся
}

// Функция для чтения из консоли числа.
func scanNum() int {
	var s string
	fmt.Println("Enter number of workers:")
	// Считывание строки до тех пор, пока не будет введено валидное положительное число
	for {
		_, err := fmt.Scan(&s)
		if err != nil {
			fmt.Println("Error while scanning")
			continue
		}
		workersNum, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Enter a valid number")
		} else if workersNum <= 0 {
			fmt.Println("Number should be positive")
		} else {
			return workersNum
		}
	}
}

// Функция для создания события
func generateEvent() interface{} {
	return struct{ x int }{x: 1}
}

// Функция для получения данных из канала и вывода в консоль
func handleEvent(ch chan any) {
	for event := range ch {
		fmt.Println(event)
	}
}
