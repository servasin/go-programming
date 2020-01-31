package main

import "fmt"

var x int
var y string
var z bool

func main()  {
    //fmt.Println(x, y, z)
    fmt.Printf("%T zero value %#v\n", x, x)
    fmt.Printf("%T zero value %#v\n", y, y)
    fmt.Printf("%T zero value %#v\n", z, z)
}
