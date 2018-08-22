package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
func alluser(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w ,"all users")
}
func newuser(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w ,"new user endpoint hit")
}
func deleteuser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"delete user endpoint hit")
}
func updateuser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"update user endpoint hit")
}
func handleRequest(){
	r := mux.NewRouter().StrictSlash(true)
	r.Handlefunc("/users",alluser).Methods("GET")
	r.Handlefunc("/user/{name}/{email}",newuser).Methods("POST")
	r.Handlefunc("/user/{name}",deleteuser).Methods("DELETE")
	r.Handlefunc("/user/{name}/{email}",updateuser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000",myRouter))
}