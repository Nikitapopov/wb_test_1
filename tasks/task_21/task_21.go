package task_21

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

func Execute() {
	makeTrip()
}

// Структура велосипед
type bicycle struct {
	model string
}

// Метод велосипеда, принимающий параметр pedals, с помощью которого выполняет логику кручения педалей
func (b *bicycle) ride(pedals iPedals) {
	pedals.turnPedals()
}

type iPedals interface {
	turnPedals()
}

// Структуры педали
type pedals struct {
	material string
}

// Метод, реализующий кручение педалей
func (p *pedals) turnPedals() {
	fmt.Println("Turning the pedals")
}

// Структура двигатель
type motor struct {
	fuelType string
}

// Метод, реализующий нажатие гашетки газа
func (m *motor) pressTrigger() {
	fmt.Println("Pressing the trigger")
}

// Структура адаптера педалей, использующий внутри структуру двигателя
type pedalsAdapter struct {
	motor motor
}

// Метод, реализующий кручения педалей, но использующий нажатие гашетки двигателя
func (a *pedalsAdapter) turnPedals() {
	fmt.Println("Adapter convert pressing the trigger to pedals rotation")
	a.motor.pressTrigger()
}

// Функция, выполняющая поздки, используя паттерн "адаптер"
func makeTrip() {
	// Объявление переменной педалей
	usualPedals := pedals{
		material: "plastic",
	}

	// Объявление переменной велосипеда и выполнение поездки, используя обычные педали usualPedals
	bike := bicycle{model: "bmx"}
	bike.ride(&usualPedals)

	// Объявление переменной двигателя
	electricMotor := motor{
		fuelType: "electricity",
	}

	// Адаптер педалей, использующий двигатель electricMotor
	adapter := pedalsAdapter{
		motor: electricMotor,
	}

	// Объявление переменной велосипеда и выполнение поездки, используя двигатель electroBike
	electroBike := bicycle{model: "tesla"}
	electroBike.ride(&adapter)
}
