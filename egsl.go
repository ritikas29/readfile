package main
import "fmt"
func main() {
	src :=[]string{"sublime","vscode","intellij","ecllipse"}
	dest := make([]string, 2)
	numElementscopied := copy(dest,src)
	fmt.Println("src=",src)
	fmt.Println("dest=",dest)
	fmt.Println("number of elements copied ",numElementscopied)
}	
