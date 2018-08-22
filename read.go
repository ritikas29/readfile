package main 
import (
	"fmt"
	"io"
	"os"
)
var path = "test.txt"
func main() {
	createFile()
	writeFile()
	readFile()
	deleteFile()
}
func createFile() {
	var _ , err = os.Stat(path)
	if os.IsNotExist(err) {
		var file,err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("file created successfully")
}
func writeFile() {
	var file,err = os.OpenFile(path,os.O_RDWR,0644)
	if isError(err) {
		return
	}
	defer file.Close()
	_, err = file.WriteString("hello\n")
	if isError(err) {
		return 
	}
	_,err = file.WriteString("World\n")
	if isError(err) {
		return
	}
	fmt.Println("file updated successfully")
}
func readFile() {
	var file,err = os.OpenFile(path,os.O_RDWR,0644)
	if isError(err){
		return 
	}
	defer file.Close()
	var text = make([]byte,1024)
	for {
		_, err = file.Read(text)
		if err == io.EOF {
			break
		}
	}
	fmt.Println("reading from file")
	fmt.Println(string(text))
}
func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return 
	}
	fmt.Println("file deleted")
}
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err !=nil)
}








