package cli

import (
	"bufio"
	"fmt"
	"mathfunctions"
	"os"
	"strconv"
)

// ascii byte space
const Space byte = 32

var (
	// результат работы функции
	Result int64
	// переменная на несколько вариантов
	Temp string
	// результат выбора интерактивного меню
	Selecter int64
	// поиск пробела
	FindSpace int
	// сканер
	GetText = bufio.NewScanner(os.Stdin)
)

// запись значений в массив
func appendValuesToArray() {
	Temp = ""
	mathfunctions.Numbers = nil
	FindSpace = 0
	fmt.Print("Введите числа через пробел: ")
	GetText.Scan()
	Temp = GetText.Text()
	if Temp > "" {
		for i := 0; i < len(Temp); i++ {
			if Temp[i] == Space {
				val, err := strconv.ParseInt(Temp[FindSpace:i], 10, 64)
				if err != nil {
					fmt.Print(err)
					panic(err)
				}
				mathfunctions.Numbers = append(mathfunctions.Numbers, val)
				FindSpace = i + 1
			} else if i+1 == len(Temp) {
				val, err := strconv.ParseInt(Temp[FindSpace:i+1], 10, 64)
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
		fmt.Scanf("%s\n", &Temp)
		if Temp != "" {
			i, err := strconv.ParseInt(Temp, 10, 64)
			if err == nil {
				Selecter = i
				break
			} else {
				fmt.Println("Это должно быть целое число из доступных (1, 2, 3, 0).")
			}
		}
	}

	if Selecter == 0 {
		os.Exit(0)
	} else {
		for {
			appendValuesToArray()
			break
		}
		if Selecter == 1 {
			Result = mathfunctions.FindMinNumber(mathfunctions.Numbers...)
		} else if Selecter == 2 {
			Result = mathfunctions.FindMaxNumber(mathfunctions.Numbers...)
		} else if Selecter == 3 {
			Result = mathfunctions.SumEvenNumber(mathfunctions.Numbers...)
		} else if Selecter == 4 {
			Result = mathfunctions.SumNoEvenNumber(mathfunctions.Numbers...)
		}
		fmt.Print(Result)
	}
	now()
}

// интерактивное дополнительное меню
func now() string {
	for {
		fmt.Print("\nВыбирите действие\n1 - еще раз\n0 - выход\n:")
		fmt.Scanf("%s\n", &Temp)
		i, err := strconv.ParseInt(Temp, 10, 32)
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
