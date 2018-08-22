package main 
import "fmt"
type example struct {
	X,Y string
}
var (
	a=example{"AUS","CAN"}
	b=&example{"JAP","KoR"}
	c= example{X:"US",Y:"Uk"}
	d= example{}
)
func main() {
	e :=b     
	b.X = "Rus"
	f := *b   
	fmt.Println("a:\t",a)
	fmt.Println("b:\t",b)
	fmt.Println("c:\t",c)
	fmt.Println("d:\t",d)
	fmt.Println("e:\t",e)
	fmt.Println("e:\t",f)	
}