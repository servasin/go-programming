package main

import "fmt"

func main()  {
    m := map[string][]string {
        "James_Bond":   []string{"football", "whiskey", "cola"},
        "Miss_Penny":   []string{"books", "coffee", "driving"},
        "Dr_Who":       []string{"computres", "golf", "sing"},
    }
    //fmt.Println(m)
    for k, v := range m {
        fmt.Println("this is record for", k)
        for i, v2 := range v {
            fmt.Println("\t", i, v2)
        }
    }
}
