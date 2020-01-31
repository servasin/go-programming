package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)         // 42
	fmt.Printf("%T\n", y)  // int
	fmt.Printf("%b\n", y)  // 101010
	fmt.Printf("%x\n", y)  // 2a
	fmt.Printf("%#x\n", y)  // 0x2a

    y = 911
    fmt.Printf("%#x\n%b\n%x\n", y, y, y) // 0x38f 111000111 38f

    s := fmt.Sprintf("%#x\n%b\n%x\n", y, y, y)
    fmt.Println(s)
}
