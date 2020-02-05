package main

import "fmt"

func main()  {
    bd := 1994
    for {
        if bd > 2020 {
            break
        }
        fmt.Println(bd)
        bd++
    }
}
