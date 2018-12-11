package main

import "fmt"

func main() {
    fruits := [3]string{"Apple", "Banana", "Cherry"}

    for i, s := range fruits {
        fmt.Printf("fruits[%d]=%s\n", i, s)
    }
}
