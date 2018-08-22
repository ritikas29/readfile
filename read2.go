package main 
import (
	"net/http"
	//"github.com/gorilla/websocket"
)	
func main() {
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request ){
	http.ServeFile(w,r ,"main.go")
     })
 http.ListenAndServe(":1234",nil)
}