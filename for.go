package main
import (
	"fmt"
	"log"
	"net/http"
)
type greeting string
func (g greeting) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,g)
}
func main() {
	err := http.ListeAndServe("localhost:4000",greeting("hello,go "))
	if err !=nil {
		log.Fatal(err)
	}
}