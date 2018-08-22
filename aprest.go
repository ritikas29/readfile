package main
import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
	
)
func read(w http.ResponseWriter,r *http.Request){
	xlsx,err := excelize.OpenFile("./ord.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	cell := xlsx.GetCellValue("Sheet1","B2")
	fmt.Println(cell)
	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
		fmt.Print(colCell,"\t")
		}
		fmt.Println()
	}
}
func write(w http.ResponseWriter,r *http.Request){
	
		xlsx := excelize.NewFile()
		// Create a new sheet.
		index := xlsx.NewSheet("Sheet2")
		// Set value of a cell.
		xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
		xlsx.SetCellValue("Sheet1", "B2", 100)
		// Set active sheet of the workbook.
		xlsx.SetActiveSheet(index)
		// Save xlsx file by the given path.
		err := xlsx.SaveAs("./ord.xlsx")
		if err != nil {
			fmt.Println(err)
		}
	
}
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/",write).Methods("PUT")
	r.HandleFunc("/",read).Methods("GET")
    if err := http.ListenAndServe("localhost:3000",r); err !=nil {
		log.Fatal(err)
	}
}