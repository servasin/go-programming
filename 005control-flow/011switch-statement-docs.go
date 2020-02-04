package main

import "fmt"

func main()  {
    switch {
    case false:
        fmt.Println("Не выведется")
    case true:
        fmt.Println("Выведется")
    }
}
