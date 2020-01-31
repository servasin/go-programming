package main

import (
	"fmt"
)

var y = 42
var s = "this is string"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	//s = 33
	//fmt.Printf("%T\n", s)
    // нельзя присвоить переменной значение другого типа

}
