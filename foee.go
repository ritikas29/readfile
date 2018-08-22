package main
import "fmt"
func main() {
	arrayone :=[3]string{"apple","mango","strawberry"}
	fmt.Println("\n-------eg1------")
	for index,element := range arrayone{
		fmt.Println(index)
		fmt.Println(element)
	}
	strdict := map[string] string {"japan":"tokyo","india": "new delhi","canada":"ottawa"}
		fmt.Println("\n------ eg2-----\n")
		for index,element := range strdict{
			fmt.Println("index:",index,"element:",element)

		}
		fmt.Println("\n------eg3----\n")
		for key :=range strdict{
			fmt.Println(key)
		}
		fmt.Println("\n-------eg4----\n")
		for _ ,value := range strdict{
			fmt.Println(value)
		}
		j:=0
		fmt.Println("\n------eg5----\n")
		for range strdict{
			fmt.Println(j)
			j++
		}
	}