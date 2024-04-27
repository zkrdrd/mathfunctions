package main

import (
	"fmt"
	"mathfunctions"
)

func main() {
	//cli.ShowMenu()
	c := mathfunctions.FindMaxNumber(0, 1, 2, 3, 2, 1, 2, 3, 1, 2)
	fmt.Print(c)
}
