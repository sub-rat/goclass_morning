package example_time_defer_sort

import (
	"fmt"
	"time"
)

func InitExampleTime() {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Time = ", t)
			}
		}
	}()

	// go timeTicker(done, ticker)

	time.Sleep(time.Minute)
	ticker.Stop()
	done <- true
	fmt.Println("Program stops")
}

// func timeTicker(done chan bool, ticker *time.Ticker) {
// 	for {
// 		select {
// 		case <-done:
// 			return
// 		case t := <-ticker.C:
// 			fmt.Println("AT TimeTicker Func Time = ", t)
// 		}
// 	}
// }
