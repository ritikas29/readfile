package main
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
type Article struct{
	title string `json:"title"`
	desc string `json:"desc"`
	content string `json:"content"`
}
type Article []Article
func allArticle(w http.ResponseWriter, r *http.Request){
	articles := Article {
		Article {Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("endpoint with all articles endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homepage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"homepage endpoint hit")
}
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/",homepage)
	myRouter.HandleFunc("/articles",allArticle)
	log.Fatal(http.ListenAndServe(":3000",myRouter))
}
func main(){
	handleRequest()
}