package main
import (
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
type Book struct {
	id string `json:"id"`
	isbn string `json:"isbn"`
	title  string `json:"id"`
	Author *Author `json:"author"`
}
type Author struct {
	firstname string `json:"firstname"`
	lastname string `json:"lastname"`
}
var books []Book
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","applicationb/json")
	json.NewEncoder(w).Encode(books)
  
}
func getBook(w http.ResponseWriter, r *http.Request) {
w.Header().Set("content-type","application/json")
params := mux.Vars(r)
for _, item := range books {
	if item.id == params["id"] {
		json.NewEncoder(w).Encode(item)
		return
	}
}
  json.NewEncoder(w).Encode(&Book{})

}
func createBook( w http.ResponseWriter, r *http.Request) {
  w.Header().Set("content-type","application/json")
  var book Book
  _ = json.NewDecoder(r.body).Decode(&book)
  book.id = strconv.Itoa(rand.Intn(100000))
  books = append(books,book)
  json.NewEncoder(w).Encode(Book)
}
func updateBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("content-type","application/json")
  params := mux.Vars(r)
  for index,item := range books {
	  if item.id == params["id"]
	 books = append(book[:index],books[idex+1:]...)
	 break
  }
  json.NewEncoder(w).Encode(books)

}
func  deleteBooks(w http.ResponseWriter, r *http.Request) {


}
func main() {
   r := mux.NewRouter()
	books = append(books,Book{id:"1",isbn:"44873",title :"book one",Author:&Author
	{firstname :"jia", lastname:"sharma"}})
	books = append(books,Book{id:"2",isbn:"44893",title :"book two",Author:&Author
		{firstname :"reema", lastname:"kumari"}})
   r.HandleFunc("/api/books",getBooks).Methods("GET")
   r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
   r.HandleFunc("/api/books",createBook).Methods("POST")
   r.HandleFunc("/api/books/{id}",getBooks).Methods("PUT")  
   r.HandleFunc("/api/books/{id}",getBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000",r))
}