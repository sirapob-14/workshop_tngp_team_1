package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	a := 40
	b := 50
	multiply(a,b)
}

func multiply(a, b int) int {
	result := a * b
	fmt.Println("a * b =",a,"*",b,"=",result)
	return result
}
