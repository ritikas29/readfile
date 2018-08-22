package main
import "fmt"
func main(){
	var empl = map[int]string{
		101: "john",
		102: "sharu",
		103: "mario",

	}
	printempl(empl,101)
	printempl(empl,104)
	if isemplexists(empl,102){
		fmt.Println("employeeid 102 found")

	}
}
func printempl(empl map[int]string,employeeid int) {
	if name,ok := empl[employeeid]; ok {
		fmt.Printf("name = %s,ok =%v\n",name,ok)
	
	 } else {
		 fmt.Printf("employee %d not found\n",employeeid)
	 }
}
func isemplexists(empl map[int]string,employeeid int)bool{
	_,ok := empl[employeeid]
	return ok
}