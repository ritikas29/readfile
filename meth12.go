package main
import "fmt"
type information interface{
	general()
	attributes()
	inventory()
}
type product struct {
	name,description string
	weight,price float64
	stock int
}
func (prd product) general(){
	fmt.Printf("\n %s",prd.name)
	fmt.Printf("\n %s\n",prd.description)
	fmt.Println(prd.price)
}
func (prd product) attributes(){
	fmt.Println(prd.weight)
}
type mobile struct{
	product
	displayfeatures []string
	processorfeatures []string
}
func (mob mobile) attributes(){
	mob.product.attributes()
	fmt.Println("\n display features:")
	for _, key := range mob.displayfeatures {
		fmt.Println(key)
	}
	fmt.Println("\n processor features:")
	for _,key := range mob.processorfeatures {
		fmt.Println(key)
	}
}
type shirts struct {
	product
	size,pattern,sleve,fabric string
}
func (shrt shirts) attributes(){
	fmt.Println("\n specification:")
	fmt.Println(shrt.size)
	fmt.Println(shrt.pattern)
	fmt.Println(shrt.sleve)
	fmt.Println(shrt.fabric)
}
func (prd product) inventory(){
	fmt.Println(prd.stock)
}
type catalog struct {
	name string
	details []information
}
func( c catalog) catalogdetails(){
	fmt.Println("________________________\n")
	fmt.Printf("\n store :%s\n",c.name)
	for _,key :=range c.details{
		fmt.Println("___******products******___")
		key.general()
		key.attributes()
		key.inventory()
		fmt.Println("_______________________")
	}
}
func main(){
	mobileprd := mobile {
		product{
			 "Apple iPhone 6s (Rose Gold, 32 GB)  (2 GB RAM)",
            "Handset, Apple EarPods with Remote and Mic, Lightning to USB Cable",
            40999,143,10,
		},
		[]string{" 2 GB RAM","4.7 inch Retina HD Display","12MP Primary Camera "},
		[]string{"32 GB ROM","4.7 inch Retina HD Display","5MP Front "},
	}
	shirtPrd :=shirts{
		product{
			"Reebok Striped Men's Round Neck Grey T-Shirt",
			"Machine Wash as per Tag, Do not Dry Clean, Do not Bleach",
			1124,400,25,
		},
		"XL","Striped","half sleeves","cotton",
	}
	category :=catalog {
		"urban",
		[]information{mobileprd,shirtPrd},
	}
	category.catalogdetails()
}
