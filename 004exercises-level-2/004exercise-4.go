package main

import "fmt"

func main()  {
    a := 42
    fmt.Printf("%d\t%b\t\t%#x\n", a, a, a)

    b := a << 1
    fmt.Printf("%d\t%b\t\t%#x\n", b, b, b)

    c := b << 1
    fmt.Printf("%d\t%b\t%#x\n", c, c, c)
}
