package main
import "fmt"
func addone( a int) int {
	return a+1
}
func addtwo(a int) int {
	return a+2
}
func printoperation(a int, f func(int)int) {
	fmt.Println(f(a))
}
func main() {
	printoperation(1,addone)
	printoperation(1,addtwo)
}