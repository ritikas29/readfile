package main  
import (
	"fmt"
	"sort"
)
type person struct {
	name string
	age int 
}
type byname []person
func(this byname) len() int {
	return len(this)
}
func(this byname) less(i,j int) bool {
	return this[i].name < this[j].name 
}
func(this byname) swap(i,j int) {
	this[i],this[j] = this[j],this[i]
}
func main() {
	kids := []person {
		{"jill",9},
		{"jack",10},
	}
	sort.Sort(byname(kids))
	fmt.Println(kids)
}