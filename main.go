package main

import (
	"fmt"
	"math"
)

func main() {
	minus(50, 40)
	squareroot(4)
	divide(100, 10)
	sum(2, 3)
	multiply(40, 50)
}
func minus(a int, b int) {
	fmt.Printf("a - b = %d + %d = %d\n", a, b, a-b)
}

func squareroot(a int) {
	sqrt := math.Sqrt(float64(a))
	fmt.Printf("Sqrt(a) = Sqrt(%d) = %.2f\n", a, sqrt)
}

func divide(a, b int) {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return
	}

	fmt.Printf("a/b = %d/%d = %d \n", a, b, a/b)
}

func sum(a int, b int) {
	sum := a + b
	fmt.Println("sum is: ", sum)
}

func multiply(a, b int) int {
	result := a * b
	fmt.Println("a * b =", a, "*", b, "=", result)
	return result
}
