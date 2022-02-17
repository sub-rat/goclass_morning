package example_time_defer_sort

import (
	"fmt"
	"sort"
)

func InitExampleSort() {
	//Sorting
	strs := []string{"d", "a", "c", "b"}
	fmt.Println("Strings = ", strs)
	sort.Strings(strs)
	fmt.Println("Sorted Strings = ", strs)

	// panic("error at mid of the program")

	ints := []int{1, 2, 3, 4}
	fmt.Println("Integers = ", ints)
	s := sort.IntsAreSorted(ints)
	if !s {
		sort.Ints(ints)
		fmt.Println("Sorted Integers = ", ints)
	}

}
