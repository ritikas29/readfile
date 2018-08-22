package main 
import "fmt"
func counter(id int,channel chan int,closer bool) {
	for i :=0 ;i<10000000;i++ {
	fmt.Println("process",id,"send",i)
	channel <- 1
	}
	if closer {close(channel)}
}
func main() {
	channel := make(chan int)
	go counter(1,channel,false)
	go counter(2,channel,true)
	x :=0
	for i := range channel {
		fmt.Println("receiving")
		x += i       
	}
	fmt.Println(x)

}