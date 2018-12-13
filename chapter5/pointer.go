package main

import "fmt"

func main() {
    var p *int

    fmt.Println(p == nil)



    var i int
    p1 := & i
    fmt.Printf("%T\n", p1)
    p2 := & p1
    fmt.Printf("%T\n", p2)

    var j int
    p3 := & j
    j = 10
    fmt.Println(*p3)
    *p3 = 15
    fmt.Println(j)


    p4 := &[3]int{1, 2, 3}
    pow(p4)
    fmt.Println(p4)
}


func pow(p *[3]int) {
    i := 0
    for i < 3 {
        (*p)[i] = (*p)[i] * (*p)[i]
        i++
    }
}

