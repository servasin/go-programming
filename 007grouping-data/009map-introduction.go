package main

import "fmt"

func main()  {
    m := map[string]int{
        "James":      32,
        "Moneypenny": 24,
    }
    fmt.Println(m)

    fmt.Println(m["James"])
    fmt.Println(m["Banana"])

    v, ok := m["Banana"]
    fmt.Println(v)
    fmt.Println(ok)

    if v, ok := m["Moneypenny"]; ok {
        fmt.Println("THIS IS IF PRINT", v)
    }
}
