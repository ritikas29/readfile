package main
import (
	"fmt"
	"errors"
)
type I interface {
    func1()(int, error)
 	func2()
	func3(string, string) bool
}
type P struct {
    q int
    err error
    j string
    k,l string
}
func( p *P ) func2(){
    fmt.Println(p.j)
}
func ( o *P) func1(){
    fmt.Println(o.q)
}
func (m *P) func3(){
 fmt.Println(m.k,m.l)
}
func main(){
	c :=&P{ q:1 }
	c.j = "hello"
	i := P{err :errors.New("explicit error")}
	c.func1()
	//b.func2()
	i.func3()
	fmt.Printf("This is the value from c instance %s\n", c.j)
}