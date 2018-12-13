package main

import "fmt"

func main() {
    m1 := make(map[int]string)
    m1[1]  = "US"
    m1[81] = "Japan"
    m1[86] = "China"

    fmt.Println(m1)


    m2 := map[int]string{1: "Taro", 2:"Hanako", 3:"Jiro"}
    fmt.Println(m2)
}
