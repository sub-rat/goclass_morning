package exapmle2array

import "fmt"

func Init() {
	// initArrayExample()
	initSliceExample()
}

func initArrayExample() { //index
	var arrayExample [3]int //[0,		1,		2]  // length - 1;
	fmt.Println(arrayExample)
	arrayExample[0] = 10
	arrayExample[1] = 11
	arrayExample[2] = 12

	fmt.Println(arrayExample)

	arrayExample2 := [5]int{10, 12, 23, 25, 40}
	arrayExample2[1] = 15
	fmt.Println(arrayExample2)

	for i := 0; i < len(arrayExample2); i++ {
		fmt.Printf("Index = %d and Value = %d\n", i, arrayExample2[i])
	}

	arrayExample3 := [3]string{"abc", "def", "ghk"}
	fmt.Println(arrayExample3)

	var arrayExample4 [2][3]int // [ [ 10,11,12 ], [ 15,24,11 ] ]
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arrayExample4[i][j] = i + j
			fmt.Println(arrayExample4[i][j])
		}
	}
	fmt.Println("Two D array = ", arrayExample4)
}
