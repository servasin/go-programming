package main

import "fmt"

func main()  {
    y := []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
    for i, v := range y {
        fmt.Println(i, v)
    }
    fmt.Printf("TYPE %T", y)
}
