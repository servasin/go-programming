package main

import "fmt"

func main()  {
    a := (42 == 42)
    b := (42 <= 43)
    c := (42 >= 43)
    d := (42 != 43)
    e := (42 < 43)
    f := (42 > 43)

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
    fmt.Println(f)
}
