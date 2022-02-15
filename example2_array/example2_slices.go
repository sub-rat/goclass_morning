package exapmle2array

import "fmt"

func initSliceExample() {
	s := make([]int, 3)
	fmt.Println(s)

	s[0] = 10
	s[1] = 20
	s[2] = 30

	fmt.Println("Slice s = ", s)
	fmt.Println("length of s = ", len(s))

	s = append(s, 40)
	s = append(s, 50)

	fmt.Println("Slice s = ", s)
	fmt.Println("length of s = ", len(s))

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println("Slice c = ", c)

	slice := s[2:5] // s[low:high] // low included // high not included
	fmt.Println("Slice 1 = ", slice)

	slice = s[:2] // [0 to 2-1]
	fmt.Println("Slice 2 = ", slice)

	slice = s[3:] // 3,4
	fmt.Println("Slice 2 = ", slice)

	slice = s[:]
	fmt.Println("Slice 2 = ", slice)

	ex2 := []string{"a", "b", "c"}
	fmt.Println(ex2)

	// var arrayExample4 [2][3]int // [ [ 10,11,12 ], [ 15,24,11 ] ]
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		arrayExample4[i][j] = i + j
	// 		fmt.Println(arrayExample4[i][j])
	// 	}
	// }

	// Two D with slices
	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		ilen := i + 1
		twoD[i] = make([]int, ilen)
		for j := 0; j < ilen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d slices = ", twoD)

}
