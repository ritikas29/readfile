package main
import "fmt"
func main() {
	daysofweek :=[7]string{"mon","tues","wed","thu","fri","sat","sun"}
	for in , value := range daysofweek {
		fmt.Printf("day %d of week=%s\n",in,value)
	}
}