package goroutinesexample

import "fmt"

//Topic Channel Direction
func initChannelWithDirection() {
	mCh1 := make(chan string, 1)
	mCh2 := make(chan string, 1)
	mCh3 := make(chan string, 1)

	mCh3 <- "channel 3 at start"

	go routine3(mCh1, "message from routine3")
	go routine4(mCh1, mCh2)
	// fmt.Println(<-mCh1)
	fmt.Println(<-mCh2)
	fmt.Println(<-mCh3)

}

// mChannel is send only type
func routine3(mChannel chan<- string, message string) {
	fmt.Println("working in routine 3")
	// <-mChannel
	mChannel <- message
}

//  mChannel 1 is receive only type and mchannel2 is send only type
func routine4(mChannel1 <-chan string, mChannel2 chan<- string) {
	fmt.Println("working in routine 4")
	msg := <-mChannel1
	mChannel2 <- msg
}
