package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("3 + 4 =", add(3, 4))
	fmt.Println("4 + 5 =", 4 + 5)
}

/*

Println is a method, fmt has sent it to the package "fmt"


*/

