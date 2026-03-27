package calculator

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func PrintResult(result int) {
	fmt.Println("Result:", result)
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}
}
