package main
import (
	"net/http"
	"log"
	"encoding/json"
)
func main() {
	makerequest()
}
func makerequest() {
	client :=http.Client{}
	request , err := http.NewRequest("GET","https://httpbin.org/get",nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp,err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}