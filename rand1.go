package main
import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)
type Maze struct {
	c,h,v []byte
	cell,hor,ver [][]byte
}
func drawmaze( rows,cols int) *Maze {
	c := make([]byte , rows*cols)
	h := bytes.Repeat([]byte{'-'},rows*cols)
	v := bytes.Repeat([]byte{'|'},rows*cols)
	cell := make([][]byte,rows)
	hor := make([][]byte,rows)
	ver := make([][]byte,rows)
	for i := range hor {
		cell[i]= c[i*cols : (i+1)*cols]
		hor[i] = h[i*cols :(i+1)*cols]
		ver[i] = v[i*cols :(i+1)*cols]
	}
	return &Maze{c,h,v,cell,hor,ver}
}
func (m *Maze) string() string {
	hwall := []byte("+----")
	hopen := []byte("+    ")
	vwall := []byte("|    ")
	vopen := []byte("     ")
	rightcorner := []byte("+\n")
	rightwall := []byte("|\n")
	var b []byte
	for r,hw := range m.hor {
		for _,h := range hw {
			if h == '-' || r == 0 {
				b = append(b,hwall...)
			} else {
				b = append(b,hopen...)
			}
		}
	}
	b = append(b, rightcorner...)
	for c,vw := range m.ver[r] {
		if vw == '|' || c == 0 {
			b = append(b,vwall...)
		} else {
			b = append(b,vopen...)
		}
		if m.cell[r][c] !=0 {
			b[len(b)-2] = m.cell[r][c]
		}
	}
	b = append(b,rightwall...)
	for _ = range m.hor[0] {
		b = append(b,hwall...)
	}
	b = append(b,rightcorner...)
	return string(b)
}
func(m *Maze) generator() {
	m.recursion(rand.Intn(len(m.cell)),rand.Intn(len(m.cell[0])))
	}
const (
	up = iota
	down
	right
	left
)	
func(m *Maze) recursion(row,col int) {
	rand.Seed(time.Now().UnixNano())
	m.cell[row][col] =' '
	for _,wall := range rand.Perm(4) {
		switch wall {
		  case up :
			if row >0 && m.cell[row-1][col]== 0{
				m.hor[row][col]=0
				m.recursion(row-1,col)
			}
			case down :
			if row < len(m.cell)-1 && m.cell[row+1][col]==0{
				m.hor[row +1][col]=0
				m.recursion(row+1,col)
			}
			case left :
			if col >0 && m.cell[row][col-1]==0{
				m.ver[row][col]=0
				m.recursion(row,col-1)
			}
		case right :
			if col <len(m.cell[0])-1 && m.cell[row][col+1]== 0{
				m.ver[row][col+1]=0
				m.recursion(row,col+1)
			}	
		}
	}
}
func main() {
	d:= drawmaze(5,7)
	d.generator()
	fmt.Print(d)
}