package main

import "fmt"

func main()  {
    x := []int{4, 5, 7, 8, 42}
    fmt.Println(x)
    x = append(x, 77, 88, 99, 1014)
    fmt.Println(x)

    y := []int{234, 456, 789}
    x = append(x, y...)
    fmt.Println(x)
}
