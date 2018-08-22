package main 
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
type movie struct {
	id bson.ObjectId `bson:"_id" json:"id"`
	name string `bson:"name" json:"name"`
	coverimage string `bson:"cover_image" json:"cover_image"`
	description string `bson:"description" json:"description"`
}
func allmoviesendpoint(w http.ResponseWriter, r *http.Request){
	
}
func findmovieendpoint(w http.ResponseWriter,r *http.Request) {

}
func createmovieendpoint(w http.ResponseWriter,r *http.Request){

}
func updatemovieendpoint(w http.ResponseWriter, r *http.Request){

}
func deletemovieendpoint(w http.ResponseWriter, r *http.Request){

}
func main(){
	r := mux.Router()
	r.HandleFunc("/movies",allmovieendpoint).Methods("GET")
	r.HandleFunc("/movies",findmovieendpoint).Methods("GET")
	r.HandleFunc("/movies",createmovieendpoint).Methods("POST")
	r.HandleFunc("/movies",updatemovieendpoint).Methods("PUT")
	r.HandleFunc("/movies",deletemovieendpoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000",r); err != nil {
		log.Fatal(err)
	}
}