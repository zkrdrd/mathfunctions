package mathfunctions

import (
	"math"
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
