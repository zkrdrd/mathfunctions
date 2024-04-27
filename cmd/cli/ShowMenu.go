package cli

import (
	"fmt"
	"mathfunctions"
	"os"
	"strconv"
)

// запись значений в массив
func appendValuesToArray() {
	mathfunctions.Temp = ""
	mathfunctions.Numbers = nil
	mathfunctions.FindSpace = 0
	fmt.Print("Введите числа через пробел: ")
	mathfunctions.GetText.Scan()
	mathfunctions.Temp = mathfunctions.GetText.Text()
	if mathfunctions.Temp > "" {
		for i := 0; i < len(mathfunctions.Temp); i++ {
			if mathfunctions.Temp[i] == mathfunctions.Space {
				val, err := strconv.ParseInt(mathfunctions.Temp[mathfunctions.FindSpace:i], 10, 64)
				if err != nil {
					fmt.Print(err)
					panic(err)
				}
				mathfunctions.Numbers = append(mathfunctions.Numbers, val)
				mathfunctions.FindSpace = i + 1
			} else if i+1 == len(mathfunctions.Temp) {
				val, err := strconv.ParseInt(mathfunctions.Temp[mathfunctions.FindSpace:i+1], 10, 64)
				if err != nil {
					fmt.Print(err)
					panic(err)
				}
				mathfunctions.Numbers = append(mathfunctions.Numbers, val)
			}
		}
	}
}

// интерактивное основное меню
func ShowMenu() {
	for {
		fmt.Print(`
1 - Найти min
2 - Найти max
Чётность в теории чисел — характеристика целого числа
3 - Сумма четных чисел (только целые числа, дробные числа округлит)
4 - Сумма нечетных чисел (только целые числа, дробные числа округлит)
0 - выход
Выбирите действие: `)
		fmt.Scanf("%s\n", &mathfunctions.Temp)
		if mathfunctions.Temp != "" {
			i, err := strconv.ParseInt(mathfunctions.Temp, 10, 64)
			if err == nil {
				mathfunctions.Selecter = i
				break
			} else {
				fmt.Println("Это должно быть целое число из доступных (1, 2, 3, 0).")
			}
		}
	}

	if mathfunctions.Selecter == 0 {
		os.Exit(0)
	} else {
		for {
			appendValuesToArray()
			break
		}
		if mathfunctions.Selecter == 1 {
			mathfunctions.Result = mathfunctions.FindMinNumber(mathfunctions.Numbers...)
		} else if mathfunctions.Selecter == 2 {
			mathfunctions.Result = mathfunctions.FindMaxNumber(mathfunctions.Numbers...)
		} else if mathfunctions.Selecter == 3 {
			mathfunctions.Result = mathfunctions.SumEvenNumber(mathfunctions.Numbers...)
		} else if mathfunctions.Selecter == 4 {
			mathfunctions.Result = mathfunctions.SumNoEvenNumber(mathfunctions.Numbers...)
		}
		fmt.Print(mathfunctions.Result)
	}
	now()
}

// интерактивное дополнительное меню
func now() string {
	for {
		fmt.Print("\nВыбирите действие\n1 - еще раз\n0 - выход\n:")
		fmt.Scanf("%s\n", &mathfunctions.Temp)
		i, err := strconv.ParseInt(mathfunctions.Temp, 10, 32)
		if err == nil {
			if i == 1 {
				ShowMenu()
			}
			if i == 0 {
				os.Exit(0)
			}
		}
	}
}
