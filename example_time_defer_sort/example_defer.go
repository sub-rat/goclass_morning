package example_time_defer_sort

import "fmt"

func InitDeferExample() {
	defer fmt.Println("Stopping the program 1") // stack1[1]
	defer fmt.Println("Stopping the program 2") // stack1[1,2]
	defer fmt.Println("Stopping the program 3") // stact1[1,2,3]
	a := 10                                     // exec 1
	b := 13                                     // exec 2
	fmt.Println(a, b)                           // exec 3  10 13
	testDefer()                                 // exec 4 go to the func testDefer()
	fmt.Println("Hello golang!")                // exec 10 Hello golang!
	//exec 11  stack1[1,2] 3 => Stopping the program 3
	//exec 12  stack1[1] 2 => Stopping the program 2
	//exec 13  stack1[] 1 => Stopping the program 1
	// return to main()
}

func testDefer() {
	defer fmt.Println("Stopping func testDefer 1") // stack2[1]
	defer fmt.Println("Stopping func testDefer 2") // stack2[1,2]
	defer fmt.Println("Stopping func testDefer 3") // stack2[1,2,3]
	fmt.Println("Runnig test Defer function")      // execute 5 Runnig test Defer function
	fmt.Println("3 * 2 =", 3*2)                    // exec 6  3 * 2 = 6
	//exec 7  stack2[1,2] 3 => Stopping func testDefer 3
	//exec 8 stack2[1] 2 => Stopping func testDefer 2
	//exec 9 stack2[] 1 => Stopping func testDefer 1
	// return to InitDeferExample()
}
