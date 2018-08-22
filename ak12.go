package main
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func main() {
	rows := readOrd("ord.csv")
	rows = calculate(rows)
	writeOrd("ord.csv",rows)
}
func readOrd(name string) [][]string {
	f,err := os.Open(name)
	if err != nil {
		log.Fatalf("cannot open '%s':'%s'\n",name,err.Error())
	}
	defer f.Close()
	r :=csv.NewReader(f)
	r.Comma = ';'
	rows,err := r.ReadAll()
	if err != nil {
		log.Fatalf("cannot read CSV data:",err.Error)
	}
	return rows
}
func calculate(rows [][]string) [][]string {
	sum := 0
	nb := 0
	for i := range rows {
		if i == 0 {
			rows[0]= append(rows[0],"Total")
		}
		item := rows[i][2]
		price,err := strconv.Atoi(strings.Replace(rows[i][3],"."," ",-1))
		if err != nil {
			log.Fatalf("cannot retrieve price of ")
		}
		qty,err := strconv.Atoi(rows[i][4])
		if err != nil {
			log.Fatalf("cannot retrieve quantity of %s: %s\n",item,err)

		}
		total := price * qty
		rows[i] = append(rows[i],intToFloatString(total))
		sum += total
		if item == "ball pen" {
			nb += qty
		}
	}
	rows = append(rows , []string{"","","","sum","",intToFloatString(sum)})
	rows = append(rows,[]string{"","","","ball pen",fmt.Sprint(nb),""})
	return rows
}
func intToFloatString(n int) string {
	intgr := n/100
	frac := n - intgr*100
	return fmt.Sprintf("%d.%d",intgr,frac)
}
func writeOrd(name string , rows [][]string) {
	f,err := os.Create(name)
	if err != nil {
		log.Fatalf("cannot open '%s': %s\n", name , err.Error())
	}
	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatalf("cannot close '%s': %s\n",name,err.Error)
		}
	}()
	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
	
}