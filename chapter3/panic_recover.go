package main

import "fmt"

func main() {
    defer func() {
        if x := recover(); x != nil {
            fmt.Println("hoge")
        }
    }()
    panic("panic!")

    fmt.Println("これは実行されない")
}
        
