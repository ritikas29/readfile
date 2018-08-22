package main 
import (
	"log"
	"encoding/csv"
	"fmt"
	"os"
	//"net/http"
)
func main() {
	filepath := "/home/ritika/qci/go-practice/ord.csv"
	openfile,err := os.Open(filepath)
	checkerror("error in opening file ", err)
	filedata,err := csv.NewReader(openfile).ReadAll()
	checkerror("error in reading file ", err)
	for i,value := range filedata {
		fmt.Printf("i = %d , avalue %s",i,value)
	}
}
func checkerror(msg string, err error) {
	if err != nil {
		log.Fatal(msg,err)
	}
}
//func log(h http.Handler ) http.Handler {
	//return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
	//	log.Prinln("Before")
	//	defer log.Println("After")
	//	h.ServeHTTP(w,r)
	//})
//}