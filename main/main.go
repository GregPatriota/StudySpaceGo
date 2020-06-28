package main

import (
	_ "chapter1/calc"
	_ "chapter2/strcon"
	"fmt"
	"log"
	"os"
)

func doPanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recover with: ", e)
		}
	}()
	panic("Just panicking for the sake of demo")
	fmt.Println("This will never be called")
}

func main() {
	_, err := os.Open("readme_log.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting to panic")
	doPanic()
	fmt.Println("Program regains control after panic recover")
}

// func main() {
// 	// x := make([]int, 1, 1)
// 	// x[0] = 10
// 	// fmt.Println(x)
// 	// fmt.Println("Length is", len(x))
// 	// fmt.Println("Capacity is", cap(x))
// 	// x = append(x, 20, 30, 40, 50)
// 	// fmt.Println(x)
// 	// fmt.Println("Length is", len(x))
// 	// fmt.Println("Capacity is", cap(x))
// 	// fmt.Println(x)
// 	// x = append(x, 60)
// 	// fmt.Println("Length is", len(x))
// 	// fmt.Println("Capacity is", cap(x))
// 	// fmt.Println(x)
// 	// x = append(x, 60, 70, 80, 90, 100)
// 	// fmt.Println("Length is", len(x))
// 	// fmt.Println("Capacity is", cap(x))
// 	// fmt.Println(x)

// 	dict := make(map[string]string)
// 	dict["go"] = "Golang"
// 	dict["cs"] = "CSharp"
// 	dict["rb"] = "Ruby"
// 	dict["py"] = "Python"
// 	dict["js"] = "JavaScript"
// 	for k, v := range dict {
// 		fmt.Printf("Key: %s Value: %s\n", k, v)
// 	}

// 	if lan, ok := dict["ts"]; ok {
// 		fmt.Println(lan, ok)
// 	}

// 	session, err := mgo.Dial("localhost") // Create MongoDB Session object
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()
// }
