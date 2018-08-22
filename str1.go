package main
import (
	"fmt"
	"strings"
)
func main() {
	var n int
	var ip string
	fmt.Println("how many strings u want to test?")
	fmt.Scanf("%d\n",&n)
	strSlice := make([]string,0)
	for i:= 0;i<n;i++{
		fmt.Printf("enter %d string\n",i+1)
		fmt.Scanf("%s\n",&ip)
		ip := strings.ToLower(ip)
		strSlice = append(strSlice ,ip)
	}
	for _,v:= range strSlice {
		fmt.Println(isP(v))
	}

}
func isP( s string) string {
	mid := len(s)/2
	last := len(s)-1
	for i :=0;i<mid;i++ {
			if s[i]!= s[last -i] {
				return "no \n "+ s + "is not palindrome"
			}
		}
	return "yes\n "+ s + "is a palindrome"
}