package main

import "fmt"

func divide(a, b int){
	if b == 0 {
		fmt.Println("Cannot divided by 0")
	}

	fmt.Printf("a/b = %d/%d = %d \n",a ,b, a/b)
 }
 
func main() {
	fmt.Println("Hello world!")
	divide(100,10)
}