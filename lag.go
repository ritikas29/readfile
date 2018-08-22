package main
import "fmt"
func main(){
	var num[50] float64
	var temp int
	fmt.Print("enter no. of elements:")
	fmt.Scanln(&temp)
	for i :=0;i<temp ;i++ {
		fmt.Print("enter the number:")
		fmt.Scan(&num[i])

	}
for j:=1;j<temp;j++{
	if(num[0]<num[j]) {
		num[0]=num[j]
	}	
}
fmt.Print("the largest number is:",num[0])
}