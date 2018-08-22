package main
import "fmt"
func zero(x int) {
	x=0
}
func main(){
	x :=34
	zero(x)
	fmt.Println(x)
}