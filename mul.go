package main
import "fmt"
type catalog interface {
	shipping() float64
	tax() float64
}
type configurable struct {
	name string
	price , qty float64
}
func ( c *configurable) tax() float64 {
	return c.price * c.qty *0.05
}
func (c *configurable) shipping() float64{
	return c.qty *5
}
type download struct{
	name string
	price,qty float64
}
func ( d *download) tax() float64{
	return d.price *d.qty *0.07
}
type simple struct{
	name string
	price , qty float64
}
func(s *simple) tax() float64{
	return s.price*s.qty*0.03
}
func(s *simple) shipping() float64{
	return s.qty*3
}
func main() {
	tshirt := configurable{}
	tshirt.price =250
	tshirt.qty = 2
	fmt.Println("configurable product")
	fmt.Println("shipping charge:",tshirt.shipping())
	fmt.Println("tax:",tshirt.tax())
	mobile := simple{"samsung ",10,25}
	fmt.Println("\n simple product")
	fmt.Println("shipping charge:",mobile.shipping())
	fmt.Println("tax:",mobile.tax())
	book := download{"golang",20,1}
	fmt.Println("\n downloadable product")
	fmt.Println("tax :",book.tax())
	
}


