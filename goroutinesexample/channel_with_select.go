package goroutinesexample

import (
	"fmt"
	"time"
)

func initChannelWithSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Millisecond)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(4 * time.Millisecond)
		c2 <- "Two"
	}()

	// fmt.Println(<-c2)
	// fmt.Println(<-c1)

	select {
	case msg1 := <-c2:
		fmt.Println(msg1)
	case msg2 := <-c1:
		fmt.Println(msg2)
	}

	select {
	case msg1 := <-c2:
		fmt.Println(msg1)
	case msg2 := <-c1:
		fmt.Println(msg2)
	case <-time.After(time.Millisecond * 3):
		fmt.Println("timeout")
		// default:
		// 	fmt.Println("default")
	}

	// same as above but by using for loop

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <-c2:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-c1:
	// 		fmt.Println(msg2)
	// 		// default:
	// 		// 	fmt.Println("default")
	// 	}
	// }

}
