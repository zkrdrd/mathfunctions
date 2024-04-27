package mathfunctions

import (
	"bufio"
	"os"
)

// ascii code space
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
