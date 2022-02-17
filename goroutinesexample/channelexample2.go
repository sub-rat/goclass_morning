package goroutinesexample

import (
	"fmt"
)

func InitChannelExample2() {

	done := make(chan bool, 2)

	fmt.Println("Starting main thread...")
	go worker(done)
	<-done
	initChannelWithDirection()
	initChannelWithSelect()
}

func worker(done chan bool) {
	fmt.Println("working on thread .....")
	sum := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10000; j++ {
			for k := 0; k < 10000; k++ {
				sum = sum + k
			}
		}
		fmt.Println("Value of i =", i)
	}

	for hour := 0; hour < 24; hour++ {
		for min := 0; min < 60; min++ {
			for sec := 0; sec < 60; sec++ {
				//sec
			}
			// min
		}
		//hour
	}

	done <- true
}
