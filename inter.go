package main
import "fmt"
func interpl(arr  []int,key int )int{
	min,max :=arr[0],arr[len(arr)-1]
	low ,high :=0,len(arr)-1
 for {
	 if key<min {
		 return low
	 }
	 if key >max {
		 return high +1

	 }
	 var guess int
	 if high == low{
		 guess = high
	 }else {
		 size := high -low 
		 offset := int(float64(size-1)*(float64(key-min))/float64(max-min))
		guess = low +offset
		}
		if arr[guess]==key{
			for guess >0 && arr[guess-1]==key{
				guess--
			}
			return guess
		}
		if arr[guess]>key {
			high = guess -1
			max = arr[high]
		}else {
			low = guess +1
			min = arr[low]
		}
 }
}
func main() {
	items := []int{1,2,7,9,13,19,7}
	fmt.Println(interpl(items,9))
}