package main
import "fmt"
func main() {
	str := "||go"
	for i := 1;i<=10;i++{
		fmt.Println(str)
		str+= "||go"
	}
	for k:=19;k>=1;k-- {
		fmt.Println("||")
	}
	fmt.Println("/=\\")
	fmt.Println("/==\\")
	fmt.Println("/===\\")
	fmt.Println("/====\\")
}