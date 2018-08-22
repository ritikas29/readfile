package main
import "fmt"
func doublefail(a int ,arr [2]int, s string) {
	a = a*2
	for i :=0;i<len[arr];i++ {
		arr[i]=arr[i]*2
	}
	s= s+s
	fmt.Println("in doublefail:",a,arr,s)
}
func main() {
	a := 1
	arr :=[2] int{2,4}
	s := "hello"
	doublefail(a,arr,s)
	fmt.Println("in main:",a,arr,s)
}