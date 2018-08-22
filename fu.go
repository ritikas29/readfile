package main 
import "fmt"
type Showprice interface {
	display(int)
}
func display(s Showprice,price int) {
	s.display(price)
}
type Car struct {
	Price int
} 
func (c Car) display(price int) {
	c.Price = price  
}
type Bike struct {
	Price int 
}
func(b *Bike) display(price int) {
	b.Price=price 
}
func main() {
	ca :=Car{Price:2000}
	bi := Bike{Price:400}
	display(ca,3000)
	display(&bi,600)
	fmt.Println("car:",ca.Price)
	fmt.Println("bike:",bi.Price)
}