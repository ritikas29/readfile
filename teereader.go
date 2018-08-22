package main
import (
	"bytes"
	"io"
	"fmt"
	"strings"
)
func main() {
	testString :=strings.NewReader("jobs,code,videos and news for go hackers.")
	var bufferRead bytes.Buffer 
	example := io.TeeReader(testString,&bufferRead)
	readerMap := make([]byte,testString.Len())
	lenght,err := example.Read(readerMap)
	fmt.Printf("\nBufferRead: %s",&bufferRead)
	fmt.Printf("\n Read : %%s",readerMap)
	fmt.Printf("\n lenght: %d,Error :%v",lenght,err)

}