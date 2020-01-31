package main

import "fmt"

var y = 43

// объявление, определение, инициализация переменной
// declare & assing = initilization
var z int = 1

func main()  {
    // fmt.Println(x)  udefined x
    x := 22
    fmt.Println(x)
    fmt.Println(y)
    foo()

    fmt.Println(z)
}

func foo()  {
    fmt.Println("again",y)
}
