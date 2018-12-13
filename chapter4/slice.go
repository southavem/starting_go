package main

import "fmt"

func main() {
    s := make([]int, 10)
    fmt.Println(s)

    a := make([]float64, 3)
    fmt.Println(a)
    a[0] = 3.14
    fmt.Println(a)
    a[1] = 6.28
    fmt.Println(a)
    fmt.Println(a[0])
    // fmt.Println(a[3])

    fmt.Println(len(a))

    s1 := make([]int, 5)
    fmt.Println(s1)
    fmt.Println(len(s1))
    fmt.Println(cap(s1))

    s2 := make([]int, 5, 10)
    fmt.Println(s2)
    fmt.Println(len(s2))
    fmt.Println(cap(s2))

    b := [...]int{1,2,3,4,5}
    fmt.Println(b[len(b)-2:])

    c := []int{1, 2, 3}
    c = append(c, 4)
    fmt.Println(c)
    c = append(c, 5, 6, 7)
    fmt.Println(c)
}
