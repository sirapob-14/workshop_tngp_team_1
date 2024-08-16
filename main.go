package main

import "fmt"
 
func main() {
	fmt.Println("Hello world!")
	sum(2, 3)
}

func sum(a int, b int) {
	sum := a+b
	fmt.Println(sum)
}