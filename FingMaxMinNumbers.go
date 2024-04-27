package mathfunctions

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
