package main 
import (
	"fmt"
)
type Runner interface {
	Run()
}
type Person string
func (p Person) Run() {
	fmt.Println(p,"is running")
}
type Dog struct {
	Name string
	Type string 
}
func(d Dog) Run() {
	fmt.Println("A",d.Type,"dog is running")
}
func start(r Runner) {
	r.Run()
}
func main() {
	p:=Person("ritika")
	start(p)
	d :=Dog{"ritika","huy"}
	start(d)
}