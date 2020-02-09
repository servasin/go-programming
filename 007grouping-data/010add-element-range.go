package main

import "fmt"

func main()  {
    m := map[string]int{
        "James":      32,
        "Moneypenny": 24,
    }
    fmt.Println(m)

    m["Tom"] = 44
    m["Anne"] = 27
    fmt.Println(m)

    for k, v := range m {
        fmt.Println(k, v)
    }

    xi := []int{4, 5, 6, 7, 42}

    for i, v := range xi {
        fmt.Println(i, v)
    }

}
