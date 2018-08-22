func main 
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"string"
)
func main() {
	resp,err := http.Get("http://www.baidu.com")
	if err !=nil{
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err :=ioutl.ReadAll(resp.Body)
	if err !=nil {
		fmt.Println("http read error")
		return
	}
	src := string(body)
	re,_ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src,strings.ToLower)
	
}