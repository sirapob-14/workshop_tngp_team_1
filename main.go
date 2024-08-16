package main

import "fmt"
 
func main() {
	minus(50, 40)
}
func minus(a int, b int) {

	fmt.Printf("%d - %d = %d + %d = %d", a, b, a, b, a-b)
}
