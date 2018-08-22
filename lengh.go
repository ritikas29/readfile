package main
import "fmt"
func main() {
	channelEx := make(chan int,5)
	fmt.Printf("\n channel:%d",len(channelEx))
	channelEx <- 0
	channelEx <- 1
	channelEx <-2
	fmt.Printf("\n Channel:%d",len(channelEx))
	var pointerEx *[5]string
	fmt.Printf("\n pointer :%d", len(pointerEx))
	mapEX := make(map[string]int)
	mapEX["A"] = 10
	mapEX["B"] = 20
	mapEX[" C"] = 30
	fmt.Printf("\n Map: %d",len(mapEX))
	sliceEx := make([]int,10)
	fmt.Printlf("\n slice :%d",len(sliceEx))
	
}