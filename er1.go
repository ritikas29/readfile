package main
import (
	"errors"
	"fmt"
)
type errorString errorString
func(e errorString) Error() string {
	return string(e)
}
func New(text string) error {
	return errorString(text)
}
var ErrNamedType =New("EOF")
var ErrStructType = errors.New("EOf")
func main() {
	if ErrNamedType == New("EOF") {
		fmt.Println("named type error")
	}
	if ErrStructType == errors.New("EOF"){
		fmt.Println("struct type error")
	}
}