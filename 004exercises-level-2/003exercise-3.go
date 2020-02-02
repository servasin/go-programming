package main

import "fmt"

const (
    a = 1
    b float64 = 2.24
)

func main()  {
    fmt.Println(a)
    fmt.Printf("%T\n", a)

    fmt.Println(b)
    fmt.Printf("%T\n", b)
}
