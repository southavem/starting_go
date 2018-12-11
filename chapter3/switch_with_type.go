package main

import "fmt"

func main() {
    var x interface{} = "string"
    
    switch x.(type) {
    case bool:
        fmt.Println("bool")
    case int, uint:
        fmt.Println("integer or unsigned integer")
    case string:
        fmt.Println("string")
    default:
        fmt.Println("don't know")
    }
}
