package main

import "fmt"


func main() {
	minus(50, 40)
	mod(5, 2)
	divide(100,10)
}
func minus(a int, b int) {
	fmt.Printf("a - b = %d + %d = %d\n", a, b, a-b)
}

func mod(a, b int) {
	result := a % b
	output := fmt.Sprintf("a %% b = %d %% %d = %d", a, b, result)
	fmt.Println(output)
}

func divide(a, b int){
	if b == 0 {
		fmt.Println("Cannot divided by 0")
	}

	fmt.Printf("a/b = %d/%d = %d \n",a ,b, a/b)
 }
 