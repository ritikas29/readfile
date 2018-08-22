package main
import(
    "fmt"

)
func main() {
	var mat1[50][50] int
	var mat2[50][50] int
	var sum[50][50] int
	var r,c int
	fmt.Print("enter no. of rows:")
	fmt.Scanln(&r)
	fmt.Print("enter no. of column:")
	fmt.Scanln(&c)
	fmt.Println()
	fmt.Println("  matrix1  ")
	fmt.Println()
	for i:=0;i<r;i++ {
		for j:=0;j<c;j++{
			fmt.Printf("enter the element of matrix1 %d %d",i+1,j+1)
			fmt.Scanln(&mat1[i][j])
		}
	}
	fmt.Println()
	fmt.Println("   Matrix2")
	fmt.Println()
	for i := 0;i<r;i++{
		for j:=0;j<c;j++ {
			fmt.Printf("enter the elements of matrix %d %d",i+1,j+1)

		}
	}
	for i:=0;i<r;i++{
		for j:=0;j<c ;j++{
			sum[i][j]=mat1[i][j]+mat2[i][j]

		}
	}
	fmt.Println()
   fmt.Println("   sum of matrix ")
   fmt.Println()
   for i:=0;i<r;i++{
	   for j:=0;j<c;j++{
		   fmt.Printf("%d",sum[i][j])
		   if(j== c-1){
			   fmt.Println("")
		   }
	   }
   }
}   