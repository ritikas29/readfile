package main
import "fmt"
func main() {
	color := make (map[string]string)
	color["R"]="Red"
	color["B"]="blue"
	color["G"]="green"
	fmt.Println("select a primary color :R/B/G")
	var colorCode string
	fmt.Scanf("%s",&colorCode)
	if value , isOK := color[colorCode]; isOK{
		fmt.Println("you love" +value) } else {
			fmt.Println("invalid entry or key not found")
		}
	}
