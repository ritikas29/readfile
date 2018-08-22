package main
import (
	"fmt"
	"sort"
)
func main() {
	mobile :=[]struct {
		brand string
		price int
	}{
		{"nokia",700},
		{"samsung",505},
		{"apple",924},
		{"sony",655},
	}
	sort.Slice(mobile, func(i, j int) bool { return mobile[i].brand < mobile[j].brand })
	fmt.Println("\n\n sorted by brand[asc]")
	for _,v :=range mobile {
		fmt.Println(v.brand,v.price)
	}
	sort.Slice(mobile ,func(i,j int) bool { return mobile[i].brand > mobile[j].brand})
	fmt.Println("\n \n sorted by brand[des]")
	for _,v := range mobile {
		fmt.Println(v.brand,v.price)
	}
	sort.Slice(mobile  ,func(i,j int) bool { return mobile[i].price >mobile[i].price})
	fmt.Println("\n \n sorted by price")
	for _,v := range mobile {
		fmt.Println(v.brand,v.price)
	}
	mobile=[] struct {
		brand string
		price int
	}{
		{"mi",900},
		{"oppo",305},
		{"iphone",924},
		{"sony",655},
	}
	sort.Slice(mobile,func(i,j int)bool { return mobile[i].brand < mobile[j].brand})
	fmt.Println("\n \n sorted by brand")
	for _,v := range mobile {
		fmt.Println(v.brand,v.price)
	}
}