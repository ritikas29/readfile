package main
import "fmt"
func main() {
	var action int
	fmt.Println("enter 1 for student and 2 for professional")
	fmt.Println(&action)
	switch action {
	case 1:
		fmt.Printf(" I am a student")
	case 2:
		fmt.Printf("I am a proffesional")
		default :
		panic(fmt.Sprintf(" I am a %d",action))
	}
	fmt.Println("")
	fmt.Println("enter 1 for us and 2 for uk")
	fmt.Scanln(&action)
	switch {
	case 1:
		fmt.Printf("US")
	case 2:
		fmt.Printf("uk")
	default:
		panic(fmt.Sprintf("i am a %d",action))
	}
}