package main 
import (
	"log"
	"os"
	"test/template"
)
type Person struct {
	Name string
	Emails []string
}
const tmpl = The name is {{.Name}}
{{range .Emails}}
 Her Email ID is {{.}}
 {{end}}
func main() {
	person := Person {
		Name: "RIA",
		Emails: []string{"abc@gmail.com","cde@gmail.com"}
	}
	t := template.New("Person template")
	t, err := t.Parse(tmpl)
	if err := t.Parse(tmpl){
		log.Fatal("Parse:",err)
		return
	}
	err = t.Execute(os.Stdout,person)
	if err != nil {
		log.Fatal("Execute:",err)
		return
	}
}
