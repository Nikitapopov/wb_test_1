package task_14

import "fmt"

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

func Execute() {
	// Объявление исходных переменных
	var var1 string = "test"
	var var2 int = 25
	var var3 bool = true
	var var4 chan any = make(chan any)

	// Вывод типов вышеобъявленных переменных
	printType(var1)
	printType(var2)
	printType(var3)
	printType(var4)
}

// Функция для определения типа переменной из списка: string, int, bool, chan any
func printType(value interface{}) {
	// switch type конструкция
	switch value.(type) {
	case string:
		fmt.Printf("variable %s has string type\n", value)
	case int:
		fmt.Printf("variable %d has int type\n", value)
	case bool:
		fmt.Printf("variable %t has bool type\n", value)
	case chan any:
		fmt.Printf("variable %v has chan(any) type\n", value)
	// Если ни один из типов не подошел, то вызов default кейса
	default:
		fmt.Printf("variable %v has unknown type\n", value)
	}
}
