package main

import "fmt"

func main() {
    var x interface{} = 3

    i, isInt := x.(int)
    f, isFloat64 := x.(float64)
    s, isString := x.(string)

    fmt.Println(i)
    fmt.Println(isInt)
    fmt.Println(f)
    fmt.Println(isFloat64)
    fmt.Println(s)
    fmt.Println(isString)

    if x == nil {
        fmt.Println("x is nil")
    } else if i, isInt := x.(int); isInt {
        fmt.Printf("x is integer : %d\n", i)
    } else if s, isString := x.(string); isString {
        fmt.Println(s)
    } else {
        fmt.Println("unsupported type!")
    }
}
