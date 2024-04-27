package mathfunctions

import (
	"bufio"
	"os"
)

// ascii byte space
const Space byte = 32

/*
**
Numbers   = массив чисел
SumNumber = сумма чисел
Maximum   = минимальное значение
Minimum   = максимальное значение
Result	  = запись результатов работы функции
Temp      = временная переменная
Selecter  = пункт меню выбора
FindSpace = поиск пробела
GetText   = сканер
**
*/
var (
	Numbers   []int64
	Result    int64
	Temp      string
	Selecter  int64
	FindSpace int
	GetText   = bufio.NewScanner(os.Stdin)
)

// сложение четных чисел
func SumEvenNumber(Numbers ...int64) int64 {
	// сумма элементов
	var SumNumber int64
	for _, num := range Numbers {
		if num%2 == 0 {
			SumNumber += num
		}
	}
	return SumNumber
}

// сложение нечетных чисел math.Mod(num, 2)
func SumNoEvenNumber(Numbers ...int64) int64 {
	// сумма элементов
	var SumNumber int64
	for _, num := range Numbers {
		if num%2 != 0 {
			SumNumber += num
		}
	}
	return SumNumber
}

// поиск максимального числа
func FindMaxNumber(Numbers ...int64) int64 {
	// максимальное число
	var Maximum int64
	for _, num := range Numbers {
		if Maximum < num {
			Maximum = num
		}
	}
	return Maximum
}

// поиск минимального числа
func FindMinNumber(Numbers ...int64) int64 {
	// минимальное число
	var Minimum int64
	for id, num := range Numbers {
		if Minimum > num || id == 0 {
			Minimum = num
		}
	}
	return Minimum
}
