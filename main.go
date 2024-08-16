package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	mod(5, 2)
}

func mod(a, b int) {
	result := a % b
	output := fmt.Sprintf("a %% b = %d %% %d = %d", a, b, result)
	fmt.Println(output)
}
