package main 
import (
	"log"
	"os"
	"fmt"
)
func readCurrentDir() {
	file,err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory %s",err)
	}
	defer file.Close()
	fileList,_ := file.Readdir(0)
	fmt.Printf("\n Name\t \tSize\tIsDirectory Last Modification\n")
	for _,files := range fileList {
		fmt.Printf("\n%-15s %-7s %-12v %v",files.Name(),files.Size(),files.IsDir(),files.ModTime())

	}
}
func main() {
	readCurrentDir()
}