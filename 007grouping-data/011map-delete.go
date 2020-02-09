package main

import "fmt"

func main()  {
    m := map[string]int{
        "James":        32,
        "Moneypenny":   24,
    }
    fmt.Println(m)

    delete(m, "James")
    fmt.Println(m)

    delete(m, "Ivan")
    fmt.Println(m)

    fmt.Println(m["Moneypenny"])
    fmt.Println(m["Ivan"])

    if v, ok := m["Moneypenny"]; ok {
        fmt.Println("value:", v)
        delete(m, "Moneypenny")
    }
    fmt.Println(m)
}
