package main
import (
	"fmt"
	"strings"
)
func main() {
	fmt.Printf("%q\n",strings.Split("a,b,c",","))
	fmt.Printf("%q\n",strings.Split("a man a plan a canal panama","a"))
	fmt.Printf("%q\n",strings.Split("xyz",""))
	fmt.Printf("%q\n",strings.Split("","bernardo o'higgins"))
} The declaration in a type switch has the same syntax as a type assertion 