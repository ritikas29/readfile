package main  
import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"~/home/ritika/qci/go-practice/main1"

)
type Person struct {
	id string `json:"id,omitempty"`
	firstname string `json:"firstname , omitempty"`
	lastname string `json:"lastname,omitempty"`
	address *address `json:"address,omitempty"`
}
type address struct {
	city string `json:"city,omitempty"`
	state string `json:"state,omitempty"`
}
var people []Person
func GetPeople(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	for _,item := range people {
		if item.id == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func CreatePerson(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.id = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	for index,item := range people{
		if item.id == params["id"] {
			people = append(people[:index],people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {

	router := mux.NewRouter()
	people = append(people, Person{id: "1", firstname: "John", lastname: "Doe", address: &address{city: "City X", state: "State X"}})
   people = append(people, Person{id: "2", firstname: "Koko", lastname: "Doe", address: &address{city: "City Z", state: "State Y"}})
    people = append(people, Person{id: "3", firstname: "Francis", lastname: "Sunday"})

	router.HandleFunc("/people",GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}",GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}",CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}",DeletePerson).Methods("DELETE")
	fmt.Println("staringg gserver");
	log.Fatal(http.ListenAndServe(":1234",router))
}
