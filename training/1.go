package main

import (
    "fmt"
)

var c, python, java, bool

func Legal(n int) (Crazy bool) {
    Crazy = n > 0 // assign named return variable
    return         // returns Crazy automatically
}

func main() {
    fmt.Println(Legal(5))  // Output: true
    fmt.Println(Legal(-3)) // Output: false
}
