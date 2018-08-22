package main
import "fmt"
func demoEllipsis(arg ...int) {
	fmt.Println("lenght of arguments:",len(arg))
	fmt.Println("arguments : %+v\n",arg)
}
func main() {
	fmt.Println("\n example 1")
	demoEllipsis([]int{1,2,3}...)
	fmt.Println("\n example 2")
	demoEllipsis(1,2,3,4,5)
	fmt.Println("\n example 3")
	demoEllipsis()
	fmt.Println("\n dynamic array")
	intArray := [...] int {1,2,3,4,5,6,7}
	fmt.Println("lenght: ",len(intArray))
	fmt.Println("capacity:",cap(intArray))
	fmt.Println(intArray)
	fmt.Println("\n interface")
	interfaceEx := []interface{}{"australia","canada","japan"}
	fmt.Println(interfaceEx...)
}