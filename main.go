package main

import "fmt"

func main() {
	minus(50, 40)
	mod(5, 2)
	divide(100, 10)
	sum(2, 3)
	multiply(40, 50)
}
func minus(a int, b int) {
	fmt.Printf("a - b = %d + %d = %d\n", a, b, a-b)
}

func mod(a, b int) {
	result := a % b
	output := fmt.Sprintf("a %% b = %d %% %d = %d", a, b, result)
	fmt.Println(output)
}

func divide(a, b int) {
	if b == 0 {
		fmt.Println("Cannot divided by 0")
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
