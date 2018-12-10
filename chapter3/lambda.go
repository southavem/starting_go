package main

import (
    "fmt"
)

func main() {
    var f func(int, int) int
    f = func(x, y int) int { return x + y }
    fmt.Println(f(2, 3))
    fmt.Printf("%T\n", f)
}
