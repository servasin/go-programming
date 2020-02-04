package main

import "fmt"

func main()  {
    if x := 42; x == 42 {
        fmt.Println("one")
    }
    fmt.Println("here's statement")
    fmt.Println("something else")
    // fmt.Println(x)  x undefined
}
