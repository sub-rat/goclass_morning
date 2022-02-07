package initials

import (
	"fmt"
	"strings"
)

const PI = 3.14 // constant global variables

var i, j, k = 2, 2, "abc" // global variables

func nit() {
	x := 10
	fmt.Println("hello go")
	var age uint8 = 100 // memory 8 bit= 1 byte ; local variable
	var B float32 = 1.1 // local variable
	fmt.Println("Age = ", age)
	fmt.Println("Value of b = ", B)
	i = 2
	i += 2 // i = i + 2
	fmt.Println(i, j, k, PI, x)

	y := 10.10
	fmt.Printf("Type of y = %T \n", y)

	example()

	example2()
	ifElseExample()
}

func ifElseExample() {
	// i = 1 j= 2   1 >= 2 ==> false
	age := 36

	if age <= 1 {
		fmt.Println("infant")
	} else if age > 1 && age <= 25 {
		fmt.Println("teen")
	} else if age > 25 && age <= 64 {
		fmt.Println("youth")
	} else {
		fmt.Println("old")
	}

	switch {
	case age <= 1:
		fmt.Println("infant")
	case age > 1 && age <= 25:
		fmt.Println("teen")
	case age > 25 && age <= 64:
		fmt.Println("youth")
	default:
		fmt.Println("old")
	}

	i := "go"
	fmt.Println(i)
	switch i {
	case "golang":
		fmt.Println("sum = 10")
	case "go":
		fmt.Println("sum = 8")
	case "abc":
		fmt.Println("sum = 5")
	default:
		fmt.Println("default")
	}

}

func example() {
	age := 80 // local variable
	fmt.Println("IN example = ", age, PI)
}

func example2() {
	/*
	   	000 = 0
	   	001 = 1
	   	010 = 2
	   	011 = 3
	   	100 = 4
	  00101 = 5  left shift  5<<2 == 01010 =   1*2^3 + 0*2^2 + 1*2^1 + 0 * 2^0
	   	110 = 6  == 1 * 2^2 + 1 * 2^1 + 0 *2^0
	   	111 = 7
	   1000 = 8
	*/

	var a = 5
	a = a << 1
	fmt.Println(a)
	var b, c = 2, 3
	d := b & c
	fmt.Println(d)
	var greeting1 = "Hello"
	var greeting2 = " World!"
	var check bool = true
	greeting := greeting1 + greeting2
	greetingUpper := strings.ToUpper(greeting)
	greetingLower := strings.ToLower(greeting)
	fmt.Println(greeting, check)
	fmt.Println(greetingUpper)
	fmt.Println(greetingLower)

	var p *int

	i := 42

	p = &i
	fmt.Println("i value = ", i)
	fmt.Println("Pointer value = ", *p)

	*p = 12
	fmt.Println("i value = ", i)
	fmt.Println("Pointer value = ", *p)

}
