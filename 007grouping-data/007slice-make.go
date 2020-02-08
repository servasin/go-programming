package main

import "fmt"

func main()  {
    x := make([]int, 10, 12)
    x[0] = 42
    x[9] = 999
    fmt.Println(x)
    fmt.Println(len(x))
    fmt.Println(cap(x))

    x = append(x, 343)

    fmt.Println(x)
    fmt.Println(len(x))
    fmt.Println(cap(x))

    x = append(x, 344)
    x = append(x, 345)

    fmt.Println(x)
    fmt.Println(len(x))
    fmt.Println(cap(x))
}
