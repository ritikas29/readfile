package main 
import "fmt"
var patchableFunc = func (i int ) int  {
	return i*2
}
func testFunc() {
	originalFunc := patchableFunc
	defer func() {
		patchableFunc= originalFunc
	}()
	fmt.Println(patchableFunc(1))
	patchableFunc = func(i int) int {
		return i*3
	}
	fmt.Println(patchableFunc(2))
}
func main() {
	fmt.Println(patchableFunc(3))
	testFunc()
	fmt.Println(patchableFunc(4))
}