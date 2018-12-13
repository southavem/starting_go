package main

import "fmt"

type Point struct {
    X int
    Y int
}

func main() {
    var pt Point
    fmt.Println(pt.X)
    fmt.Println(pt.Y)

    pt.X = 10
    pt.Y = 8
    fmt.Println(pt.X)
    fmt.Println(pt.Y)
}
