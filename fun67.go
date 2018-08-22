package main 
import (
	"fmt"
	"reflect"
)
type Num interface{}
type List func(...Num) Num
func (list List) Apply(a Num) List {
	return func(values ...Num) Num {
		values = append([]Num{a},values...)
		return list(values...)
	}
}
func wrapper(fn Num) List {
	return func(a ...Num) Num {
		return caller(fn,a)
	}
}
func caller(fn Num,args []Num) Num {
	var vs []reflect.Value
	for _, arg := range args {
		vs = append(vs,reflect.ValueOf(arg))
	} 
	fun := reflect.ValueOf(fn)
	return fun.Call(vs)[0].Interface()
}
func mixture(elems []Num,fn Num) List {
	var list List = wrapper(fn)
	for _ ,elem := range elems {
		list = list.Apply(elem)
	}
	return list 
}
func addition(l int,m int,n int,o int) int {
	return l * m +n * o 
}
func main() {
	fmt.Println(mixture([]Num{1,2,3},addition)(4))
	fmt.Println(mixture([]Num{1,2,3,4},addition)())
}


