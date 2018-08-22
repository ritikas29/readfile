package main
import "fmt"
func compressLZW(teststr string)[]int {
	code := 256
	dictionary := make(map[string]int)
	for i :=0 ;i<256;i++{
		dictionary[string(i)] = i 
	}
	currChar := " "
	result := make([]int, 0)
	for _ ,c := range []byte(teststr){
		phrase := currChar + string(c)
		if _, isTrue := dictionary[phrase]; isTrue {
			currChar = phrase
		} else {
			result = append(result,dictionary[currChar])
			dictionary[phrase]= code
			code++
			currChar = string(c)
		}
	}
	if currChar !=" " {
		result = append(result,dictionary[currChar])
	}
	return result
}
func decompressLZW(compressed []int) string{
	code := 256
	dictionary := make(map[int]string)
	for i:=0;i<256 ;i++{
		dictionary[i] = string(i)

	}
	currChar := string(compressed[0])
	result := currChar
	for _ , element := range compressed[1:]{
		var word string
		if x,ok := dictionary[element]; ok {
			word =x
		} else if element == code {
			word = currChar + currChar[:1]
		} else {
			panic(fmt.Sprintf("bad compressed element:%d",element))
		}
		result +=word
		dictionary[code] = currChar + word[:1]
		code ++
		currChar = word
	}
	return  result
}
func main() {
	fmt.Print("enter any string:")
	var teststr string
	fmt.Scanln(&teststr)
	compressed :=compressLZW(teststr)
	fmt.Println("\n after compression:",compressed)
	uncompression := decompressLZW(compressed)
	fmt.Println("\n after uncompression",uncompression)
}