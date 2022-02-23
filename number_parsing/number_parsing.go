package number_parsing

import (
	"fmt"
	"strconv"
)

func InitNumberParsing() {
	testValue := "100" // base 2 == binary(1,0) ; decimal => base 10 (0-9)
	fmt.Println(testValue)
	fmt.Printf("Type of value = %T\n", testValue)
	value, err := strconv.ParseInt(testValue, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
	fmt.Printf("Type of value = %T\n", value)

	testValue = "1.1"
	fmt.Println(testValue)
	fmt.Printf("Type of value = %T\n", testValue)
	value1, err := strconv.ParseFloat(testValue, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value1)
	fmt.Printf("Type of value = %T\n", value1)

	v1, _ := strconv.Atoi("11")
	fmt.Println(v1)
	initRandomNumbers()
}
