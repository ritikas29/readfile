package main
import ("net/http";"io")
func hello(res http.ResponseWriter,req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"Text/html",
	)
	io.WriteString(
		res ,
		<DOCTYPE html>
		<html>
		<head>
		<title>hello world</title>
		</head>
		<body>
		hello everyone!!!!!
		</body>
		</html>,
	)
}
func main(){
	http.HandleFunc("/hello",hello)
	http.ListenAndServe(":9000",nil)
}