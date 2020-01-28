package even

import (
	"fmt"
)
type Text struct{
	Name string
	id int
}
var A Text
func init() {
	A.Name = "hello"
	A.id = 10
}
func Say_Hello_s(){
	fmt.Println("hello")
}

func say_hello(){
	fmt.Println("hello")
}
func main(){
	fmt.Println("sads")
}