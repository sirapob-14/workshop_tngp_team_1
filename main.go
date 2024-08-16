package main

import "fmt"

func main() {
	minus(50, 40)
	mod(5, 2)
}
func minus(a int, b int) {
	fmt.Printf("%d - %d = %d + %d = %d", a, b, a, b, a-b)
}

func mod(a, b int) {
	result := a % b
	output := fmt.Sprintf("a %% b = %d %% %d = %d", a, b, result)
	fmt.Println(output)
}
