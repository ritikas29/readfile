package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	response , err :=http.Get("https://reqres.in/api/users/2")
	if err != nil {
      fmt.Printf("the http request failed with error %s\n",err)
	} else {
		data,_:=ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	jsonData := map[string]string{"firstname":"ria","lastname":"sharma"}
	jsonValue,_ := json.Marshal(jsonData)
	request,_:=http.NewRequest
	//response,err = http.Post("https://reqres.in/api/users/","application/json",bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("the http request failed with error %s\n",err)
	  } else {
		  data,_:=ioutil.ReadAll(response.Body)
		  fmt.Println(string(data))
	  }
}