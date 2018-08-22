package main
import "fmt"
func divandremainder(a int ,b int) (int,int) {
	return a/b,a%b
}
func main() {
	div,remainder := divandremainder(2,3)
	fmt.Println(div,remainder)
	div,_ = divandremainder(10,4)
	fmt.Println(div)
	_,remainder = divandremainder(100,-100)
	fmt.Println(remainder)
	divandremainder(-1,20)
}
