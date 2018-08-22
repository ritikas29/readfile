package main 
import(
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"io"
)
type Person struct {
	firstname string `json:"firstname"`
	lastname string `json:"lastname"`
	Address *Address`json:"address,omitempty"`
}
type Address struct {
	city string `json:"city"`
	state string `json:"state"`
}
func main() {
	csvFile,_ := os.Open("people.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []Person
	for {
		line,error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		people = append(people, Person{
			firstname: line[0],
			lastname: line[1],
			Address: &Address{
				city: line[2],
				state: line[3],
			},
		})

	}
	peopleJson, _ := json.Marshal(people)
	fmt.Println(string(peopleJson))
}