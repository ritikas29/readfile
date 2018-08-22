package main
import (
	"fmt"
	"strings"
)
func main() {
	var is string
	fmt.Println("enter string:")
	fmt.Scanf("%s\n",&is)
	is = strings.ToLower(is)
	fmt.Println(ip(is))
}
func ip( s string) string {
	mid :=len(s)/2
	last := len(s)-1
	for i :=0 ;i<mid ;i++{
		if s[i] != s[last -i]{
			return "no it's not a palimdrome"
		}
	}
	return "yes you've entered a palindrome"
}