package examplefunctions

import "fmt"

func InitExampleClosures() {

	// anonymous functions
	bar := func() {
		fmt.Println("====================")
	}
	fmt.Println("hello")
	bar()
	fmt.Println("world")
	bar()
	// bar()
	// bar()
	// bar()

	seq := intSeq()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	bar()
	seq1 := intSeq()
	fmt.Println(seq1())
	fmt.Println(seq1())

	func() {
		fmt.Println("i am inside anonymous function")
	}()

}

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
