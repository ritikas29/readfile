package main
import(
	"fmt"
	"math/rand"
	"time"
)
func main(){
	slice := generateslice(20)
	fmt.Println("\n....	Unsorted....\n\n",slice)
	selecsort(slice)
	fmt.Println("\n....Sorted...\n\n",slice,"\n")
}
func generateslice(size int) []int {
	slice := make([]int,size,size)
	rand.Seed(time.Now().UnixNano())
	for i :=0;i<size;i++{
		slice[i]=rand.Intn(999)-rand.Intn(999)

	}
	return slice
}
func selecsort(items []int){
	var n = len(items)
	for i:=0;i<n;i++{
		var minIdx =i 
		for j:=1;j<n;j++{
			if items[j]<items[minIdx] {
			minIdx=j
		}
			
		} 
		items[i],items[minIdx]=items[minIdx],items[i]
	}
}
