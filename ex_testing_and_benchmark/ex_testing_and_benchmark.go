package extestingandbenchmark

import (
	"errors"
)

func InitTestingAndBenchmark() {
}

// calculate sum of a and b also , a and b must be postive, return error if a or b is negative
func Sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a and b must be postive")
	}
	return a + b, nil
}

func Loop() {
	for i := 0; i < 10; i++ {
		// fmt.Println("tesing")
	}
}
