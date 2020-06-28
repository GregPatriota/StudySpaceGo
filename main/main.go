package main

import (
	_ "chapter1/calc"
	_ "chapter2/strcon"
	"fmt"
)

func main() {
	x := make([]int, 1, 1)
	x[0] = 10
	fmt.Println(x)
	fmt.Println("Length is", len(x))
	fmt.Println("Capacity is", cap(x))
	x = append(x, 20, 30, 40, 50)
	fmt.Println(x)
	fmt.Println("Length is", len(x))
	fmt.Println("Capacity is", cap(x))
	fmt.Println(x)
	x = append(x, 60)
	fmt.Println("Length is", len(x))
	fmt.Println("Capacity is", cap(x))
	fmt.Println(x)
	x = append(x, 60, 70, 80, 90, 100)
	fmt.Println("Length is", len(x))
	fmt.Println("Capacity is", cap(x))
	fmt.Println(x)
}
