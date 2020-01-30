package main

import "fmt"

func main()  {
  fmt.Println("Hello World")
  foo()

  fmt.Println("somthing here")

  for i := 0; i < 10; i++ {
    if i % 2 == 0 {
        fmt.Println(i)
        }
    }
}

func foo()  {
  fmt.Println("This is foo function")
}
