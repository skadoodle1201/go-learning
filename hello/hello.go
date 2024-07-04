package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var myNum int = 10
	fmt.Println(myNum)

	var difFloat float64 = 1233211.9
	fmt.Println(difFloat)

	const pi float32 = 3.14
	fmt.Println(pi)

	multiply := difFloat * float64(pi)
	fmt.Println(multiply)

}
