package main
import "fmt"
type Address struct {
	city, state string
}
type person struct {
	name string
	age int
	address Address
	
}
func main(){
	var p person
	p.name ="ritika"
	p.age = 23
	p.address = Address {
		city :"india",
		state :"new delhi",
	}
	fmt.Println("name",p.name)
	fmt.Println("age:",p.age)
	fmt.Println("city:",p.address.city)
	fmt.Println("state:",p.address.state)
}