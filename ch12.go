package main
import (
	"fmt"
	"time"
)	
func hello(done chan bool) {
	fmt.Println("hello go rountine is going to sleep")
	time.Sleep(4*time.Second)
	fmt.Println("hello go rountine awake and going to write done")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("main going to call hello go ")
	go hello(done)
	<-done
	fmt.Println("main received data")
}