package main
import (
	"fmt"
	"runtime"
)
func stackExample() {
	stackSlice := make([]byte,512)
	s := runtime.Stack(stackSlice,false)
	fmt.Printf("\n %s",stackSlice[0:s])
}
func first() {
	second()
}
func second() {
	third()
}
func third() {
	for c:=0;c<5;c++ {
		fmt.Println(runtime.Caller(c))
	}
}
func main() {
	fmt.Println("####stack###")
	stackExample()
	fmt.Println("\n####caller###")
	first()
}