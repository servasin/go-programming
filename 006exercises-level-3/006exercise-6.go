package main

import "fmt"

func main()  {
    n := 20
    for n < 100 {
        if n % 2 == 0 {
            fmt.Printf("%v четное\n", n)
        }
        n++
    }

    x := "James Bond"

	if x == "James Bond" {
		fmt.Println(x)
	}
}
