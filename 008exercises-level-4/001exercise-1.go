package main

import "fmt"

func main()  {
    x := [5]int{11, 22, 33, 44, 55}

    for i, v := range x {
        fmt.Println(i, v)
    }

    fmt.Printf("TYPE %T\n", x)
}
