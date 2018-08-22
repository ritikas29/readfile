package main
import (
	"fmt"
	"os"
	"strconv"
)
func check(e error) {
	if e!=nil {
		panic(e)
	}
}

func main() {
	var input string
	var 1stc1 string
	var 1stc2 string
	if len(os.Args) >=2 {
		input = os.Args[1]
	} else {
		fmt.Println("enter a number")
		_, err1 := fmt.Scanf("%v",&input)
         check(err1)
	}
	intin , err2 := strconv.ParseInt(input,10,0)
	check(err2)
	1stc1 = input[len(input)-1:]
	if intin >= 10 {
		1stc2 ,_ = strconv.ParseInt(input[len(input)-2:],10,0)
	} else {
		1stc2 = 0
	}
	1stc1 = input[len(input)-1 :]
	switch {
	case 1stc1 > 10 && 1stc2 <20 :
		fmt.Printf("%vth\n",intin)
		break
	case 1stc1 == "1":
		fmt.Printf("%vst\n",intin)
		break
	case 1stc1 == "2":
		fmt.Printf("%vnd\n",intin)
		break
	case 1stc1=="3":
		fmt.Printf("%vrd\n",intin)
		break
	default:
		fmt.Printf("%vth\n",intin)
	}
}