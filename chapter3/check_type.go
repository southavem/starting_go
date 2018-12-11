package main

import "fmt"

func main() {
    var x interface{} = 3
    i := x.(int)
    f := x.(float64)
    fmt.Println(i)
    fmt.Println(f)
}
