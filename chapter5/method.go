package main

import "fmt"

type Point struct{ X, Y int }

func (p *Point) Render() {
    fmt.Printf("<%d,%d>\n", p.X, p.Y)
}


type IntPair [2]int

func (ip IntPair) First() int {
    return ip[0]
}

func (ip IntPair) Last() int {
    return ip[1]
}

type Strings []string

func (s Strings) Join(d string) string {
    sum := ""
    for _, v := range s {
        if sum != "" {
            sum += d
        }
        sum += v
    }
    return sum
}

func main() {
    p := &Point{X: 5, Y: 12}
    p.Render()

    ip := IntPair{1, 2}
    fmt.Println(ip.First())
    fmt.Println(ip.Last())

    fmt.Println(Strings{"A", "B", "C"}.Join(","))
}

