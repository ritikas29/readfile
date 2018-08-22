package main
import "fmt"
type I interface {
    func1()(int, error)
 	func2()
	func3(string, string) bool
}
type P struct{
	q int
	err error
	j,k,l string
}
func ( p P) func2(){
fmt.Println(p.j)
}
func( o P) func1(){
	fmt.Println(o.q)
	fmt.Println(o.err)
}
func( m P ) func3(){
	fmt.Println( m.k,m.l)
}
func main() {
 c := P{ q:1 }
 c{ j :"hello"}  
var i I = P {k:"hello", l:"welcome"}
 c.func1() 
 b.func2()
 i.func3()

}