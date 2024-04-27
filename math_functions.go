package mathfunctions

import (
	"bufio"
	"math"
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
	Numbers   []float64
	SumNumber float64
	Maximum   float64
	Minimum   float64
	Result    float64
	Temp      string
	Selecter  int
	FindSpace int
	GetText   = bufio.NewScanner(os.Stdin)
)

// сложение четных чисел
func SumEvenNumber(Numbers ...float64) float64 {
	for _, num := range Numbers {
		if int(math.Round(num))%2 == 0 {
			SumNumber += num
		}
	}
	return SumNumber
}

// сложение нечетных чисел math.Mod(num, 2)
func SumNoEvenNumber(Numbers ...float64) float64 {
	for _, num := range Numbers {
		if int(math.Round(num))%2 != 0 {
			SumNumber += num
		}
	}
	return SumNumber
}

// поиск максимального числа
func FingMaxNumber(Numbers ...float64) float64 {
	for _, num := range Numbers {
		if Maximum < num {
			Maximum = num
		}
	}
	return Maximum
}

// поиск минимального числа
func FingMinNumber(Numbers ...float64) float64 {
	for id, num := range Numbers {
		if Minimum > num || id == 0 {
			Minimum = num
		}
	}
	return Minimum
}
