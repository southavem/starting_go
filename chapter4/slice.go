package main

import "fmt"

func main() {
    s := make([]int, 10)
    fmt.Println(s)

    a := make([]float64, 3)
    fmt.Println(a)
    a[0] = 3.14
    fmt.Println(a)
    a[1] = 6.28
    fmt.Println(a)
    fmt.Println(a[0])
    // fmt.Println(a[3])

    fmt.Println(len(a))
}
