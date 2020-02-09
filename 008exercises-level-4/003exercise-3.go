package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
    fmt.Println(x)

    a := append(x[:5])
    fmt.Println(a)

    b := append(x[5:])
    fmt.Println(b)

    c := append(x[2:7])
    fmt.Println(c)

    d := append(x[1:6])
    fmt.Println(d)
}
