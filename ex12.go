package main
 
import (
    "fmt"
    "sort"
)
 
func main() {
    a := []int{15, 4, 33, 52, 551, 90, 8, 16, 15, 105}    // unsorted
    sort.Sort(sort.Reverse(sort.IntSlice(a)))
    fmt.Println("\n",a)
     
    a = []int{-15, -4, -33, -52, -551, -90, -8, -16, -15, -105}     // unsorted
    sort.Sort(sort.Reverse(sort.IntSlice(a)))
    fmt.Println("\n",a)
     
     
    b := []string{"Montana","Alaska","Indiana","Nevada","Washington","Ohio","Texas"}   // unsorted
    sort.Sort(sort.Reverse(sort.StringSlice(b)))
    fmt.Println("\n",b)
     
    b = []string{"ALASKA","indiana","OHIO","Nevada","Washington","TEXAS","Montana"}  // unsorted
    sort.Sort(sort.Reverse(sort.StringSlice(b)))
    fmt.Println("\n",b)
     
    c := []float64{90.10, 80.10, 160.15, 40.15, 8.95} //    unsorted
    sort.Sort(sort.Reverse(sort.Float64Slice(c)))
    fmt.Println("\n",c)
     
    c = []float64{-90.10, -80.10, -160.15, -40.15, -8.95} // unsorted
    sort.Sort(sort.Reverse(sort.Float64Slice(c)))
    fmt.Println("\n",c)
}