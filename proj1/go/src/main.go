package main 
import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"proj1/ablum"
)
func main() {
	router := ablum.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET","POST","DELETE","PUT"})
	log.Fatal(http.ListenAndServe(":9000",handlers.CORS(allowedOrigins,allowedMethods)(router)))
}