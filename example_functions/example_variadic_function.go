package examplefunctions

import "fmt"

func InitVariadicFunction() {
	nums := []int{1, 2, 3, 4, 5}
	a := []int{6, 7}
	nums = append(nums, a...)
	calculate_sum(nums...)
	calculate_sum(1, 2, 3, 4, 5, 6)
	calculate_sum(5, 5)
	calculate_sum(5, 6, 7, 8)

	testInterface(1, "abc", 1.1, 20, true)
}

func calculate_sum(nums ...int) { // receive n no of arguments
	sum := 0
	for _, num := range nums {
		sum += num // sum = sum + num
	}
	fmt.Println("sum = ", sum)
}

func testInterface(a ...interface{}) {
	for _, value := range a {
		fmt.Println(value)
	}
}
