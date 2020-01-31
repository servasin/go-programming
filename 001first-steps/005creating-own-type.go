package main

import (
	"fmt"
)

var a int
type myType int
var b myType

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	//a = b
	//fmt.Println(a)
	//fmt.Printf("%T\n", a)
	//cannot use b (type myType) as type int in assignment
}
