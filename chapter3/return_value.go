package main

import (
    "fmt"
)

func doSomething() (a int) {
    a = 10
    return
}

func main() {
    fmt.Println(doSomething())
}
