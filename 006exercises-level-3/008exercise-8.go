package main

import "fmt"

func main()  {
        switch {
        case true:
            fmt.Println("This is print")
        case false:
            fmt.Println("This is no print")
        case false:
            fmt.Println("This is also no print")
        default:
            fmt.Println("Are this print ?")
        }
}
