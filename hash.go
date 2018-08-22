package main
import "fmt"
func main(){
	var count map[int]string
	count = make(map[int]string)
	count[1]="INDIA"
	count[2]="CHINA"
	count[3]="AUSTRIA"
	count[4]="GERM"
	count[5]="MALYISYA"
	count[6]="CANANDA"
	for i,j := range count {
		fmt.Printf("key:%d value:%s\n",i,j)
	}
}