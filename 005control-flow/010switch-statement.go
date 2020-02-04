package main

import "fmt"

func main()  {
    name := "James"
    switch name {
    case "Bond":
        fmt.Println("Bond")
    case "James":
        fmt.Println("James")
    default:
        fmt.Println("James Bond")
    }
    fmt.Println("done.")
}
