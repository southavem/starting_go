package main

import (
    "fmt"
)

func plus(x, y int) int {
    return x + y
}

var plusAlias = plus

func main() {
    var f func(int, int) int
    f = func(x, y int) int { return x + y }
    fmt.Println(f(2, 3))
    fmt.Printf("%T\n", f)

    fmt.Printf("%#v\n", func(x, y int) int { return x + y } )
    fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3) )

    fmt.Println(plusAlias(5, 4))
}
