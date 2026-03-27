package main

import (
	"fmt"

	"test1/calculator"
)

func main() {
	var age int = 10
	var name string = "John"
	var isStudent bool = true
	fmt.Println(age, name, isStudent)

	result := calculator.Add(3, 5)
	fmt.Println("3 + 5 =", result)
}
