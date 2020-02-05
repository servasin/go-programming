package main

import "fmt"

func main()  {
    for i := 10; i <= 100; i++ {
        if i % 4 == 0 {
            fmt.Println(i)
        }
    }

    for j := 10; j <= 100; j++ {
        fmt.Printf("Когда %v делиться на 4, остаток (modulus) %v\n", j, j%4)
    }
}
