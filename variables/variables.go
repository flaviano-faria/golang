package main

import "fmt"

func main() {
	x := 10
	y := "hello world"

	fmt.Println(x, y)

	//print variable value and type
	fmt.Printf("x: %v, %T\n", x, x)

	//println returns two values: numberOfBytes and errors
	numberOfBytes, errors := fmt.Printf("x: %v, %T\n", x, x)
	fmt.Println(numberOfBytes, errors)

	byteNumber, _ := fmt.Println("teste")
	fmt.Println(byteNumber)
}
