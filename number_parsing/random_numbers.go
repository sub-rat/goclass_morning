package number_parsing

import (
	"fmt"
	"math/rand"
)

func initRandomNumbers() {
	fmt.Println("======== Random Numbers =========")
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(1000))
	fmt.Println(rand.Intn(55))
	fmt.Println(rand.Intn(700))

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64()*5 + 5)
	fmt.Println(rand.Float64())

}
