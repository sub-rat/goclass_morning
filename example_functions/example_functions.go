package examplefunctions

import (
	"errors"
	"fmt"
)

func InitExampleFunctions() {
	s, err := sum(5, 10)
	if err != nil {
		fmt.Println("Error occured = ", err)
	} else {
		fmt.Println("sum = ", s)
	}
	s, err = sum(-6, 10)
	if err != nil {
		fmt.Println("Error occured = ", err)
	} else {
		fmt.Println("sum = ", s)
	}
	s, _ = sum(8, 10)
	fmt.Println("sum = ", s)
	x, y := "kathmandu", "Dhulikhel"
	fmt.Println(x, y)
	// x, y = swap(x, y)
	fmt.Println(x, y)
}

func sum(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a and b cannot be 0")
	}
	sum := a + b
	return sum, nil
}

// func swap(x, y string) (string, string) {
// 	return y, x
// }
