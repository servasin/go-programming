package main

import "fmt"

var x bool

func main()  {
    fmt.Println(x)
    x = true
    fmt.Println(x)

    a := 7
    b := 43
    fmt.Println(a <= b)

    c := 66
    d := 54
    fmt.Println(c >= d)
}
