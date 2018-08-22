package main 
import (
	"encoding/json"
	"fmt"
)
type key interface {
	less(key) bool
	eq(key) bool
}
type Node struct {
	data key 
	balance int
	link [2]*Node 
}
func opp(dir int) int {
	return 1 - dir 
}
func single(root *Node, dir int) *Node {
    save := root.Link[opp(dir)]
    root.Link[opp(dir)] = save.Link[dir]
    save.Link[dir] = root
    return save
}
func double(root *Node,dir int) *Node {
	save := root.Link[opp(dir)].Link[dir]
	root.Link[opp(dir)].Link[dir]=save.Link[opp(dir)]
	save.Link[opp(dir)] = root.Link[opp(dir)]
	root.Link[opp(dir)]= save
	save = root.Link[opp(dir)]
	root.Link[opp(dir)]= save.Link[dir]
	save.Link[dir] = root
	return save
}
func adjustBalance(root *Node,dir,bal int) {
	n := root.Link[dir]
	nn := n.Link[opp(dir)]
	switch nn.Balance {
	case 0 :
		root.Balance = 0
		n.Balance = 0
	case bal :
		root.Balance = -bal
		n.Balance = 0
		default :
		root.Balance = 0
		n.Balance = bal
	}
	nn.Balance = 0
}
func insertBalance(root *Node,dir int) *Node {
	n := root.Link[dir]
	bal := 2 *dir -1
	if n.Balance == bal {
		root.Balance = 0
		n.Balance = 0
		return single(root,opp(dir))
	}
	adjustBalance(root,dir ,bal)
	return double(root, opp(dir))
}
func insetR(root *Node , data key) (*Node , bool) {
	if root == nil {
		return &Node{Data: data},false
	}
	dir := 0
	if root.Data.less(data) {
		dir = 1
	}
	var done bool 
	root.Link[dir],done = insertR(root.Link[dir],data)
	if done {
		return root,true  
	}
	root.Balance += 2*dir -1
	switch root.Balance {
	case 0:
		return root,true
	case 1,-1 :
		return root,false 
	}
	return insertBalance(root,dir),true  
}
func Insert(tree **Node,data key) {
	*tree,_ = insertR(*tree,data)
}
func Remove(tree **Node,data key) {
	*tree,_ = removeR(*tree,data )
}
func removeBalance(root *Node,dir int) (*Node,bool){
	n:= root.Link[opp(dir)]
	bal :=2*dir -1
	switch n.Balance {
	case -bal:
		root.Balance =0
		n.Balance =0
		return single(root,dir),false
	case bal :
		adjustBalance(root,opp(dir),-bal)
		return double(root,dir),false
	}
	root.Balance = -bal
	n.Balance = bal
	return single(root,dir),true  
}
func removeR(root *Node,data key)(*Node,bool){
	if root == nil {
		return nil,false
	}
	if root.Data.Eq(data){
		switch {
		case root.Link[0] == nil:
			return root.Link[1],false
		case root.Link[1] == nil :
			return root.Link[0],false 
		}
		heir := root.Link[0]
		for heir.Link[1] !=nil {
			heir = heir.Link[1]
		}
		root.Data = heir.Data 
		data = heir.Data 
	}
	dir := 0
	if root.Data.Less(data){
		dir = 1
	}
	var done bool
	root.Link[dir],done = removeR(root.Link[dir],data)
	if done {
		return root,true  
	}
	root.Balance += 1 - 2*dir
	switch root.Balance{
	case 1,-1 :
		return root,true
	case 0 :
		return root,false 
	}
	return removeBalance(root,dir)
}
type intkey int
func (k intkey) less(k2 key) bool { return k <k2.(intkey)}
func (k intkey) Eq(k2 key) bool {return k == k2.(intkey)}
func main() {
	var tree *Node
	fmt.Println("empty tree:")
	avl,_ :=json.MarshalIndent(tree,"","  ")
	fmt.Println(string(avl))
	fmt.Println("\n insert tree:")
	Insert(&tree, intkey(4))
	Insert(&tree, intkey(2))
	Insert(&tree, intkey(7))
	Insert(&tree, intkey(6))
	Insert(&tree, intkey(6))
	Insert(&tree, intkey(9))
	avl,_ = json.MarshalIndent(tree,"","  ")
	fmt.Println(string(avl))
	fmt.Println("\n Remove tree:")
	Remove(&tree,intkey(4))
	Remove(&tree,intkey(6))
	avl,_ = json.MarshalIndent(tree,"","  ")
	fmt.Println(string(avl))	
}


