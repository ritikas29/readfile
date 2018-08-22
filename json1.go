package main 
import (
	"os"
	"fmt"
	"encoding/json"

)
type Config struct {
    database struct {
		host string `json:"host"`
		port string `json:"port"`
    } `json:"database`
		host string `json:"host"`
		port string `json:"port"`
}
func loadconfiguration(filename string ) (Config,error) {
	var config Config
	configFile,err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config,err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config,err
}
func main() {
	fmt.Println("starting the application...")
	config,_ :=loadconfiguration("config.json")
	fmt.Println(config)
}

