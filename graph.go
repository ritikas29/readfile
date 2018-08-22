package main 
import (
	"fmt"
	"math"
)
type graph struct {
	to int 
	wt float64
}
func floywarshall(g [][]graph) [][]float64 {
	dist := make([][]float64,len(g))
	for i := range dist {
		di := make([]float64,len(g))
		for j:= range di {
			di[j]=math.Inf(1)
		}
		di[i]=0
		dist[i]=di 
	}
	
}