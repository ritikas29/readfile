package main
import "fmt"
func makeadder(b int)func(int)int{
	return func(a int)int{
		return a+b
	}
}
func makedouble(f func(int)int)func(int)int {
	return func(a int)int{
		b:= f(a)
		return b*2
	}
}
func main() {
	addone :=makeadder(1)
	doubleaddone := makedouble(addone)
	fmt.Println(addone(1))
	fmt.Println(doubleaddone(1))
}