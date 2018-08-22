package main    
import (
	"html/template"
	"net/http"
)
var testTemplate *template.Template 
type Widget struct{
	name  string
	price int 
}
type ViewData struct {
	name string
	Widgets []Widget
}
func main() {
	var err error
	testTemplate,err = template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000",nil)
}
func handler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","text/html")
	err := testTemplate.Execute(w , ViewData{
		name :"riya" ,
		Widgets: []Widget{
			Widget{"Blue Widget",12},
			Widget{"Red Widget",12},
			Widget{"Green Widget",12},
		},
	})
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}