package main

import "fmt"

func main()  {
    jb := []string{"James", "Bond", "Chocolate", "Martini"}
    fmt.Println(jb)
    mp := []string{"Miss", "Moneypenny", "Strawberry", "Hezlnut"}
    fmt.Println(mp)

    xp := [][]string{jb, mp}
    fmt.Println(xp)
}
