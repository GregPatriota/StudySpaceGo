package main

import (
	calc "chapter1/calc"
	strcon "chapter2/strcon"
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
	fmt.Println(calc.Add(1, 2))
	s := strcon.SwapCase("Gopher")
	fmt.Println("Converted string is :", s)
}
