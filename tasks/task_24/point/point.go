package point

import "math"

// Интерфейс точки
type IPoint interface {
	GetCoordinates() (int, int)
	GetDistance(anotherPoint IPoint) float64
}

// Структура точки
type point struct {
	x int
	y int
}

// Конструтктор с координатами
func NewPoint(x, y int) IPoint {
	return &point{
		x: x,
		y: y,
	}
}

// Метод для получения координат
func (p point) GetCoordinates() (int, int) {
	return p.x, p.y
}

// Метод для вычисления расстояния между точками с округлением до тысячных
func (p point) GetDistance(p2 IPoint) float64 {
	// Координаты первой точки имеются в p

	// Координаты второй точки p2
	x, y := p2.GetCoordinates()

	// Вычисление расстояния между p и p2
	distance := math.Sqrt(math.Pow(float64(x-p.x), 2) + math.Pow(float64(y-p.y), 2))

	// Возврат округленного до тысячных результата
	return math.Round(distance*1000) / 1000
}
