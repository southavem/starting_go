package main

import "fmt"

func runDefer() {
    defer fmt.Println("defer")
    fmt.Println("done")
}

func main() {
    fmt.Println("main開始")
    runDefer()
    fmt.Println("main終了")

    file, err := os.Open("./test.txt")
    if err != nil {
        return
    }

    defer file.Close()
}
