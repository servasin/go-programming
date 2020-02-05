package main

import "fmt"

func main()  {
    n := 20
    for n < 80 {
        if n % 15 == 0 {
            fmt.Printf("%v делится и на 3 и на 5\n", n)
        } else if n % 5 == 0 {
            fmt.Printf("%v делится на 5\n", n)
        } else if n % 3 == 0{
            fmt.Printf("%v делится на 3\n", n)
        } else {
            fmt.Printf("%v не делится на 3 или 5\n", n)
        }
        n++
    }
}
