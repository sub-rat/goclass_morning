package goroutinesexample

import (
	"fmt"
	"time"
)

func InitGorountineExample() {
	routine("init routine example")

	go routine("goroutine")

	go func() {
		fmt.Println("annomyous function")
	}()

	go routine("testing")

	time.Sleep(time.Millisecond * 10)
	fmt.Println("done ")

	initChannelExample()
}

func routine(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, " ", i)
	}
}
