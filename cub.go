package main
import "fmt"
func cuboiddraw(drawx,drawy,drawz int) {
	fmt.Printf("cuboid %d %d\n",drawx,drawy,drawz)
	cubeline(drawy+1,drawx,0,"+-")
	for i :=1;i<=drawy ;i++{
		cubeline(drawy-i+1,drawx,i-1,"/|")
	}
	cubeline(0,drawx,drawy,"+-|")
	for i :=4 *drawz - drawy -2 ;i>0;i--{
		cubeline(0,drawx,drawy,"||")
	}
	cubeline(0,drawx,drawy,"|+")
	for i :=1 ;i<= drawy ;i++{
		cubeline(0,drawx,drawy-i,"|/")
	}
	cubeline(0,drawx,0,"+-\n")
}
func cubeline(n,drawx,drawy int,cubedraw string){
	fmt.Printf("%*s",n+1,cubedraw[:1])
	for d := 9*drawx-1 ;d>0;d--{
		fmt.Print(cubedraw[1:2])

	}
	fmt.Print(cubedraw[:1])
	fmt.Printf("%*s\n",drawy +1,cubedraw[2:])
}
func main() {
	fmt.Println("enter 3 dimensions of cubiod")
	var l,b,h int
	fmt.Scanln(&l)
	fmt.Scanln(&b)
	fmt.Scanln(&h)
	cuboiddraw(l,b,h )
}