package exampleerrorhandling

import (
	"errors"
	"fmt"
)

func InitExampleErrorHandling() {
	n, _ := fmt.Println("Error Example")
	fmt.Println("value of n=", n)

	pNo, err := postiveNo(10)
	panicErrorHandle(err)
	fmt.Println("postive no = ", pNo)

	pNo, err = postiveNo(-1)
	panicErrorHandle(err)
	fmt.Println("postive no = ", pNo)
}

func panicErrorHandle(err error) {
	if err != nil {
		panic(err)
	}
}

func postiveNo(arg int) (int, error) {
	if arg < 0 {
		return 0, errors.New(" args cannot be negative ")
	}
	return arg, nil
}
