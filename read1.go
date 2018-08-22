package main 
import (
	"net/http"
    "fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
	"log"
	"gopkg.in/mgo.v2"
)
type do struct {
	Server string
	Database string
}
var db *mgo.Database
const (
	COLLECTION = "projectx"
)

func read(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "result")
	xlsx,err := excelize.OpenFile("./price.xlsx")
    if err != nil {
		fmt.Println(err)
		return
	}
	cell := xlsx.GetCellValue("Sheet1","C3")
	fmt.Println(cell)
	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
		fmt.Print(colCell,"\t")
		}
		fmt.Println()
	}
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",read).Methods("GET")
	log.Fatal(http.ListenAndServe(":4567",myRouter))
}
func ( m *do ) Connect() {
	
	session ,err := mgo.Dial(" http://localhost:4567/")
	//m.Server
	if err != nil {
		log.Fatal(err)
	}
   db = session.DB(m.Database)
}
func main() {
	fmt.Println("server")
	//db ,err := mgo.Dial("http://localhost:4567/")
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//db.SetMode(mgo.Monotonic,true)
	//c := db.DB(database).C(Collection)

	//if err != nil {
	//	log.Fatal(err)
	//}
	//return 

	handleRequests()
}

	