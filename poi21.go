package main
import "fmt"
func setto10( a *int){
	a = new(int)
	*a =10
}
func main(){
	a :=20
	fmt.Println(a)
	setto10(&a)
	fmt.Println(a)
}