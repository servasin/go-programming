package main

import "fmt"

func main()  {
    favSport := "football"

    switch favSport {
    case "football":
        fmt.Println("I like", favSport)
    case "boxing":
        fmt.Println("I like", favSport)
    case "golf":
        fmt.Println("I like", favSport)
    default:
        fmt.Println("I like another sport")
    }
}
