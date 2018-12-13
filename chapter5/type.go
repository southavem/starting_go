package main

import "fmt"

type (
    IntPair      [2]int
    Strings      []string
    AreaMap      map[string][2]float64
 )

func main() {
    type MyInt int

    var n1 MyInt = 5
    n2 := MyInt(7)
    fmt.Println(n1)
    fmt.Println(n2)

    pair := IntPair{1, 2}
    fmt.Println(pair)
    strs := Strings{"Apple", "Banana", "Cherry"}
    fmt.Println(strs)
    amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
    fmt.Println(amap)
}
