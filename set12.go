package main
import "fmt"
type user struct{
	id int
	name,loc string

}
func(u *user) greetings() string {
	return fmt.Sprintf("hi %s from %s",u.name,u.loc)
}
func newuser(id int, name , loc string) *user{
	id++
	return &user{id,name,loc}
}
func main() {
	u :=newuser(23,"riya","ind")
	fmt.Println(u.greetings())
}	