package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main()  {
    s := fmt.Sprintf("%v %T\t %v %T\t %v %T\t", x, x, y, y, z, z)
    fmt.Println(s)
    fmt.Printf("type of \"s\" is %T", s)
}
