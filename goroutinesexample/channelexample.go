package goroutinesexample

import "fmt"

func initChannelExample() {

	mesChannel := make(chan string, 2)
	mesChannel <- "Main Thread"
	// go func() {
	// 	mesChannel <- "message from annynomous go routine"
	// }()

	go routine2(mesChannel)

	msg := <-mesChannel
	fmt.Println(msg)
	msg = <-mesChannel
	fmt.Println(msg)

}

func routine2(mChannel chan string) {
	mChannel <- "message from annynomous go routine"
}
