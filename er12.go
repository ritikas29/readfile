package main 
import (
	"errors"
	"fmt"
	"time"
)
type CustomError struct {
	message string
	function string
	timestamp time.Time 
}
func(e *CustomError) Error() string{
	return e.message
}
func(e *CustomError) GetHost() time.Time {
	return e.timestamp
}
func main() {
	errs :=[]interface{} {
		CustomError{"custom error.","host01.local",time.Now()},
		errors.New("Generic error"),
	}
	for _, e := range errs {
        if v, ok := e.(CustomError); ok {
            fmt.Println("Error: ", v.Error())
            fmt.Println("  -> This error is of type CustomError")        
        } else {
            fmt.Println("Error: ", v.Error())
            fmt.Println("  -> This error is generic")
        }
    }
}