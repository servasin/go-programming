package main

import "fmt"

type myType int
var x myType

func main()  {
    fmt.Println(x)
    fmt.Printf("%v %T\n", x, x)
    x = 42
    fmt.Println(x)
    fmt.Printf("%v %T\n", x, x)
    //x = true
    //fmt.Printf("%v %T\n", x, x)
    //cannot use true (type bool) as type myType in assing
}
