package main

import "fmt"

func main() {
LOOP:
    for {
        for {
            for {
                fmt.Println("開始")
                break LOOP
            }
            fmt.Println("ここは通らない")
        }
        fmt.Println("ここは通らない")
    }
    fmt.Println("終了")
}
