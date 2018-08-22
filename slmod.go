package main
import "fmt"
func main() {
	a :=[7]string{"mon","tues","wed","thus","fri","sat","sun"}
	slice1 :=a[1:]
	slice2 := a[3:]
	fmt.Println("..........before modification....")
	fmt.Println("a =",a)
	fmt.Println("slice1 =",slice1)
	fmt.Println("slice2 =",slice2)
	slice1[0] ="tues"
	slice1[1] ="wed"
	slice1[2]="thus"
	slice2[1]="fri"
	fmt.Println("\n........after modifications....")
	fmt.Println("a =",a)
	fmt.Println("slice1 =",slice1)
	fmt.Println("slice2 =",slice2)



}