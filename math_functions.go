package mathfunctions

// массив чисел
var Numbers []int64

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
