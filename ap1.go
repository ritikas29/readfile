package main
import (
	"io/ioutil"
	"log"
	"fmt"
	"os"
)
func appendfile() {
	file, err :=os.OpenFile("test.txt",os.O_WRONGLY|os.O_APPEND,0644)
	if err !=nil{
		log.Fatalf("failed opening file :%s",err)
	}
	defer file.Close()
	len,err := file.WriteString("world is beautiful")
	if err != nil {
		log.Fatalf("failed  writing to file :%s",err)
	}
	fmt.Printf("\n lenght:%d bytes",len)
	fmt.Printf("\n file name :%s",file.Name())
}
func ReadFile() {
	data , err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Panicf("failed reading data from file :%s",err)
	}
	fmt.Printf("\n lenght: %d bytes",len(data))
	fmt.Printf("\n data :%s",data)
	fmt.Printf("\n error :%v",err)
}
func main() {
	fmt.Printf("########append file####")
	appendfile()
	fmt.Printf("\n \n ###### read file####")
	ReadFile()
}