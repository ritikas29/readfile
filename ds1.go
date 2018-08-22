package main
import "fmt"
type node struct {
	prev *node
	next *node
	key interface{}
}
type list struct {
	head *node
	tail *node
}
func(l *list) insert(key interface{}){
	list :=&node {
		next: l.head,
		key: key,
	}
	if l.head !=nil {
		l.head.prev = list
	}
	l.head = list
	L := l.head
	for L.next !=nil {
		L=L.next
	}
	l.tail =L
}
func (L *list ) display() {
	list := L.head
	for list !=nil {
		fmt.Printf("%+v ->",list.key)
		list = list.next
	}
	fmt.Println()
}
func display(list *node){
	for list !=nil{
		fmt.Printf("%v ->",list.key)
		list=list.next
	}
	fmt.Println()
}
func showbackwards(list *node) {
	for list !=nil {
		fmt.Printf("%v <-",list.key)
		list = list.prev
	}
	fmt.Println()
}
func(L *list) reverse(){
	curr := L.head
	var prev *node
	L.tail = L.head
	for curr != nil{
		next :=curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	L.head = prev
	display(L.head)
}
func main(){
	link := list {}
	link.insert(5)
	link.insert(2)
	link.insert(7)
	link.insert(21)
	link.insert(8)
	link.insert(11)
	fmt.Println("\n ---------------")
	fmt.Printf("head:%v\n",link.head.key)
	fmt.Printf("tail: %v\n",link.tail.key)
	link.display()
	fmt.Println("\n-------")
	fmt.Printf("head: %v\n",link.head.key)
	fmt.Printf("tail: %v\n",link.tail.key)
	link.reverse()
	fmt.Println("\n---------")
}