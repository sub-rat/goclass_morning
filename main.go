package main

import "fmt"

const PI = 3.14 // constant global variables

var i, j, k = 1, 2.2, "abc" // global variables

func main() {
	x := 10
	fmt.Println("hello go")
	var age uint8 = 100 // memory 8 bit= 1 byte ; local variable
	var B float32 = 1.1 // local variable
	fmt.Println("Age = ", age)
	fmt.Println("Value of b = ", B)
	i = 2
	fmt.Println(i, j, k, PI, x)

	y := 10.10
	fmt.Printf("Type of y = %T \n", y)

	example()
}

func example() {
	age := 80 // local variable
	fmt.Println("IN example = ", age, PI)
}
