package main
import (
	"log"
	"os"
)
type job struct {
	command string
	*log.Logger

}
func newjob(command string)*job {
	return &job{command,log.New(os.Stderr,"job",log.Ldate)}
}
func main(){
	newjob("demo").Print("starting now...")
}